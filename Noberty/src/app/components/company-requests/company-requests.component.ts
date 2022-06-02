import { Component, OnInit } from '@angular/core';
import { CompanyResponseDto } from 'src/app/interfaces/company-response-dto';
import { CompanyService } from 'src/app/services/company-service/company.service';

@Component({
  selector: 'app-company-requests',
  templateUrl: './company-requests.component.html',
  styleUrls: ['./company-requests.component.css']
})
export class CompanyRequestsComponent implements OnInit {

  items! : CompanyResponseDto[];
  constructor(private companyService : CompanyService) { }

  ngOnInit(): void {
    this.companyService.getAllPendingCompanies().subscribe((res) => {this.items = res});
  }

  approveRequest(id : number) {
    this.companyService.approveRequest(id).subscribe({
      next: (requests : CompanyResponseDto[]) => {
        this.items = requests;

      }});

  }

  rejectRequest(id : number){
    this.companyService.rejectRequest(id).subscribe({
      next: (requests : CompanyResponseDto[]) => {
        this.items = requests;

      }});

  }

}
