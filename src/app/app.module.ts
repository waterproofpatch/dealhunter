import { NgModule } from '@angular/core';
import { MatSliderModule } from '@angular/material/slider'
import { MatSelectModule } from '@angular/material/select'
import { MatTooltipModule } from '@angular/material/tooltip'
import { BrowserModule } from '@angular/platform-browser';
import { MatDialogModule } from '@angular/material/dialog';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatMenuModule } from '@angular/material/menu'
import { AuthInterceptor } from './auth-interceptor.service';
import { ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
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
import { ProfileComponent } from './profile/profile.component';
import { ErrorDialogComponent } from './error-dialog/error-dialog.component';

@NgModule({
  declarations: [
    AppComponent,
    TruncatePipe,
    DealFormComponent,
    DealsComponent,
    DealComponent,
    AuthenticationComponent,
    ProfileComponent,
    ErrorDialogComponent
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
    FormsModule,
    MatSliderModule,
    MatTooltipModule,
    MatDialogModule,
    MatSelectModule,
    BrowserAnimationsModule
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
