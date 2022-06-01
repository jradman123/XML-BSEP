import { Component, Input, OnInit } from '@angular/core';
import { IJobOffer } from 'src/app/interfaces/job-offer';

@Component({
  selector: 'app-job-offer-list-view',
  templateUrl: './job-offer-list-view.component.html',
  styleUrls: ['./job-offer-list-view.component.css']
})
export class JobOfferListViewComponent implements OnInit {
  @Input()
  jobOffer!: IJobOffer
  constructor() {

  }

  ngOnInit(): void {
  }
}
