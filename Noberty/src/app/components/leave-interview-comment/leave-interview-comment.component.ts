import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { IInterview } from 'src/app/interfaces/interview';
import { CompanyService } from 'src/app/services/company-service/company.service';

@Component({
  selector: 'app-leave-interview-comment',
  templateUrl: './leave-interview-comment.component.html',
  styleUrls: ['./leave-interview-comment.component.css']
})
export class LeaveInterviewCommentComponent implements OnInit {

  createForm!: FormGroup;
  interview!: IInterview;
  cid!: string;
  constructor(
    private _formBuilder: FormBuilder,
    private companyService: CompanyService,
    private router: Router,
    public dialogRef: MatDialogRef<LeaveInterviewCommentComponent>,
    private _snackBar: MatSnackBar
    ) { 
      this.interview = {} as IInterview
    this.createForm = this._formBuilder.group({
      Comment: new FormControl('',[Validators.required]),
      Difficulty : new FormControl('',[Validators.required]),
      Rating : new FormControl('',[Validators.required,Validators.min(1),Validators.max(5)])
    })
  }

  ngOnInit(): void {
    console.log(this.router.url);
    this.cid = this.router.url.substring(9);
  }

  submitRequest(): void {

    this.createInterviewRequest();
    console.log(this.interview);

    this.companyService.CreateInterview(this.interview).subscribe({
      next: (res) => {
        console.log(res);

        this.clearForm();
        this.dialogRef.close({ event: "Created interview comment", data: res });
        this._snackBar.open(
          'You have created a interview comment.',
          'Dismiss'
        );
      },
      error: (err: HttpErrorResponse) => {
        this.clearForm();
        this._snackBar.open(err.error.message + "!", 'Dismiss');
      },
      complete: () => console.info('complete')
    });
  }
  createInterviewRequest() {
    console.log(this.createForm.value.Name);

    this.interview.comment = this.createForm.value.Comment;
    this.interview.userUsername = localStorage.getItem("username")!;
    this.interview.companyID = parseInt(this.cid);
    this.interview.difficulty = this.createForm.value.Difficulty;
    this.interview.rating = this.createForm.value.Rating;
  }
  clearForm() {
    this.createForm.reset()
  }

}
