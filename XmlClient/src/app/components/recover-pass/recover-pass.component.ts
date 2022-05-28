import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../../services/user-service/user.service'
import { UserData } from '../../interfaces/subject-data'
import { MatSnackBar } from '@angular/material/snack-bar';
import { HttpErrorResponse } from '@angular/common/http';
import { NewPass } from 'src/app/interfaces/new-pass';

@Component({
  selector: 'app-recover-pass',
  templateUrl: './recover-pass.component.html',
  styleUrls: ['./recover-pass.component.css']
})
export class RecoverPassComponent implements OnInit {

  errorMessage!: string;
  createForm!: FormGroup;
  formData!: FormData;
  uploaded: boolean = false;
  passMatch: boolean = false;
  recoverPass!: NewPass;
  constructor(
    private formBuilder: FormBuilder,
    private _snackBar: MatSnackBar,
    private router: Router,
    private authService: UserService
  ) { 
    this.recoverPass = {} as NewPass;
  }

  ngOnInit(): void {
    this.createForm = this.formBuilder.group({
      Username: new FormControl('', [
        Validators.required,
        Validators.pattern('^[A-Za-z][A-Za-z0-9_]{7,29}$'),
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
       
        
        this.router.navigate(['/']);
        this._snackBar.open(
          'Your registration request has been sumbitted. Please check your email and confirm your email adress to activate your account.',
          'Dismiss'
        );
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss');
      }

    }
    //this.authService.registerUser(this.newUser).subscribe(registerObserver)
    this.router.navigate(['http://localhost:4200/login']);
  }

  createRequest(): void {

    this.recoverPass.password = this.createForm.value.Password;
    this.recoverPass.username = this.createForm.value.Username;
    this.recoverPass.code = this.createForm.value.Code;
  }
}
