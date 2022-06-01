import { Component, Input, OnInit } from '@angular/core';
import { ICompanyInfo } from 'src/app/interfaces/company-info';

@Component({
  selector: 'app-company-list-view',
  templateUrl: './company-list-view.component.html',
  styleUrls: ['./company-list-view.component.css']
})
export class CompanyListViewComponent implements OnInit {
  @Input() 
  items!:ICompanyInfo[]
  constructor() { }

  ngOnInit(): void {
  }

}
