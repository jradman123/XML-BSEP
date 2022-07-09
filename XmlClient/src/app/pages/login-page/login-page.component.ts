import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { ILoginRequest } from 'src/app/interfaces/login-request';
import { IUsername } from 'src/app/interfaces/username';
import { AuthService } from 'src/app/services/auth-service/auth.service';


@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {

  aFormGroup!: FormGroup;
  formPL! : FormGroup;
  passwordless = false;
  siteKey!: string;
  loginReq: ILoginRequest;
  usernameReq!: IUsername;
  username: string = "";
  showCode = false;

  constructor(
    private authService: AuthService,
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

    this.formPL =  this.formBuilder.group({
      username: new FormControl('', [
        Validators.required,
      ]),
      code : new FormControl('', 
      Validators.required)
    });
  }

  submitPL(){

    const plObserver = {
      next: () => {
        this._router.navigate(['myProfile']);
        this._snackBar.open(
          'Welcome!',
          '',
          {duration : 3000,panelClass: ['snack-bar']}
        );
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open("Error happend" + "!", '', {duration : 3000,panelClass: ['snack-bar']});
      }
    }

    const sendCodeObserver = {
      next: (x: any) => {
        this._snackBar.open("Code is sent to your mail!", '', {duration : 3000,panelClass: ['snack-bar']});
        this.showCode = true;
      },
      error: (err: HttpErrorResponse) => {

        this._snackBar.open("Error happend" + "!", '', {duration : 3000,panelClass: ['snack-bar']});
      },
    };

    if(!this.showCode)
      this.authService.passwordlessLoginRequest(this.formPL.value.username).subscribe(sendCodeObserver);
    else 
      this.authService.passwordlessLogin(this.formPL.value.code).subscribe(plObserver)
  }

  pswrdless(){
    this.passwordless = true;
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
        this._snackBar.open("Invalid credentials. Please try again.", "",  {
          duration: 3000,
          panelClass: ['snack-bar']
        });
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
        this._router.navigate(['/feed']);
        this._snackBar.open("Welcome!", "", {
          duration: 3000,
          panelClass: ['snack-bar']
        });
      },
      error: (err: HttpErrorResponse) => {

        this._snackBar.open("Invalid credentials. Please try again.", '', {
          duration: 3000,
          panelClass: ['snack-bar']
        });
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
