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
  skills! : SkillDto[];
  educations! : EducationDto[];
  interests! : InterestDto[];
  experiences! : ExperienceDto[];
  newSkill! : SkillDto;
  newExperience! : ExperienceDto;
  newEducation! : EducationDto;
  newInterest! : InterestDto;
  date! : any;
  constructor(private dialogRef: MatDialogRef<UserInfoComponent>,private userService : UserService,public dialog: MatDialog,
    private _snackBar : MatSnackBar, private _router: Router) { 
      this.skills = [] as SkillDto[];
      this.educations = [] as EducationDto[];
      this.experiences = [] as ExperienceDto[];
      this.interests = [] as InterestDto[];
      this.newSkill = {} as SkillDto;
      this.newEducation = {} as EducationDto;
      this.newExperience = {} as ExperienceDto;
      this.newInterest = {} as InterestDto;
    
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
    biography: new FormControl('', Validators.required),
    skill : new FormControl('',null),
    experience : new FormControl('',null),
    interest : new FormControl('',null),
    education : new FormControl('',null)
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
        this.skills = data.skills;
        this.interests = data.interests;
        this.experiences = data.experiences;
        this.educations = data.educations;
        this.userDetails = data
        this.username = data.username;
        this.email = data.email;
        this.initialDetails = JSON.parse(JSON.stringify(data)); 
        this.userDetailsForm.controls['firstName'].setValue(data.firstName)
        this.userDetailsForm.controls['lastName'].setValue(data.lastName)
        this.userDetailsForm.controls['phoneNumber'].setValue(data.phoneNumber)
        this.userDetailsForm.controls['gender'].setValue(data.gender)
        this.userDetailsForm.controls['biography'].setValue(data.biography)
        this.date = new Date(data.dateOfBirth.substring(0,10));     
      },
    });

  }

  save() {
    this.createUserDetails();
    console.log(this.userDetails);
    const registerObserver = {
      next: () => {
       
        
        this.dialogRef.close();
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
    this.userDetails.dateOfBirth = this.userDetailsForm.value.dateOfBirth;
    this.userDetails.biography = this.userDetailsForm.value.biography;
    this.newSkill.skill = this.userDetailsForm.value.skill;
    this.skills.push(this.newSkill);
    this.userDetails.skills = this.skills;
    this.newEducation.education = this.userDetailsForm.value.education;
    this.educations.push(this.newEducation);
    this.userDetails.educations =this.educations;
    this.newExperience.experience = this.userDetailsForm.value.experience;
    this.experiences.push(this.newExperience);
    this.userDetails.experiences = this.experiences;
    this.newInterest.interest = this.userDetailsForm.value.interest;
    this.interests.push(this.newInterest);
    this.userDetails.interests = this.interests;
  }

}

 
