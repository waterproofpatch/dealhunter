import { Injectable } from '@angular/core';
import { BehaviorSubject, map } from 'rxjs';

import { Deal } from './models/deal';
import { LocationService } from './location.service';

@Injectable({
  providedIn: 'root'
})
export class DealsService {

  private deals$: BehaviorSubject<Deal[]> = new BehaviorSubject<Deal[]>([])
  private deals: Deal[] = []

  constructor(private locationService: LocationService) { }

  public addDeal(deal: Deal) {
    this.deals.push(deal)
    this.deals$.next(this.deals)
  }

  public getDeals() {
    return this.deals$.pipe(
      map(deals => deals.sort((a, b) => this.locationService.calculateDistance(a.location) - this.locationService.calculateDistance(b.location)))
    );
  }

}
