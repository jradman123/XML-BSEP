import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
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

  GetAllReactionsForPost(Id: string) {
    return this._http.get<any>(
      'http://localhost:9090/post/' + Id + '/reactions'
    );
  }
  GetAllCommentsForPost(link: any) {
    return this._http.get<any>(
      'http://localhost:9090' + link
    );
  }

  CreatePost(newPost: IPostRequest) {
    return this._http.post<any>(
      'http://localhost:9090/post',
      newPost
    );
  }
  LikePost(Username: string, link: any) {
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


}
