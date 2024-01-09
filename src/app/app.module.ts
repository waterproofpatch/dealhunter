import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatMenuModule } from '@angular/material/menu'
import { AuthInterceptor } from './auth-interceptor.service';
import { ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatCardModule } from '@angular/material/card'
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { TruncatePipe } from './pipes/truncate-pipe';
import { MatIconModule } from '@angular/material/icon'

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { DealFormComponent } from './deal-form/deal-form.component';
import { DealsComponent } from './deals/deals.component';
import { DealComponent } from './deal/deal.component';
import { AuthenticationComponent } from './authentication/authentication.component';

@NgModule({
  declarations: [
    AppComponent,
    TruncatePipe,
    DealFormComponent,
    DealsComponent,
    DealComponent,
    AuthenticationComponent
  ],
  imports: [
    HttpClientModule,
    MatMenuModule,
    MatIconModule,
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
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
