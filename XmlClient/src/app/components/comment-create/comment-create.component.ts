import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';


@Component({
  selector: 'app-comment-create',
  templateUrl: './comment-create.component.html',
  styleUrls: ['./comment-create.component.css']
})
export class CommentCreateComponent implements OnInit {
  createForm!: FormGroup;
 
 
  constructor(
    private _formBuilder: FormBuilder,
  ) {
    this.createForm = this._formBuilder.group({
      comment: new FormControl('',Validators.required)
    })
   }

  ngOnInit(): void {
  }
  submitRequest(): void {}
  
}
