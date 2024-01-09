// authentication.service.ts
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BaseHttpService } from './base-http.service';

import { User } from './models/user';
import { BehaviorSubject, tap } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService extends BaseHttpService {
  private isAuthenticated$: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
  private rawToken: string | null = null

  constructor(private http: HttpClient) { super(); }

  public getToken(): string | null {
    return this.rawToken
  }

  public signIn(email: string | null, password: string | null) {
    // Replace with your actual sign-in API endpoint
    const url = '/auth/signin';

    return this.http.post<any>(`${this.apiUrl}${url}`, { email, password }).pipe(
      tap(token => {
        this.rawToken = token
        this.isAuthenticated$.next(true);
      })
    );
  }

  public signUp(email: string | null, password: string | null) {
    // Replace with your actual sign-up API endpoint
    const url = '/auth/signup';

    return this.http.post<any>(`${this.apiUrl}${url}`, { email, password }).pipe(
      tap(token => {
        this.rawToken = token
        this.isAuthenticated$.next(true);
      })
    );
  }
}
