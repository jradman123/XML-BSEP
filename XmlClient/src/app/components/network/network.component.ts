import { Component, OnInit } from '@angular/core';
import { UserDetails } from 'src/app/interfaces/user-details';

@Component({
  selector: 'app-network',
  templateUrl: './network.component.html',
  styleUrls: ['./network.component.css']
})
export class NetworkComponent implements OnInit {
  searchText : string = "";
  profiles! : UserDetails[];
  constructor() { }

  ngOnInit(): void {
  }

  handleMe(searchText : string){
    this.searchText = searchText;
  }
}
