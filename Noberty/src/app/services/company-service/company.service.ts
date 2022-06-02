import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { IComment } from 'src/app/interfaces/comment';
import { ICompanyInfo } from 'src/app/interfaces/company-info';
import { IInterview } from 'src/app/interfaces/interview';
import { IJobOffer } from 'src/app/interfaces/job-offer';
import { NewCompanyRequestDto } from 'src/app/interfaces/new-company-request-dto';
import { ISalaryComment } from 'src/app/interfaces/salary-comment';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class CompanyService {
  
  CreateJobOffer(jobOffer: IJobOffer): Observable<ICompanyInfo> {
    throw new Error('Method not implemented.');
  }
  UpdateInfo(description: string): Observable<ICompanyInfo> {
    throw new Error('Method not implemented.');
  }

  private apiServerUrl = environment.apiBaseUrl;
  RegisterCompany(company: NewCompanyRequestDto):Observable<any> {
    return this.http.post(`${this.apiServerUrl}/company/new`,company,{
      responseType: 'text',
    });
     
  }

  constructor(private http : HttpClient) { }

  getAlCompaniesForUser() : Observable<any>{
    return this.http.get(`${this.apiServerUrl}/company/getAllForUser`);

  }

  getOffersForCompany(id : string) : Observable<IJobOffer[]>{
    return this.http.get<IJobOffer[]>(`${this.apiServerUrl}/offer/all/` + id);
  }

  getCommentsForCompany(id: string): Observable<IComment[]> {
    return this.http.get<IComment[]>(`${this.apiServerUrl}/company/` + id + `/comments`);
  }
  getInterviewsForCompany(id: string): Observable<IInterview[]> {
    return this.http.get<IInterview[]>(`${this.apiServerUrl}/company/` + id+ `/interviews`);
  }
  getSalaryCommentsForCompany(id: string) : Observable<ISalaryComment[]> {
    return this.http.get<ISalaryComment[]>(`${this.apiServerUrl}/company/` + id+ `/salaryComments`);
  }
  

}
