import { HttpErrorResponse } from '@angular/common/http';
import { Component, Inject, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { PostService } from 'src/app/services/post-service/post.service';
import {MAT_DIALOG_DATA} from '@angular/material/dialog';

@Component({
  selector: 'app-posts-comments-view',
  templateUrl: './posts-comments-view.component.html',
  styleUrls: ['./posts-comments-view.component.css']
})
export class PostsCommentsViewComponent implements OnInit {
  username:string=""
  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any,
    public _matDialog: MatDialog,
    private _snackBar: MatSnackBar,
    private service: PostService,
  ) { 
    this.username = localStorage.getItem('username')!
    console.log(data);
    
    const reactionsObserver = {
      next: (res: any) => {
       console.log(res);
       
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', { duration: 3000 });
      },
    }
    this.service.GetAllCommentsForPost(data).subscribe(reactionsObserver)
  }

  ngOnInit(): void {
  }

}
