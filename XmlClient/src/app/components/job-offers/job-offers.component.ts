import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { JobOffer } from 'src/app/interfaces/job-offer';
import { JobOfferService } from 'src/app/services/job-offer-service/job-offer.service';
import { NewJobOfferComponent } from '../new-job-offer/new-job-offer.component';

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
  constructor(private jobOfferService : JobOfferService, private _router : Router, private _matDialog: MatDialog,
    ) { }

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

    this.jobOfferService.getMyJobOffers(this.username!).subscribe(
      res => {
        this.myJobOffers = res.JobOffers
      }
    );
  }

  handleMe(searchText : string){
    this.searchText = searchText;
  }


  openDialog() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    const dialogRef = this._matDialog.open(NewJobOfferComponent, dialogConfig);

    dialogRef.afterClosed().subscribe({
      next: (res) => {
        console.log(res.data);
        this.myJobOffers.push(res.data)
      }
    })

  }
  
}
