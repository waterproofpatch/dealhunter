import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { DealFormComponent } from './deal-form/deal-form.component';
import { DealsComponent } from './deals/deals.component';

const routes: Routes = [
  { path: 'deal-form', component: DealFormComponent },
  { path: 'deals', component: DealsComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
