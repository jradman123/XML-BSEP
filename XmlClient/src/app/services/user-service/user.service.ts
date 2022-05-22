import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, map, Observable } from 'rxjs';
import { LogedUser } from 'src/app/interfaces/logged-user';
import { UserData } from 'src/app/interfaces/subject-data';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private currentUserSubject: BehaviorSubject<LogedUser>;
  public currentUser: Observable<LogedUser>;
  private user! : LogedUser;

  constructor(private _http: HttpClient) { 
    this.currentUserSubject = new BehaviorSubject<LogedUser>(
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

  login(model: any): Observable<LogedUser> {
    return this._http.post(`http://localhost:8082/login`, model).pipe(
      map((response: any) => {
        if (response && response.token) {
          localStorage.setItem('token', response.token.accessToken);
          localStorage.setItem('currentUser', JSON.stringify(response));
          localStorage.setItem('role' ,response.role)
          localStorage.setItem('email' ,response.email)
          this.currentUserSubject.next(response);
        }
        return this.user;
      })
    );
  }

  logout() {
    localStorage.removeItem('token');
    localStorage.removeItem('currentUser');
    localStorage.removeItem('role');
    localStorage.removeItem('email');
  }

  public get currentUserValue(): LogedUser {
    return this.currentUserSubject.value;
  }

  loggedIn(): boolean {
    const token = localStorage.getItem('token');
    return true;
  }

}
