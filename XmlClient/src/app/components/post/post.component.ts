import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IComment } from 'src/app/interfaces/comment';
import { IPost } from 'src/app/interfaces/post-request';
import { IReactions } from 'src/app/interfaces/reactions';
import { PostService } from 'src/app/services/post-service/post.service';
import { CommentCreateComponent } from '../comment-create/comment-create.component';
import { PostsCommentsViewComponent } from '../posts-comments-view/posts-comments-view.component';
@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent implements OnInit {
  @Input()
  item: IPost
  isLiked: boolean = false
  isDisliked: boolean = false
  createForm!: FormGroup;
  imageSrc!: any;
  username: string = ""
  reactions!: IReactions
  likesConst!: number
  dislikedConst!: number
  panelOpenState = false;

  constructor(
    private _formBuilder: FormBuilder,
    public _matDialog: MatDialog,
    private _snackBar: MatSnackBar,
    private service: PostService,

  ) {
    this.username = localStorage.getItem('username')!
    this.createForm = this._formBuilder.group({
      comment: new FormControl('', Validators.required)
    })
    this.item = {} as IPost
    this.getAllReactionsForPostAsync()
    this.getValueWithAsync()
  }
  async setImgPath() {
    this.imageSrc = 'data:image/jpeg;base64,' + this.item.ImagePaths
  }
  async getAllReactionsForPostAsync() {
    const value = <number>await this.resolveAfter2Seconds(10);
    console.log(`async result: ${value}`);
    this.getAllReactionsForPost()
  }
  async getValueWithAsync() {
    const value = <number>await this.resolveAfter2Seconds(5);
    console.log(`async result: ${value}`);
    this.setImgPath()
  }

  resolveAfter2Seconds(x: any) {
    return new Promise(resolve => {
      setTimeout(() => {
        resolve(x);
      }, 2000);
    });
  }
  ngOnInit(): void {
  }

  likePost() {
    this.isLiked = !this.isLiked
    this.isDisliked = false

    if (this.reactions.LikesNumber == this.likesConst) {
      this.reactions.LikesNumber = this.reactions.LikesNumber + 1
    } else {
      this.reactions.LikesNumber = this.reactions.LikesNumber - 1
    }

    const likeObserver = {
      next: () => {
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', { duration: 3000 });
      },
    }

    this.service.LikePost(this.username, this.item.Links.Like).subscribe(likeObserver)

  }
  dislikePost() {
    this.isDisliked = !this.isDisliked
    this.isLiked = false

    if (this.reactions.DislikesNumber == this.dislikedConst) {
      this.reactions.DislikesNumber = this.reactions.DislikesNumber + 1
    } else {
      this.reactions.DislikesNumber = this.reactions.DislikesNumber - 1
    }

    const dislikeObserver = {
      next: () => {
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', { duration: 3000 });
      },
    }

    this.service.DislikePost(this.username, this.item.Links.Dislike).subscribe(dislikeObserver)
  }

  leaveComment() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = '300px';
    dialogConfig.width = '500px';
    const dialogRef = this._matDialog.open(CommentCreateComponent, dialogConfig);
    dialogRef.afterClosed().subscribe({
      next: (res) => {
        console.log(res.data);

        const commentObserver = {
          next: () => {
          },
          error: (err: HttpErrorResponse) => {
            this._snackBar.open(err.error.message + "!", 'Dismiss', { duration: 3000 });
          },
        }
        this.service.CommentPost(res.data, this.item.Links.Comment).subscribe(commentObserver)
      }
    })

  }
  getAllReactionsForPost() {
    const reactionsObserver = {
      next: (res: any) => {
        this.reactions = res
        this.likesConst = res.LikesNumber
        this.dislikedConst = res.DislikesNumber
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', { duration: 3000 });
      },
    }

    this.service.GetAllReactionsForPost(this.item.Id).subscribe(reactionsObserver)
  }
  seeComments() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = '300px';
    dialogConfig.width = '500px';
    dialogConfig.data = this.item.Links.Comment
    
    const dialogRef = this._matDialog.open(PostsCommentsViewComponent, dialogConfig,);
  }
}

