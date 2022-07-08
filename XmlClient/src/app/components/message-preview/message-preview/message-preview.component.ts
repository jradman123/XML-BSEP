import { Component, Input, OnInit } from '@angular/core';
import { IMesssage } from 'src/app/interfaces/message';

@Component({
  selector: 'app-message-preview',
  templateUrl: './message-preview.component.html',
  styleUrls: ['./message-preview.component.css']
})
export class MessagePreviewComponent implements OnInit {
  @Input()
  item: IMesssage
  constructor() {
    this.item = {} as IMesssage
   }

  ngOnInit(): void {
   
  }

}
