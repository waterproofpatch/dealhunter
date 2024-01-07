import { Component, OnInit } from '@angular/core';
import { Input } from '@angular/core';
import { Location } from '../models/location';

import { Deal } from '../models/deal';
@Component({
  selector: 'app-deal',
  templateUrl: './deal.component.html',
  styleUrl: './deal.component.css'
})
export class DealComponent implements OnInit {

  @Input() deal: Deal | undefined
  @Input() location: Location | undefined
  ngOnInit(): void {
  }

  private toRadians(degrees: number): number {
    return degrees * Math.PI / 180;
  }

  public calculateDistance(location1: Location, location2: Location): number {
    const R = 3958.8; // Radius of the Earth in miles
    const lat1 = this.toRadians(location1.latitude);
    const lon1 = this.toRadians(location1.longitude);
    const lat2 = this.toRadians(location2.latitude);
    const lon2 = this.toRadians(location2.longitude);

    const dlon = lon2 - lon1;
    const dlat = lat2 - lat1;

    const a = Math.sin(dlat / 2) * Math.sin(dlat / 2) +
      Math.cos(lat1) * Math.cos(lat2) *
      Math.sin(dlon / 2) * Math.sin(dlon / 2);
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));

    return R * c;
  }
}
