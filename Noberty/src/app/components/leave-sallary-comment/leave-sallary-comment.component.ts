import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { ISalaryComment } from 'src/app/interfaces/salary-comment';
import { CompanyService } from 'src/app/services/company-service/company.service';

@Component({
  selector: 'app-leave-sallary-comment',
  templateUrl: './leave-sallary-comment.component.html',
  styleUrls: ['./leave-sallary-comment.component.css']
})
export class LeaveSallaryCommentComponent implements OnInit {

  comment!: ISalaryComment;
  createForm!: FormGroup;
  cid!: string;
  constructor(
    private _formBuilder: FormBuilder,
    private companyService: CompanyService,
    private router: Router,
    public dialogRef: MatDialogRef<LeaveSallaryCommentComponent>,
    private _snackBar: MatSnackBar
  ) { 
    this.comment = {} as ISalaryComment
    this.createForm = this._formBuilder.group({
      Salary: new FormControl('',[
        Validators.required,
        Validators.pattern('^(?=.*[a-zA-Z])(?=.*[0-9]).*$'
        )
      ]),
      Position : new FormControl('',[Validators.required]),
    })
  }

  ngOnInit(): void {
    console.log(this.router.url);
    this.cid = this.router.url.substring(9);
  }

  submitRequest(): void {
    
    if (this.createForm.invalid)
        return;
        
    this.createCommentRequest();
    console.log(this.comment);

    this.companyService.CreateSalaryComment(this.comment).subscribe({
      next: (res) => {
        console.log(res);

        this.clearForm();
        this.dialogRef.close({ event: "Created salary comment", data: res });
        this._snackBar.open(
          'You have created a salary comment.',
          '', {
            duration: 3000
          });
      },
      error: (err: HttpErrorResponse) => {
        this.clearForm();
        this._snackBar.open(err.error.message + "!", '', {
          duration: 3000
        });
      },
      complete: () => console.info('complete')
    });
  }
  createCommentRequest() {

    this.comment.position = this.createForm.value.Position;
    this.comment.salary = this.createForm.value.Salary;
    this.comment.userUsername = localStorage.getItem("username")!;
    this.comment.companyID = parseInt(this.cid);
  }
  clearForm() {
    this.createForm.reset()
  }

}
