import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { MaterialModule } from './material/material.module';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { UserHomeComponent } from './pages/user-home/user-home.component';
import { AuthGuard } from './AuthGuard/AuthGuard';

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
    path: "userHome",
    component: UserHomeComponent , canActivate: [AuthGuard]
  }
];

@NgModule({

  imports: [RouterModule.forRoot(routes),
    MaterialModule,
  ],
  exports: [RouterModule]
})
export class AppRoutingModule { }
