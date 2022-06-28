import { Component, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';

@Component({
  selector: 'app-testing',
  templateUrl: './testing.component.html',
  styleUrls: ['./testing.component.css']
})
export class TestingComponent implements OnInit {

  constructor() { }
  createForm!: FormGroup;

  ngOnInit(): void {
  }

  onSubmit(){
    
  }

  onPasswordInput(){}

}
