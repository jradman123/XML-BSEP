import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LandingPageComponent } from './components/landing-page/landing-page.component';
import { AdminHomeComponent } from './components/admin-home/admin-home.component';
import { ClientHomeComponent } from './components/client-home/client-home.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NavbarComponent } from './components/navbar/navbar.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { ClientNavbarComponent } from './components/client-navbar/client-navbar.component';
 
@NgModule({
  declarations: [
    AppComponent,
    LandingPageComponent,
    AdminHomeComponent,
    ClientHomeComponent,
    NavbarComponent,
    ClientNavbarComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    NgbModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
