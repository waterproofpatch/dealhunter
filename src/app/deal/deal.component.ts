import { Component, OnInit } from '@angular/core';
import { Input } from '@angular/core';
import { Location } from '../models/location';

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

}
