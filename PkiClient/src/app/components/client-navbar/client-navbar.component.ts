import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { UserService } from 'src/app/services/UserService/user.service';
import { ChangePasswordComponent } from '../change-password/change-password.component';

@Component({
  selector: 'app-client-navbar',
  templateUrl: './client-navbar.component.html',
  styleUrls: ['./client-navbar.component.css']
})
export class ClientNavbarComponent implements OnInit {

  constructor(private userService : UserService) { }

  ngOnInit(): void {
  }


}
