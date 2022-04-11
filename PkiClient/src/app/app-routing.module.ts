import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LandingPageComponent } from './components/landing-page/landing-page.component';
import { AdminHomeComponent } from './components/admin-home/admin-home.component';
import { ClientHomeComponent } from './components/client-home/client-home.component';
import { AllCertificatesComponent } from './components/all-certificates/all-certificates.component';
import { CreateCertificateComponent } from './components/create-certificate/create-certificate.component';
import { CreateSubjectComponent } from './components/create-subject/create-subject.component';
import { CertificateComponent } from './components/certificate/certificate.component';
import { CreateCertificateUserComponent} from './components/create-certificate-user/create-certificate-user.component';

const routes: Routes = [
  { path: '', component: LandingPageComponent },
  { path: 'ahome', component: AdminHomeComponent },
  { path: 'allCertificates', component: AllCertificatesComponent },
  { path: 'chome', component: ClientHomeComponent },
  { path: 'createCertificate', component: CreateCertificateComponent },
  { path: 'createSubject', component: CreateSubjectComponent },
  { path: 'certificate', component: CertificateComponent },
  { path: 'createCertificateUser', component: CreateCertificateUserComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule { }
export const routingComponents = [
  LandingPageComponent,
  AdminHomeComponent,
  ClientHomeComponent,
];
