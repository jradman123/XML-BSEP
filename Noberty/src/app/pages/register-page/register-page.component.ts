import { DatePipe } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { NewUser } from 'src/app/interfaces/new-user';
import { UserServiceService } from 'src/app/services/UserService/user-service.service';

@Component({
  selector: 'app-register-page',
  templateUrl: './register-page.component.html',
  styleUrls: ['./register-page.component.css']
})
export class RegisterPageComponent implements OnInit {

  errorMessage!: string;
  createForm!: FormGroup;
  formData!: FormData;
  uploaded: boolean = false;
  passMatch: boolean = false;
  newUser!: NewUser;
  

  constructor(private formBuilder: FormBuilder,
    private _snackBar: MatSnackBar,
    private router: Router,
    private userService : UserServiceService,
    private datePipe: DatePipe
    ) { 
      this.newUser = {} as NewUser;
    }

  ngOnInit(): void {
    this.createForm = this.formBuilder.group({
      firstName: new FormControl('', [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),
      lastName: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),
      username: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[a-zA-Z0-9]([._-](?![._-])|[a-zA-Z0-9]){3,18}[a-zA-Z0-9]$'),
      ]),
      phoneNumber: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[0]{1}[0-9]{8}$'),
      ]),
      gender: new FormControl(null, [Validators.required]),
      dateOfBirth: new FormControl(null, [Validators.required]),
      email: new FormControl(null, [Validators.required, Validators.email]),
      recoveryEmail: new FormControl(null, [Validators.required, Validators.email]),
      password: new FormControl(null, [
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(30),
        Validators.pattern(
          '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
        )]),
      passConfirmed: new FormControl(null, [
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(30),
        Validators.pattern(
          '^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!"#$@%&()*<>+_|~]).*$'
        )]),
    });

  }

  onPasswordInput(): void {
    this.passMatch =
      this.createForm.value.password === this.createForm.value.passConfirmed;
  }

  onSubmit(): void {
      this.createUser();
      this.userService.registerUser(this.newUser).subscribe(
        (res) => {
          console.log(res);
          this.router.navigate(['/']);
          this._snackBar.open(
            'Your registration request has been sumbitted. Please check your email and confirm your email adress to activate your account.',
            'Dismiss'
          );
          
        },
        (err) => {
          let parts = err.error.split(':');
          let mess = parts[parts.length - 1];
          let description = mess.substring(1, mess.length - 4);
          this._snackBar.open(description, 'Dismiss');
        }
      );
  }

  createUser(): void {
    
      this.newUser.firstName = this.createForm.value.firstName;
      this.newUser.lastName = this.createForm.value.lastName;
      this.newUser.phoneNumber = this.createForm.value.phoneNumber;
      this.newUser.username = this.createForm.value.username;
      this.newUser.gender = this.createForm.value.gender;
      this.newUser.email = this.createForm.value.email;
      this.newUser.recoveryEmail = this.createForm.value.recoveryEmail;
      this.newUser.password = this.createForm.value.password;
      const date = this.datePipe.transform(this.createForm.value.dateOfBirth, 'MM/dd/yyyy');
      this.newUser.dateOfBirth = date;
  }
}


