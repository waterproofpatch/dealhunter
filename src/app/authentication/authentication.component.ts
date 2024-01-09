// authentication.component.ts
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { FormGroup, FormControl, Validators } from '@angular/forms';

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

  constructor(private route: ActivatedRoute) {
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
    } else if (this.mode === 'signup' && this.signUpForm.valid) {
      // Handle sign up
      console.log('Sign up:', this.signUpForm.value);
    }
  }
}
