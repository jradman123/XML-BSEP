import { Component, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
import { NewPassword } from 'src/app/interfaces/new-password';

@Component({
  selector: 'app-reset-password',
  templateUrl: './reset-password.component.html',
  styleUrls: ['./reset-password.component.css'],
})
export class ResetPasswordComponent implements OnInit {
  passMatch: boolean = false;
  createForm!: FormGroup;
  newPasswordDto!: NewPassword;

  constructor(private formBuilder: FormBuilder) {}

  ngOnInit(): void {
    this.newPasswordDto = {} as NewPassword;

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

  onPasswordInput(): void {
    this.passMatch =
      this.createForm.value.password === this.createForm.value.passConfirmed;
  }

  onSubmit(): void {
    //this.newPasswordDto.email = localStorage.getItem('email');
    this.newPasswordDto.newPassword = this.createForm.value.password;
  }
}
