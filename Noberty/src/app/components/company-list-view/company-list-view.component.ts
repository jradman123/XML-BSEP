import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CompanyResponseDto } from 'src/app/interfaces/company-response-dto';

@Component({
  selector: 'app-company-list-view',
  templateUrl: './company-list-view.component.html',
  styleUrls: ['./company-list-view.component.css']
})
export class CompanyListViewComponent implements OnInit {
  @Input()
  item!:CompanyResponseDto

  
  constructor(private router:Router) { 
  
  }

  seeMore() {
    this.router.navigate(['/company/' + this.item.companyId]);
  }
  ngOnInit(): void {
  }

}
