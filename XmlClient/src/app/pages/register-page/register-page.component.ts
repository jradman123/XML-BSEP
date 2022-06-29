import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../../services/user-service/user.service'
import { UserData } from '../../interfaces/subject-data'
import { MatSnackBar } from '@angular/material/snack-bar';
import { HttpErrorResponse } from '@angular/common/http';

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
    private userService: UserService) {
    this.newUser = {} as UserData;
  }

  ngOnInit(): void {
    this.createForm = this.formBuilder.group({
      Username: new FormControl('', [
        Validators.required,
        Validators.pattern('^[A-Za-z][A-Za-z0-9_]{7,29}$'),
      ]),
      Gender: new FormControl(''),
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
      DateOfBirth: new FormControl(null,
        [
          Validators.required,
        ]),
    });

  }

  onPasswordInput(): void {
    this.passMatch =
      this.createForm.value.Password === this.createForm.value.passConfirmed;
  }

  onSubmit(): void {

    this.createUser();
    console.log(this.newUser);
    const registerObserver = {
      next: () => {
       
        
        this.router.navigate(['/activate']);
        this._snackBar.open(
          'Your registration request has been sumbitted. Please check your email.',
          'Dismiss',
          {duration : 3000}
        );
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', {duration : 3000});
      }

    }
    this.userService.registerUser(this.newUser).subscribe(registerObserver)
  
  }

  createUser(): void {

    this.newUser.username = this.createForm.value.Username;
    this.newUser.password = this.createForm.value.Password;
    this.newUser.email = this.createForm.value.Email;
    this.newUser.recoveryEmail = this.createForm.value.RecoveryEmail;

    this.newUser.phoneNumber = this.createForm.value.PhoneNumber;
    this.newUser.firstName = this.createForm.value.FirstName;
    this.newUser.lastName = this.createForm.value.LastName;
    this.newUser.dateOfBirth = this.createForm.value.DateOfBirth;
    this.newUser.gender = this.createForm.value.Gender;
  }
}
