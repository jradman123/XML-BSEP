import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { JobOffer } from 'src/app/interfaces/job-offer';
import { JobOfferService } from 'src/app/services/job-offer-service/job-offer.service';

@Component({
  selector: 'app-job-offer',
  templateUrl: './job-offer.component.html',
  styleUrls: ['./job-offer.component.css']
})
export class JobOfferComponent implements OnInit {

  jobOffers! : JobOffer[];
  searchText = '';

  constructor(
    private jobOfferService : JobOfferService
  ) { }

  ngOnInit(): void {
    this.jobOfferService.getAllJobOffers().subscribe(
      res => {
        this.jobOffers = res.JobOffers
      }
    )
  }

  
}
