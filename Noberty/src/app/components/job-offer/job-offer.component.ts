import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IJobOffer } from 'src/app/interfaces/job-offer';
import { CompanyService } from 'src/app/services/company-service/company.service';

@Component({
  selector: 'app-job-offer',
  templateUrl: './job-offer.component.html',
  styleUrls: ['./job-offer.component.css']
})

export class JobOfferComponent implements OnInit {
  jobOffer: IJobOffer;
  createForm!: FormGroup;
  requirements!: string[];

  constructor(
    private _snackBar: MatSnackBar,
    private _formBuilder: FormBuilder,
    private companyService: CompanyService) {

    this.jobOffer = {} as IJobOffer
    this.requirements = [];
    this.createForm = this._formBuilder.group({
      Name: new FormControl(''),
      Requirements: new FormControl(''),
    })
  }

  ngOnInit(): void {
  }
  submitRequest(): void {
    this.companyService.CreateJobOffer(this.jobOffer).subscribe({
      next: () => {
        this._snackBar.open(
          'Your request has been successfully submitted.',
          'Dismiss'
        );
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss');
      },
      complete: () => console.info('complete')
    });
  }
  addItem() {
    console.log(this.requirements.length);
    this.requirements.push(this.createForm.value.Requirements)
  }

}
