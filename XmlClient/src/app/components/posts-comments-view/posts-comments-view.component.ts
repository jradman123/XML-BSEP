import { HttpErrorResponse } from '@angular/common/http';
import { Component, Inject, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { PostService } from 'src/app/services/post-service/post.service';
import {MAT_DIALOG_DATA} from '@angular/material/dialog';
import { IComment } from 'src/app/interfaces/comment';

@Component({
  selector: 'app-posts-comments-view',
  templateUrl: './posts-comments-view.component.html',
  styleUrls: ['./posts-comments-view.component.css']
})
export class PostsCommentsViewComponent implements OnInit {
  username:string=""
  comments : IComment[] = []
  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any,
    public _matDialog: MatDialog,
    private _snackBar: MatSnackBar,
    private service: PostService,
  ) { 
    this.username = localStorage.getItem('username')!
  
    const reactionsObserver = {
      next: (res: any) => {
       console.log(res);
       this.comments = res.Comments;
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open("Error happend" + "!", '', {duration : 3000,panelClass: ['snack-bar']});
      },
    }
    this.service.GetAllCommentsForPost(data).subscribe(reactionsObserver)
  }

  ngOnInit(): void {
  }

}
