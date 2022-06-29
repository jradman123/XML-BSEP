import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { BehaviorSubject, Observable, map, Subscription, delay, of } from 'rxjs';
import { IAuthenticate } from 'src/app/interfaces/authenticate';
import { LoggedUser } from 'src/app/interfaces/logged-user';
import { ILoginRequest } from 'src/app/interfaces/login-request';
import { UserData } from 'src/app/interfaces/subject-data';
import { IUsername } from 'src/app/interfaces/username';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private currentUserSubject: BehaviorSubject<LoggedUser>;
  public currentUser: Observable<LoggedUser>;
  private user!: LoggedUser;
  jwtToken!: any;
  timeout!: any;
  tokenSubscription = new Subscription();

  constructor(
    private _http: HttpClient,
    private jwtHelper: JwtHelperService,
    private router: Router) {
    this.currentUserSubject = new BehaviorSubject<LoggedUser>(
      JSON.parse(localStorage.getItem('currentUser')!)
    );
    this.currentUser = this.currentUserSubject.asObservable();
  }

  auth(loginReq: ILoginRequest): Observable<any> {
    return this._http
      .post(`http://localhost:9090/users/auth/user`, loginReq)
      .pipe(
        map((response: any) => {
          if (response) {
            localStorage.setItem('username', response.username);
            this.currentUserSubject.next(response);

          }
          return response;
        })
      );
  }

  login(loginRegularRequest: IUsername): Observable<LoggedUser> {
    return this._http
      .post(`http://localhost:9090/users/auth/user/regular`, loginRegularRequest)
      .pipe(
        map((response: any) => {
          if (response) {
            this.storeUserData(response)

          }
          return response;

        })
      );
  }

  authenticate2FA(request: IAuthenticate): Observable<any> {
    return this._http.post<any>(
      'http://localhost:9090/2fa/authenticate',
      request
    ).pipe(
      map((response: any) => {
        if (response) {
          this.storeUserData(response)
        }
        return response;
      })
    );
  }
  passwordlessLoginRequest(username: any) {
    return this._http.post<any>(
      'http://localhost:9090/users/login/passwordless',
      { username }
    );
  }

  passwordlessLogin(code: any) {
    return this._http.get<any>(
      'http://localhost:9090/users/login/passwordless/' + code
    )
      .pipe(
        map((response: any) => {
          if (response) {
            this.storeUserData(response)
          }
          return this.user;
        })
      );
  }
  storeUserData(response: any) {

    this.jwtToken = this.jwtHelper.decodeToken(response.token);
    console.log(this.jwtToken);

    localStorage.setItem('token', response.token);
    localStorage.setItem('currentUser', JSON.stringify(response));
    localStorage.setItem('role', this.jwtToken.roles[0]);
    localStorage.setItem('email', response.email);
    localStorage.setItem('username', this.jwtToken.username);
    var now = new Date().valueOf();

    console.log(this.jwtToken.exp*1000 )
    console.log(now);

    this.timeout = this.jwtToken.exp*1000 - now;
    this.expirationCounter(this.timeout)
    this.currentUserSubject.next(response);
  }
  expirationCounter(timeout: any) {
    this.tokenSubscription.unsubscribe();
    this.tokenSubscription = of(null).pipe(delay(timeout)).subscribe((expired) => {
      console.log(expired);

      console.log('EXPIRED!!');

      this.logout();
      this.router.navigate(["/login"]);
    });
  }

  logout() {
    localStorage.removeItem('token');
    localStorage.removeItem('currentUser');
    localStorage.removeItem('role');
    localStorage.removeItem('email');
    localStorage.removeItem('username');
  }

  public get currentUserValue(): LoggedUser {
    return this.currentUserSubject.value;
  }

  loggedIn(): boolean {
    const token = localStorage.getItem('token');
    return true;
  }

}

