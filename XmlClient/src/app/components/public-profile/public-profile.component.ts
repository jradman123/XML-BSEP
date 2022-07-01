import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-public-profile',
  templateUrl: './public-profile.component.html',
  styleUrls: ['./public-profile.component.css']
})
export class PublicProfileComponent implements OnInit {
  sub! : Subscription;
  user! : UserDetails;
  initialDetails: any;
  id!: number;
  searchText : string = "";
  username! : string;
  isLoggedIn = localStorage.getItem('token') !== null


  constructor(private userService : UserService, private _router : Router) {
    this.username = this._router.url.substring(16) ?? '';
    this.user = {} as UserDetails;
    this.getUserDetails();
  }

  ngOnInit(): void {
  }

  handleMe(searchText : string){
    this.searchText = searchText;
  }
  
  getUserDetails() {
    this.sub = this.userService.getUserDetails(this.username).subscribe({
      next: (data: UserDetails) => {
        const variable  = data.dateOfBirth.substring(0,10);
        data.dateOfBirth = variable;
        this.user = data;
        this.initialDetails = JSON.parse(JSON.stringify(data)); 
      },
    });
  }

}
