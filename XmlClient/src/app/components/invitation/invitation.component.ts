import { Component, Input, OnInit } from '@angular/core';
import { UserDetails } from 'src/app/interfaces/user-details';
import { ConnectionService } from 'src/app/services/connection-service/connection.service';

@Component({
  selector: 'app-invitation',
  templateUrl: './invitation.component.html',
  styleUrls: ['./invitation.component.css']
})
export class InvitationComponent implements OnInit {
 
  @Input()
  profile! : UserDetails;
  accepted = false;

  constructor(private _connectionService : ConnectionService) { }

  ngOnInit(): void {
  }

  acceptInvite(username: string){
    this._connectionService.acceptConnection(localStorage.getItem('username')!, username).subscribe(
      res => {
        this.accepted = true;
      }
    );
  }
}