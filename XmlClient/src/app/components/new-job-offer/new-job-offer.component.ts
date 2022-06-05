import { DatePipe } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
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
    private jobOfferService: JobOfferService
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
    this.requirements.push(this.createForm.value.Requirements);
  }

  submitRequest() {
    //if (this.createForm.invalid) return;
    this.createJobOffer();
    console.log('jede govna');
    this.jobOfferService.createJobOffer(this.newJobOffer).subscribe();
  }

  createJobOffer() {
    this.newJobOffer.Position = this.createForm.value.Position;
    this.newJobOffer.Requirements = this.requirements;
    this.newJobOffer.Publisher = localStorage.getItem('username')!;
    this.newJobOffer.JobDescription = this.createForm.value.JobDescription;
    this.newJobOffer.Duration = this.createForm.value.DueDate;
    let ahora = new Date();
    let stringVal = this.datepipe.transform(ahora, 'yyyy-MM-dd')!;
    this.newJobOffer.DatePosted = new Date(stringVal);
  }
}
