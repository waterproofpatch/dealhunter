import { Component, OnInit } from '@angular/core';

import { DealsService } from '../deals.service';
import { LocationService } from '../location.service';

@Component({
  selector: 'app-deals',
  templateUrl: './deals.component.html',
  styleUrl: './deals.component.css'
})
export class DealsComponent implements OnInit {
  public distanceToDeal: number = 5
  public sortOption: string = ""
  constructor(
    public dealsService: DealsService,
    public locationService: LocationService,
  ) {
  }

  ngOnInit(): void {
    this.sortOption = localStorage.getItem('sortOption') || ""
    this.distanceToDeal = +(localStorage.getItem('distance') || 5)
  }

  public saveSortOption(event: any): void {
    console.log(`Saving sortOption: ${event.value}`)
    localStorage.setItem('sortOption', event.value)
  }

  public saveDistanceOption(event: any): void {
    console.log(`Saving distance: ${event}`)
    localStorage.setItem('distance', event)
  }
}
