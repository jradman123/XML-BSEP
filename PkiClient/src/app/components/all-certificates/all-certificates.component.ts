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
  email! : any;
  admin! : boolean;

  constructor(private certificateService : CertificateService,private _snackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.email=localStorage.getItem('email');
    if(localStorage.getItem('role') == 'admin'){
      this.admin = true;
    }else{
      this.admin =false;
    }
    if(this.admin){
      this.getAllCertificates();
    }else{
      this.getAllUsersCertificates();
    }
 
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
getAllUsersCertificates() {
  console.log("usao1") ; 
  this.certificateService.getAllUsersCertificates(this.email).subscribe(
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


