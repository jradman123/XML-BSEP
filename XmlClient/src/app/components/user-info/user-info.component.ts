import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { EducationDto } from 'src/app/interfaces/education-dto';
import { ExperienceDto } from 'src/app/interfaces/experience-dto';
import { InterestDto } from 'src/app/interfaces/interest-dto';
import { SkillDto } from 'src/app/interfaces/skill-dto';
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
  id!: number;
  username! : string;
  email! : string;
  constructor(private dialogRef: MatDialogRef<UserInfoComponent>,private userService : UserService,public dialog: MatDialog,
    private _snackBar : MatSnackBar, private _router: Router) { 
    
  }

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
    this.userDetailsForm.controls['firstName'].setValue(this.initialDetails.firstName)
    this.userDetailsForm.controls['lastName'].setValue(this.initialDetails.lastName)
    this.userDetailsForm.controls['email'].setValue(this.initialDetails.email)
    this.userDetailsForm.controls['phoneNumber'].setValue(this.initialDetails.phoneNumber)
    this.userDetailsForm.controls['username'].setValue(this.initialDetails.username)
    this.userDetailsForm.controls['dateOfBirth'].setValue(this.initialDetails.dateOfBirth)
    this.userDetailsForm.controls['gender'].setValue(this.initialDetails.gender)
    this.userDetailsForm.controls['biography'].setValue(this.initialDetails.biography)
    this.dialogRef.close();
  }

  getUserDetails() {
    this.sub = this.userService.getUserDetails(localStorage.getItem('username')).subscribe({
      next: (data: UserDetails) => {
        this.userDetails = data
        this.username = data.username;
        this.email = data.email;
        this.initialDetails = JSON.parse(JSON.stringify(data)); 
        this.userDetailsForm.controls['firstName'].setValue(data.firstName)
        this.userDetailsForm.controls['lastName'].setValue(data.lastName)
        this.userDetailsForm.controls['phoneNumber'].setValue(data.phoneNumber)
        this.userDetailsForm.controls['gender'].setValue(data.gender)
        this.userDetailsForm.controls['biography'].setValue(data.biography)
        this.userDetailsForm.controls['dateOfBirth'].setValue(data.dateOfBirth)
      },
    });

  }

  save() {
    this.createUserDetails();
    console.log(this.userDetails);
    const registerObserver = {
      next: () => {
       
        
        this._router.navigate(['/']);
        this._snackBar.open(
          'Your registration request has been sumbitted. Please check your email and confirm your email adress to activate your account.',
          'Dismiss'
        );
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss');
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
    this.userDetails.dateOfBirth = this.userDetailsForm.value.dateOfBirth;
    this.userDetails.biography = this.userDetailsForm.value.biography;
    this.userDetails.skills = [] as SkillDto[];
    this.userDetails.educations = [] as EducationDto[];
    this.userDetails.experiences = [] as ExperienceDto[];
    this.userDetails.interests = [] as InterestDto[];
  }
}

 
