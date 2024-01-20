// authentication.service.ts
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BaseHttpService } from './base-http.service';
import { BehaviorSubject, tap } from 'rxjs';

import { JwtAccessToken } from './models/tokens';
import { MatDialog } from '@angular/material/dialog';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService extends BaseHttpService {
  public isAuthenticated$: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
  public jwtAccessToken$: BehaviorSubject<JwtAccessToken | null> = new BehaviorSubject<JwtAccessToken | null>(null)
  public userId$: BehaviorSubject<number> = new BehaviorSubject<number>(0)

  constructor(private http: HttpClient, private _dialog: MatDialog) {
    super(_dialog);
    // display access tokens as they arrive
    this.jwtAccessToken$.subscribe((x: JwtAccessToken | null) => {
      if (!x) {
        return
      }
      if (x.isExpired()) {
        console.log("Expired token.")
        // this.signOut()
        // this.jwtAccessToken$.next(null)
        // return
      }
      this.userId$.next(x.id)
      this.isAuthenticated$.next(true);
      console.log(`parsed JwtAccessToken: id=${x.id}, email=${x.email}, exp=${x.exp}, is expired? ${x.isExpired()}, expires in ${x.secondsUntilExpiration()}`)
    })

    // if we had a previously cached access token, use that one
    let existingToken = localStorage.getItem('token')
    if (existingToken) {
      console.log(`Have cached accessToken ${existingToken}...`)
      this.jwtAccessToken$.next(new JwtAccessToken(existingToken))
    }
  }

  public signOut(): void {
    console.log("Logging out...")
    this.isAuthenticated$.next(false)
    localStorage.removeItem("token")
    this.jwtAccessToken$.next(null)
    this.userId$.next(0)
  }

  public signIn(email: string | null, password: string | null) {
    // Replace with your actual sign-in API endpoint
    const url = '/auth/signin';

    return this.withLoading(this.http.post<any>(`${this.apiUrl}${url}`, { email, password })).pipe(
      tap(token => {
        this.jwtAccessToken$.next(new JwtAccessToken(token.AccessToken))
        localStorage.setItem('token', token.AccessToken); // Save token to local storage
      })
    );
  }

  public signUp(email: string | null, password: string | null) {
    // Replace with your actual sign-up API endpoint
    const url = '/auth/signup';

    return this.withLoading(this.http.post<any>(`${this.apiUrl}${url}`, { email, password })).pipe(
      tap(token => {
        this.jwtAccessToken$.next(new JwtAccessToken(token.AccessToken))
        localStorage.setItem('token', token.AccessToken); // Save token to local storage
      })
    );
  }
  public refreshToken() {
    console.log("Issuing refresh request...")
    // Replace with your actual sign-up API endpoint
    const url = '/auth/refresh';
    return this.withLoading(this.http.get<any>(`${this.apiUrl}${url}`)).pipe(
      tap(token => {
        console.log(`Got refreshed access token ${token}`)
        this.jwtAccessToken$.next(new JwtAccessToken(token.AccessToken))
        localStorage.setItem('token', token.AccessToken); // Save token to local storage
      })
    );
  }
}
