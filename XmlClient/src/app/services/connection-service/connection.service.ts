import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ConnectionService {

  constructor(private _http: HttpClient) { }

  getUsersConnections(username: string) {
    return this._http.get<any>(
      'http://localhost:9000/connection/connected/' + username
    );
  }

  getUsersInvitations(username: string) {
    return this._http.get<any>(
      'http://localhost:9000/connection/requests/' + username
    );
  }

  getUsersRecommendation(username: string) {
    return this._http.get<any>(
      'http://localhost:9000/connection/recommended/' + username
    );
  }

  connectUsers(senderUsername: string, recieverUsername: string){
    return this._http.post<any>(
      'http://localhost:9000/connection/new', {
          "userSender": senderUsername,
          "userReceiver": recieverUsername,
      }
    );
  }
   
  acceptConnection(senderUsername: string, recieverUsername: string){
    return this._http.post<any>(
      'http://localhost:9000/connection/accepted', {
        "userSender": senderUsername,
        "userReceiver": recieverUsername,
    }
    );
  }

  connectionStatus(senderUsername: string, recieverUsername: string){
    return this._http.post<any>(
      'http://localhost:9000/connection/status', {
          "userSender": senderUsername,
          "userReceiver": recieverUsername,
      }
    );
  }

  blockUser(senderUsername: string, recieverUsername: string){
      return this._http.post<any>(
        'http://localhost:9000/connection/block', {
            "userSender": senderUsername,
            "userReceiver": recieverUsername,
        }
      );
  }
}