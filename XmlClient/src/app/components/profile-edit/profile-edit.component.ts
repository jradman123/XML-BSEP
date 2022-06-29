import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Subscription } from 'rxjs';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-profile-edit',
  templateUrl: './profile-edit.component.html',
  styleUrls: ['./profile-edit.component.css']
})
export class ProfileEditComponent implements OnInit {


 sub!: Subscription;
  userDetails! : UserDetails;
  initialDetails: any;
  id!: number;
  date : any;
  initialDate : any;
  username! : string;
  email! : string;

  constructor(private userService : UserService,private _snackBar : MatSnackBar) {
     this.date = "";
   }

  ngOnInit(): void {
    this.getUserDetails();
  }
  getUserDetails() {
    this.sub = this.userService.getUserDetails(localStorage.getItem('username')).subscribe({
      next: (data: UserDetails) => {
        this.username = data.username;
        this.email = data.email;
        this.userDetails = data
        this.initialDetails = JSON.parse(JSON.stringify(data)); 
        this.userDetailsForm.controls['firstName'].setValue(data.firstName)
        this.userDetailsForm.controls['lastName'].setValue(data.lastName)
        this.userDetailsForm.controls['phoneNumber'].setValue(data.phoneNumber)
        this.userDetailsForm.controls['gender'].setValue(data.gender)
        this.userDetailsForm.controls['biography'].setValue(data.biography)
        this.date = new Date(data.dateOfBirth.substring(0,10));
        this.initialDate = this.date;
      },
    });
  }

  userDetailsForm = new FormGroup({
    firstName: new FormControl('', Validators.required),
    lastName: new FormControl('', Validators.required),
    phoneNumber: new FormControl('', Validators.required),
    dateOfBirth: new FormControl('', Validators.required),
    gender: new FormControl('', Validators.required),
    biography: new FormControl('', Validators.required),
  })

  cancel() {
    this.userDetailsForm.controls['firstName'].setValue(this.initialDetails.firstName)
    this.userDetailsForm.controls['lastName'].setValue(this.initialDetails.lastName)
    this.userDetailsForm.controls['phoneNumber'].setValue(this.initialDetails.phoneNumber)
    this.userDetailsForm.controls['dateOfBirth'].setValue(this.initialDate)
    this.userDetailsForm.controls['gender'].setValue(this.initialDetails.gender)
    this.userDetailsForm.controls['biography'].setValue(this.initialDetails.biography)
  }

  save() {
    this.createUserDetails();
    console.log(this.userDetails);
    const registerObserver = {
      next: () => {
        this._snackBar.open(
          'Success!.',
          'Dismiss',
          {
            duration : 3000
          }
        );
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss',{
          duration : 3000
        });
      }

    }
    this.userService.updateUser(this.userDetails).subscribe(registerObserver)

  }

  createUserDetails(): void {
    this.userDetails.username = this.username;
    this.userDetails.email = this.email;
    this.userDetails.phoneNumber = this.userDetailsForm.value.phoneNumber;
    this.userDetails.firstName = this.userDetailsForm.value.firstName;
    this.userDetails.lastName = this.userDetailsForm.value.lastName;
    this.userDetails.gender = this.userDetailsForm.value.gender;
    if(this.userDetailsForm.value.dateOfBirth == ""){
      this.userDetails.dateOfBirth = this.date;
    }else{
      const d = new Date( this.userDetailsForm.value.dateOfBirth.getTime() -  this.userDetailsForm.value.dateOfBirth.getTimezoneOffset() * 60000)
      this.userDetails.dateOfBirth = d.toISOString();
    }
    this.userDetails.biography = this.userDetailsForm.value.biography;
  }

}
