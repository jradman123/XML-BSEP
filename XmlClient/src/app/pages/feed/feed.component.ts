import { Component, OnInit } from '@angular/core';
import { PostService } from 'src/app/services/post-service/post.service';
import { UserService } from 'src/app/services/user-service/user.service';
import { IPost } from 'src/app/interfaces/post-request';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { PostCreateFileComponent } from 'src/app/components/post-create-file/post-create-file.component';
import { ConnectionService } from 'src/app/services/connection-service/connection.service';
import { UserDetails } from 'src/app/interfaces/user-details';

@Component({
  selector: 'app-feed',
  templateUrl: './feed.component.html',
  styleUrls: ['./feed.component.css']
})
export class FeedComponent implements OnInit {

  searchText : string = "";
  posts! : IPost[];
  username = localStorage.getItem('username');
  suggested! : UserDetails[];
  myInfo! : UserDetails;

  constructor(private _postService : PostService, private _userService : UserService, 
    private _matDialog : MatDialog, private _connectionService : ConnectionService) { }

  ngOnInit(): void {

    this._userService.getUserDetails(this.username!).subscribe( 
      res => {
        this.myInfo = res;
      }
    )

    this._postService.getUsersFeed(this.username!).subscribe(
      res => {
        this.posts = res.Feed
      }
    )

    this._connectionService.getUsersRecommendation(this.username!).subscribe(
      res => {
        this.suggested = res.users
        if(this.suggested.length > 5) this.suggested = this.suggested.slice(0, 5);
      }
    )
  }

  handleMe(searchText : string){
    this.searchText = searchText;
  }

  openCreatePost(){
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    const dialogRef = this._matDialog.open(PostCreateFileComponent, dialogConfig);
    dialogRef.afterClosed().subscribe({
      next: (res) => {
        console.log(res.data);
        this.posts.unshift(res.data);
      }
    })
  }
}
