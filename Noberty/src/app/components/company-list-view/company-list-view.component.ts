import { Component, Input, OnInit } from '@angular/core';
import { CompanyResponseDto } from 'src/app/interfaces/company-response-dto';

@Component({
  selector: 'app-company-list-view',
  templateUrl: './company-list-view.component.html',
  styleUrls: ['./company-list-view.component.css']
})
export class CompanyListViewComponent implements OnInit {
  @Input()
  item!:CompanyResponseDto

  
  constructor() { 
  
  }

  ngOnInit(): void {
  }

}
