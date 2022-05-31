import { Component, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
export interface ICompany {
  description:string;

}
@Component({
  selector: 'app-company-register',
  templateUrl: './company-register.component.html',
  styleUrls: ['./company-register.component.css']
})
export class CompanyRegisterComponent implements OnInit {
  company: ICompany;
  constructor(private _snackBar: MatSnackBar) { 
    this.company = {} as ICompany
  }

  ngOnInit(): void {
  }
  doTextareaValueChange(ev: any) {
    this.company.description = ev.target.value;
  }
  
  submitRequest(): void {
    if (this.company && this.company.description != '') {
      this._snackBar.open(
        'Your feedback has been successfully submitted.',
        'Dismiss'
      );
    }
  }

}








