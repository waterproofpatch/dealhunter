import { Injectable } from '@angular/core';
import { environment } from '../environments/environment';
import { BehaviorSubject, Observable } from 'rxjs';
import { catchError, finalize } from 'rxjs/operators';
import { MatDialog } from '@angular/material/dialog';
import { ErrorDialogComponent } from './error-dialog/error-dialog.component'; // Import your Error Dialog Component

@Injectable({
  providedIn: 'root'
})
export class BaseHttpService {
  public apiUrl: string = environment.apiUrl
  public isLoading$: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false)

  constructor(private dialog: MatDialog) { } // Inject MatDialog

  public withLoading<T>(httpCall: Observable<T>): Observable<T> {
    this.isLoading$.next(true);
    return httpCall.pipe(
      catchError(err => {
        this.dialog.open(ErrorDialogComponent, { // Open the Error Dialog
          data: {
            message: 'An error occurred: ' + err.error,
            status_code: err.status
          }
        });
        throw new Error(err.error)
      }),
      finalize(() => this.isLoading$.next(false))
    );
  }
}
