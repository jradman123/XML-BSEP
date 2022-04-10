import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { LogedUser } from '../interfaces/loged-user';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  
  constructor(private http: HttpClient
    ) { }

    login(model: LogedUser): any {
      return this.http.post('http://localhost:8443/api/login', model);
    }
}
