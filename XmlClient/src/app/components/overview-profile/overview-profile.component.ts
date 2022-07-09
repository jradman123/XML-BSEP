import { Component, Input, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserPersonalDetails } from 'src/app/interfaces/user-personal-details';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-overview-profile',
  templateUrl: './overview-profile.component.html',
  styleUrls: ['./overview-profile.component.css']
})
export class OverviewProfileComponent implements OnInit {

  sub!: Subscription;
  @Input()
  user! : UserPersonalDetails;
  @Input()
  email! : string;
  initialDetails: any;
  id!: number;
  constructor(private userService : UserService) { }

  ngOnInit(): void {
    this.setUserDetails();
  }
  setUserDetails() {
  
        this.initialDetails = this.user; 
  }

}
