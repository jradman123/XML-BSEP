import { Component, OnInit } from '@angular/core';
import { LoggedUserDto } from 'src/app/interfaces/logged-user-dto';
import { UserServiceService } from 'src/app/services/UserService/user-service.service';

@Component({
  selector: 'app-auth-header',
  templateUrl: './auth-header.component.html',
  styleUrls: ['./auth-header.component.css']
})
export class AuthHeaderComponent implements OnInit {

  currentUser : LoggedUserDto;
  isOwner : boolean = false;
  constructor(private userService : UserServiceService) {
    this.currentUser = {} as LoggedUserDto;
   }

  ngOnInit(): void {
   this.currentUser = this.userService.getUserValue();
   if(this.currentUser.role === "OWNER"){
      this.isOwner = true;
   }
  }

  logout() : void {
    this.userService.logout();
  }

}
