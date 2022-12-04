import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiKeyService {

  constructor(private http: HttpClient) {

   }

   generateApiKey(username : string) : Observable<any> {
     return this.http.post("http://localhost:9000/users/token/generate", {username});
   }
}
