import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { UserServiceService } from 'src/app/services/UserService/user-service.service';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {

  public form!: FormGroup;
  usernamee!: string;
  constructor( private formBuilder: FormBuilder,
    private _router: Router,
    private _userService: UserServiceService,
    private _snackBar: MatSnackBar) { }

    ngOnInit(): void {
      this.form = this.formBuilder.group({
        username:'',
        password:''
      });
    }

    forgotPass() {
      this._userService.sendCode(this.usernamee).subscribe();
      localStorage.setItem('usernamee', this.usernamee);
      this._router.navigate(['/resetPassword']);
  }

      submit():void{
        const loginObserver = {
          next: (x:any) => {
             this._snackBar.open("     Welcome","Dismiss");
             
                this._router.navigate(['/user/landing']);
          },
           error: (err:any) => {
             this._snackBar.open("Username or password are incorrect.Try again,please.","Dismiss"); 
           
           }};
        
        this._userService.login(this.form.getRawValue()).subscribe(loginObserver);
       }
   
    }


