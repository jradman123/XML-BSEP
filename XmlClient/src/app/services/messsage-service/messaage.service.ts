import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { IMesssage } from 'src/app/interfaces/message';

@Injectable({
  providedIn: 'root'
})
export class MessageService {
 
  constructor(private _http: HttpClient) { }

  SendMessage(newMessaage: IMesssage) {
    return this._http.post<any>(
      'http://localhost:9000/messages/send' ,
        newMessaage
    );
  }
  GetSentMessages() {
    return this._http.get<any>(
      'http://localhost:9000/messages/' + localStorage.getItem('username') + "/sent",
    );
  }
  GetReceivedMessages() {
    return this._http.get<any>(
      'http://localhost:9000/messages/' + localStorage.getItem('username') + "/received",
    );
  }

}
