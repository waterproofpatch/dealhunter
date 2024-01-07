import { Component, OnInit } from '@angular/core';

import { Location } from '../models/location';
import { DealsService } from '../deals.service';
import { LocationService } from '../location.service';

@Component({
  selector: 'app-deals',
  templateUrl: './deals.component.html',
  styleUrl: './deals.component.css'
})
export class DealsComponent implements OnInit {
  constructor(public dealsService: DealsService, public locationService: LocationService) {

  }
  ngOnInit(): void {
  }
}
