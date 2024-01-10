// authentication.service.ts
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BaseHttpService } from './base-http.service';
import { BehaviorSubject, tap } from 'rxjs';

import { JwtAccessToken } from './models/tokens';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService extends BaseHttpService {
  public isAuthenticated$: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
  public jwtAccessToken$: BehaviorSubject<JwtAccessToken | null> = new BehaviorSubject<JwtAccessToken | null>(null)

  constructor(private http: HttpClient) {
    super();
    // display access tokens as they arrive
    this.jwtAccessToken$.subscribe((x: JwtAccessToken | null) => {
      if (!x) {
        return
      }
      this.isAuthenticated$.next(true);
      console.log(`parsed JwtAccessToken: id=${x.id}, email=${x.email}`)
    })

    // if we had a previously cached access token, use that one
    let existingToken = localStorage.getItem('token')
    if (existingToken) {
      console.log(`Have cached accessToken ${existingToken}...`)
      this.jwtAccessToken$.next(new JwtAccessToken(existingToken))
    }
  }

  public signIn(email: string | null, password: string | null) {
    // Replace with your actual sign-in API endpoint
    const url = '/auth/signin';

    return this.http.post<any>(`${this.apiUrl}${url}`, { email, password }).pipe(
      tap(token => {
        this.jwtAccessToken$.next(new JwtAccessToken(token.AccessToken))
        localStorage.setItem('token', token.AccessToken); // Save token to local storage
      })
    );
  }

  public signUp(email: string | null, password: string | null) {
    // Replace with your actual sign-up API endpoint
    const url = '/auth/signup';

    return this.http.post<any>(`${this.apiUrl}${url}`, { email, password }).pipe(
      tap(token => {
        this.jwtAccessToken$.next(new JwtAccessToken(token.AccessToken))
        localStorage.setItem('token', token.AccessToken); // Save token to local storage
      })
    );
  }
}
