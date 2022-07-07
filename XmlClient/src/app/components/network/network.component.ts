import { Component, OnInit } from '@angular/core';
import { UserDetails } from 'src/app/interfaces/user-details';
import { ConnectionService } from 'src/app/services/connection-service/connection.service';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-network',
  templateUrl: './network.component.html',
  styleUrls: ['./network.component.css']
})
export class NetworkComponent implements OnInit {
  searchText : string = "";
  connections! : UserDetails[];
  invitations! : UserDetails[];
  people! : UserDetails[];

  constructor(private _connectionService : ConnectionService, private userService : UserService) { }

  ngOnInit(): void {

    this._connectionService.getUsersRecommendation(localStorage.getItem('username')!).subscribe(
      res => {
        this.people = res.users
      }
    )

    this._connectionService.getUsersConnections(localStorage.getItem('username')!).subscribe(
      res => {
        this.connections = res.users
      }
    );

    this._connectionService.getUsersInvitations(localStorage.getItem('username')!).subscribe(
      res => {
        this.invitations = res.users;
      }
      
    )
  }

  handleMe(searchText : string){
    this.searchText = searchText;
  }
}
