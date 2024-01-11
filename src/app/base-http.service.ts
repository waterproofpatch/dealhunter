import { Injectable } from '@angular/core';
import { environment } from '../environments/environment';
import { BehaviorSubject, Observable, finalize } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { of } from 'rxjs'

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
      catchError(err => {
        console.error('An error occurred:', err.error);
        throw new Error("Some error")
      }),
      finalize(() => this.isLoading$.next(false))
    );
  }


}
