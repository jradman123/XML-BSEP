import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { ChangePasswordComponent } from 'src/app/components/change-password/change-password.component';
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
  showMe! : boolean;

   constructor(public matDialog: MatDialog,private userService : UserServiceService) {
     this.showMe = localStorage.getItem('role') != 'ADMIN' ? true : false;
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

  openModalForChangePassword(): void {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '500px';
    this.matDialog.open(ChangePasswordComponent, dialogConfig);
  }
}
