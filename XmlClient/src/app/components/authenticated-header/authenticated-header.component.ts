import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-authenticated-header',
  templateUrl: './authenticated-header.component.html',
  styleUrls: ['./authenticated-header.component.css']
})
export class AuthenticatedHeaderComponent implements OnInit {

  constructor(private userService : UserService,private router : Router) { }

  ngOnInit(): void {
  }

  logout() {
    this.userService.logout();
    this.router.navigate(['login']);
  }

}
