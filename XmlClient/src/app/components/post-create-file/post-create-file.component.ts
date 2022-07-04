import { Component, OnInit } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IPostRequest } from 'src/app/interfaces/post-request';
import { PostService } from 'src/app/services/post-service/post.service';
import Pusher from 'pusher-js'

@Component({
  selector: 'app-post-create-file',
  templateUrl: './post-create-file.component.html',
  styleUrls: ['./post-create-file.component.css']
})
export class PostCreateFileComponent implements OnInit {
  newPost: IPostRequest
  imageSrc!: string;
  defaultImg = true
  uploaded: boolean = false;
  fileToUpload!: File;
  createForm!: FormGroup;

  username!: string | null;
  imgPath!: string;

  constructor(
    private _formBuilder: FormBuilder,
    private service: PostService,
    private _snackBar: MatSnackBar
  ) {
    this.newPost = {} as IPostRequest
    this.username = localStorage.getItem('username');
    this.createForm = this._formBuilder.group({
      file: new FormControl('', [Validators.required]),
      postText: new FormControl()
    })
  }
  ngOnInit(): void {

    Pusher.logToConsole = true;

    const pusher = new Pusher('dd3ce2a9c4a58e3577a4', {
      cluster: 'eu'
    });

    const channel = pusher.subscribe('notification');
      channel.bind('message', (data: never) => {
        alert(data)
        console.log(data)
    });

  }

  onFileChange(event: any) {
    this.fileToUpload = <File>event.target.files[0];
    const reader = new FileReader();
    reader.readAsDataURL(this.fileToUpload);
    reader.onload = () => {
      this.imageSrc = reader.result as string;
    }

    this.toBase64(this.fileToUpload).then(
      (res) => {
        console.log("ISSSSS HEREEEEEEEEEEE !");
        this.imgPath = res as string

      }
    );

  }
  toBase64 = (file: Blob) =>
    new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = () => resolve(reader.result);
      reader.onerror = (error) => reject(error);
    });

  submitPost() {
    if (this.createForm.invalid) {
      return;
    }
    this.newPost.Username = this.username!
    this.newPost.ImagePaths = this.imgPath
    this.newPost.PostText = this.createForm.value.postText
    console.log(this.newPost);
    const postObserver = {
      next: () => {
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', { duration: 3000 });
      },
    }
    this.service.CreatePost(this.newPost).subscribe(postObserver)

  }
}