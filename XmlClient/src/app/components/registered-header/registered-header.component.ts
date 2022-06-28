import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-registered-header',
  templateUrl: './registered-header.component.html',
  styleUrls: ['./registered-header.component.css']
})
export class RegisteredHeaderComponent implements OnInit {

  username : any;
  visible : boolean = false;
  constructor() { }

  ngOnInit(): void {
    this.username=localStorage.getItem('username');
  }

  click() {
    console.log('usao');
    if(this.visible){
      console.log('true')
      this.visible=false;
    }else{
      console.log('false')
      this.visible=true;
    }
  }

}
