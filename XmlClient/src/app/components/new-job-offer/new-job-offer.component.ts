import { DatePipe } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { JobOffer } from 'src/app/interfaces/job-offer';
import { JobOfferService } from 'src/app/services/job-offer-service/job-offer.service';

@Component({
  selector: 'app-new-job-offer',
  templateUrl: './new-job-offer.component.html',
  styleUrls: ['./new-job-offer.component.css'],
})
export class NewJobOfferComponent implements OnInit {
  newJobOffer!: JobOffer;
  createForm!: FormGroup;
  requirements!: string[];

  constructor(
    private _formBuilder: FormBuilder,
    private datepipe: DatePipe,
    private jobOfferService: JobOfferService,
    private dialogRef: MatDialogRef<NewJobOfferComponent>,

  ) {
    this.newJobOffer = {} as JobOffer;
    this.requirements = [];
    this.createForm = this._formBuilder.group({
      Requirements: new FormControl('', [Validators.required]),
      Position: new FormControl('', [Validators.required]),
      Description: new FormControl('', [Validators.required]),
      DueDate: new FormControl('', [Validators.required]),
    });
  }

  ngOnInit(): void {}

  addItem() {
    if(this.createForm.value.Requirements !== '') 
      this.requirements.push(this.createForm.value.Requirements);
  }

  submitRequest() {
    if (this.createForm.invalid) return;
    this.createJobOffer();
    this.jobOfferService.createJobOffer(this.newJobOffer).subscribe(res => {
      this.dialogRef.close({ data: this.newJobOffer});

    });
  }

  createJobOffer() {
    this.newJobOffer.Position = this.createForm.value.Position;
    this.newJobOffer.Requirements = this.requirements;
    this.newJobOffer.Publisher = localStorage.getItem('username')!;
    this.newJobOffer.JobDescription = this.createForm.value.Description;
    this.newJobOffer.Duration = this.createForm.value.DueDate;
    let ahora = new Date();
    let stringVal = this.datepipe.transform(ahora, 'yyyy-MM-dd HH:mm:ss')!;
    this.newJobOffer.DatePosted = new Date(stringVal);
  }
}
