import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { MaterialModule } from './material/material.module';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { UserHomeComponent } from './pages/user-home/user-home.component';
import { AuthGuard } from './AuthGuard/AuthGuard';
import { RecoverPassRequestComponent } from './components/recover-pass-request/recover-pass-request.component';
import { RecoverPassComponent } from './components/recover-pass/recover-pass.component';
import { ActivateAccountComponent } from './components/activate-account/activate-account.component';
import { EditUserComponent } from './pages/edit-user/edit-user.component';

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
    component: UserHomeComponent  //, canActivate: [AuthGuard]
  },
  {
    path: "editUser",
    component: EditUserComponent //, canActivate: [AuthGuard]
  }
];

@NgModule({

  imports: [RouterModule.forRoot(routes),
    MaterialModule,
  ],
  exports: [RouterModule]
})
export class AppRoutingModule { }
