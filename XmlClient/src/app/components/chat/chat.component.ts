import { DatePipe } from '@angular/common';
import { Component, ElementRef, Input, OnChanges, OnInit, SimpleChanges } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import Pusher from 'pusher-js';
import { IMesssage } from 'src/app/interfaces/message';
import { MessageService } from 'src/app/services/messsage-service/messaage.service';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css']
})
export class ChatComponent implements OnInit, OnChanges  {
 
  @Input()
  receiver : string = '';
  sender : string = localStorage.getItem('username') ?? ''; 
  msgText : string = '';
  messages : IMesssage[] = []
  form!: FormGroup;


  constructor(private _router : Router, private _service : MessageService, private _datepipe: DatePipe,
    private _formBuilder: FormBuilder) { 

    this.messages.forEach((m : any) => m = {} as IMesssage)
  }


  ngOnChanges(changes: SimpleChanges) {
    console.log('its a change ' + changes)
    this.messages = []
    this.getData();
  
}

  ngOnInit(): void {
    this.form = this._formBuilder.group({
      msgText: ['', Validators.required],
    });
    
    this.messages = []

    Pusher.logToConsole = true;

    const pusher = new Pusher('e49d7a86a937f12da028', {
      cluster: 'eu'
    });

    console.log('evo ti ga pushhher ')
    console.log(pusher)

    const channel = pusher.subscribe('messages');
    channel.bind('message', (data: never) => {
    

      let receivedMsg = data as IMesssage
      if(receivedMsg.ReceiverUsername == this.sender) {
        this.messages.push(receivedMsg);
      }
    });
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
  }

  sendMessage() {

    let newMsg : IMesssage = {
      Id : '', 
      SenderUsername: this.sender,
      ReceiverUsername : this.receiver,
      MessageText : this.form.value.msgText,
      TimeSent :this._datepipe.transform(new Date(), 'yyyy-MM-dd HH:mm:ss')!
    }

    this._service.SendMessage(newMsg).subscribe(res => {
      this.messages.push(newMsg);
      this.form.reset()
    });

  }
  
}