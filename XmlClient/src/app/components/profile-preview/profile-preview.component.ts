import { Component, Input, OnInit } from '@angular/core';
import { UserDetails } from 'src/app/interfaces/user-details';

@Component({
  selector: 'app-profile-preview',
  templateUrl: './profile-preview.component.html',
  styleUrls: ['./profile-preview.component.css']
})
export class ProfilePreviewComponent implements OnInit {

  @Input()
  profile! : UserDetails;
  isLoggedIn! : boolean;
  constructor() { }

  ngOnInit(): void {
    this.isLoggedIn = localStorage.getItem('token') !== null;
  }

}
