import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CompaniesListComponent } from './pages/companies-list/companies-list.component';

import { HomePageComponent } from './pages/home-page/home-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { ResetPasswordComponent } from './pages/reset-password/reset-password.component';
import { UserLandingPageComponent } from './pages/user-landing-page/user-landing-page.component';

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
    path: "user/landing",
    component: UserLandingPageComponent
  },
  { path: 'resetPassword', 
    component: ResetPasswordComponent 
  },
  {
    path: "companies",
    component: CompaniesListComponent
  },
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
