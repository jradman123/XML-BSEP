import { Component, Input, OnInit } from '@angular/core';
import { JobOffer } from 'src/app/interfaces/job-offer';

@Component({
  selector: 'app-job-offer-view',
  templateUrl: './job-offer-view.component.html',
  styleUrls: ['./job-offer-view.component.css']
})
export class JobOfferViewComponent implements OnInit {

  @Input()
  jo! : JobOffer
  constructor() { }

  ngOnInit(): void {
  }

}
