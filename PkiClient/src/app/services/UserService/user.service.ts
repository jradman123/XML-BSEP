import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';
import { LogedUser } from 'src/app/interfaces/loged-user';
import { SubjectData } from 'src/app/interfaces/subject-data';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private _http: HttpClient) { }

  createSubject(newSubject: SubjectData): Observable<any> {
    return this._http.post<any>(
      'http://localhost:8443/api/createSubject',
      newSubject
    );
  }

  login(model: LogedUser): any {
    return this._http.post('http://localhost:8443/api/login', model);
  }
}
