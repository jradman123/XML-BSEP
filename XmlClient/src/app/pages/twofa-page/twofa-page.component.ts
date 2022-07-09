import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { IAuthenticate } from 'src/app/interfaces/authenticate';
import { AuthService } from 'src/app/services/auth-service/auth.service';

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
    private authService: AuthService,
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
        if (res == null) {
          this._snackBar.open("Invalid code!", '',
            {duration : 3000,panelClass: ['snack-bar']}
          );
          return
        }
        this._router.navigate(['/myProfile']);
        this._snackBar.open("Welcome!", '',{duration : 3000,panelClass: ['snack-bar']});
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open("Error happend!", '',{duration : 3000,panelClass: ['snack-bar']});
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
