import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IPosts } from 'src/app/interfaces/post-request';
import { UserDetails } from 'src/app/interfaces/user-details';
import { PostService } from 'src/app/services/post-service/post.service';

@Component({
  selector: 'app-posts-view',
  templateUrl: './posts-view.component.html',
  styleUrls: ['./posts-view.component.css']
})
export class PostsViewComponent implements OnInit {
  
  @Input()
  username! : string;
  Posts: IPosts;

  
  constructor(
    private _service: PostService,
    private _snackBar: MatSnackBar,

  ) { 
    this.Posts = {} as IPosts
  }

  ngOnInit(): void {
    
    const getPostsObserver = {
      next: (res: IPosts) => {
        console.log(res);
  
        if (res.Posts.length == 0) {
          return
        }
  
        this.Posts = res
       
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', { duration: 3000 });
      },
    }
    console.log(this.username);
    this._service.GetAllPosts(this.username).subscribe(getPostsObserver);
  }
 
}

