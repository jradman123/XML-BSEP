import { Component, OnInit } from '@angular/core';
import { PageEvent } from '@angular/material/paginator';
import { ICompanyInfo } from 'src/app/interfaces/company-info';

@Component({
  selector: 'app-companies-list',
  templateUrl: './companies-list.component.html',
  styleUrls: ['./companies-list.component.css']
})
export class CompaniesListComponent implements OnInit {
  items!: ICompanyInfo[]
  constructor() { }

  ngOnInit(): void {
    this.items = [{
      name: "Endava",
      site: "string",
      headquaters: "string",
      founded: "string",
      industry: "Software Outsourcing",
      employees: 1000,
      origin: "string",
      offices: "Beograd, Novi Sad, Kragujevac, Čača"
    },
    
    {
      name: "Levi9 Technology Services",
      site: "string",
      headquaters: "string",
      founded: "string",
      industry: "Software Outsourcing",
      employees: 800,
      origin: "string",
      offices: "Beograd, Novi Sad, Zrenjanin"
    },
    {
      name: "Synechron",
      site: "string",
      headquaters: "string",
      founded: "string",
      industry: "IT Services",
      employees: 500,
      origin: "string",
      offices: "Novi Sad, Beograd"
    }
    ]
  }


}
