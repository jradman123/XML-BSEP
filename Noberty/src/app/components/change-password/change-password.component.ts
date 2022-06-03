import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { ChangePasswordDto } from 'src/app/interfaces/change-password-dto';
import { UserServiceService } from 'src/app/services/UserService/user-service.service';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent implements OnInit {

  passMatch: boolean = false;
  changedPassword! : ChangePasswordDto;
  constructor(private userService : UserServiceService, private router : Router, private _snackBar : MatSnackBar,private dialogRef: MatDialogRef<ChangePasswordComponent>) {
    this.changedPassword = {} as ChangePasswordDto;
   }

  ngOnInit(): void {
  }

  form = new FormGroup({
    currentPassword:new FormControl(null, [
      Validators.required,
      Validators.minLength(10),
      Validators.maxLength(30),
      Validators.pattern(
        '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
      )]),
    newPassword: new FormControl(null, [
      Validators.required,
      Validators.minLength(10),
      Validators.maxLength(30),
      Validators.pattern(
        '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
      )]),
    repeatedPassword: new FormControl(null, [
      Validators.required,
      Validators.minLength(10),
      Validators.maxLength(30),
      Validators.pattern(
        '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
      )]),
  })

  onPasswordInput(): void {
    this.passMatch =
      this.form.value.newPassword === this.form.value.repeatedPassword;
  }

  createNewPassword(): void {
    this.changedPassword.newPassword = this.form.value.newPassword;
    this.changedPassword.currentPassword = this.form.value.currentPassword;
  }

  submit() {
    this.createNewPassword();
      this.userService.changePassword(this.changedPassword).subscribe(
        (res) => {
          this.logout();
          this._snackBar.open(
            'Password is changed!',
            'Dismiss' , {
              duration: 3000
            });
            
          this.logout();
          this.dialogRef.close();
        },
        (err) => {
          let parts = err.error.split(':');
          let mess = parts[parts.length - 1];
          let description = mess.substring(1, mess.length - 4);
          this._snackBar.open(description, 'Dismiss', {
            duration: 3000
          });
        }
      );
    this.userService.changePassword(this.form.value).subscribe();
    
  }

  logout() {
    this.userService.logout();
  }




}
