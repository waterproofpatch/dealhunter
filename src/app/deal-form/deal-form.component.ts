import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ValidationErrors, AbstractControl, FormControl, FormGroup, Validators } from '@angular/forms';
import { DealsService } from '../deals.service';
import { LocationService } from '../location.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-deal-form',
  templateUrl: './deal-form.component.html',
  styleUrl: './deal-form.component.css',
})
export class DealFormComponent implements OnInit {
  dealForm = new FormGroup({
    retailPrice: new FormControl('', Validators.required),
    actualPrice: new FormControl('', Validators.required),
    storeName: new FormControl('', Validators.required),
    itemName: new FormControl('', Validators.required),
    address: new FormControl('', Validators.required),
  }, { validators: this.priceValidator });
  private subscription: Subscription = new Subscription;

  constructor(
    private dealsService: DealsService,
    private router: Router,
    public locationService: LocationService) {

  }

  ngOnInit(): void {
    console.log("Requesting location update...")
    this.subscription = this.locationService.address$.subscribe(address => {
      this.dealForm.controls.address.setValue(address);
    });
    this.locationService.refreshLocation()
  }
  ngOnDestroy() {
    if (this.subscription) {
      this.subscription.unsubscribe();
    }
  }

  priceValidator(control: AbstractControl): ValidationErrors | null {
    const group = control as FormGroup;
    const retailPrice = group.controls['retailPrice'].value;
    const actualPrice = group.controls['actualPrice'].value;

    return retailPrice > actualPrice ? null : { notValid: true };
  }

  onSubmit() {
    if (this.dealForm.valid) {
      console.log(this.dealForm.value);
      if (this.dealForm.controls.storeName.value && this.dealForm.controls.itemName.value && this.dealForm.controls.address.value) {
        this.dealsService.addDeal({
          RetailPrice: Number(this.dealForm.controls.retailPrice.value),
          ActualPrice: Number(this.dealForm.controls.actualPrice.value),
          ItemName: this.dealForm.controls.itemName.value,
          StoreName: this.dealForm.controls.storeName.value,
          Location: { Latitude: 0, Longitude: 0 }, // authoritative
          CreatedAt: "", // authoritative
          LastUpvoteTime: "", // authoritative
          Upvotes: 0, // authoritative
          ID: 0, // authoritative
          User: { ID: 0 },
        }, this.dealForm.controls.address.value)
      } else {
        console.log('storeName and itemName are required');
      }
      this.router.navigate(['/deals']);
    }
  }
}
