import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { JobOffer } from 'src/app/interfaces/job-offer';
import { JobOfferService } from 'src/app/services/job-offer-service/job-offer.service';

@Component({
  selector: 'app-job-offers',
  templateUrl: './job-offers.component.html',
  styleUrls: ['./job-offers.component.css']
})
export class JobOffersComponent implements OnInit {

  searchText = ''
  jobOffers! : JobOffer[];
  myJobOffers : JobOffer[] = [];
  suggestedJobOffers! : JobOffer[];
  username = localStorage.getItem('username');
  constructor(private jobOfferService : JobOfferService, private _router : Router) { }

  ngOnInit(): void {
    this.jobOfferService.getAllJobOffers().subscribe(
      res => {
        this.jobOffers = res.JobOffers
      }
    )

    this.jobOfferService.getSuggestedJobOffers(this.username!).subscribe(
      res => {
        this.suggestedJobOffers = res.offers
      }
    );
  }

  openCreateJobOffer(){
    this._router.navigate(['newJobOffer']);

  }

  handleMe(searchText : string){
    this.searchText = searchText;
  }
}
