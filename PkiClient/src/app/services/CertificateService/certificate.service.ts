import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CertificateView } from 'src/app/interfaces/certificate-view';
import { IssuerData } from 'src/app/interfaces/issuer-data';

@Injectable({
  providedIn: 'root',
})
export class CertificateService {

  
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
  getUsersChainCertificates(email : string) : Observable<CertificateView[][][]> {
    return this.http.get<CertificateView[][][]>("http://localhost:8443/api/certificate/chains/"+"a@gmail.com");
  }
}
