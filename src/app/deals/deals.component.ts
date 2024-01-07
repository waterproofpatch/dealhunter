import { Component, OnInit } from '@angular/core';

import { Location } from '../models/location';

@Component({
  selector: 'app-deals',
  templateUrl: './deals.component.html',
  styleUrl: './deals.component.css'
})
export class DealsComponent implements OnInit {
  location: Location = { latitude: 0, longitude: 0 }

  ngOnInit(): void {
    this.getLocation()
  }
  getLocation(): void {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition((position) => {
        const longitude = position.coords.longitude;
        const latitude = position.coords.latitude;
        this.location.longitude = longitude
        this.location.latitude = latitude
      });
    } else {
      console.log("No support for geolocation")
    }
  }
}
