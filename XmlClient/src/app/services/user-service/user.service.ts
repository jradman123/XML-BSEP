import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ActivateAccount } from 'src/app/interfaces/activate-account';
import { NewPass } from 'src/app/interfaces/new-pass';
import { UserData } from 'src/app/interfaces/subject-data';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserPersonalDetails } from 'src/app/interfaces/user-personal-details';
import { UserProfessionalDetails } from 'src/app/interfaces/user-professional-details';


@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(private _http: HttpClient,) {
  }

  registerUser(registerRequest: UserData): Observable<any> {
    return this._http.post<any>(
      'http://localhost:9090/users/register/user',
      registerRequest
    );
  }
  
  recoverPass(recoverPass: NewPass) {
    return this._http.post<any>(
      'http://localhost:9090/users/recover/user',
      recoverPass
    );
  }

  recoverPassRequest(recoverPass: any) {
    return this._http.post<any>(
      'http://localhost:9090/users/recoveryRequest/user',
      recoverPass
    );
  }

  passIsPwned(pass: any) {
    return this._http.post<any>(
      'http://localhost:9090/users/pwnedPassword/user',
      pass
    );
  }

  activateAccount(activateData: ActivateAccount) {
    return this._http.post<any>(
      'http://localhost:9090/users/activate/user',
      activateData
    );
  }
  
  enable2FA(username: string): Observable<any> {
    return this._http.post<any>(
      'http://localhost:9090/2fa/enable',
      { username }
    );
  }
  
  disable2FA(username: string) {
    return this._http.post<any>(
      'http://localhost:9090/2fa/disable',
      { username }
    );
  }

  check2FAStatus(username: string): Observable<any> {
    return this._http.post<any>(
      'http://localhost:9090/2fa/check',
      { username }
    );
  }

  getUserDetails(username: string | null) {
    console.log('evo me u metodi, username je ovde ' + username)
    return this._http.post<UserDetails>(
      'http://localhost:9090/users/user/details', {
      username
    }
    );
  }


  getUsers() {
    return this._http.get<any>(
      'http://localhost:9090/users'
    );
  }

  updateUser(user: UserDetails) {
    return this._http.post<UserDetails>('http://localhost:9090/users/user/edit',
      user
    )
  }

  updateUserPersonalDetails(user : UserPersonalDetails){
    return this._http.post<UserPersonalDetails>('http://localhost:9090/users/user/editPersonal',
      user
    )
  }

  updateUserProfessionalDetails(user : UserProfessionalDetails){
    return this._http.post<UserProfessionalDetails>('http://localhost:9090/users/user/editProfessional',
      user
    )
  }
}
