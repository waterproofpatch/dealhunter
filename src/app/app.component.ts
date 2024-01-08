import { Component } from '@angular/core';
import { environment } from '../environments/environment';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'deals-app';
  constructor() {
    console.log("Production? " + environment.production); // Logs false for development environment
  }
}
