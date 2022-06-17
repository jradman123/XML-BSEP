import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { IAuthenticate } from 'src/app/interfaces/authenticate';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-twofa-page',
  templateUrl: './twofa-page.component.html',
  styleUrls: ['./twofa-page.component.css']
})
export class TwofaPageComponent implements OnInit {
  createForm!: FormGroup;
  request!: IAuthenticate;
  username: string = ""

  constructor(
    private _snackBar: MatSnackBar,
    private authService: UserService,
    private _router: Router,
    private formBuilder: FormBuilder) { }

  ngOnInit(): void {
    this.createForm = this.formBuilder.group({
      code: new FormControl('', [
        Validators.required,
        Validators.pattern('^[0-9]{6}$'),
      ]),
    });
  }
  onSbmit() {

    if (this.createForm.invalid) {
      return;
    }
    const loginObserver = {
      next: (res: any) => {
        if (res == null){
          this._snackBar.open("Invalid code!", 'Dismiss');
          return
        }
        this._router.navigate(['/editUser']);
        this._snackBar.open("Welcome!", "Dismiss");
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error, 'Dismiss');
      },
    };
    this.username = localStorage.getItem('username')!
    if (this.username != null) {
      this.request = {
        username: this.username,
        token: this.createForm.value.code
      }

      this.authService.authenticate2FA(this.request).subscribe(loginObserver);
    }
  }
}
