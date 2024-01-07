import { Component, OnInit } from '@angular/core';
import { Input } from '@angular/core';

import moment from 'moment'

import { Deal } from '../models/deal';
import { LocationService } from '../location.service';
@Component({
  selector: 'app-deal',
  templateUrl: './deal.component.html',
  styleUrl: './deal.component.css'
})
export class DealComponent implements OnInit {

  @Input() deal: Deal | undefined
  constructor(public locationService: LocationService) {

  }

  ngOnInit(): void {
  }

  public getSecondsSinceCreation(): number {
    // const dateString = '2024-01-07T00:42:46.786403-05:00';
    if (!this.deal) {
      return 0
    }
    const dateString = this.deal.CreatedAt
    const date = moment(dateString);
    const now = moment();

    const diffInMinutes = now.diff(date, 'minutes');
    return diffInMinutes
  }

}
