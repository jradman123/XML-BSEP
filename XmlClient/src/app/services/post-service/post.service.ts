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
      'http://localhost:9090/post/user/' + username,
    );
  }
  GetUserReactionToPost(username: string, Id: string) {
    return this._http.get<any>(
      'http://localhost:9090/post/' + Id + "/" + username + '/reaction'
    );
  }
  GetAllReactionsForPost(Id: string) {
    return this._http.get<any>(
      'http://localhost:9090/post/' + Id + '/reactions'
    );
  }
  GetAllCommentsForPost(Id: any) {
    return this._http.get<any>(
      'http://localhost:9090/post/' + Id + '/comments'
    );
  }

  CreatePost(newPost: IPostRequest) : Observable<any> {
    return this._http.post<any>(
      'http://localhost:9090/post',
      newPost
    );
  }
  LikePost(Username: string, link: any)  {
    return this._http.post<any>(
      'http://localhost:9090' + link,
      { Username }
    );
  }
  DislikePost(Username: string, link: any) {
    return this._http.post<any>(
      'http://localhost:9090' + link,
      { Username }
    );
  }
  CommentPost(Comment: any, link: any) {
    return this._http.post<any>(
      'http://localhost:9090' + link,
      Comment
    );
  }

  GetPost(Id: any) {
    return this._http.get<any>(
      'http://localhost:9090/post/' + Id 
    );
  }

  getUsersFeed(username : string ){
    return this._http.get<any>(
      'http://localhost:9090/users/' + username + '/feed'
    );
  }


}
