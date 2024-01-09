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
  private user: User | null = null;
  private isAuthenticated$: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);

  constructor(private http: HttpClient) { super(); }

  public getUser(): User | null {
    return this.user;
  }

  public signIn(email: string, password: string) {
    // Replace with your actual sign-in API endpoint
    const url = '/auth/signin';

    return this.http.post<User>(`${this.apiUrl}/${url}`, { email, password }).pipe(
      tap(user => {
        this.user = user;
        this.isAuthenticated$.next(true);
      })
    );
  }

  public signUp(email: string, password: string) {
    // Replace with your actual sign-up API endpoint
    const url = '/auth/signup';

    return this.http.post<User>(`${this.apiUrl}/${url}`, { email, password }).pipe(
      tap(user => {
        this.user = user;
        this.isAuthenticated$.next(true);
      })
    );
  }
}
