import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IPosts } from 'src/app/interfaces/post-request';
import { UserDetails } from 'src/app/interfaces/user-details';
import { PostService } from 'src/app/services/post-service/post.service';
import { PostCreateFileComponent } from '../post-create-file/post-create-file.component';

@Component({
  selector: 'app-posts-view',
  templateUrl: './posts-view.component.html',
  styleUrls: ['./posts-view.component.css']
})
export class PostsViewComponent implements OnInit {
  
  @Input()
  username! : string;
  Posts: IPosts;
  showButton = false;

  
  constructor(
    private _service: PostService,
    private _snackBar: MatSnackBar,
    private _matDialog : MatDialog

  ) { 
    this.Posts = {} as IPosts
  }

  ngOnInit(): void {
    
    this.showButton = this.username === localStorage.getItem('username')

   this.getPosts();
  }

  getPosts(){
    const getPostsObserver = {
      next: (res: IPosts) => {
        console.log(res);
  
        if (res.Posts.length == 0) {
          return
        }
  
        this.Posts = res
       
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open("Error happend" + "!", '', {duration : 3000,panelClass: ['snack-bar']});
      },
    }
    console.log(this.username);
    this._service.GetAllPosts(this.username).subscribe(getPostsObserver);
  }
 
  openPostDialog(){
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    const dialogRef = this._matDialog.open(PostCreateFileComponent, dialogConfig);
    dialogRef.afterClosed().subscribe({
      next: (res) => {
        console.log(res.data);
        this.getPosts()
      }
    })
  }
}

