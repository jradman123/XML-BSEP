import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ICompanyInfo } from 'src/app/interfaces/company-info';
import { NewCompanyRequestDto } from 'src/app/interfaces/new-company-request-dto';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class CompanyService {

  private apiServerUrl = environment.apiBaseUrl;
  RegisterCompany(company: NewCompanyRequestDto):Observable<any> {
    return this.http.post(`${this.apiServerUrl}/company/new`,company,{
      responseType: 'text',
    });
    
  }

  constructor(private http : HttpClient) { }
}
