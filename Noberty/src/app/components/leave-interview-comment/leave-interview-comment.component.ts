import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-leave-interview-comment',
  templateUrl: './leave-interview-comment.component.html',
  styleUrls: ['./leave-interview-comment.component.css']
})
export class LeaveInterviewCommentComponent implements OnInit {

  createForm!: FormGroup;
  constructor(private _formBuilder : FormBuilder) { 
    this.createForm = this._formBuilder.group({
      Comment: new FormControl('',[Validators.required]),
      Difficulty : new FormControl('',[Validators.required]),
      Rating : new FormControl('',[Validators.required,Validators.min(1),Validators.max(5)])
    })
  }

  ngOnInit(): void {
  }

}
