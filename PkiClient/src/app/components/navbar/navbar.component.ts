import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { NavigationStart, Router } from '@angular/router';
import { UserService } from 'src/app/services/UserService/user.service';
import { TwoFactorAuthComponent } from '../two-factor-auth/two-factor-auth.component';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  role! : string;
  show! : boolean ;
  constructor(private userService : UserService,
    public matDialog: MatDialog) { }

  ngOnInit(): void {
   this.role  = localStorage.getItem('role')!;
   if(this.role === 'Admin'){
      this.show = true;
      console.log('admin');
   } else {
     this.show = false;
   }
  }

  logout() : void{
    this.userService.logout();
  }

  tfa() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '800px';
    this.matDialog.open(TwoFactorAuthComponent, dialogConfig);
  }
   

}
