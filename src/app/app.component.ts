import { Component } from '@angular/core';
import { environment } from '../environments/environment';
import { AuthenticationService } from './authentication.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'deals-app';
  public sliderValue: number = 0
  constructor(public authenticationService: AuthenticationService) {
  }
}
