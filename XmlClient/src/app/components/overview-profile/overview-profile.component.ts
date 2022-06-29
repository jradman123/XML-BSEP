import { Component, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-overview-profile',
  templateUrl: './overview-profile.component.html',
  styleUrls: ['./overview-profile.component.css']
})
export class OverviewProfileComponent implements OnInit {

  sub!: Subscription;
  userDetails! : UserDetails;
  initialDetails: any;
  id!: number;
  constructor(private userService : UserService) { }

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
