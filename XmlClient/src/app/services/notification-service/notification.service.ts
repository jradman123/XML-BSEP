import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class NotificationService {

  constructor(private _http: HttpClient) { }

  getUsersNotifications(username: string) {
    return this._http.get<any>(
      'http://localhost:9090/notification/user/' + username
    );
  }

  markAsRead(id: string) {
    let idAsStr = "\""  + id + "\"" 
    console.log(idAsStr)
    return this._http.post<any>(
      'http://localhost:9090/notification/read', idAsStr
    );
  }

}
