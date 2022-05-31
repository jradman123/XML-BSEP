import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { UserLandingPageComponent } from './pages/user-landing-page/user-landing-page.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { UnauthHeaderComponent } from './components/unauth-header/unauth-header.component';
import { FooterComponent } from './components/footer/footer.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { AuthHeaderComponent } from './components/auth-header/auth-header.component';
import { MatCardModule } from '@angular/material/card';
import { CompaniesListComponent } from './pages/companies-list/companies-list.component';
import { MatDialogModule } from '@angular/material/dialog';
import { CompanyRegisterComponent } from './components/company-register/company-register.component';
import {MatSnackBarModule} from '@angular/material/snack-bar';
import { MaterialModule } from './material/material.module';



@NgModule({
  declarations: [
    AppComponent,
    RegisterPageComponent,
    LoginPageComponent,
    UserLandingPageComponent,
    HomePageComponent,
    UnauthHeaderComponent,
    FooterComponent,
    AuthHeaderComponent,
    CompaniesListComponent,
    CompanyRegisterComponent,

  ],
  imports: [
    BrowserModule,
    MatCardModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    NgbModule,
    MatDialogModule,
    MatSnackBarModule,
    MaterialModule
  ],

  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
