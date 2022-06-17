import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { ILoginRequest } from 'src/app/interfaces/login-request';
import { IUsername } from 'src/app/interfaces/username';
import { UserService } from 'src/app/services/user-service/user.service';


@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {

  aFormGroup!: FormGroup;
  siteKey!: string;
  loginReq: ILoginRequest;
  usernameReq!: IUsername;
  username: string = ""

  constructor(
    private authService: UserService,
    private _snackBar: MatSnackBar,
    private _router: Router,
    private formBuilder: FormBuilder) {
    this.siteKey = "6Lcht3AgAAAAABAzzhnpricVg4eDjZmX-HBFbm6u"
    this.loginReq = {} as ILoginRequest
    //secret key
    //6Lcht3AgAAAAAJXjRp64PpHTEaT5G8Ygp-sMO0Pi
  }

  ngOnInit(): void {

    this.aFormGroup = this.formBuilder.group({
      recaptcha: ['', Validators.required],
      username: ['', Validators.required],
      password: ['', Validators.required]
    });

  }

  onSubmit() {
    if (this.aFormGroup.invalid) {
      return;
    }

    const loginObserver = {
      next: (res: any) => {
        if (res.twofa == true) {
          this.twoFaLogin()

        } else {
          this.regularLogin()
        }
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error, 'Dismiss');
      },
    };
    this.loginReq = {
      username: this.aFormGroup.value.username,
      password: this.aFormGroup.value.password
    }

    this.authService.auth(this.loginReq).subscribe(loginObserver);
  }
  twoFaLogin() {
    this._router.navigate(['/twofa']);
  }
  regularLogin() {
    const nextLoginOserver = {
      next: (res: any) => {
        if (res == null) {
          return
        }
        this._router.navigate(['/editUser']);
        this._snackBar.open("Welcome!", "Dismiss");
      },
      error: (err: HttpErrorResponse) => {

        this._snackBar.open(err.error, 'Dismiss');
      },
    }

    this.username = localStorage.getItem('username')!
    if (this.username != null) {
      this.usernameReq = {
        username: this.username
      }
      this.authService.login(this.loginReq).subscribe(nextLoginOserver);
    }
  }

}
