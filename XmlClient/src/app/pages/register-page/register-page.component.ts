import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import {UserService} from '../../services/user-service/user.service'
import {UserData} from '../../interfaces/subject-data'
import { MatSnackBar } from '@angular/material/snack-bar';

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
  newUser!: UserData;

  constructor(private formBuilder: FormBuilder,
    private _snackBar: MatSnackBar,
    private router: Router,
    private authService : UserService) { 
      this.newUser = {} as UserData;
    }

  ngOnInit(): void {
    this.createForm = this.formBuilder.group({
      commonName: new FormControl('', [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),
      organization: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),
      organizationUnit: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),
      locality: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),
      country: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),
      email: new FormControl(null, [Validators.required, Validators.email]),
      recoveryMail: new FormControl(null, [Validators.required, Validators.email]),
      password: new FormControl(null, [
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(30),
        Validators.pattern(
          '^(?=.[a-z])(?=.[A-Z])(?=.[0-9])(?=.[!"#$@%&()<>+_|~]).$'
        )]),
      passConfirmed: new FormControl(null, [
        Validators.required,
        Validators.minLength(10),
        Validators.maxLength(30),
        Validators.pattern(
          '^(?=.[a-z])(?=.[A-Z])(?=.[0-9])(?=.[!"#$@%&()<>+_|~]).$'
        )]),
    });

  }

  onPasswordInput(): void {
    this.passMatch =
      this.createForm.value.password === this.createForm.value.passConfirmed;
  }

  onSubmit(): void {
      this.createUser();
      this.authService.registerUser(this.newUser).subscribe(
        (res) => {
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
    
      this.newUser.commonName = this.createForm.value.commonName;
      this.newUser.organization = this.createForm.value.organization;
      this.newUser.organizationUnit = this.createForm.value.organizationUnit;
      this.newUser.locality = this.createForm.value.locality;
      this.newUser.country = this.createForm.value.country;
      this.newUser.email = this.createForm.value.email;
      this.newUser.recoveryMail = this.createForm.value.recoveryMail;
      this.newUser.password = this.createForm.value.password;
  }
}
