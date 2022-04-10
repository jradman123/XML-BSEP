import { Component, OnInit } from '@angular/core';
import { CertificateView } from 'src/app/interfaces/certificate-view';
import { CertificateService } from 'src/app/services/CertificateService/certificate.service';

@Component({
  selector: 'app-all-certificates',
  templateUrl: './all-certificates.component.html',
  styleUrls: ['./all-certificates.component.css']
})
export class AllCertificatesComponent implements OnInit {


  certificates!:CertificateView[]; 

  constructor(private certificateService : CertificateService) { }

  ngOnInit(): void {
    console.log("usao ng Init");
    this.getAllCertificates();
  }

  revokeCertificate( serialNumber:string) : void {
    console.log(serialNumber);
    this.certificateService.revokeCertificate(serialNumber).subscribe({
      next: (result) => {
        this.certificates = result;
        
      },
      error: data => {
        if (data.error && typeof data.error === "string")
        console.log("desila se greska")
      }
    });
  }

  getAllCertificates() { console.log("usao2") ; this.certificateService.getAllCertificates().subscribe(
    {
      next: (result) => {
        this.certificates = result;
      },
      error: data => {
        if (data.error && typeof data.error === "string")
        console.log("desila se greska")
      }
    }
  );
}

}
