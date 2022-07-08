import { DatePipe } from '@angular/common';
import { AfterViewInit, Component, ElementRef, Input, OnChanges, OnInit, SimpleChanges, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { IMesssage } from 'src/app/interfaces/message';
import { MessageService } from 'src/app/services/messsage-service/messaage.service';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css']
})
export class ChatComponent implements OnInit, AfterViewInit, OnChanges  {

  @ViewChild('scrollframe', {static: false}) scrollFrame!: ElementRef;
  private scrollContainer: any;
  @Input()
  receiver : string = '';
  sender : string = localStorage.getItem('username') ?? ''; 
  msgText : string = '';
  messages : IMesssage[] = []

  constructor(private _router : Router, private _service : MessageService, private _datepipe: DatePipe,) { 
    this.messages.forEach((m : any) => m = {} as IMesssage)
  }

  ngAfterViewInit() {
    this.scrollContainer = this.scrollFrame.nativeElement;  
  }

  ngOnChanges(changes: SimpleChanges) {
    console.log('its a change ' + changes)
    this.messages = []
    this.getData();
    
}

  ngOnInit(): void {
    this.messages = []
    this.scrollToBottom();

  }

  getData(){
    this._service.GetSentMessages().subscribe(res => {
      console.log('sent')
        console.log(res.Messages)
      res.Messages.forEach((message  : any )=> {
        if(message.ReceiverUsername == this.receiver) {
          this.messages.push(message);
        }
      });
      this.sortUs()
    })

    this._service.GetReceivedMessages().subscribe(res => {
      console.log('received')
        console.log(res.Messages)
      res.Messages.forEach((message  : any )=> {
        if(message.SenderUsername == this.receiver) {
          this.messages.push(message);
        }
      });
      this.sortUs()
    })
  }


  sortUs() {
    this.messages =  this.messages.sort((a: IMesssage, b: IMesssage) => {
      let r = new Date(a.TimeSent)
      let q = new Date(b.TimeSent)

      if (r > q) {
        return 1;
      } else if (r < q) {
        return -1;
      } else {
        return 0;
      }

    })

    console.log("nakon sorta")
    console.log(this.messages)
  }

  sendMessage() {
    let newMsg : IMesssage = {
      Id : '', 
      SenderUsername: this.sender,
      ReceiverUsername : this.receiver,
      MessageText : this.msgText,
      TimeSent :this._datepipe.transform(new Date(), 'yyyy-MM-dd HH:mm:ss')!
    }
    this._service.SendMessage(newMsg).subscribe(res => {
      this.messages.push(newMsg);
      this.msgText = '';
      this.scrollToBottom();
    });
  }

  private scrollToBottom(): void {
    this.scrollContainer.scroll({
      top: this.scrollContainer.scrollHeight,
      left: 0,
      behavior: 'smooth'
    });
  }
}