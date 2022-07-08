import { DatePipe } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { IMesssage, IMesssages } from 'src/app/interfaces/message';
import { UserDetails } from 'src/app/interfaces/user-details';
import { ConnectionService } from 'src/app/services/connection-service/connection.service';
import { MessageService } from 'src/app/services/messsage-service/messaage.service';

@Component({
  selector: 'app-chatbox',
  templateUrl: './chatbox.component.html',
  styleUrls: ['./chatbox.component.css']
})
export class ChatboxComponent implements OnInit {

  receiver : string = '';
  sender : string = localStorage.getItem('username') ?? ''; 
  connections : UserDetails[] = []
  searchText = ''
  searchConnection = ''

  constructor(private _connectionService : ConnectionService) { 
  }

  ngOnInit(): void {
    this._connectionService.getUsersConnections(localStorage.getItem('username')!).subscribe(
      res => {
        this.connections = res.users
      }
    )
  }

  handleMe(searchText : string){
    this.searchText = searchText;
  }

  openChat(username : string) {
    this.receiver = username;
  }
}
