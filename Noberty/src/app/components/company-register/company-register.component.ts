import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { NewCompanyRequestDto } from 'src/app/interfaces/new-company-request-dto';
import { CompanyService } from 'src/app/services/company-service/company.service';

@Component({
  selector: 'app-company-register',
  templateUrl: './company-register.component.html',
  styleUrls: ['./company-register.component.css']
})
export class CompanyRegisterComponent implements OnInit {
  company: NewCompanyRequestDto;
  createForm!: FormGroup;
  constructor(
    private _snackBar: MatSnackBar,
    private _formBuilder: FormBuilder,
    private companyService: CompanyService,
    private dialogRef: MatDialogRef<CompanyRegisterComponent>
  ) {
    this.company = {} as NewCompanyRequestDto

    this.createForm = this._formBuilder.group({
      Name: new FormControl('', [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ0-9 ]*$'),
      ]),
      Site: new FormControl('', [
        Validators.required,
        Validators.pattern('(https?://)?([\\da-z.-]+)\\.([a-z.]{2,6})[/\\w .-]*/?')
      ]),
      Hedquaters: new FormControl('', [
        Validators.required,
      ]),
      Founded: new FormControl('', [Validators.required]),
      Industry: new FormControl('', [Validators.required]),
      Employees: new FormControl('', [
        Validators.required,
        Validators.pattern('^((?!(0))[0-9]{1,4})$'),
      ]),
      ///^-?(0|[1-9]\d*)?$/)]),
      Origin: new FormControl('', [Validators.required]),
      Offices: new FormControl('', [Validators.required]),
      CompanyPolicy: new FormControl('', [Validators.required]),

    });
  }

  ngOnInit(): void {
  }

  submitRequest(): void {

    this.createCompany();

    this.companyService.RegisterCompany(this.company).subscribe({
      next: () => {
        this._snackBar.open(
          'You have successfully created a request that has been sent to administration for approval.',
          '', {
            duration: 3000
          });
        this.dialogRef.close();
        

      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', {
          duration: 3000
        });
      },
      complete: () => console.info('complete')
    });

  }
  createCompany(): void {
    this.company.noOfEmpl = this.createForm.value.Employees
    this.company.companyName = this.createForm.value.Name
    this.company.companyWebsite = this.createForm.value.Site
    this.company.founded = this.createForm.value.Founded
    this.company.headquarters = this.createForm.value.Hedquaters
    this.company.industry = this.createForm.value.Industry
    this.company.offices = this.createForm.value.Offices
    this.company.countryOfOrigin = this.createForm.value.Origin
    this.company.companyPolicy = this.createForm.value.CompanyPolicy
    this.company.ownerUsername = localStorage.getItem('username')

  }

}







