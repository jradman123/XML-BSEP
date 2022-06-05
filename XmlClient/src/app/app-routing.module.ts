import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './AuthGuard/AuthGuard';
import { ActivateAccountComponent } from './components/activate-account/activate-account.component';
import { JobOfferComponent } from './components/job-offer/job-offer.component';
import { NewJobOfferComponent } from './components/new-job-offer/new-job-offer.component';
import { RecoverPassRequestComponent } from './components/recover-pass-request/recover-pass-request.component';
import { RecoverPassComponent } from './components/recover-pass/recover-pass.component';
import { MaterialModule } from './material/material.module';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { UserHomeComponent } from './pages/user-home/user-home.component';

const routes: Routes = [
  {
    path: '',
    component: HomePageComponent,
  },
  {
    path: "login",
    component: LoginPageComponent
  },
  {
    path: "register",
    component: RegisterPageComponent
  },
  {
    path: "recoverRequest",
    component: RecoverPassRequestComponent 
  },
  {
    path: "recover",
    component: RecoverPassComponent 
  },
  {
    path: "activate",
    component: ActivateAccountComponent 
  },
  {
    path: "userHome",
    component: UserHomeComponent , canActivate: [AuthGuard]
  },
   {
    path: "jobOffers", 
    component: JobOfferComponent
  },
  {
    path: "newJobOffer",
    component: NewJobOfferComponent
  }
];

@NgModule({

  imports: [RouterModule.forRoot(routes),
    MaterialModule,
  ],
  exports: [RouterModule]
})
export class AppRoutingModule { }
