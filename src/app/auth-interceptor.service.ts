import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpRequest, HttpHandler } from '@angular/common/http';
import { AuthenticationService } from './authentication.service';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  accessToken: string | undefined = undefined
  constructor(private authenticationService: AuthenticationService) {
    authenticationService.jwtAccessToken$.subscribe((x) => {
      this.accessToken = x?.accessToken
    })
  }

  intercept(req: HttpRequest<any>, next: HttpHandler) {
    if (this.accessToken) {
      // Clone the request and replace the original headers with
      // cloned headers, updated with the authorization.
      console.log("We are authenticated, sending accessToken")
      const authReq = req.clone({
        headers: req.headers.set('Authorization', this.accessToken)
      });
      // Send cloned request with header to the next handler.
      return next.handle(authReq);
    } else {
      return next.handle(req)
    }
  }
}
