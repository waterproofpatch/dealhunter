import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { DealsService } from '../deals.service';

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
  });

  constructor(private dealsService: DealsService) {

  }

  onSubmit() {
    if (this.dealForm.valid) {
      console.log(this.dealForm.value);
      if (this.dealForm.controls.storeName.value && this.dealForm.controls.itemName.value) {
        this.dealsService.addDeal({
          retailPrice: Number(this.dealForm.controls.retailPrice.value),
          actualPrice: Number(this.dealForm.controls.actualPrice.value),
          itemName: this.dealForm.controls.itemName.value,
          storeName: this.dealForm.controls.storeName.value,
          location: { latitude: 0, longitude: 0 }, // authoritative
        })
      } else {
        console.log('storeName and itemName are required');
      }
    }
  }
}
