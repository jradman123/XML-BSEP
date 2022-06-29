import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-unauthenticated-header',
  templateUrl: './unauthenticated-header.component.html',
  styleUrls: ['./unauthenticated-header.component.css']
})
export class UnauthenticatedHeaderComponent implements OnInit {

  mobNavIcon = 'bi bi-list mobile-nav-toggle';
  mobNav = "landing-navbar";
  constructor() { }

  ngOnInit(): void {
  }

  changeClass(){
    if (this.mobNav === "landing-navbar") this.mobNav = "landing-navbar landing-navbar-mobile";
    else this.mobNav = "landing-navbar";

    if ( this.mobNavIcon === 'bi bi-list mobile-nav-toggle') this.mobNavIcon = 'bi mobile-nav-toggle bi-x'
    else this.mobNavIcon = 'bi bi-list mobile-nav-toggle'
   }

}
