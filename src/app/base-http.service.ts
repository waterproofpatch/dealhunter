import { Injectable } from '@angular/core';
import { environment } from '../environments/environment';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BaseHttpService {
  public apiUrl: string = environment.apiUrl
  public isLoading$: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false)

  constructor() { }
}
