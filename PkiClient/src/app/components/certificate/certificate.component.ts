import { Component, Input, OnInit } from '@angular/core';
import { FoodNode } from '../certificate-chain/certificate-chain.component';

@Component({
  selector: 'app-certificate',
  templateUrl: './certificate.component.html',
  styleUrls: ['./certificate.component.css']
})
export class CertificateComponent implements OnInit {
  @Input()
  items!: any;
  constructor() { }

  ngOnInit(): void {
    console.log(this.items);
    
  }

}
