import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { UserLandingPageComponent } from './pages/user-landing-page/user-landing-page.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { UnauthHeaderComponent } from './components/unauth-header/unauth-header.component';
import { FooterComponent } from './components/footer/footer.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { AuthHeaderComponent } from './components/auth-header/auth-header.component';
import { MatCardModule } from '@angular/material/card';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { JwtInterceptor } from './JwtInterceptor/JwtInterceptor';
import { DatePipe } from '@angular/common';
import { ResetPasswordComponent } from './pages/reset-password/reset-password.component';
import { CompaniesListComponent } from './pages/companies-list/companies-list.component';
import { MatDialogModule } from '@angular/material/dialog';
import { CompanyRegisterComponent } from './components/company-register/company-register.component';
import {MatSnackBarModule} from '@angular/material/snack-bar';
import { MaterialModule } from './material/material.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CompanyListViewComponent } from './components/company-list-view/company-list-view.component';
import { CompanyProfileComponent } from './pages/company-profile/company-profile.component';
import { JobOfferComponent } from './components/job-offer/job-offer.component';
import { JobOfferListViewComponent } from './components/job-offer-list-view/job-offer-list-view.component';
import { CompanyRequestsComponent } from './components/company-requests/company-requests.component';
import { CompanyRequestsPageComponent } from './pages/company-requests-page/company-requests-page.component';
import { CommentComponent } from './components/comment/comment.component';
import { InterviewComponent } from './components/interview/interview.component';
import { SalaryCommentComponent } from './components/salary-comment/salary-comment.component';



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
    JobOfferListViewComponent

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
