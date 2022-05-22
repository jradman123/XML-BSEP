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
      Username: new FormControl('', [
        Validators.required,
        Validators.pattern('^[A-Za-z][A-Za-z0-9_]{7,29}$'),
      ]),
      PhoneNumber: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$'),
      ]),
      FirstName: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),
      LastName: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[A-ZŠĐŽČĆ][a-zšđćčžA-ZŠĐŽČĆ ]*$'),
      ]),
      Email: new FormControl(null, [Validators.required, Validators.email]),
      RecoveryEmail: new FormControl(null, [Validators.required, Validators.email]),
       Password: new FormControl(null
      , [
        Validators.required,
        Validators.pattern(
          '^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{10,30}$'
        )]),
       passConfirmed: new FormControl(null, 
        [
          Validators.required,
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
    
      this.newUser.Username = this.createForm.value.Username;
      this.newUser.Password = this.createForm.value.Password;
      this.newUser.Email = this.createForm.value.Email;
      this.newUser.RecoveryEmail = this.createForm.value.RecoveryEmail;
 
      this.newUser.PhoneNumber = this.createForm.value.PhoneNumber;
      this.newUser.FirstName = this.createForm.value.FirstName;
      this.newUser.LastName = this.createForm.value.LastName;
      // this.newUser.DateOfBirth = this.createForm.value.DateOfBirth;
      // this.newUser.Gender = this.createForm.value.Gender;
  }
}
