import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { CommentCreateComponent } from '../comment-create/comment-create.component';
@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent implements OnInit {
  isLiked: boolean = false
  createForm!: FormGroup;

 
  constructor(
    private _formBuilder: FormBuilder,
    public matDialog: MatDialog
  ) {
    this.createForm = this._formBuilder.group({
      comment: new FormControl('',Validators.required)
    })
   }

  ngOnInit(): void {
  }
  likePost() {
    this.isLiked = !this.isLiked
    console.log(this.isLiked);

  }
  leaveComment() {

    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = '300px';
    dialogConfig.width = '500px';
    const dialogRef = this.matDialog.open(CommentCreateComponent, dialogConfig);
    dialogRef.afterClosed().subscribe({
      next: () => {

      }
    })
  }
}
