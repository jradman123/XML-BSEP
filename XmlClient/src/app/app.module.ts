import { DatePipe } from '@angular/common';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatBadgeModule } from '@angular/material/badge';
import { MatBottomSheetModule } from '@angular/material/bottom-sheet';
import { MatButtonModule } from '@angular/material/button';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatCardModule } from '@angular/material/card';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatChipsModule } from '@angular/material/chips';
import { MatNativeDateModule, MatRippleModule } from '@angular/material/core';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatDialogModule } from '@angular/material/dialog';
import { MatDividerModule } from '@angular/material/divider';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatListModule } from '@angular/material/list';
import { MatMenuModule } from '@angular/material/menu';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatRadioModule } from '@angular/material/radio';
import { MatSelectModule } from '@angular/material/select';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatSliderModule } from '@angular/material/slider';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatSortModule } from '@angular/material/sort';
import { MatStepperModule } from '@angular/material/stepper';
import { MatTableModule } from '@angular/material/table';
import { MatTabsModule } from '@angular/material/tabs';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatTreeModule } from '@angular/material/tree';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule, NoopAnimationsModule } from '@angular/platform-browser/animations';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ActivateAccountComponent } from './components/activate-account/activate-account.component';
import { FooterComponent } from './components/footer/footer.component';
import { JobOfferComponent } from './components/job-offer/job-offer.component';
import { NewJobOfferComponent } from './components/new-job-offer/new-job-offer.component';
import { RecoverPassRequestComponent } from './components/recover-pass-request/recover-pass-request.component';
import { RecoverPassComponent } from './components/recover-pass/recover-pass.component';
import { UnauthenticatedHeaderComponent } from './components/unauthenticated-header/unauthenticated-header.component';
import { MaterialModule } from './material/material.module';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { FiterPipePipe } from './pipes/fiter-pipe.pipe';
import { NgxCaptchaModule } from 'ngx-captcha';
import { ConfirmDialogComponent } from './components/confirm-dialog/confirm-dialog.component';
import { QRCodeModule } from 'angular2-qrcode';
import { TwofaPageComponent } from './pages/twofa-page/twofa-page.component';
import { JwtInterceptor } from './JwtInterceptor/JwtInterceptor';
import { JwtHelperService, JWT_OPTIONS } from '@auth0/angular-jwt';
import { PostComponent } from './components/post/post.component';
import { CommentCreateComponent } from './components/comment-create/comment-create.component';
import { PostCreateFileComponent } from './components/post-create-file/post-create-file.component';
import { PostsViewComponent } from './components/posts-view/posts-view.component';
import { PostsCommentsViewComponent } from './components/posts-comments-view/posts-comments-view.component';
import { MyProfileComponent } from './components/my-profile/my-profile.component';
import { RegisteredHeaderComponent } from './components/registered-header/registered-header.component';
import { ProfileEditComponent } from './components/profile-edit/profile-edit.component';
import { SettingsComponent } from './components/settings/settings.component';
import { ChangePasswordComponent } from './components/change-password/change-password.component';
import { OverviewProfileComponent } from './components/overview-profile/overview-profile.component';
import { ProfileAboutComponent } from './components/profile-about/profile-about.component';
import { EducationDialogComponent } from './components/dialogs/education-dialog/education-dialog.component';
import { ProfilePreviewComponent } from './components/profile-preview/profile-preview.component';
import { ProfileListComponent } from './components/profile-list/profile-list.component';
import { ProfileSearchPipe } from './pipes/profile-search.pipe';
import { PublicProfileComponent } from './components/public-profile/public-profile.component';



@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    LoginPageComponent,
    RegisterPageComponent,
    UnauthenticatedHeaderComponent,
    FooterComponent,
    RecoverPassRequestComponent,
    RecoverPassComponent,
    ActivateAccountComponent,
    JobOfferComponent,
    FiterPipePipe,
    NewJobOfferComponent,
    ConfirmDialogComponent,
    TwofaPageComponent,
    PostComponent,
    CommentCreateComponent,
    PostCreateFileComponent,
    PostsViewComponent,
    PostsCommentsViewComponent,
    MyProfileComponent,
    RegisteredHeaderComponent,
    ProfileEditComponent,
    SettingsComponent,
    ChangePasswordComponent,
    OverviewProfileComponent,
    ProfileAboutComponent,
    EducationDialogComponent,
    ProfilePreviewComponent,
    ProfileListComponent,
    ProfileSearchPipe,
    PublicProfileComponent,
  ],
  imports: [
    BrowserModule,
    MaterialModule,
    AppRoutingModule,
    NoopAnimationsModule,
    NgbModule,
    HttpClientModule,
    FormsModule,
    BrowserAnimationsModule,
    MatCardModule,
    MatAutocompleteModule,
    MatBadgeModule,
    MatBottomSheetModule,
    MatButtonModule,
    MatButtonToggleModule,
    MatCardModule,
    MatCheckboxModule,
    MatChipsModule,
    MatStepperModule,
    MatDatepickerModule,
    MatDialogModule,
    MatDividerModule,
    MatExpansionModule,
    MatGridListModule,
    MatIconModule,
    MatInputModule,
    MatListModule,
    MatMenuModule,
    MatNativeDateModule,
    MatPaginatorModule,
    MatProgressBarModule,
    MatProgressSpinnerModule,
    MatRadioModule,
    MatRippleModule,
    MatSelectModule,
    MatSidenavModule,
    MatSliderModule,
    MatSlideToggleModule,
    MatSnackBarModule,
    MatSortModule,
    MatTableModule,
    MatTabsModule,
    MatToolbarModule,
    MatTooltipModule,
    MatTreeModule,
    MatFormFieldModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatRippleModule,
    ReactiveFormsModule,
    NgxCaptchaModule,
    QRCodeModule,
  ],
  exports:[
    MaterialModule,
    MatIconModule
  ],
  providers: [HttpClientModule,
    { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true },DatePipe,{ provide: JWT_OPTIONS, useValue: JWT_OPTIONS }, JwtHelperService],
  bootstrap: [AppComponent]
})
export class AppModule { }
