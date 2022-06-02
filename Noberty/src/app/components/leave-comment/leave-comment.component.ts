import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { IComment } from 'src/app/interfaces/comment';
import { CompanyService } from 'src/app/services/company-service/company.service';

@Component({
  selector: 'app-leave-comment',
  templateUrl: './leave-comment.component.html',
  styleUrls: ['./leave-comment.component.css']
})
export class LeaveCommentComponent implements OnInit {

  comment!: IComment;
  createForm!: FormGroup;
  constructor(private _formBuilder: FormBuilder,
    private companyService: CompanyService) { 
    this.createForm = this._formBuilder.group({
      Comment: new FormControl('',Validators.required)
    })
  }

  ngOnInit(): void {
  }



}
