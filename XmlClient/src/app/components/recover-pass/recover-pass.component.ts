import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../../services/user-service/user.service'
import { MatSnackBar } from '@angular/material/snack-bar';
import { HttpErrorResponse } from '@angular/common/http';
import { NewPass } from 'src/app/interfaces/new-pass';

@Component({
  selector: 'app-recover-pass',
  templateUrl: './recover-pass.component.html',
  styleUrls: ['./recover-pass.component.css']
})
export class RecoverPassComponent implements OnInit {

  enterNew = false;
  showCode = false;
  createForm!: FormGroup;
  formData!: FormData;
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

    if (!this.showCode && !this.enterNew){
      // send code request
      console.log(this.createForm.value.Username);
      
      this.authService.recoverPassRequest(this.createForm.value.Username).subscribe(
        res => this.showCode = true);
    }
    else if(!this.enterNew && this.showCode) {
        // verify code entered
        if ( this.createForm.value.Code) this.enterNew = true;
    }
    else if(this.enterNew && this.showCode) {
      console.log('udjem ovde')
      this.createRequest();
      
      const registerObserver = {
        next: () => {
          this.router.navigate(['/login']);
          this._snackBar.open(
            'Your password has been changed!',
            '',{duration : 3000,panelClass: ['snack-bar']}
          );
        },
        error: (err: HttpErrorResponse) => {
          this._snackBar.open("Error happend" + "!", '', {duration : 3000,panelClass: ['snack-bar']});
        }
  
      }

      this.authService.recoverPass(this.createForm.value.Username).subscribe(registerObserver)
    }

   
  }

  createRequest(): void {

    this.recoverPass.password = this.createForm.value.Password;
    this.recoverPass.username = this.createForm.value.Username;
    this.recoverPass.code = this.createForm.value.Code;
  }
}
