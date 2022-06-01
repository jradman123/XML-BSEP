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
import { MaterialModule } from './material-module';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { JwtInterceptor } from './JwtInterceptor/JwtInterceptor';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { DatePipe } from '@angular/common';
import { ResetPasswordComponent } from './pages/reset-password/reset-password.component';


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
    ResetPasswordComponent
  ],
  imports: [
    BrowserModule,
    MatCardModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    NgbModule,
    MaterialModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  providers: [ HttpClientModule,
    { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true },
  DatePipe],
  bootstrap: [AppComponent]
})
export class AppModule { }
