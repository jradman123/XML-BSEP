import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { CompanyRegisterComponent } from 'src/app/components/company-register/company-register.component';
import { UserInformationResponseDto } from 'src/app/interfaces/user-information-response-dto';
import { UserServiceService } from 'src/app/services/UserService/user-service.service';


@Component({
  selector: 'app-user-landing-page',
  templateUrl: './user-landing-page.component.html',
  styleUrls: ['./user-landing-page.component.css']
})
export class UserLandingPageComponent implements OnInit {

  userInfo : UserInformationResponseDto;
  constructor(public matDialog: MatDialog,private userService : UserServiceService) {
    this.userInfo = {} as UserInformationResponseDto;
   }

  ngOnInit(): void {
    this.userService.getUserInformation().subscribe((res) => {this.userInfo = res});
  }
  openModal(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '500px';
    this.matDialog.open(CompanyRegisterComponent, dialogConfig);
  }
}
