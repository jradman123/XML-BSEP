import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { CompanyRegisterComponent } from 'src/app/components/company-register/company-register.component';
import { JobOfferComponent } from 'src/app/components/job-offer/job-offer.component';
import { IComment } from 'src/app/interfaces/comment';
import { ICompanyInfo } from 'src/app/interfaces/company-info';
import { IInterview } from 'src/app/interfaces/interview';
import { IJobOffer } from 'src/app/interfaces/job-offer';
import { ISalaryComment } from 'src/app/interfaces/salary-comment';
import { CompanyService } from 'src/app/services/company-service/company.service';

@Component({
  selector: 'app-company-profile',
  templateUrl: './company-profile.component.html',
  styleUrls: ['./company-profile.component.css']
})
export class CompanyProfileComponent implements OnInit {
  description!: string;
  item!: ICompanyInfo;
  editable!: boolean;
  newDescription!: string
  jobOffer!: IJobOffer[]
  comments!: IComment[]
  interviews!: IInterview[]
  salaryComments!: ISalaryComment[]

  constructor(
    private _snackBar: MatSnackBar,
    private companyService: CompanyService,
    public matDialog: MatDialog
  ) {
    this.jobOffer = [{
      name: "Senior develper for outsourcing firm",
      requirements: ["5+ years experience", "c++ development", "java development"]
    },
    {
      name: "Junior develper ",
      requirements: ["c++ development", "c# development"]
    }
    ]
    this.description = " Company Info: orem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore"
      + "magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo"
      + " consequat...";
    this.item = {
      name: "Levi9 Technology Services",
      site: "https://www.levi9.com/",
      headquaters: "Novi Sad",
      founded: "2018",
      industry: "Software Outsourcing",
      employees: 800,
      origin: "SRB",
      offices: "Beograd, Novi Sad, Zrenjanin"
    }
    this.editable = false;
    this.newDescription = this.description
    this.comments = [{
      comment:"asasas",
      companyId  : 1,
      userUsername : "usernameeebebebooo"
    },{
      comment:"eeeeeeeeeeeee",
      companyId  : 1,
      userUsername : "zjuuu"
    }]

    this.interviews = [{
      comment : "aaashodhsuhd",
      companyID : 1,
      difficulty : "HARD",
      rating : 5,
      userUsername : "SHDOSHD"
    },{
      comment : "FGDGDG",
      companyID : 1,
      difficulty : "HARD",
      rating : 4,
      userUsername : "SHDOSSADDDDHD"
    }]

    this.salaryComments = [{
      companyId : 1, 
      position : "position",
      salary : "120e",
      userUsername : " dsdhsddhsh"
    },{
      companyId : 1, 
      position : "another one ",
      salary : "12000e",
      userUsername : " heheheheh"
    }]
  }

  ngOnInit(): void {
    
  }
  enableEdit() {
    this.editable = true
  }
  updateInfo() {

    this.editable = false;

    if (this.description == this.newDescription) {
      return;
    }
    console.log("SENDING REQUEST");
    this.description = this.newDescription

    this.companyService.UpdateInfo(this.description).subscribe({
      next: () => {
        this._snackBar.open(
          'Your request has been successfully submitted.',
          'Dismiss'
        );

      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss');
      },
      complete: () => console.info('complete')
    });

  }
  openModal() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '500px';
    this.matDialog.open(JobOfferComponent, dialogConfig); //TODO: OVDJE JOB OFFER
  }

}
