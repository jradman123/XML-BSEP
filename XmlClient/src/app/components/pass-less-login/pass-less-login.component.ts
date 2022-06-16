import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-pass-less-login',
  templateUrl: './pass-less-login.component.html',
  styleUrls: ['./pass-less-login.component.css']
})
export class PassLessLoginComponent implements OnInit {

  errorMessage!: string;
  createForm!: FormGroup;
  formData!: FormData;
  constructor(
    private formBuilder: FormBuilder,
    private _snackBar: MatSnackBar,
    private router: Router,
    private authService: UserService
  ) {
   }

  ngOnInit(): void {
    this.createForm = this.formBuilder.group({
      Code: new FormControl(null, [
        Validators.required,
        Validators.pattern('^[a-zA-Z0-9]*$'),
      ]),
    });
  }

  onSubmit(): void {

    const registerObserver = {
      next: () => {
       
        
        this.router.navigate(['jobOffers']);
        this._snackBar.open(
          'Welcome!',
          'Dismiss'
        );
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error + "!", 'Dismiss');
      }

    }
    this.authService.passwordlessLogin( this.createForm.value.Code).subscribe(registerObserver)
  }

  

}
