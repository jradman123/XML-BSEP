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
import { ProfileListComponent } from './components/profile-list/profile-list.component';
import { MessagesPageComponent } from './pages/messages-page/messages-page/messages-page.component';
import { PostsViewComponent } from './components/posts-view/posts-view.component';
import { PostCreateFileComponent } from './components/post-create-file/post-create-file.component';
import { PublicProfileComponent } from './components/public-profile/public-profile.component';
import { NetworkComponent } from './components/network/network.component';
import { NotFoundComponent } from './pages/not-found/not-found.component';
import { JobOffersComponent } from './components/job-offers/job-offers.component';
import { MessageCreateComponent } from './components/message-create/message-create/message-create.component';

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
    path : 'public-profile/:username',
    component: PublicProfileComponent
  },
  {
    path : 'posts',
    component: PostsViewComponent,
  },
  {
    path : 'post-create',
    component: PostCreateFileComponent,
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
  {
    path: 'myMessages',
    component: MessagesPageComponent,

  },
  {
    path: 'network',
    component: NetworkComponent, canActivate: [AuthGuardRegular]
  },
  {
    path: '404',
    component: NotFoundComponent,
  },
   {
    path: 'job-offers',
    component: JobOffersComponent
   }, 
   {
    path: 'send-message/:username',
    component: MessageCreateComponent
   }
];

@NgModule({
  imports: [RouterModule.forRoot(routes), MaterialModule],
  exports: [RouterModule],
})
export class AppRoutingModule { }
