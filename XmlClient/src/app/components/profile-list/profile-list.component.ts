import { Component, OnInit } from '@angular/core';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-profile-list',
  templateUrl: './profile-list.component.html',
  styleUrls: ['./profile-list.component.css']
})
export class ProfileListComponent implements OnInit {

  profiles! : UserDetails[];
  searchText : string = "";
  isLoggedIn = localStorage.getItem('token') != null;
  
  constructor(private _service : UserService) { }

  ngOnInit(): void {
    this._service.getUsers().subscribe(
      res => {
        this.profiles =res.users.filter( (user: any ) =>
          !(user.username === localStorage.getItem('username'))
        );
      }
    )
  }

  handleMe(searchText : string){
    this.searchText = searchText;
  }

}
