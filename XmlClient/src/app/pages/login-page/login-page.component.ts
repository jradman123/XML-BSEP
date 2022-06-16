import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { ILoginRequest } from 'src/app/interfaces/login-request';
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
      next: (x: any) => {
        this._router.navigate(['/']);
        this._snackBar.open("Welcome!", "Dismiss");
       
      },
      error: (err: HttpErrorResponse) => {
      
        this._snackBar.open(err.error, 'Dismiss');
      },
    };
    this.loginReq = {
      username: this.aFormGroup.value.username,
      password: this.aFormGroup.value.password
    }
    
    this.authService.login(this.loginReq).subscribe(loginObserver);
  }

}

