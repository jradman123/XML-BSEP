import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ChangeSettingsRequest } from 'src/app/interfaces/change-settings-request';

@Injectable({
  providedIn: 'root'
})
export class NotificationService {

  constructor(private _http: HttpClient) { }

  getUsersNotifications(username: string) {
    return this._http.get<any>(
      'http://localhost:9000/notification/user/' + username
    );
  }

  markAsRead(id: string) {
    let idAsStr = "\""  + id + "\"" 
    console.log(idAsStr)
    return this._http.post<any>(
      'http://localhost:9000/notification/read', idAsStr
    );
  }

  getUsersNotificationsSettings(username: string) {
    return this._http.get<any>(
      'http://localhost:9000/notification/settings/' + username
    );
  }

  changeNotificationSettings(changeSettingsRequest : ChangeSettingsRequest) {
    return this._http.post<any>(
      'http://localhost:9000/notification/change-settings/' + changeSettingsRequest.username,
      changeSettingsRequest.settings
    );
  }

}
