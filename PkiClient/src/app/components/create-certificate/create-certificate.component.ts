import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { IssuerData } from 'src/app/interfaces/issuer-data';

@Component({
  selector: 'app-create-certificate',
  templateUrl: './create-certificate.component.html',
  styleUrls: ['./create-certificate.component.css'],
})
export class CreateCertificateComponent implements OnInit {
  isLinear = false;
  firstFormGroup!: FormGroup;
  secondFormGroup!: FormGroup;
  cType!: String;
  startDate!: string;
  endDate!: string;
  todayDate: Date = new Date();
  potentialIssuers: IssuerData[];


  constructor(private _formBuilder: FormBuilder) {
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
  }

  changeType(value: String) {
    this.cType = value;
    console.log('type is ' + value);
  }

  typeNext(): void {
    // if( cType == "root" ) prebaci na subject stranu, nema poziva getCert
    //else getCertificatesForSigning based on cType & start-end date
    // dobijeni sert idu u this.potentialIssuers
  }

  dateRangeChange(
    dateRangeStart: HTMLInputElement,
    dateRangeEnd: HTMLInputElement
  ) {
    this.startDate = dateRangeStart.value;
    this.endDate = dateRangeEnd.value;
  }
}
