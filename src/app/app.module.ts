import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatToolbarModule } from '@angular/material/toolbar';
import { ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatCardModule } from '@angular/material/card'
import { HttpClientModule } from '@angular/common/http';
import { TruncatePipe } from './pipes/truncate-pipe';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DealFormComponent } from './deal-form/deal-form.component';
import { DealsComponent } from './deals/deals.component';
import { DealComponent } from './deal/deal.component';

@NgModule({
  declarations: [
    AppComponent,
    TruncatePipe,
    DealFormComponent,
    DealsComponent,
    DealComponent
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    MatCardModule,
    AppRoutingModule,
    MatToolbarModule,
    MatInputModule,
    MatButtonModule,
    MatFormFieldModule,
    ReactiveFormsModule,
    BrowserAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
