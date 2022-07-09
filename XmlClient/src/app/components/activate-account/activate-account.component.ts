import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../../services/user-service/user.service'
import { MatSnackBar } from '@angular/material/snack-bar';
import { HttpErrorResponse } from '@angular/common/http';
import { ActivateAccount } from 'src/app/interfaces/activate-account';

@Component({
  selector: 'app-activate-account',
  templateUrl: './activate-account.component.html',
  styleUrls: ['./activate-account.component.css']
})
export class ActivateAccountComponent implements OnInit {

  errorMessage!: string;
  createForm!: FormGroup;
  formData!: FormData;
  activateAccount!: ActivateAccount;
  constructor(
    private formBuilder: FormBuilder,
    private _snackBar: MatSnackBar,
    private router: Router,
    private authService: UserService
  ) {
    this.activateAccount = {}  as ActivateAccount;
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
    });
  }

  onSubmit(): void {
    if(this.createForm.invalid){
      this._snackBar.open("Fields cannot be empty" + "!", '',{duration : 3000,panelClass: ['snack-bar']});
      return;
    }
    this.createRequest();
    console.log(this.activateAccount);
    const registerObserver = {
      next: () => {
       
        
        this.router.navigate(['/login']);
        this._snackBar.open(
          'Your account has been activated!',
          '',{
            duration : 3000,
            panelClass: ['snack-bar ']
          }
        );
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", '',{duration : 3000,panelClass: ['snack-bar']});
      }

    }
    this.authService.activateAccount(this.activateAccount).subscribe(registerObserver)
  }

  createRequest(): void {
    this.activateAccount.username = this.createForm.value.Username;
    this.activateAccount.code = this.createForm.value.Code;
  }

}
