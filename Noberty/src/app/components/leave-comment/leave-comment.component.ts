import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { IComment } from 'src/app/interfaces/comment';
import { CompanyService } from 'src/app/services/company-service/company.service';

@Component({
  selector: 'app-leave-comment',
  templateUrl: './leave-comment.component.html',
  styleUrls: ['./leave-comment.component.css']
})
export class LeaveCommentComponent implements OnInit {

  comment!: IComment;
  createForm!: FormGroup;
  cid!: string;
  constructor(
    private _formBuilder: FormBuilder,
    private companyService: CompanyService,
    private router: Router,
    public dialogRef: MatDialogRef<LeaveCommentComponent>,
    private _snackBar: MatSnackBar) { 
      
      this.comment = {} as IComment
      this.createForm = this._formBuilder.group({
      Comment: new FormControl('',Validators.required)
    })
  }

  ngOnInit(): void {
    console.log(this.router.url);
    this.cid = this.router.url.substring(9);
  }


  submitRequest(): void {

    this.createCommentRequest();
    console.log(this.comment);

    this.companyService.CreateComment(this.comment).subscribe({
      next: (res) => {
        console.log(res);

        this.clearForm();
        this.dialogRef.close({ event: "Created comment", data: res });
        this._snackBar.open(
          'You have created a comment.',
          'Dismiss', {
            duration: 3000
          });
      },
      error: (err: HttpErrorResponse) => {
        this.clearForm();
        this._snackBar.open(err.error.message + "!", 'Dismiss', {
          duration: 3000
        });
      },
      complete: () => console.info('complete')
    });
  }
  createCommentRequest() {
    console.log(this.createForm.value.Name);

    this.comment.comment = this.createForm.value.Comment;
    this.comment.userUsername = localStorage.getItem("username")!;
    this.comment.companyId = parseInt(this.cid);
  }
  clearForm() {
    this.createForm.reset()
  }


}
