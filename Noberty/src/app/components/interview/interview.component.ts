import { Component, Input, OnInit } from '@angular/core';
import { IInterview } from 'src/app/interfaces/interview';

@Component({
  selector: 'app-interview',
  templateUrl: './interview.component.html',
  styleUrls: ['./interview.component.css']
})
export class InterviewComponent implements OnInit {

  @Input()
  interview! : IInterview;

  constructor() {}
  ngOnInit(): void {
  }

}
