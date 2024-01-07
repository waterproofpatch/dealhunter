import { Injectable } from '@angular/core';

import { UserMeta } from './models/user-meta';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiUrl = 'http://localhost:8000'; // replace with your API endpoint
  public userMeta$: BehaviorSubject<UserMeta> = new BehaviorSubject<UserMeta>({ Token: "", CreatedAt: "", ID: 0 })

  constructor(private http: HttpClient) {
    this.getUserMetaHttp().subscribe((userMeta: UserMeta) => {
      this.userMeta$.next(userMeta)
    })
  }

  private getUserMetaHttp(): Observable<any> {
    return this.http.get(`${this.apiUrl}/user-meta`);
  }
}
