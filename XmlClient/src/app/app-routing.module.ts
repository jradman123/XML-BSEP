import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './AuthGuard/AuthGuard';
import { ActivateAccountComponent } from './components/activate-account/activate-account.component';
import { ApiKeyComponent } from './components/api-key/api-key.component';
import { JobOfferComponent } from './components/job-offer/job-offer.component';
import { NewJobOfferComponent } from './components/new-job-offer/new-job-offer.component';
import { RecoverPassRequestComponent } from './components/recover-pass-request/recover-pass-request.component';
import { RecoverPassComponent } from './components/recover-pass/recover-pass.component';
import { MaterialModule } from './material/material.module';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { UserHomeComponent } from './pages/user-home/user-home.component';
import { EditUserComponent } from './pages/edit-user/edit-user.component';
import { PassLessReqComponent } from './components/pass-less-req/pass-less-req.component';
import { PassLessLoginComponent } from './components/pass-less-login/pass-less-login.component';
import { AuthGuardRegular } from './AuthGuard/AuthGuardRegular';

const routes: Routes = [
  {
    path: '',
    component: HomePageComponent,
    /*children: [
      {
        path: 'jobOffers',
        component: JobOfferComponent,canActivate: [AuthGuard]
      },
      {
        path: 'genApiKey',
        component: ApiKeyComponent, canActivate: [AuthGuard]
      },
      {
        path: 'newJobOffer',
        component: NewJobOfferComponent,canActivate: [AuthGuard]
      }
    ],*/
  },
  {
    path: 'login',
    component: LoginPageComponent,
  },
  {
    path: 'register',
    component: RegisterPageComponent,
  },
  {
    path: 'recoverRequest',
    component: RecoverPassRequestComponent,
  },
  {
    path: 'recover',
    component: RecoverPassComponent,
  },
  {
    path: 'activate',
    component: ActivateAccountComponent,
  },
  {
    path: 'passwordlessReq',
    component: PassLessReqComponent,
  },
  {
    path: 'passwordlessLogin',
    component: PassLessLoginComponent,
  },
 {
    path: 'editUser',
    component: EditUserComponent,
    canActivate: [AuthGuard],
  },
  {
    path: 'jobOffers',
    component: JobOfferComponent,canActivate: [AuthGuard]
  },
  {
    path: 'genApiKey',
    component: ApiKeyComponent, canActivate: [AuthGuardRegular]
  },
  {
    path: 'newJobOffer',
    component: NewJobOfferComponent,canActivate: [AuthGuardRegular]
  },
  {
    path: 'userHome',
    component: UserHomeComponent,canActivate: [AuthGuard]
  }

  
];

@NgModule({
  imports: [RouterModule.forRoot(routes), MaterialModule],
  exports: [RouterModule],
})
export class AppRoutingModule {}
