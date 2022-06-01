import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { BehaviorSubject, Observable } from 'rxjs';
import { LoggedUserDto } from 'src/app/interfaces/logged-user-dto';
import { NewUser } from 'src/app/interfaces/new-user';
import { environment } from 'src/environments/environment';
import { map } from 'rxjs/operators';
import { UserInformationResponseDto } from 'src/app/interfaces/user-information-response-dto';

@Injectable({
  providedIn: 'root'
})
export class UserServiceService {

  private currentUserSubject: BehaviorSubject<LoggedUserDto>;
  public currentUser: Observable<LoggedUserDto>;
  private user! : LoggedUserDto;
  private loginStatus = new BehaviorSubject<boolean>(false);

  private apiServerUrl = environment.apiBaseUrl;
  constructor(private http: HttpClient, private router : Router) {
    this.currentUserSubject = new BehaviorSubject<LoggedUserDto>(
      JSON.parse(localStorage.getItem('currentUser')!)
    );
    this.currentUser = this.currentUserSubject.asObservable();
   }

   public get currentUserValue(): LoggedUserDto {
    return this.currentUserSubject.value;
}

public getUserValue() : LoggedUserDto {
  console.log("Token" + this.currentUserSubject.value.token.accessToken);
  return this.currentUserSubject.value;
}

registerUser(newUser: NewUser) {
return this.http.post(`${this.apiServerUrl}/api/signup`, newUser, {
  responseType: 'text',
});
}

loggedIn(): boolean {
  const token = localStorage.getItem('token');
  return true;
}

get isLoggedIn() {
  return this.loginStatus.asObservable();
}

login(model: any): Observable<LoggedUserDto> {
  return this.http.post(`${this.apiServerUrl}/api/login`, model).pipe(
    map((response: any) => {
      if (response && response.token) {
        this.loginStatus.next(true);
        localStorage.setItem('token', response.token.accessToken);
        localStorage.setItem('currentUser', JSON.stringify(response));
        localStorage.setItem('role' ,response.role)
        localStorage.setItem('username' ,response.username)
        this.currentUserSubject.next(response);
      }
      return this.user;
    })
  );
}

logout() {
  this.loginStatus.next(false);
  localStorage.removeItem('token');
  localStorage.removeItem('role');
  localStorage.removeItem('currentUser');
  localStorage.removeItem('username');
  this.router.navigate(['/']);
}

sendCode(email: string): Observable<any> {
  return this.http.post<any>(`${this.apiServerUrl}/api/sendCode`, email);
}

resetPassword(newPassword: string): Observable<any> {
  return this.http.post<any>(`${this.apiServerUrl}/api/resetPassword`, {
    username: localStorage.getItem('usernamee'),
    newPassword: newPassword,
  });
}

checkCode(verCode: string): Observable<any> {
  return this.http.post<any>(`${this.apiServerUrl}/api/checkCode`, {
    username: localStorage.getItem('usernamee'),
    code: verCode,
  });
}

getUserInformation() : Observable<any>{
  return this.http.get(`${this.apiServerUrl}/api/users/getUserInformation`);
}
}
