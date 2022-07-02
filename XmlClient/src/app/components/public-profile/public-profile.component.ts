import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { UserDetails } from 'src/app/interfaces/user-details';
import { ConnectionService } from 'src/app/services/connection-service/connection.service';
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
  isLoggedIn = localStorage.getItem('token') !== null;
  buttonText = ""


  constructor(private userService : UserService, private _router : Router, private _connectionService : ConnectionService) {
    this.username = this._router.url.substring(16) ?? '';
    this.user = {} as UserDetails;
    this.getUserDetails();
  }

  ngOnInit(): void {
    this._connectionService.connectionStatus(localStorage.getItem('username')!, this.username).subscribe(
      res => {
        if (res.connectionStatus === "BLOCKED"){
          this._router.navigate(['/404']);
        }
        else if (res.connectionStatus === "CONNECTED"){
          this.buttonText = "Following";
        }
        else if (res.connectionStatus === "REQUEST_SENT"){
          this.buttonText = "Pending"
        }
        else this.buttonText = "Connect"
      }
    );
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

  connect(username:string){
    this._connectionService.connectUsers(localStorage.getItem('username')!, username).subscribe(
      res => {
        let status = res.connectionStatus;
        status === "CONNECTED" ? this.buttonText = "Following" : this.buttonText = "Pending";

      }
    )
  }

  block(username:string){
    // this._connectionService.blockUser(localStorage.getItem('username')!, username).subscribe(
    //   res => {
    //     // mat snack we did it
    //     this._router.navigate(['/feed']);
    //   }
    // )
  }

}
