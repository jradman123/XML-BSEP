import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { IPostRequest } from 'src/app/interfaces/post-request';

@Injectable({
  providedIn: 'root'
})
export class PostService {


  constructor(private _http: HttpClient,) {
  }
  GetAllPosts(username : string) {
    return this._http.get<any>(
      'http://localhost:9000/post/user/' + username,
    );
  }
  GetUserReactionToPost(username: string, Id: string) {
    return this._http.get<any>(
      'http://localhost:9000/post/' + Id + "/" + username + '/reaction'
    );
  }
  GetAllReactionsForPost(Id: string) {
    return this._http.get<any>(
      'http://localhost:9000/post/' + Id + '/reactions'
    );
  }
  GetAllCommentsForPost(Id: any) {
    return this._http.get<any>(
      'http://localhost:9000/post/' + Id + '/comments'
    );
  }

  CreatePost(newPost: IPostRequest) : Observable<any> {
    return this._http.post<any>(
      'http://localhost:9000/post',
      newPost
    );
  }
  LikePost(Username: string, link: any)  {
    return this._http.post<any>(
      'http://localhost:9000' + link,
      { Username }
    );
  }
  DislikePost(Username: string, link: any) {
    return this._http.post<any>(
      'http://localhost:9000' + link,
      { Username }
    );
  }
  CommentPost(Comment: any, link: any) {
    return this._http.post<any>(
      'http://localhost:9000' + link,
      Comment
    );
  }

  GetPost(Id: any) {
    return this._http.get<any>(
      'http://localhost:9000/post/' + Id 
    );
  }

  getUsersFeed(username : string ){
    return this._http.get<any>(
      'http://localhost:9000/users/' + username + '/feed'
    );
  }


}
