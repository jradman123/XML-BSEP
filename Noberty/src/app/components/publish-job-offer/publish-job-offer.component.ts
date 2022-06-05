import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { IJobOfferPublish } from 'src/app/interfaces/ijob-offer-publish';
import { JobOfferWithPublisher } from 'src/app/interfaces/job-offer-with-publisher';
import { JobOfferService } from 'src/app/services/joboffer-service/job-offer.service';
import { JobOfferComponent } from '../job-offer/job-offer.component';

@Component({
  selector: 'app-publish-job-offer',
  templateUrl: './publish-job-offer.component.html',
  styleUrls: ['./publish-job-offer.component.css']
})
export class PublishJobOfferComponent implements OnInit {

  publishing! : IJobOfferPublish;
  ApiKey! : String;
  createForm!: FormGroup;

  constructor(@Inject(MAT_DIALOG_DATA) public data: any, private _formBuilder: FormBuilder, private jobOfferService : JobOfferService,
  public dialogRef: MatDialogRef<JobOfferComponent>,
  ) {
    this.createForm = this._formBuilder.group({
     ApiKey: new FormControl('',Validators.required)
    });
    this.publishing = {} as IJobOfferPublish;
    this.publishing.jobOffer = {} as JobOfferWithPublisher;
     this.createPublishing();
   }


  ngOnInit(): void {
  }

  publishOffer(){ 

    if (this.createForm.invalid) return;
    
    this.publishing.apiToken = this.createForm.value.ApiKey;

    this.jobOfferService.publishJobOfferOnDislinkt(this.publishing).subscribe(
      res => {
        console.log('poslato');
        this.dialogRef.close();
      },
    );
    
  }

  createPublishing() {
    this.publishing.jobOffer.Publisher = this.data.companyName;
    //this.publishing.jobOffer.DatePosted = 237635;
    //this.publishing.jobOffer.Duration ="2374";
    this.publishing.jobOffer.JobDescription = this.data.jobDescription;
    this.publishing.jobOffer.Position = this.data.position;
    this.publishing.jobOffer.Requirements = this.data.requirements;
  }
}

