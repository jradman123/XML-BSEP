import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Subscription } from 'rxjs';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-user-infos',
  templateUrl: './user-infos.component.html',
  styleUrls: ['./user-infos.component.css']
})
export class UserInfosComponent implements OnInit {

  sub!: Subscription;
  userDetails! : UserDetails;
  initialDetails: any;
  editMode = false;
  id!: number;
  constructor(private userService : UserService,public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getUserDetails();
  }
  getUserDetails() {
    this.sub = this.userService.getUserDetails(localStorage.getItem('username')).subscribe({
      next: (data: UserDetails) => {
        const variable  = data.dateOfBirth.substring(0,10);
        data.dateOfBirth = variable;
        this.userDetails = data
        this.initialDetails = JSON.parse(JSON.stringify(data)); 
      },
    });

}
}
