import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Subscription } from 'rxjs';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/UserService/user.service';
@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css'],
})
export class LandingPageComponent implements OnInit {
  sub!: Subscription;
  constructor(private authService: UserService, private _router: Router) {}
  emaill!: string;

  ngOnInit(): void {}

  forgotPass() {
    console.log(this.emaill);
    this.authService.sendCode(this.emaill).subscribe();
    localStorage.setItem('email', this.emaill);
    this._router.navigate(['/resetPassword']);
  }

  onSubmit(f: NgForm) {
    const loginObserver = {
      next: (x: any) => {
        console.log(x);
        console.log('uuuuuuuuuuu');
        if (x === 'user') {
          this._router.navigate(['/chome']);
          console.log('aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa');
        } else {
          this._router.navigate(['/ahome']);
          console.log('bbbbbbbbbbbbbbbbbbbb');
        }
      },
      error: (err: any) => {
        console.log(err);
      },
    };
    this.authService.login(f.value).subscribe({
      next: (x: any) => {
        localStorage.setItem('email', x.email);
        localStorage.setItem('role', x.role);
        console.log(x);
        if (x.role === 'user') {
          this._router.navigate(['/chome']);
          console.log('aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa');
        } else {
          this._router.navigate(['/ahome']);
          console.log('bbbbbbbbbbbbbbbbbbbb');
        }
      },
    });
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
