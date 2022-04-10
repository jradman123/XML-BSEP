import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { CertificateView } from 'src/app/interfaces/certificate-view';
import { CertificateService } from 'src/app/services/CertificateService/certificate.service';

@Component({
  selector: 'app-all-certificates',
  templateUrl: './all-certificates.component.html',
  styleUrls: ['./all-certificates.component.css']
})
export class AllCertificatesComponent implements OnInit {


  certificates!:CertificateView[]; 

  constructor(private certificateService : CertificateService,private _snackBar: MatSnackBar) { }

  ngOnInit(): void {
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

  downloadCertificate( serialNumber:string) : void {
    console.log(serialNumber);
    this.certificateService.downloadCertificate(serialNumber).subscribe(
      (res) => {this._snackBar.open(
        'Certificate downloaded successfully.Check out your Downloads folder.',
        'Dismiss'
      );
    
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
function res(res: any): import("rxjs").PartialObserver<any> | ((value: any) => void) | null | undefined {
  throw new Error('Function not implemented.');
}

