import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Subscription } from 'rxjs';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/UserService/user.service';
import { MatSnackBar } from '@angular/material/snack-bar';
@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css'],
})
export class LandingPageComponent implements OnInit {
  sub!: Subscription;
  constructor(
    private authService: UserService,
    private _router: Router,
    private _snackBar: MatSnackBar
  ) {}
  emaill!: string;
  email!: string;
  password! : string;
  code! : string;
  tfaEnabled = false;
  passwordless = false;
  ngOnInit(): void {}

  forgotPass() {
    console.log(this.emaill);
    this.authService.sendCode(this.emaill).subscribe();
    localStorage.setItem('emailForReset', this.emaill);
    this._router.navigate(['/resetPassword']);
  }
  onSubmit(f: NgForm) {
    const loginObserver = {
      next: (x: any) => {
        this._snackBar.open('     Welcome', 'Dismiss',{duration : 3000});
        if (localStorage.getItem('role') == 'Admin') {
          this._router.navigate(['/ahome']);
        } else {
          this._router.navigate(['/chome']);
        }
      },
      error: (err: any) => {
        this._snackBar.open(
          'Email or password are incorrect.Try again,please.',
          'Dismiss',{
            duration : 3000
           }
        );
      },
    };

    this.authService.login(f.value).subscribe(loginObserver);
  }

  check2FAStatus(){
    this.authService.check2FAStatus(this.emaill).subscribe(
      res =>{
        this.tfaEnabled = res
        console.log(this.tfaEnabled)
      }
    )
  }

  enablePL() {
    this.passwordless = true;
  }

  pswrdless(f2: NgForm ){
    this.authService.sendMagicLink(this.email).subscribe( 
      res => {
        this._snackBar.open('Check your email. We sent you a magic link to log-in to your account.', '', {
          duration: 3000,
        });
      },
      err => {
        this._snackBar.open('User with this username does not exist.', '', {
          duration: 3000,
        });
      }
    )
  }
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
