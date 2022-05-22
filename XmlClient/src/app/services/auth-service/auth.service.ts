import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { BehaviorSubject, map, Observable } from 'rxjs';
import { LoggedUser } from 'src/app/interfaces/logged-user';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private currentUserSubject: BehaviorSubject<LoggedUser>;
  public currentUser: Observable<LoggedUser>;
  private user: LoggedUser | undefined;
  private router: Router | undefined;

  constructor(private http: HttpClient) {

    this.currentUserSubject = new BehaviorSubject<LoggedUser>(
      JSON.parse(localStorage.getItem('currentUser')!)
    );
    this.currentUser = this.currentUserSubject.asObservable();
  }
  
  login(value: any) {
    return this.http.post('http://localhost:8082/login', value).pipe(
      map((response: any) => {
        if (response && response.jwt) {
          localStorage.setItem('token', response.jwt);
          localStorage.setItem('currentUser', JSON.stringify(response));
          this.currentUserSubject.next(response);
          window.location.href = 'http://localhost:4200/home';
        }
        return this.user;
      })
    );
  }

}
