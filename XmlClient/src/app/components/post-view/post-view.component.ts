import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { IPost } from 'src/app/interfaces/post-request';
import { PostService } from 'src/app/services/post-service/post.service';

@Component({
  selector: 'app-post-view',
  templateUrl: './post-view.component.html',
  styleUrls: ['./post-view.component.css']
})
export class PostViewComponent implements OnInit {

  searchText : string = "";
  post! : IPost;
  postId : string;
  constructor(private _router : Router, private _service : PostService) { 
    this.postId = this._router.url.substring(6) ?? ''
  }

  ngOnInit(): void {
    this._service.GetPost(this.postId).subscribe(
      res => {
        console.log(res);
        this.post = res.Post
      }
    );
  }

  handleMe(searchText : string){
    this.searchText = searchText;
  }

}
