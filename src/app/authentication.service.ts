import { Injectable } from '@angular/core';
import { BaseHttpService } from './base-http.service';

import { User } from './models/user';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService extends BaseHttpService {

  private user: User | null = null
  private isAuthenticated$: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false)

  constructor() { super() }

  public getUser(): User | null {
    return this.user
  }
}
