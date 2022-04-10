import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CertificateView } from 'src/app/interfaces/certificate-view';

@Injectable({
  providedIn: 'root'
})
export class CertificateService {
  revokeCertificate(serialNumber: string) {
    return this.http.post<any>('http://localhost:8443/api/certificate/revoke',serialNumber);
  }

  constructor(private http: HttpClient) { }

  getAllCertificates() : Observable<CertificateView[]> {
    return this.http.get<CertificateView[]>('http://localhost:8443/api/certificate/');
  }

  downloadCertificate(serialNumber : string) {
    return this.http.post<any>('http://localhost:8443/api/certificate/downloadCertificate',serialNumber);
  }
}
