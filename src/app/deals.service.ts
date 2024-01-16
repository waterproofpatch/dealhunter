import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, BehaviorSubject, map } from 'rxjs';

import { Deal } from './models/deal';
import { LocationService } from './location.service';

import { environment } from '../environments/environment';
import { BaseHttpService } from './base-http.service';
import { MatDialog } from '@angular/material/dialog';

@Injectable({
  providedIn: 'root'
})
export class DealsService extends BaseHttpService {
  private deals$: BehaviorSubject<Deal[]> = new BehaviorSubject<Deal[]>([])

  constructor(private locationService: LocationService, private http: HttpClient, private _dialog: MatDialog) {
    super(_dialog)
    this.withLoading(this.getDealsHttp()).subscribe((deals: Deal[]) => {
      this.deals$.next(deals)
    })
  }

  public addDeal(deal: Deal, address: string) {
    this.withLoading(this.createDealHttp(deal, address)).subscribe((deals: Deal[]) => {
      this.deals$.next(deals);
    });
  }

  public deleteDeal(deal: Deal) {
    this.withLoading(this.deleteDealHttp(deal)).subscribe((deals: Deal[]) => {
      this.deals$.next(deals)
    })
  }

  public downvoteDeal(deal: Deal) {
    this.withLoading(this.downvoteDealHttp(deal)).subscribe((deals: Deal[]) => {
      this.deals$.next(deals)
    })
  }
  public upvoteDeal(deal: Deal) {
    this.withLoading(this.upvoteDealHttp(deal)).subscribe((deals: Deal[]) => {
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

  private createDealHttp(deal: Deal, address: string): Observable<any> {
    /* fill in location */
    deal.Location = this.locationService.location
    return this.http.post(`${this.apiUrl}/deals?address=${address}`, deal);
  }
  private downvoteDealHttp(deal: Deal): Observable<any> {
    return this.http.put(`${this.apiUrl}/deals/${deal.ID}?vote=down`, {});
  }
  private upvoteDealHttp(deal: Deal): Observable<any> {
    return this.http.put(`${this.apiUrl}/deals/${deal.ID}?vote=up`, {});
  }
  private deleteDealHttp(deal: Deal): Observable<any> {
    return this.http.delete(`${this.apiUrl}/deals/${deal.ID}`, {});

  }


}
