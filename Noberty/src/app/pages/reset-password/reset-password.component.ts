import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/UserService/user.service';

@Component({
  selector: 'app-reset-password',
  templateUrl: './reset-password.component.html',
  styleUrls: ['./reset-password.component.css']
})
export class ResetPasswordComponent implements OnInit {

  passMatch: boolean = false;
  createForm!: FormGroup;
  newPasswordDto!: string;
  divVisible: boolean = false;
  kodic!: string;

  constructor(
    private formBuilder: FormBuilder,
    private userService: UserService,
    private snackBar : MatSnackBar,
    private router: Router
  ) {}

  ngOnInit(): void {
   

    this.createForm = this.formBuilder.group({
      password: new FormControl(null, [
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(30),
        Validators.pattern(
          '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
        ),
      ]),
      passConfirmed: new FormControl(null, [
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(30),
        Validators.pattern(
          '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
        ),
      ]),
    });
  }

  verify() {
    this.userService.checkCode(this.kodic).subscribe(
      (res) => {
        this.divVisible = true;
      },
      err => {
        this.snackBar.open(err.error, '', {
          duration: 3000
        })
      });
  }
  onCodeInput(event: any): void {}

  onPasswordInput(): void {
    this.passMatch =
      this.createForm.value.password === this.createForm.value.passConfirmed;
  }

  onSubmit(): void {
    this.userService
      .resetPassword(this.createForm.value.password)
      .subscribe(
      (res) => {
    this.router.navigate(['/login']);
      });
  }

}
