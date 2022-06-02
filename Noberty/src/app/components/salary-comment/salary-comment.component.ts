import { Component, Input, OnInit } from '@angular/core';
import { ISalaryComment } from 'src/app/interfaces/salary-comment';

@Component({
  selector: 'app-salary-comment',
  templateUrl: './salary-comment.component.html',
  styleUrls: ['./salary-comment.component.css']
})
export class SalaryCommentComponent implements OnInit {

  @Input()
  salaryComment! : ISalaryComment;
  constructor() { }

  ngOnInit(): void {
  }

}
