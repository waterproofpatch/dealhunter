import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatToolbarModule } from '@angular/material/toolbar';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DealFormComponent } from './deal-form/deal-form.component';
import { DealsComponent } from './deals/deals.component';
import { DealComponent } from './deal/deal.component';

@NgModule({
  declarations: [
    AppComponent,
    DealFormComponent,
    DealsComponent,
    DealComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MatToolbarModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
