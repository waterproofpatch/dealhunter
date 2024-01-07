import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

import { Deal } from './models/deal';

@Injectable({
  providedIn: 'root'
})
export class DealsService {

  private deals$: BehaviorSubject<Deal[]> = new BehaviorSubject<Deal[]>([])
  private deals: Deal[] = []

  constructor() { }

  public addDeal(deal: Deal) {
    this.deals.push(deal)
    this.deals$.next(this.deals)
  }

  public getDeals(): BehaviorSubject<Deal[]> {

    return this.deals$
  }
}
