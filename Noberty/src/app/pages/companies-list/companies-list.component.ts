import { Component, OnInit } from '@angular/core';
import { CompanyResponseDto } from 'src/app/interfaces/company-response-dto';
import { CompanyService } from 'src/app/services/company-service/company.service';

@Component({
  selector: 'app-companies-list',
  templateUrl: './companies-list.component.html',
  styleUrls: ['./companies-list.component.css']
})
export class CompaniesListComponent implements OnInit {
  items!: CompanyResponseDto[]
  constructor(private companyService : CompanyService) { }

  ngOnInit(): void {
    this.companyService.getAlCompaniesForUser().subscribe((res) => {this.items = res});
  }


}
