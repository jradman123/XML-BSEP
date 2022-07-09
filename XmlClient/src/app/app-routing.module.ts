import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ActivateAccountComponent } from './components/activate-account/activate-account.component';
import { NewJobOfferComponent } from './components/new-job-offer/new-job-offer.component';
import { RecoverPassRequestComponent } from './components/recover-pass-request/recover-pass-request.component';
import { RecoverPassComponent } from './components/recover-pass/recover-pass.component';
import { MaterialModule } from './material/material.module';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { TwofaPageComponent } from './pages/twofa-page/twofa-page.component';
import { AuthGuardRegular } from './AuthGuard/AuthGuardRegular';
import { MyProfileComponent } from './components/my-profile/my-profile.component';
import { ProfileListComponent } from './components/profile-list/profile-list.component';
import { PusherComponent } from './components/pusher/pusher.component';
import { PostsViewComponent } from './components/posts-view/posts-view.component';
import { PostCreateFileComponent } from './components/post-create-file/post-create-file.component';
import { PublicProfileComponent } from './components/public-profile/public-profile.component';
import { NetworkComponent } from './components/network/network.component';
import { NotFoundComponent } from './pages/not-found/not-found.component';
import { NotificationComponent } from './components/notification/notification.component';
import { PostViewComponent } from './components/post-view/post-view.component';
import { MessagesPageComponent } from './pages/messages-page/messages-page/messages-page.component';
import { JobOffersComponent } from './components/job-offers/job-offers.component';
import { MessageCreateComponent } from './components/message-create/message-create/message-create.component';
import { ChatboxComponent } from './components/chatbox/chatbox.component';
import { LandingPageComponent } from './pages/landing-page/landing-page.component';
import { FeedComponent } from './pages/feed/feed.component';

const routes: Routes = [
  {
    path: '',
    component: LandingPageComponent,
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
    component: PostCreateFileComponent,canActivate: [AuthGuardRegular]
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
    path: 'newJobOffer',
    component: NewJobOfferComponent,canActivate: [AuthGuardRegular]
  },
  {
    path: 'myProfile',
    component: MyProfileComponent, canActivate: [AuthGuardRegular]
  },
  {
    path: 'pusher',
    component: PusherComponent
  },
  {
    path: 'myMessages',
    component: MessagesPageComponent,canActivate: [AuthGuardRegular]
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
    path: 'noti',
    component: NotificationComponent,canActivate: [AuthGuardRegular]
  },
  {
    path: 'post/:id',
    component: PostViewComponent
  },

   {
    path: 'job-offers',
    component: JobOffersComponent
   }, 
   {
    path: 'send-message/:username',
    component: MessageCreateComponent,canActivate: [AuthGuardRegular]
   }, 
   {
    path: 'chatbox',
    component: ChatboxComponent,canActivate: [AuthGuardRegular]
   },
   {
    path : 'feed',
    component : FeedComponent,canActivate: [AuthGuardRegular]
   }
];

@NgModule({
  imports: [RouterModule.forRoot(routes), MaterialModule],
  exports: [RouterModule],
})
export class AppRoutingModule { }
