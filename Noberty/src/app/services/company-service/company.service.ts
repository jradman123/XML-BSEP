import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ICompanyInfo } from 'src/app/interfaces/company-info';

@Injectable({
  providedIn: 'root'
})
export class CompanyService {
  RegisterCompany(company: ICompanyInfo):Observable<ICompanyInfo> {
    throw new Error('Method not implemented.');
  }

  constructor() { }
}
