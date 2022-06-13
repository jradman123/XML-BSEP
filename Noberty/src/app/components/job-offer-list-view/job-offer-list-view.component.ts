import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { CompanyResponseDto } from 'src/app/interfaces/company-response-dto';
import { IJobOffer } from 'src/app/interfaces/job-offer';
import { CompanyService } from 'src/app/services/company-service/company.service';
import { PublishJobOfferComponent } from '../publish-job-offer/publish-job-offer.component';

@Component({
  selector: 'app-job-offer-list-view',
  templateUrl: './job-offer-list-view.component.html',
  styleUrls: ['./job-offer-list-view.component.css']
})
export class JobOfferListViewComponent implements OnInit {
  @Input()
  jobOffer!: IJobOffer
  isVisible = false

  
  constructor(public matDialog: MatDialog, private companyService : CompanyService
    ) {
      this.companyService.getAllUsersCompanies(localStorage.getItem('username')!).subscribe(
        res => {
          console.log(res)
          res.forEach((c: CompanyResponseDto) => {
            if(c.companyId == this.jobOffer.companyId)
              this.isVisible = true;
          })
        }
      )
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
