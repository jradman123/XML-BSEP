import { DatePipe } from '@angular/common';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatDialogModule } from '@angular/material/dialog';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AuthHeaderComponent } from './components/auth-header/auth-header.component';
import { ChangePasswordComponent } from './components/change-password/change-password.component';
import { CommentComponent } from './components/comment/comment.component';
import { CompanyListViewComponent } from './components/company-list-view/company-list-view.component';
import { CompanyRegisterComponent } from './components/company-register/company-register.component';
import { CompanyRequestsComponent } from './components/company-requests/company-requests.component';
import { FooterComponent } from './components/footer/footer.component';
import { InterviewComponent } from './components/interview/interview.component';
import { JobOfferListViewComponent } from './components/job-offer-list-view/job-offer-list-view.component';
import { JobOfferComponent } from './components/job-offer/job-offer.component';
import { LeaveCommentComponent } from './components/leave-comment/leave-comment.component';
import { LeaveInterviewCommentComponent } from './components/leave-interview-comment/leave-interview-comment.component';
import { LeaveSallaryCommentComponent } from './components/leave-sallary-comment/leave-sallary-comment.component';
import { SalaryCommentComponent } from './components/salary-comment/salary-comment.component';
import { UnauthHeaderComponent } from './components/unauth-header/unauth-header.component';
import { JwtInterceptor } from './JwtInterceptor/JwtInterceptor';
import { MaterialModule } from './material/material.module';
import { CompaniesListComponent } from './pages/companies-list/companies-list.component';
import { CompanyProfileComponent } from './pages/company-profile/company-profile.component';
import { CompanyRequestsPageComponent } from './pages/company-requests-page/company-requests-page.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { MyCompaniesListComponent } from './pages/my-companies-list/my-companies-list.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { ResetPasswordComponent } from './pages/reset-password/reset-password.component';
import { UserLandingPageComponent } from './pages/user-landing-page/user-landing-page.component';
import { PublishJobOfferComponent } from './components/publish-job-offer/publish-job-offer.component';
import { TfaComponent } from './components/tfa/tfa.component';
import { ConfirmDialogComponent } from './components/confirm-dialog/confirm-dialog.component';
import { PasswordlessLoginComponent } from './components/passwordless-login/passwordless-login.component';




@NgModule({
  declarations: [
    AppComponent,
    RegisterPageComponent,
    LoginPageComponent,
    UserLandingPageComponent,
    HomePageComponent,
    UnauthHeaderComponent,
    FooterComponent,
    AuthHeaderComponent,
    ResetPasswordComponent,
    CompaniesListComponent,
    CompanyRegisterComponent,
    CompanyListViewComponent,
    CompanyProfileComponent,
    JobOfferComponent,
    JobOfferListViewComponent,
    CompanyRequestsComponent,
    CompanyRequestsPageComponent,
    CommentComponent,
    InterviewComponent,
    SalaryCommentComponent,
    JobOfferListViewComponent,
    LeaveCommentComponent,
    LeaveInterviewCommentComponent,
    LeaveSallaryCommentComponent,
    MyCompaniesListComponent,
    ChangePasswordComponent,
    PublishJobOfferComponent,
    TfaComponent,
    ConfirmDialogComponent,
    PasswordlessLoginComponent,

  ],
  imports: [
    BrowserModule,
    MatCardModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    NgbModule,
    MaterialModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    MatDialogModule,
    MatSnackBarModule,
    MaterialModule,
    FormsModule,
    ReactiveFormsModule
  ],

  providers: [ HttpClientModule,
    { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true },
  DatePipe],
  bootstrap: [AppComponent]
})
export class AppModule { }
