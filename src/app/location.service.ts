import { Injectable } from '@angular/core';
import { Location } from './models/location';
import { BaseHttpService } from './base-http.service';
import { MatDialog } from '@angular/material/dialog';
import { BehaviorSubject, Observable } from 'rxjs';

import { Address } from './models/address';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class LocationService extends BaseHttpService {
  public location: Location = { Latitude: 0, Longitude: 0 }
  public address$: BehaviorSubject<string> = new BehaviorSubject<string>("")

  constructor(private _dialog: MatDialog, private http: HttpClient) {
    super(_dialog)
    this.getLocation()
  }

  private getLocation(): boolean {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition((position) => {
        const longitude = position.coords.longitude;
        const latitude = position.coords.latitude;
        this.location.Longitude = longitude;
        this.location.Latitude = latitude;
        this.getAddress(this.location)
        return true
      });
    } else {
      console.log('No support for geolocation');
    }
    return false
  }

  private toRadians(degrees: number): number {
    return degrees * Math.PI / 180;
  }

  public calculateDistance(location2: Location): number {
    const R = 3958.8; // Radius of the Earth in miles
    const lat1 = this.toRadians(this.location.Latitude);
    const lon1 = this.toRadians(this.location.Longitude);
    const lat2 = this.toRadians(location2.Latitude);
    const lon2 = this.toRadians(location2.Longitude);

    const dlon = lon2 - lon1;
    const dlat = lat2 - lat1;

    const a = Math.sin(dlat / 2) * Math.sin(dlat / 2) +
      Math.cos(lat1) * Math.cos(lat2) *
      Math.sin(dlon / 2) * Math.sin(dlon / 2);
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));

    return R * c;
  }

  /**
   * given lat/lon, get the address right now
   */
  private getAddress(location: Location): void {
    this.getAddressHttp(location.Latitude, location.Longitude).subscribe((address: Address) => {
      console.log(`Address is ${address.Address}`)
      this.address$.next(address.Address)
    });
  }

  private getAddressHttp(lat: number, lon: number): Observable<any> {
    return this.http.get(`${this.apiUrl}/address?lat=${lat}&lon=${lon}`)
  }
}
