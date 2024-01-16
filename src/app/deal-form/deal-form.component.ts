import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { DealsService } from '../deals.service';
import { LocationService } from '../location.service';

@Component({
  selector: 'app-deal-form',
  templateUrl: './deal-form.component.html',
  styleUrl: './deal-form.component.css',
})
export class DealFormComponent {
  dealForm = new FormGroup({
    retailPrice: new FormControl('', Validators.required),
    actualPrice: new FormControl('', Validators.required),
    storeName: new FormControl('', Validators.required),
    itemName: new FormControl('', Validators.required),
    address: new FormControl('', Validators.required),
  });

  constructor(
    private dealsService: DealsService,
    private router: Router,
    public locationService: LocationService) {

  }

  onSubmit() {
    if (this.dealForm.valid) {
      console.log(this.dealForm.value);
      if (this.dealForm.controls.storeName.value && this.dealForm.controls.itemName.value) {
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
        })
      } else {
        console.log('storeName and itemName are required');
      }
      this.router.navigate(['/deals']);
    }
  }
}
