import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-leave-sallary-comment',
  templateUrl: './leave-sallary-comment.component.html',
  styleUrls: ['./leave-sallary-comment.component.css']
})
export class LeaveSallaryCommentComponent implements OnInit {

  createForm!: FormGroup;
  constructor(private _formBuilder : FormBuilder) { 
    this.createForm = this._formBuilder.group({
      Comment: new FormControl('',[Validators.required]),
      Position : new FormControl('',[Validators.required]),
    })
  }

  ngOnInit(): void {
  }

}
