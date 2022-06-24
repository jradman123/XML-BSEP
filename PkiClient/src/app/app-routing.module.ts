import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LandingPageComponent } from './components/landing-page/landing-page.component';
import { AdminHomeComponent } from './components/admin-home/admin-home.component';
import { ClientHomeComponent } from './components/client-home/client-home.component';
import { AllCertificatesComponent } from './components/all-certificates/all-certificates.component';
import { CreateCertificateComponent } from './components/create-certificate/create-certificate.component';
import { CreateSubjectComponent } from './components/create-subject/create-subject.component';
import { CertificateComponent } from './components/certificate/certificate.component';
import { CertificateChainComponent } from './components/certificate-chain/certificate-chain.component';
import { ResetPasswordComponent } from './components/reset-password/reset-password.component';
import { CreateCertificateUserComponent } from './components/create-certificate-user/create-certificate-user.component';
import { RegistrationComponent } from './components/registration/registration.component';
import { ChangePasswordComponent } from './components/change-password/change-password.component';
import { AuthGuard } from './AuthGuard/AuthGuard';
import { TwoFactorAuthComponent } from './components/two-factor-auth/two-factor-auth.component';
import { PasswordlessLoginComponent } from './components/passwordless-login/passwordless-login.component';

const routes: Routes = [
  { path: '', component: LandingPageComponent },
  { path: 'registration', component: RegistrationComponent },
  { path: 'passwordless-login/:token', component: PasswordlessLoginComponent },
  {
    path: 'ahome',
    component: AdminHomeComponent,
    children: [
      {
        path: '',
        component: AllCertificatesComponent,
         canActivate: [AuthGuard]
      },
      { path: 'createCertificate', component: CreateCertificateComponent, canActivate: [AuthGuard] },
      { path: 'createSubject', component: CreateSubjectComponent, canActivate: [AuthGuard] },
      { path: 'certificate', component: CertificateComponent , canActivate: [AuthGuard]},
      { path: 'chain', component: CertificateChainComponent ,  canActivate: [AuthGuard] },
      {
        path: 'createCertificateUser',
        component: CreateCertificateUserComponent,
        canActivate: [AuthGuard]
      },
      { path: 'changePassword', component: ChangePasswordComponent, canActivate: [AuthGuard] },
      { path: 'two-factor-auth', component: TwoFactorAuthComponent, canActivate: [AuthGuard] }
    ],
  },

  {
    path: 'chome',
    component: ClientHomeComponent,
    children: [
      { path: '', component: AllCertificatesComponent, canActivate: [AuthGuard] },
      { path: 'changePassword', component: ChangePasswordComponent , canActivate: [AuthGuard]},
    ],
  },
  { path: 'createCertificate', component: CreateCertificateComponent, canActivate: [AuthGuard] },
  { path: 'createSubject', component: CreateSubjectComponent , canActivate: [AuthGuard]},
  { path: 'certificate', component: CertificateComponent, canActivate: [AuthGuard] },
  { path: 'chain', component: CertificateChainComponent, canActivate: [AuthGuard] },
  { path: 'createCertificateUser', component: CreateCertificateUserComponent , canActivate: [AuthGuard]},
  { path: 'resetPassword', component: ResetPasswordComponent },
  { path: 'two-factor-auth', component: TwoFactorAuthComponent, canActivate: [AuthGuard] }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
export const routingComponents = [
  LandingPageComponent,
  AdminHomeComponent,
  ClientHomeComponent,
];
