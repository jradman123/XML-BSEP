import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LandingPageComponent } from './components/landing-page/landing-page.component';
import { AdminHomeComponent } from './components/admin-home/admin-home.component';
import { ClientHomeComponent } from './components/client-home/client-home.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NavbarComponent } from './components/navbar/navbar.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { ClientNavbarComponent } from './components/client-navbar/client-navbar.component';
import { AllCertificatesComponent } from './components/all-certificates/all-certificates.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './material-module';
import { MatCardModule } from '@angular/material/card';
import { CreateCertificateComponent } from './components/create-certificate/create-certificate.component';
import { CreateSubjectComponent } from './components/create-subject/create-subject.component';
import { CertificateComponent } from './components/certificate/certificate.component';
import { CertificateChainComponent } from './components/certificate-chain/certificate-chain.component';

import { CreateCertificateUserComponent } from './components/create-certificate-user/create-certificate-user.component';
import { ResetPasswordComponent } from './components/reset-password/reset-password.component';
import { JwtInterceptor } from './JwtInterceptor/jwt-interceptor';
import { RegistrationComponent } from './components/registration/registration.component';
import { ChangePasswordComponent } from './components/change-password/change-password.component';
import { TwoFactorAuthComponent } from './components/two-factor-auth/two-factor-auth.component';
import { ConfirmDialogComponent } from './components/confirm-dialog/confirm-dialog.component';
import { PasswordlessLoginComponent } from './components/passwordless-login/passwordless-login.component';

@NgModule({
  declarations: [
    AppComponent,
    LandingPageComponent,
    AdminHomeComponent,
    ClientHomeComponent,
    NavbarComponent,
    ClientNavbarComponent,
    AllCertificatesComponent,
    CreateCertificateComponent,
    CreateSubjectComponent,
    CertificateComponent,
    CertificateChainComponent,
    CreateCertificateUserComponent,
    ResetPasswordComponent,
    RegistrationComponent,
    ChangePasswordComponent,
    TwoFactorAuthComponent,
    ConfirmDialogComponent,
    PasswordlessLoginComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    NgbModule,
    BrowserAnimationsModule,
    MaterialModule,
    MatCardModule,
  ],
  providers: [
    HttpClientModule,
    { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
