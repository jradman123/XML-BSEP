import { HttpErrorResponse } from '@angular/common/http';
import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Subscription } from 'rxjs';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserPersonalDetails } from 'src/app/interfaces/user-personal-details';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-profile-edit',
  templateUrl: './profile-edit.component.html',
  styleUrls: ['./profile-edit.component.css']
})
export class ProfileEditComponent implements OnInit {

@Output() newEvent : EventEmitter<UserPersonalDetails> = new EventEmitter()
 sub!: Subscription;
  userDetails! : UserDetails;
  initialDetails: any;
  id!: number;
  date : any;
  initialDate : any;
  username! : string;
  email! : string;
  userPersonalDetails! : UserPersonalDetails;

  constructor(private userService : UserService,private _snackBar : MatSnackBar) {
     this.date = "";
     this.userPersonalDetails = {} as UserPersonalDetails;
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
    console.log(this.userPersonalDetails);
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
        this._snackBar.open("Error happend" + "!", '',{duration : 3000,panelClass: ['snack-bar']});
      }

    }
    this.userService.updateUserPersonalDetails(this.userPersonalDetails).subscribe(registerObserver)
    this.sendToParent(this.userPersonalDetails)

  }

  sendToParent(userPersonalDetails : UserPersonalDetails){
    console.log("usao u send to parent")
    this.newEvent.emit(userPersonalDetails)
  }
  createUserDetails(): void {
    this.userPersonalDetails.username = this.username;
    this.userPersonalDetails.phoneNumber = this.userDetailsForm.value.phoneNumber;
    this.userPersonalDetails.firstName = this.userDetailsForm.value.firstName;
    this.userPersonalDetails.lastName = this.userDetailsForm.value.lastName;
    this.userPersonalDetails.gender = this.userDetailsForm.value.gender;
    if(this.userDetailsForm.value.dateOfBirth == ""){
      this.userPersonalDetails.dateOfBirth = this.date;
    }else{
      const d = new Date( this.userDetailsForm.value.dateOfBirth.getTime() -  this.userDetailsForm.value.dateOfBirth.getTimezoneOffset() * 60000)
      this.userPersonalDetails.dateOfBirth = d.toISOString();
    }
    this.userPersonalDetails.biography = this.userDetailsForm.value.biography;
  }

}
