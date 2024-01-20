import { Injectable } from '@angular/core';
import { BaseHttpService } from './base-http.service';
import { MatDialog } from '@angular/material/dialog';
import { HttpClient } from '@angular/common/http';

import { User } from './models/user';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProfileService extends BaseHttpService {

  public user$: BehaviorSubject<User> = new BehaviorSubject<User>({ ID: 0, Reputation: 0 })

  constructor(private http: HttpClient, private _dialog: MatDialog) {
    super(_dialog)
    this.withLoading(this.getProfileHttp()).subscribe((user: User) => {
      this.user$.next(user)
    })
  }

  private getProfileHttp(): Observable<User> {
    return this.http.get<User>(this.apiUrl + "/profile")
  }
}
