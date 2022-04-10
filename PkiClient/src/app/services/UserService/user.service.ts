import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';
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
}
