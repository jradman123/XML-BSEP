import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { Subscription } from 'rxjs';
import { UserDetails } from 'src/app/interfaces/user-details';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-user-info',
  templateUrl: './user-info.component.html',
  styleUrls: ['./user-info.component.css']
})
export class UserInfoComponent implements OnInit {

  sub!: Subscription;
  userDetails! : UserDetails;
  initialDetails: any;
  editMode = false;
  id!: number;
  constructor(private userService : UserService,public dialog: MatDialog) { }

  ngOnInit(): void {
    this.getUserDetails();
  }

  userDetailsForm = new FormGroup({
    firstName: new FormControl('', Validators.required),
    lastName: new FormControl('', Validators.required),
    email: new FormControl('',Validators.required),
    phoneNumber: new FormControl('', Validators.required),
    username: new FormControl('', Validators.required),
    dateOfBirth: new FormControl('', Validators.required),
    gender: new FormControl('', Validators.required),
    biography: new FormControl('', Validators.required)
  })

  cancel() {
    this.editMode = false
    this.userDetailsForm.controls['firstName'].setValue(this.initialDetails.firstName)
    this.userDetailsForm.controls['lastName'].setValue(this.initialDetails.lastName)
    this.userDetailsForm.controls['email'].setValue(this.initialDetails.email)
    this.userDetailsForm.controls['phoneNumber'].setValue(this.initialDetails.phoneNumber)
    this.userDetailsForm.controls['username'].setValue(this.initialDetails.username)
    this.userDetailsForm.controls['dateOfBirth'].setValue(this.initialDetails.dateOfBirth)
    this.userDetailsForm.controls['gender'].setValue(this.initialDetails.gender)
    this.userDetailsForm.controls['biography'].setValue(this.initialDetails.biography)
  }

  edit() {}

  changePassword(){}

  getUserDetails() {
    this.sub = this.userService.getUserDetails(localStorage.getItem('username')).subscribe({
      next: (data: UserDetails) => {
        this.userDetails = data
        this.initialDetails = JSON.parse(JSON.stringify(data)); 
        this.userDetailsForm.controls['firstName'].setValue(data.firstName)
        this.userDetailsForm.controls['lastName'].setValue(data.lastName)
        this.userDetailsForm.controls['email'].setValue(data.email)
        this.userDetailsForm.controls['phoneNumber'].setValue(data.phoneNumber)
        this.userDetailsForm.controls['username'].setValue(data.username)
        this.userDetailsForm.controls['gender'].setValue(data.gender)
        this.userDetailsForm.controls['biography'].setValue(data.biography)
        this.userDetailsForm.controls['dateOfBirth'].setValue(data.dateOfBirth)
      },
    });

  }

}
