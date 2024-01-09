import { Injectable } from '@angular/core';
import { environment } from '../environments/environment';
import { BehaviorSubject, Observable, finalize } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BaseHttpService {
  public apiUrl: string = environment.apiUrl
  public isLoading$: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false)

  constructor() { }

  public withLoading<T>(httpCall: Observable<T>): Observable<T> {
    this.isLoading$.next(true);
    return httpCall.pipe(
      finalize(() => this.isLoading$.next(false))
    );
  }

}
