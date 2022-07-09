import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IComment } from 'src/app/interfaces/comment';


@Component({
  selector: 'app-comment-create',
  templateUrl: './comment-create.component.html',
  styleUrls: ['./comment-create.component.css']
})
export class CommentCreateComponent implements OnInit {
  createForm!: FormGroup;
  newComment: IComment
  username!: string

  constructor(
    private _formBuilder: FormBuilder,
    public dialogRef: MatDialogRef<CommentCreateComponent>,
    private _snackBar: MatSnackBar
  ) {
    this.newComment = {} as IComment
    this.username = localStorage.getItem('username')!
    this.createForm = this._formBuilder.group({
      comment: new FormControl('', Validators.required)
    })
  }

  ngOnInit(): void {
  }
  submitRequest(): void {
    if(this.createForm.invalid){
      this._snackBar.open(
        'You cannot create a empty comment.',
        '', {
        duration: 3000
      });
      return;
    }
    this.newComment.Username = this.username
    this.newComment.CommentText = this.createForm.value.comment
    this.clearForm();
    this.dialogRef.close({ event: "Created comment", data: this.newComment });
    this._snackBar.open(
      'You have created a comment.',
      '', {
      duration: 3000
    });
  }
  
  clearForm() {
    this.createForm.reset()
  }
}


