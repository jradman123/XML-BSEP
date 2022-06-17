import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, Subject } from 'rxjs';
import { LogedUser } from 'src/app/interfaces/loged-user';
import { NewPassword } from 'src/app/interfaces/new-password';
import { SubjectData } from 'src/app/interfaces/subject-data';
import { map } from 'rxjs/operators';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private currentUserSubject: BehaviorSubject<LogedUser>;
  public currentUser: Observable<LogedUser>;
  private user!: LogedUser;

  checkCode(verCode: string): Observable<any> {
    return this._http.post<any>('http://localhost:8443/api/checkCode', {
      email: localStorage.getItem('emailForReset'),
      code: verCode,
    });
  }

  sendCode(email: string): Observable<any> {
    return this._http.post<any>('http://localhost:8443/api/sendCode', email);
  }

  constructor(private _http: HttpClient,private router : Router) {
    this.currentUserSubject = new BehaviorSubject<LogedUser>(
      JSON.parse(localStorage.getItem('currentUser')!)
    );
    this.currentUser = this.currentUserSubject.asObservable();
  }

  createSubject(newSubject: SubjectData): Observable<any> {
    return this._http.post<any>(
      'http://localhost:8443/api/createSubject',
      newSubject
    );
  }

  login(model: any): Observable<LogedUser> {
    return this._http.post(`http://localhost:8443/api/login`, model).pipe(
      map((response: any) => {
        if (response && response.token) {
          localStorage.setItem('token', response.token.accessToken);
          localStorage.setItem('currentUser', JSON.stringify(response));
          localStorage.setItem('role', response.role);
          localStorage.setItem('email', response.email);
          this.currentUserSubject.next(response);
        }
        return this.user;
      })
    );
  }

  changePassword(data: any) {
    return this._http.put(`http://localhost:8443/api/changePassword`, data);
  }

  resetPassword(newPassword: string): Observable<any> {
    return this._http.post<any>('http://localhost:8443/api/resetPassword', {
      email: localStorage.getItem('emailForReset'),
      newPassword: newPassword,
    });
  }

  public get currentUserValue(): LogedUser {
    return this.currentUserSubject.value;
}

logout() {
  localStorage.removeItem('token');
  localStorage.removeItem('role');
  localStorage.removeItem('currentUser');
  localStorage.removeItem('email');
  this.router.navigate(['/']);
}

}
