import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, NgForm, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-pass-less-req',
  templateUrl: './pass-less-req.component.html',
  styleUrls: ['./pass-less-req.component.css']
})
export class PassLessReqComponent implements OnInit {
  aFormGroup!: FormGroup;
  
  constructor(
    private authService: UserService,
    private _snackBar: MatSnackBar,
    private _router: Router,
    private formBuilder: FormBuilder
    ) { }

  ngOnInit(): void {
    this.aFormGroup = this.formBuilder.group({
    
      username: ['', Validators.required],
    
    });

  }

  onSubmit() {

    if (this.aFormGroup.invalid) {
      return;
    }
    console.log(this.aFormGroup.value.username);
    
    const loginObserver = {
      next: (x: any) => {
        console.log(x);
        this._router.navigate(['passwordlessLogin']);
        this._snackBar.open("Code is sent to your mail!", "Dismiss",{duration : 3000});
      },
      error: (err: HttpErrorResponse) => {
       
        this._snackBar.open(err.error + "!", 'Dismiss', {duration : 3000});
      },
    };
    this.authService.passwordlessLoginRequest(this.aFormGroup.value.username).subscribe(loginObserver);
  }

}
