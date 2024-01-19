import { Component } from '@angular/core';
import { AuthenticationService } from '../authentication.service';
import { LocationService } from '../location.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css'
})
export class ProfileComponent {

  constructor(public authenticationService: AuthenticationService, public locationService: LocationService) {

  }

}
