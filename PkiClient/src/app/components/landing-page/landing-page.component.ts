import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Subscription } from 'rxjs';
import { LoginService } from 'src/app/services/login.service';
import { Router } from '@angular/router';
@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css']
})
export class LandingPageComponent implements OnInit {

  sub! : Subscription;
  constructor(private authService: LoginService,private _router: Router,) { 
    
  }

  ngOnInit(): void {
  }
  onSubmit(f: NgForm) {

    const loginObserver = {
      next: (x: any) => {
        console.log(x);
        console.log("uuuuuuuuuuu");
        if(x === "user"){
          this._router.navigate(['/chome']);
          console.log("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
        }else {this._router.navigate(['/ahome']);
        console.log("bbbbbbbbbbbbbbbbbbbb");
        }

      },
      error: (err: any) => {
        console.log(err);
      },
    };
    this.authService.login(f.value).subscribe({
      next: (x:any) =>{
        localStorage.setItem('email',x.email);
        localStorage.setItem('role',x.role);
        console.log(x);
        if(x.role === "user"){
          this._router.navigate(['/chome']);
          console.log("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
        }else {this._router.navigate(['/ahome']);
        console.log("bbbbbbbbbbbbbbbbbbbb");
        }
      }
    }
    );
  }
  // login(f: NgForm) {

  //   const loginObserver = {
  //     next: (x: any) => {
  //       console.log(x);
  //       if(x == 'user'){
  //         this._router.navigate(['/chome']);
  //       }else this._router.navigate(['/ahome']);
  //     },
  //     error: (err: any) => {
  //       console.log(err);
  //     },
  //   };
  //   this.authService.login(f.value).subscribe(loginObserver);
  // }
}
