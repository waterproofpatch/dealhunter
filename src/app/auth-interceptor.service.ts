import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpRequest, HttpHandler } from '@angular/common/http';
import { AuthenticationService } from './authentication.service';
import { Observable, catchError, switchMap, throwError } from 'rxjs';
import { HttpEvent, HttpErrorResponse } from '@angular/common/http';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  accessToken: string | undefined = undefined
  constructor(private authenticationService: AuthenticationService) {
    this.authenticationService.jwtAccessToken$.subscribe((x) => {
      this.accessToken = x?.accessToken
    })
  }

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    if (this.accessToken) {
      const authReq = req.clone({
        headers: req.headers.set('Authorization', `Bearer ${this.accessToken}`),
        withCredentials: true,
      });
      return next.handle(authReq).pipe(
        catchError((error: HttpErrorResponse) => {
          if (error.status === 419) { // Assuming 419 is the status code for expired tokens
            return this.authenticationService.refreshToken().pipe(
              switchMap((token: any) => {
                console.log(`Issuing re-request for ${req.url}`)
                const refreshedReq = req.clone({
                  headers: req.headers.set('Authorization', `Bearer ${token}`),
                  withCredentials: true,
                });
                return next.handle(refreshedReq);
              })
            );
          } else {
            // something wrong with refresh token, need to force user to 
            // reauth
            // this.authenticationService.signOut()
            return throwError(error);
          }
        })
      );
    } else {
      const authReq = req.clone({
        withCredentials: true,
      });
      return next.handle(authReq);
    }
  }

}
