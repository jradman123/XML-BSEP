import { HttpErrorResponse } from '@angular/common/http';
import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { IJobOffer } from 'src/app/interfaces/job-offer';
import { IJobOfferRequest } from 'src/app/interfaces/job-offer-request';
import { IJobOfferResponse } from 'src/app/interfaces/job-offer-response';
import { CompanyService } from 'src/app/services/company-service/company.service';


@Component({
  selector: 'app-job-offer',
  templateUrl: './job-offer.component.html',
  styleUrls: ['./job-offer.component.css']
})

export class JobOfferComponent implements OnInit {

  jobOfferRequest: IJobOfferRequest;
  createForm!: FormGroup;
  requirements!: string[];
  allOffers: IJobOfferResponse;
  cid!: string;
  @Output() newItemEvent = new EventEmitter<IJobOffer[]>();

  constructor(
    public dialogRef: MatDialogRef<JobOfferComponent>,
    private _snackBar: MatSnackBar,
    private router: Router,
    private _formBuilder: FormBuilder,
    private companyService: CompanyService) {

    this.jobOfferRequest = {} as IJobOfferRequest
    this.allOffers = {} as IJobOfferResponse
    this.requirements = [];
    this.createForm = this._formBuilder.group({
      Requirements: new FormControl('', [
        Validators.required,
        Validators.pattern('^[a-zA-Z0-9_.-]*$')
      ]),
      Position: new FormControl('', [
        Validators.required,
        Validators.pattern('^[a-zA-Z0-9_.-]*$')
      ]),
      Description: new FormControl('', [
        Validators.required
      ]),
    })

  }

  ngOnInit(): void {
    console.log(this.router.url);
    this.cid = this.router.url.substring(9);
  }
  submitRequest(): void {

    this.createJobOfferRequest();
    console.log(this.jobOfferRequest);

    this.companyService.CreateJobOffer(this.jobOfferRequest).subscribe({
      next: (res) => {
        this.clearForm();
        this.dialogRef.close({ event: "Created Job offer", data: res });
        this._snackBar.open(
          'You have created a job offer.',
          'Dismiss', {
            duration: 3000
          });
      },
      error: (err: HttpErrorResponse) => {
        this.clearForm();
        this._snackBar.open(err.error.message + "!", 'Dismiss', {
          duration: 3000
        });
      },
      complete: () => console.info('complete')
    });
  }
  addItem() {
    console.log(this.requirements.length);
    this.requirements.push(this.createForm.value.Requirements)
  }
  createJobOfferRequest() {
    console.log(this.createForm.value.Name);

    this.jobOfferRequest.requirements = this.requirements;
    this.jobOfferRequest.position = this.createForm.value.Position;
    this.jobOfferRequest.jobDescription = this.createForm.value.Description;
    this.jobOfferRequest.companyId = parseInt(this.cid);
  }
  clearForm() {
    this.createForm.reset()
    this.requirements = []
  }
}


