import { HttpErrorResponse } from '@angular/common/http';
import { error } from '@angular/compiler/src/util';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ChangeEmailRequest } from 'src/app/interfaces/change-email-request';
import { ChangeUsernameRequest } from 'src/app/interfaces/change-username-request';
import { EmailRequest } from 'src/app/interfaces/email-request';
import { EmailUsernameResponse } from 'src/app/interfaces/email-username-response';
import { UsernameRequest } from 'src/app/interfaces/username-request';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-edit-email-username',
  templateUrl: './edit-email-username.component.html',
  styleUrls: ['./edit-email-username.component.css']
})
export class EditEmailUsernameComponent implements OnInit {

  changeUsernameRequest! : ChangeUsernameRequest;
  changeEmailRequest! : ChangeEmailRequest;
  constructor(private userService : UserService,private _snackBar : MatSnackBar) {
    this.changeUsernameRequest = {} as ChangeUsernameRequest
    this.changeUsernameRequest.username = {} as UsernameRequest
    this.changeEmailRequest = {} as ChangeEmailRequest
    this.changeEmailRequest.email = {} as EmailRequest
   }

  ngOnInit(): void {
    this.getEmailUsername();
  }

  getEmailUsername() {
    this.userService.getEmailUsername(localStorage.getItem('username')).subscribe({
      next: (data: EmailUsernameResponse) => {
        this.emailForm.controls['email'].setValue(data.emailUsername.email)
        this.usernameForm.controls['username'].setValue(data.emailUsername.username)
        localStorage.setItem('userId',data.userId)
      },
    });
  }

  changeEmail() {
    if(this.emailForm.invalid){
      return;
    }
    const changeEmailObserver = {
      next: () => {
        this._snackBar.open("Success" + "!", 'Dismiss', { duration: 3000 });
        localStorage.setItem('email' ,this.emailForm.value.email )
        this.getEmailUsername()
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open("Email already exists" + "!", 'Dismiss', { duration: 3000 });
      },
    }
    this.changeEmailRequest.userId = localStorage.getItem('userId')
    this.changeEmailRequest.email.email = this.emailForm.value.email
    this.userService.changeEmail(this.changeEmailRequest).subscribe(changeEmailObserver);

  }

  changeUsername() {
    if(this.usernameForm.invalid){
      return;
    }
    const changeUsernameObserver = {
      next: () => {
        this._snackBar.open("Success" + "!", 'Dismiss', { duration: 3000 });
        localStorage.setItem('username' ,this.usernameForm.value.username )
        this.getEmailUsername()
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open("Username already exists" + "!", 'Dismiss', { duration: 3000 });
      },
    }
    this.changeUsernameRequest.userId = localStorage.getItem('userId')
    this.changeUsernameRequest.username.username = this.usernameForm.value.username
    this.userService.changeUsername(this.changeUsernameRequest).subscribe(changeUsernameObserver);

  }

  emailForm = new FormGroup({
    email: new FormControl(null, [Validators.required, Validators.email]),
  })

  usernameForm = new FormGroup({
    username: new FormControl('', [
      Validators.required,
      Validators.pattern('^[A-Za-z][A-Za-z0-9_]{7,29}$'),
    ]),
  })

}
