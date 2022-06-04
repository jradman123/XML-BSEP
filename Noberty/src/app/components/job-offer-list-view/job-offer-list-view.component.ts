import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { IJobOffer } from 'src/app/interfaces/job-offer';
import { PublishJobOfferComponent } from '../publish-job-offer/publish-job-offer.component';

@Component({
  selector: 'app-job-offer-list-view',
  templateUrl: './job-offer-list-view.component.html',
  styleUrls: ['./job-offer-list-view.component.css']
})
export class JobOfferListViewComponent implements OnInit {
  @Input()
  jobOffer!: IJobOffer

  
  constructor(public matDialog: MatDialog
    ) {
      

  }

  ngOnInit(): void {
  
  }

  publish() {
    
    console.log(this.jobOffer);
    this.matDialog.open(PublishJobOfferComponent, {
      width: '800px',
      height: 'fit-content',
      disableClose : false,
      data: this.jobOffer
    });
  }

}
