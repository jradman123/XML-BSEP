import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ConnectionService {

  constructor(private _http: HttpClient) { }

  getUsersConnections(username: string) {
    return this._http.get<any>(
      'http://localhost:9090/connection/connected/' + username
    );
  }

  getUsersInvitations(username: string) {
    return this._http.get<any>(
      'http://localhost:9090/connection/requests/' + username
    );
  }

  connectUsers(senderUsername: string, recieverUsername: string){
    return this._http.post<any>(
      'http://localhost:9090/connection/new', {
          "userSender": senderUsername,
          "userReceiver": recieverUsername,
      }
    );
  }
   
  acceptConnection(senderUsername: string, recieverUsername: string){
    return this._http.post<any>(
      'http://localhost:9090/connection/accepted', {
        "userSender": senderUsername,
        "userReceiver": recieverUsername,
    }
    );
  }

}
