import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, map, Observable } from 'rxjs';
import { ActivateAccount } from 'src/app/interfaces/activate-account';
import { LoggedUser } from 'src/app/interfaces/logged-user';
import { LoginRequest } from 'src/app/interfaces/login-request';
import { NewPass } from 'src/app/interfaces/new-pass';
import { UserData } from 'src/app/interfaces/subject-data';
import { UserDetails } from 'src/app/interfaces/user-details';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private currentUserSubject: BehaviorSubject<LoggedUser>;
  public currentUser: Observable<LoggedUser>;
  private user!: LoggedUser;


  constructor(private _http: HttpClient) {
    this.currentUserSubject = new BehaviorSubject<LoggedUser>(
      JSON.parse(localStorage.getItem('currentUser')!)
    );
    this.currentUser = this.currentUserSubject.asObservable();


  }
  registerUser(registerRequest: UserData): Observable<any> {

    return this._http.post<any>(
      'http://localhost:9090/users/register/user',
       registerRequest);
  }

  login(loginRequest: LoginRequest): Observable<LoggedUser> {
    return this._http.post(`http://localhost:9090/users/login/user`, loginRequest).pipe(
      map((response: any) => {

        if (response && response.token) {
          localStorage.setItem('token', response.token.accessToken);
          localStorage.setItem('currentUser', JSON.stringify(response));
          localStorage.setItem('role', response.role)
          localStorage.setItem('email', response.email)
          localStorage.setItem('username', response.username)
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

  public get currentUserValue(): LoggedUser {
    return this.currentUserSubject.value;
  }

  loggedIn(): boolean {
    const token = localStorage.getItem('token');
    return true;
  }

  recoverPass(recoverPass: NewPass) {
    return this._http.post<any>(
      'http://localhost:9090/users/recover/user',
       recoverPass);
  }

  recoverPassRequest(recoverPass: any) {
    return this._http.post<any>(
      'http://localhost:9090/users/recoveryRequest/user',
       recoverPass);
  }

  passIsPwned(pass: any) {
    return this._http.post<any>(
      'http://localhost:9090/users/pwnedPassword/user',
       pass);
  }
  
  activateAccount(activateData: ActivateAccount) {
    return this._http.post<any>(
      'http://localhost:9090/users/activate/user',
       activateData);
  }

  getUserDetails(username : any) {
    return this._http.post<UserDetails>(
      'http://localhost:9090/users/user/details', {
        username
      });
  }

  updateUser(user : UserDetails) {
    return this._http.post<UserDetails>('http://localhost:9090/users/user/edit',{
      user
    })
  }


}
