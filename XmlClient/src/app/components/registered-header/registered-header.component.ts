import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth-service/auth.service';

@Component({
  selector: 'app-registered-header',
  templateUrl: './registered-header.component.html',
  styleUrls: ['./registered-header.component.css']
})
export class RegisteredHeaderComponent implements OnInit {

  username : any;
  visible : boolean = false;
  constructor(private authService : AuthService,private router : Router) { }

  ngOnInit(): void {
    this.username=localStorage.getItem('username');
  }

  click() {
    console.log('usao');
    if(this.visible){
      console.log('true')
      this.visible=false;
    }else{
      console.log('false')
      this.visible=true;
    }
  }

  logout() {
    this.authService.logout();
    this.router.navigate(['login']);
  }

  showProfile() {
    this.router.navigate(['myProfile']);
  }

}
