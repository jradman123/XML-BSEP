import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserDetails } from 'src/app/interfaces/user-details';

@Component({
  selector: 'app-recommended-profiles',
  templateUrl: './recommended-profiles.component.html',
  styleUrls: ['./recommended-profiles.component.css']
})
export class RecommendedProfilesComponent implements OnInit {

  @Input()
  profile! : UserDetails;
  isLoggedIn! : boolean;
  constructor(private _router: Router) { }

  ngOnInit(): void {
    this.isLoggedIn = localStorage.getItem('token') !== null;

  }

  openProfile(){
    let path = '/public-profile/' + this.profile.username;
    this._router.navigate([path]);
  }


}
