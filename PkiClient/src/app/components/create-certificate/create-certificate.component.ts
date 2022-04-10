import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { IssuerData } from 'src/app/interfaces/issuer-data';
import { CertificateService } from 'src/app/services/CertificateService/certificate.service';

@Component({
  selector: 'app-create-certificate',
  templateUrl: './create-certificate.component.html',
  styleUrls: ['./create-certificate.component.css'],
})
export class CreateCertificateComponent implements OnInit {
  isLinear = false;
  firstFormGroup!: FormGroup;
  secondFormGroup!: FormGroup;
  thirdFormGroup!: FormGroup;
  fourthFormGroup!: FormGroup;
  cType!: String;
  startDate!: string;
  endDate!: string;
  todayDate: Date = new Date();
  potentialIssuers: IssuerData[];
  enableIssuerStep = true;

  constructor(
    private _formBuilder: FormBuilder,
    private certificateService: CertificateService
  ) {
    this.potentialIssuers = [] as IssuerData[];
  }

  ngOnInit(): void {
    this.firstFormGroup = this._formBuilder.group({
      firstCtrl: ['', Validators.required],
    });
    this.secondFormGroup = this._formBuilder.group({
      secondCtrl: ['', Validators.required],
      start: ['', Validators.required],
      end: ['', Validators.required],
    });
    this.thirdFormGroup = this._formBuilder.group({
      thirdCtrl: ['', Validators.required],
    });
    this.fourthFormGroup = this._formBuilder.group({
      fourthCtrl: [''],
    });
  }

  changeType(value: String) {
    this.cType = value;
    console.log('type is ' + value);
  }

  dateRangeChange(
    dateRangeStart: HTMLInputElement,
    dateRangeEnd: HTMLInputElement
  ) {
    this.startDate = dateRangeStart.value;
    this.endDate = dateRangeEnd.value;
  }

  dateNext(): void {
    // if( cType == "root" ) prebaci na subject stranu, nema poziva getCert
    // else getCertificatesForSigning based on cType & start-end date
    // dobijeni sert idu u this.potentialIssuers
    // service.getCAsForDateRange(this.startDate, this.endDate);
    if (this.cType == 'client' || this.cType == 'intermediate') {
      this.certificateService
        .getSignersForDateRange(this.startDate, this.endDate)
        .subscribe((res) => {
          this.potentialIssuers = res;
        });
    } else {
      this.enableIssuerStep = false;
    }
  }

  issuerSelected(matOpr: any) {
    console.log(this.potentialIssuers);
  }
}
