import { Component, OnInit } from '@angular/core';
import { NavigationStart, Router } from '@angular/router';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  role! : string;
  show! : boolean ;
  constructor() { }

  ngOnInit(): void {
   this.role  = localStorage.getItem('role')!;
   if(this.role === 'Admin'){
      this.show = true;
      console.log('admin');
   } else {
     this.show = false;
   }
  }
   

}
