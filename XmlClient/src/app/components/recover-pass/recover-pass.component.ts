import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../../services/user-service/user.service'
import { MatSnackBar } from '@angular/material/snack-bar';
import { HttpErrorResponse } from '@angular/common/http';
import { NewPass } from 'src/app/interfaces/new-pass';
import { PwnedResponse } from 'src/app/interfaces/pwned-response';
import { Password } from 'src/app/interfaces/password';
@Component({
  selector: 'app-recover-pass',
  templateUrl: './recover-pass.component.html',
  styleUrls: ['./recover-pass.component.css']
})
export class RecoverPassComponent implements OnInit {

  createForm!: FormGroup;
  formData!: FormData;
  passMatch: boolean = false;
  recoverPass!: NewPass;
  alertMessage!: string;
  pwned!: number;
  passwordForCheck!: Password;
  constructor(
    private formBuilder: FormBuilder,
    private _snackBar: MatSnackBar,
    private router: Router,
    private authService: UserService
  ) { 
    this.recoverPass = {} as NewPass;
    this.alertMessage = ""
    this.pwned = 0
    this.passwordForCheck = {} as Password;
  }

  ngOnInit(): void {
    this.createForm = this.formBuilder.group({
      Username: new FormControl('', [
        Validators.required,
        Validators.pattern('^[A-Za-z][A-Za-z0-9_]{6,29}$'),
      ]),
      Code: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[0-9]*$'),
      ]),
       Password: new FormControl(null
        , [
          Validators.required,
          Validators.minLength(10),
          Validators.pattern(
            '^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{10,30}$'
          )]),
      passConfirmed: new FormControl(null,
        [
          Validators.required,
          Validators.minLength(10),
          Validators.pattern(
            '^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{10,30}$'

          )]),
    });
  }
  onPasswordInput(): void {
    this.passMatch =
      this.createForm.value.Password === this.createForm.value.passConfirmed;
  }

  onSubmit(): void {

    this.createRequest();
    console.log(this.recoverPass);
    const registerObserver = {
      next: () => {
       
        
        this.router.navigate(['/login']);
        this._snackBar.open(
          'Your password has been changed!',
          'Dismiss'
        );
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss');
      }

    }
    this.authService.recoverPass(this.recoverPass).subscribe(registerObserver)
  }

  checkPwned(){
    console.log(this.pwned);
    const registerObserver = {
      next: (pwnedResponse: PwnedResponse) => { 
        if(pwnedResponse.pwned) {
          this.pwned = 2;
          this.alertMessage = pwnedResponse.message;
        }else{
          this.pwned = 1;
          this.alertMessage = pwnedResponse.message;
        }
        this._snackBar.open(
          pwnedResponse.message,
          'Dismiss'
        );
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error + "!", 'Dismiss');
      }

    }
    this.passwordForCheck.password = this.createForm.value.Password;
    this.authService.passIsPwned(this.passwordForCheck).subscribe(registerObserver)
  }

  createRequest(): void {

    this.recoverPass.password = this.createForm.value.Password;
    this.recoverPass.username = this.createForm.value.Username;
    this.recoverPass.code = this.createForm.value.Code;
  }
}
