// authentication.component.ts
import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-authentication',
  templateUrl: './authentication.component.html',
  styleUrls: ['./authentication.component.css']
})
export class AuthenticationComponent {
  public mode: string = "";

  constructor(private route: ActivatedRoute) {
    this.route.queryParams.subscribe(params => {
      this.mode = params['mode'];
    });
  }
}
