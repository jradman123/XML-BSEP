import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { IComment } from 'src/app/interfaces/comment';
import { IInterview } from 'src/app/interfaces/interview';
import { IsUsersCompanyDto } from 'src/app/interfaces/is-users-company-dto';
import { IJobOffer } from 'src/app/interfaces/job-offer';
import { IJobOfferRequest } from 'src/app/interfaces/job-offer-request';
import { NewCompanyRequestDto } from 'src/app/interfaces/new-company-request-dto';
import { ISalaryComment } from 'src/app/interfaces/salary-comment';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class CompanyService {
  CreateComment(comment: IComment): Observable<any> {
    return this.http.post(`${this.apiServerUrl}/api/company/comment`, comment);
  }
  CreateInterview(interview: IInterview) : Observable<any> {
    return this.http.post(`${this.apiServerUrl}/api/company/interview`, interview);
  }
  CreateSalaryComment(comment: ISalaryComment) : Observable<any> {
    return this.http.post(`${this.apiServerUrl}/api/company/salary-comment`, comment);
  }

  CreateJobOffer(jobOffer: IJobOfferRequest): Observable<any> {
    return this.http.post(`${this.apiServerUrl}/api/company/create-offer`, jobOffer);
  }
  UpdateInfo(company: any): Observable<any> {
    return this.http.put(`${this.apiServerUrl}/api/company/edit/` + company.companyId
      , company);
  }

  private apiServerUrl = environment.apiBaseUrl;
  RegisterCompany(company: NewCompanyRequestDto): Observable<any> {
    return this.http.post(`${this.apiServerUrl}/api/company/new`, company, {
      responseType: 'text',
    });

  }

  constructor(private http: HttpClient) { }

  getAlCompaniesForUser(): Observable<any> {
    return this.http.get(`${this.apiServerUrl}/api/company/search-companies`);

  }

  getAllPendingCompanies(): Observable<any> {
    return this.http.get(`${this.apiServerUrl}/api/company/pending`);

  }

  approveRequest(id: number): Observable<any> {
    return this.http.get(`${this.apiServerUrl}/api/company/approve/` + id);
  }

  rejectRequest(id: number): Observable<any> {
    return this.http.get(`${this.apiServerUrl}/api/company/reject/` + id);
  }
  getAllUsersCompanies(username: string): Observable<any> {
    return this.http.get(`${this.apiServerUrl}/api/company/users-company/` + username);
  }

  getById(id: any): Observable<any> {
    return this.http.get(`${this.apiServerUrl}/api/company/` + id);
  }
  getOffersForCompany(id: string): Observable<IJobOffer[]> {
    return this.http.get<IJobOffer[]>(`${this.apiServerUrl}/api/offer/all/` + id);
  }

  getCommentsForCompany(id: string): Observable<IComment[]> {
    return this.http.get<IComment[]>(`${this.apiServerUrl}/api/company/` + id + `/comments`);
  }
  getInterviewsForCompany(id: string): Observable<IInterview[]> {
    return this.http.get<IInterview[]>(`${this.apiServerUrl}/api/company/` + id + `/interviews`);
  }
  getSalaryCommentsForCompany(id: string): Observable<ISalaryComment[]> {
    return this.http.get<ISalaryComment[]>(`${this.apiServerUrl}/api/company/` + id + `/salary-comments`);
  }

  isUsersCompany(id : string) : Observable<any>{
    return this.http.get<IsUsersCompanyDto>(`${this.apiServerUrl}/api/company/isUsersCompany/` + id);
  }
  

}
