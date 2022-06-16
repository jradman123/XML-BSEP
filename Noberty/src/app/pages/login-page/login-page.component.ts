import { Component, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/UserService/user.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css'],
})
export class LoginPageComponent implements OnInit {
  form!: FormGroup;
  formPL! : FormGroup;

  usernamee!: string;
  passwordless = false;
  tfaEnabled = false;
  constructor(
    private formBuilder: FormBuilder,
    private _router: Router,
    private _userService: UserService,
    private _snackBar: MatSnackBar
  ) {}

  ngOnInit(): void {
    
    this.form = this.formBuilder.group({
      username: new FormControl('', [
        Validators.required,
        Validators.pattern(
          '^[a-zA-Z0-9]([._-](?![._-])|[a-zA-Z0-9]){3,18}[a-zA-Z0-9]$'
        ),
      ]),
      password: new FormControl('', [
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(30),
        Validators.pattern(
          '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
        ),
      ]),
      code: new FormControl('', []),
    });

    this.formPL =  this.formBuilder.group({
      username: new FormControl('', [
        Validators.required,
        Validators.pattern(
          '^[a-zA-Z0-9]([._-](?![._-])|[a-zA-Z0-9]){3,18}[a-zA-Z0-9]$'
        ),
      ]),
    });
  }

  submitPL() {
    this._userService.sendMagicLink(this.formPL.value.username).subscribe(
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
    );
  }

  enablePL() {
    this.passwordless = true;
  }

  check2FAStatus() {
    this._userService.check2FAStatus(this.form.value.username).subscribe(
      res => this.tfaEnabled = res
    )
  }

  forgotPass() {
    if (
      this.form.value.username == '' ||
      this.form.value.username == undefined
    ) {
      this._snackBar.open('Please enter your username.', '', {
        duration: 3000,
      });
      return;
    }
    this._userService.sendCode(this.usernamee).subscribe(
      (res) => {
        localStorage.setItem('usernamee', this.usernamee);
        this._router.navigate(['/resetPassword']);
      },
      (err) => {
        this._snackBar.open('User with this usename does not exist!', '', {
          duration: 3000,
        });
      }
    );
  }

  submit(): void {
    if (this.form.invalid) return;

    const loginObserver = {
      next: (x: any) => {
        this._snackBar.open('Welcome!', '', {
          duration: 3000,
        });

        this._router.navigate(['/user/landing']);
      },
      error: (err: any) => {
        this._snackBar.open(
          'Username or password are incorrect.Try again,please.',
          '',
          {
            duration: 3000,
          }
        );
      },
    };

    this._userService.login(this.form.getRawValue()).subscribe(loginObserver);
  }
}
