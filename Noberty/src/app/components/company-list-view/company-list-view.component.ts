import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ICompanyInfo } from 'src/app/interfaces/company-info';

@Component({
  selector: 'app-company-list-view',
  templateUrl: './company-list-view.component.html',
  styleUrls: ['./company-list-view.component.css']
})
export class CompanyListViewComponent implements OnInit {
  @Input() 
  item!:ICompanyInfo
  constructor(private router:Router) { }

  seeMore() {
    this.router.navigate(['/company/' + this.item.id]);
  }
  ngOnInit(): void {
  }

}
