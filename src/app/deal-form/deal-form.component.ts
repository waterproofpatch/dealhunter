import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

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

  onSubmit() {
    if (this.dealForm.valid) {
      console.log(this.dealForm.value);
    }
  }
}
