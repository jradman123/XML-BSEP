import { Component, OnInit } from '@angular/core';
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
  ) { 
    // const jo1 : JobOffer = {
    //    Position : "Senior .NET Developer",
    //    JobDescription : "Job Description should be longer than this.",
    //    Duration : "May 22 2022",
    //    Publisher : "Marble IT",
    //    DatePosted : "Apr 14 2022",
    //    Requirements : ["Dobar", "Pametan", "Ne pije"]
    //   };
    //   const jo2 : JobOffer = {
    //     Position : "Junior Java Developer",
    //     Duration : "May 22 2022",
    //     JobDescription : "Probacu u sledecoj nmg sad majke mi odvratnoje",
    //     Publisher : "ICodeFactory",
    //     DatePosted : "Apr 10 2022",
    //     Requirements : ["Jedan", "Dva", "Samo da je lud"]
    //    };   
    //    this.jobOffers = [jo1, jo2]
  }

  ngOnInit(): void {
    this.jobOfferService.getAllJobOffers().subscribe(
      res => this.jobOffers = res
    )
  }

  
}
