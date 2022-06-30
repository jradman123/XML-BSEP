import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { JobOffer } from 'src/app/interfaces/job-offer';
import { AuthService } from 'src/app/services/auth-service/auth.service';
import { JobOfferService } from 'src/app/services/job-offer-service/job-offer.service';

@Component({
  selector: 'app-job-offer',
  templateUrl: './job-offer.component.html',
  styleUrls: ['./job-offer.component.css']
})
export class JobOfferComponent implements OnInit {

  jobOffers! : JobOffer[];
  searchText = '';
  username : any;

  constructor(
    private jobOfferService : JobOfferService, private router : Router, private authService : AuthService
  ) {
      this.username = localStorage.getItem('username');
   }

  ngOnInit(): void {
    this.jobOfferService.getAllJobOffers().subscribe(
      res => {
        this.jobOffers = res.JobOffers
      }
    )
  }

  logout() {
    this.authService.logout();
    this.router.navigate(['login']);
  }

  showProfile() {
    this.router.navigate(['myProfile']);
  }

  openCreateJobOffer() {
    this.router.navigate(['newJobOffer']);
  }

  
}
