import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { JobOffer } from 'src/app/interfaces/job-offer';

@Injectable({
  providedIn: 'root'
})

export class JobOfferService {

  constructor(private http : HttpClient) {
   }

   getAllJobOffers() : Observable<any> {
    return this.http.get("http://localhost:9000/job_offer");
   }

   createJobOffer(newjo : JobOffer) : Observable<any> {
     return this.http.post("http://localhost:9000/job_offer", newjo);
   }

   getSuggestedJobOffers(username : string) : Observable<any> {
    return this.http.get("http://localhost:9000/jobOffers/recommended/" + username);
   }

   getMyJobOffers(username : string) : Observable<any> {
    return this.http.get("http://localhost:9000/job_offer/" + username);
   }
}
