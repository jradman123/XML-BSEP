import { Component, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserPersonalDetails } from 'src/app/interfaces/user-personal-details';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-my-profile',
  templateUrl: './my-profile.component.html',
  styleUrls: ['./my-profile.component.css']
})
export class MyProfileComponent implements OnInit {

  sub!: Subscription;
  userDetails! : UserDetails;
  initialuserPersonalDetails! : UserPersonalDetails;
  initialDetails: any;
  id!: number;
  constructor(private userService : UserService) {
    this.initialuserPersonalDetails = {} as UserPersonalDetails;
   }

  ngOnInit(): void {
    this.getUserDetails();
  }
  getUserDetails() {
    this.sub = this.userService.getUserDetails(localStorage.getItem('username')).subscribe({
      next: (data: UserDetails) => {
        const variable  = data.dateOfBirth.substring(0,10);
        data.dateOfBirth = variable;
        this.userDetails = data
        this.initialuserPersonalDetails.firstName = this.userDetails.firstName;
        this.initialuserPersonalDetails.lastName=this.userDetails.lastName;
        this.initialuserPersonalDetails.biography= this.userDetails.lastName;
        this.initialuserPersonalDetails.phoneNumber = this.userDetails.phoneNumber;
        this.initialuserPersonalDetails.biography = this.userDetails.biography;
        this.initialuserPersonalDetails.dateOfBirth = this.userDetails.dateOfBirth.substring(0,10);
        this.initialuserPersonalDetails.gender = this.userDetails.gender;
        this.initialDetails = JSON.parse(JSON.stringify(data)); 
      },
    });
  }

  refreshPersonalData(userPersonalDetails : UserPersonalDetails)  {
    this.initialuserPersonalDetails.username = userPersonalDetails.username;
    this.initialuserPersonalDetails.firstName = userPersonalDetails.firstName;
    this.initialuserPersonalDetails.lastName = userPersonalDetails.lastName;
    this.initialuserPersonalDetails.biography = userPersonalDetails.biography;
    this.initialuserPersonalDetails.dateOfBirth = userPersonalDetails.dateOfBirth.substring(0,10);
    this.initialuserPersonalDetails.phoneNumber = userPersonalDetails.phoneNumber;
    this.initialuserPersonalDetails.gender = userPersonalDetails.gender;
  }

  refreshEmail(email : string) {
    this.userDetails.email = email
  }

}
