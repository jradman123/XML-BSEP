import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminGuard } from './AuthGuard/AdminGuard';
import { AuthGuard } from './AuthGuard/AuthGuard';
import { OwnerGuard } from './AuthGuard/OwnerGuard';
import { TfaComponent } from './components/tfa/tfa.component';
import { CompaniesListComponent } from './pages/companies-list/companies-list.component';
import { CompanyProfileComponent } from './pages/company-profile/company-profile.component';
import { CompanyRequestsPageComponent } from './pages/company-requests-page/company-requests-page.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { MyCompaniesListComponent } from './pages/my-companies-list/my-companies-list.component';
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
    path: "user/landing", canActivate:[AuthGuard],
    component: UserLandingPageComponent
  },
  { path: 'resetPassword', 
    component: ResetPasswordComponent 
  },
  {
    path: "companies", canActivate:[AuthGuard],
    component: CompaniesListComponent
  },
  {
    path: "mycompanies", canActivate:[OwnerGuard],
    component: MyCompaniesListComponent
  },
  {
    path: "company/:id", canActivate:[AuthGuard],
    component: CompanyProfileComponent
  },
  {
    path: "companyRequests", canActivate:[AdminGuard],
    component: CompanyRequestsPageComponent
  },
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
