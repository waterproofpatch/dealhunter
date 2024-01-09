// authentication.component.ts
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
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
    private authenticationService: AuthenticationService) {
    this.route.queryParams.subscribe(params => {
      this.mode = params['mode'];
    });
  }

  ngOnInit() {
  }

  onSubmit() {
    if (this.mode === 'signin' && this.signInForm.valid) {
      // Handle sign in
      console.log('Sign in:', this.signInForm.value);
      this.authenticationService.signIn(this.signInForm.controls.email.value, this.signInForm.controls.password.value).subscribe((token: any) => {
        console.log(`Got jwt token ${token}`)
      })
    } else if (this.mode === 'signup' && this.signUpForm.valid) {
      // Handle sign up
      console.log('Sign up:', this.signUpForm.value);
      this.authenticationService.signUp(this.signUpForm.controls.email.value, this.signUpForm.controls.password.value).subscribe((token: any) => {
        console.log(`Got jwt token ${token}`)
      })
    }
  }
}
