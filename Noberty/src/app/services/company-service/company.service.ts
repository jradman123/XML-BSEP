import { Injectable } from '@angular/core';
import { observable, Observable } from 'rxjs';
import { ICompanyInfo } from 'src/app/interfaces/company-info';
import { IJobOffer } from 'src/app/interfaces/job-offer';

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
  RegisterCompany(company: ICompanyInfo): Observable<ICompanyInfo> {
    throw new Error('Method not implemented.');
  }

  constructor() { }
}
