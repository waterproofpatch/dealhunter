import { Component, OnInit } from '@angular/core';
import { Input } from '@angular/core';

import moment from 'moment'

import { Deal } from '../models/deal';
import { LocationService } from '../location.service';
import { DealsService } from '../deals.service';
import { AuthenticationService } from '../authentication.service';
import { Address } from '../models/address';
@Component({
  selector: 'app-deal',
  templateUrl: './deal.component.html',
  styleUrl: './deal.component.css'
})
export class DealComponent implements OnInit {
  secondsSinceCreation: number = 0
  secondsSinceLastUpvote: number = 0
  dealAddress: string = "Fetching address..."

  @Input() deal: Deal | undefined
  constructor(
    public locationService: LocationService,
    public dealsService: DealsService,
    public authenticationService: AuthenticationService) {

  }

  ngOnInit(): void {
    this.getSecondsSinceCreation()
    this.getSecondsSinceLastUpvote()
    if (this.deal?.Location) {
      this.locationService.getAddressForLocation(this.deal?.Location).subscribe((address: Address) => {
        this.dealAddress = address.Address
      })
    }
  }

  public getSecondsSinceDate(dateString: string): number {
    // const dateString = '2024-01-07T00:42:46.786403-05:00';
    if (!this.deal) {
      return 0
    }
    const date = moment(dateString);
    const now = moment();

    const diffInMinutes = now.diff(date, 'minutes');
    return diffInMinutes
  }

  private getSecondsSinceLastUpvote(): void {
    if (!this.deal) {
      return
    }
    this.secondsSinceLastUpvote = this.getSecondsSinceDate(this.deal.LastUpvoteTime)
  }

  private getSecondsSinceCreation(): void {
    if (!this.deal) {
      return
    }
    this.secondsSinceCreation = this.getSecondsSinceDate(this.deal.CreatedAt)
  }

}
