import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ActivateAccountComponent } from './components/activate-account/activate-account.component';
import { JobOfferComponent } from './components/job-offer/job-offer.component';
import { NewJobOfferComponent } from './components/new-job-offer/new-job-offer.component';
import { RecoverPassRequestComponent } from './components/recover-pass-request/recover-pass-request.component';
import { RecoverPassComponent } from './components/recover-pass/recover-pass.component';
import { MaterialModule } from './material/material.module';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { TwofaPageComponent } from './pages/twofa-page/twofa-page.component';
import { AuthGuardRegular } from './AuthGuard/AuthGuardRegular';
import { MyProfileComponent } from './components/my-profile/my-profile.component';
import { ProfilePreviewComponent } from './components/profile-preview/profile-preview.component';
import { ProfileListComponent } from './components/profile-list/profile-list.component';

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
    path: 'people',
    component: ProfileListComponent,
  },
  {
    path: 'login',
    component: LoginPageComponent,
  }, 
  {

    path: 'twofa',
    component: TwofaPageComponent,
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
    path: 'jobOffers',
    component: JobOfferComponent,// canActivate: [AuthGuard]
  },
  {
    path: 'newJobOffer',
    component: NewJobOfferComponent,canActivate: [AuthGuardRegular]
  },
  {
    path: 'myProfile',
    component: MyProfileComponent, canActivate: [AuthGuardRegular]
  },


];

@NgModule({
  imports: [RouterModule.forRoot(routes), MaterialModule],
  exports: [RouterModule],
})
export class AppRoutingModule { }
