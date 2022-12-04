import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { IJobOfferPublish } from 'src/app/interfaces/ijob-offer-publish';

@Injectable({
  providedIn: 'root'
})
export class JobOfferService {

  constructor(private http: HttpClient) { 
  }

  publishJobOfferOnDislinkt(jobOffer : IJobOfferPublish) : Observable<any> {
    return this.http.post('http://localhost:9000/users/share/jobOffer', jobOffer);
  }
  
}
