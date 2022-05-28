import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {

  

  constructor(
    private authService: UserService,
    private _snackBar: MatSnackBar,
    private _router: Router) { }

  ngOnInit(): void {
  }

  onSubmit(f: NgForm) {

    const loginObserver = {
      next: (x: any) => {
        console.log(x);
        this._router.navigate(['/']);
        this._snackBar.open("Welcome!", "Dismiss");
      },
      error: (err: HttpErrorResponse) => {
       
        this._snackBar.open(err.error.message + "!", 'Dismiss');
      },
    };
    this.authService.login(f.value).subscribe(loginObserver);
  }

}

