import { Component, Input, OnInit } from '@angular/core';
import { IMesssage } from 'src/app/interfaces/message';
@Component({
  selector: 'app-sent-message-preview',
  templateUrl: './sent-message-preview.component.html',
  styleUrls: ['./sent-message-preview.component.css']
})
export class SentMessagePreviewComponent implements OnInit {
  @Input()
  item: IMesssage
  constructor() {
    this.item = {} as IMesssage
   }

  ngOnInit(): void {
   
  }

}
