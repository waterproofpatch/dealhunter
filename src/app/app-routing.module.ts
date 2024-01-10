import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { DealFormComponent } from './deal-form/deal-form.component';
import { DealsComponent } from './deals/deals.component';
import { AuthenticationComponent } from './authentication/authentication.component';
import { ProfileComponent } from './profile/profile.component';

const routes: Routes = [
  { path: 'deal-form', component: DealFormComponent },
  { path: 'deals', component: DealsComponent },
  { path: 'profile', component: ProfileComponent },
  { path: 'authentication', component: AuthenticationComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
