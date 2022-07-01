import { Component, EventEmitter, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-nonregistered-search-header',
  templateUrl: './nonregistered-search-header.component.html',
  styleUrls: ['./nonregistered-search-header.component.css']
})
export class NonregisteredSearchHeaderComponent implements OnInit {
 
  @Output() searchInput : EventEmitter<string> = new EventEmitter();
  constructor() { 
  }

  ngOnInit(): void {
  }

  emitMe( searchText : any){
    this.searchInput.emit(searchText.target.value);
  }
  

}
