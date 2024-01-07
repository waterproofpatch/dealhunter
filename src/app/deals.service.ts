import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, BehaviorSubject, map } from 'rxjs';

import { Deal } from './models/deal';
import { LocationService } from './location.service';

@Injectable({
  providedIn: 'root'
})
export class DealsService {
  private apiUrl = 'http://localhost:8000'; // replace with your API endpoint
  private deals$: BehaviorSubject<Deal[]> = new BehaviorSubject<Deal[]>([])

  constructor(private locationService: LocationService, private http: HttpClient) {
    this.getDealsHttp().subscribe((deals: Deal[]) => {
      this.deals$.next(deals)
    })
  }

  public addDeal(deal: Deal) {
    this.createDealHttp(deal).subscribe((deals: Deal[]) => {
      this.deals$.next(deals)
    })
  }

  public upvoteDeal(deal: Deal) {
    this.upvoteDealHttp(deal).subscribe((deals: Deal[]) => {
      this.deals$.next(deals)
    })
  }

  public getDeals() {
    return this.deals$.pipe(
      map(deals => deals.sort((a, b) => this.locationService.calculateDistance(a.Location) - this.locationService.calculateDistance(b.Location)))
    );
  }

  private getDealsHttp(): Observable<any> {
    return this.http.get(`${this.apiUrl}/deals`);
  }

  private createDealHttp(deal: Deal): Observable<any> {
    /* fill in location */
    deal.Location = this.locationService.location
    return this.http.post(`${this.apiUrl}/deals`, deal);
  }
  private upvoteDealHttp(deal: Deal): Observable<any> {
    return this.http.put(`${this.apiUrl}/deals/${deal.ID}?vote=up`, {});
  }


}
