import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CertificateView } from 'src/app/interfaces/certificate-view';
import { IssuerData } from 'src/app/interfaces/issuer-data';
import { NewCertificate } from 'src/app/interfaces/new-certificate';
import { SubjectData } from 'src/app/interfaces/subject-data';

@Injectable({
  providedIn: 'root',
})
export class CertificateService {
  createCertificateByUser(newCertificate: NewCertificate) {
    return this.http.post<any>(
      'http://localhost:8443/api/certificate/generateByClient',
      newCertificate
    );
  }
  getSignersForDateRangeByUser(
    startDate: String,
    endDate: String
  ): Observable<IssuerData[]> {
    return this.http.get<IssuerData[]>(
      'http://localhost:8443/api/certificate/getCAsForSigningClientsCertificatesInDateRange?email='+ localStorage.getItem('email') +'&startDate=' +
        startDate +
        '&endDate=' +
        endDate
    );
  }
  
  constructor(private http: HttpClient) {}

  getAllCertificates(): Observable<CertificateView[]> {
    return this.http.get<CertificateView[]>(
      'http://localhost:8443/api/certificate/'
    );
  }

  revokeCertificate(serialNumber: string) {
    return this.http.post<any>('http://localhost:8443/api/certificate/revoke',serialNumber);
  }
  
  getSignersForDateRange(
    startDate: String,
    endDate: String
  ): Observable<IssuerData[]> {
    return this.http.get<IssuerData[]>(
      'http://localhost:8443/api/certificate/getCAsForSigning?startDate=' +
        startDate +
        '&endDate=' +
        endDate
    );
  }

  downloadCertificate(serialNumber : string) {
    return this.http.post<any>('http://localhost:8443/api/certificate/downloadCertificate',serialNumber);
  }

  getAllUsersCertificates(email : string) : Observable<CertificateView[]> {
    return this.http.get<CertificateView[]>('http://localhost:8443/api/certificate/getAllUsersCertificates/' + email);
  
  }
  getSubjects(): Observable<SubjectData[]> {
    return this.http.get<SubjectData[]>('http://localhost:8443/api/users');
  }

  createCertificate(newCertificate: NewCertificate): Observable<any> {
    return this.http.post<any>(
      'http://localhost:8443/api/certificate/generate',
      newCertificate
    );
    }
}
