import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IPost } from 'src/app/interfaces/post-request';
import { IReactions } from 'src/app/interfaces/reactions';
import { IUserReaction } from 'src/app/interfaces/user-reaction';
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
  userReaction: IUserReaction

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
    this.userReaction = {} as IUserReaction
    this.getAllReactionsForPostAsync()
    this.getValueWithAsync()
    this.getUserReactionToPostAsync()
  }

  async getAllReactionsForPostAsync() {
    const value = <number>await this.resolveAfter2Seconds(10);
    console.log(`async result: ${value}`);
    this.getAllReactionsForPost()
  }
  async getUserReactionToPostAsync() {
    const value = <number>await this.resolveAfter2Seconds(10);
    console.log(`async result: ${value}`);
    this.getUserReactionToPost()
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
  setImgPath() {
    this.imageSrc = 'data:image/jpeg;base64,' + this.item.ImagePaths
  }
  getUserReactionToPost() {
    const reactionObserver = {
      next: (res: any) => {
        console.log(res);
        this.userReaction = res
        this.isLiked = res.Liked
        this.isDisliked = res.Disliked

      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', { duration: 3000 });
      },
    }

    if( this.username ) this.service.GetUserReactionToPost(this.username, this.item.Id).subscribe(reactionObserver)
  }
  likePost() {

    if (this.isLiked == true && this.isDisliked == false) {
      this.reactions.LikesNumber = this.reactions.LikesNumber - 1
      this.isLiked = !this.isLiked
    }
    else if (this.isLiked == false && this.isDisliked == false) {
      this.reactions.LikesNumber = this.reactions.LikesNumber + 1
      this.isLiked = !this.isLiked

    } else if (this.isLiked == false && this.isDisliked == true) {
      this.reactions.DislikesNumber = this.reactions.DislikesNumber - 1
      this.reactions.LikesNumber = this.reactions.LikesNumber + 1
      this.isLiked = !this.isLiked
      this.isDisliked = false
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

    if (this.isDisliked == true && this.isLiked == false) {
      this.reactions.DislikesNumber = this.reactions.DislikesNumber - 1
      this.isDisliked = !this.isDisliked
    }
    else if (this.isDisliked == false && this.isLiked == false) {
      this.reactions.DislikesNumber = this.reactions.DislikesNumber + 1
      this.isDisliked = !this.isDisliked
    }
    else if(this.isDisliked == false && this.isLiked == true){
      this.reactions.LikesNumber = this.reactions.LikesNumber - 1
      this.reactions.DislikesNumber = this.reactions.DislikesNumber + 1
      this.isDisliked = !this.isDisliked
      this.isLiked = false
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
    dialogConfig.height = '600px';
    dialogConfig.width = '900px';
    dialogConfig.data = this.item.Id

    this._matDialog.open(PostsCommentsViewComponent, dialogConfig);
  }
}

