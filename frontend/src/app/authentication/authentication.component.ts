// authentication.component.ts
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { AuthenticationService } from '../authentication.service';

@Component({
  selector: 'app-authentication',
  templateUrl: './authentication.component.html',
  styleUrls: ['./authentication.component.css']
})
export class AuthenticationComponent implements OnInit {
  mode: string = "";
  signInForm = new FormGroup({
    'email': new FormControl(null, [Validators.required, Validators.email]),
    'password': new FormControl(null, Validators.required)
  });

  signUpForm = new FormGroup({
    'email': new FormControl(null, [Validators.required, Validators.email]),
    'password': new FormControl(null, Validators.required),
    'confirmPassword': new FormControl(null, Validators.required)
  });

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private authenticationService: AuthenticationService) {
    this.route.queryParams.subscribe(params => {
      this.mode = params['mode'];
      if (this.mode == "signout") {
        this.authenticationService.signOut()
        this.router.navigate(["/deals"])
      }
    });
  }

  ngOnInit() {
  }

  onSubmit() {
    if (this.mode === 'signin' && this.signInForm.valid) {
      // Handle sign in
      this.authenticationService.signIn(this.signInForm.controls.email.value, this.signInForm.controls.password.value).subscribe((token: any) => {
        this.router.navigate(["/deals"])
      })
    } else if (this.mode === 'signup' && this.signUpForm.valid) {
      // Handle sign up
      this.authenticationService.signUp(this.signUpForm.controls.email.value, this.signUpForm.controls.password.value).subscribe((token: any) => {
        this.router.navigate(["/deals"])
      })
    }
  }
}
