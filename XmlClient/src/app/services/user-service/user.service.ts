import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { LoggedUser } from 'src/app/interfaces/logged-user';
import { UserData } from 'src/app/interfaces/subject-data';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private currentUserSubject: BehaviorSubject<LoggedUser>;
  public currentUser: Observable<LoggedUser>;
  private user! : LoggedUser;

  constructor(private _http: HttpClient) { 
    this.currentUserSubject = new BehaviorSubject<LoggedUser>(
      JSON.parse(localStorage.getItem('currentUser')!)
    );
    this.currentUser = this.currentUserSubject.asObservable();

  }
  registerUser(registerRequest: UserData): Observable<any> {
    return this._http.post<any>(
      'http://localhost:8082/register',
      registerRequest
    );
  }

}
