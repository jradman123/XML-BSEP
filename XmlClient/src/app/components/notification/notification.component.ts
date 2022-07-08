import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { INotification } from 'src/app/interfaces/notification';
import { NotificationService } from 'src/app/services/notification-service/notification.service';

@Component({
  selector: 'app-notification',
  templateUrl: './notification.component.html',
  styleUrls: ['./notification.component.css']
})
export class NotificationComponent implements OnInit {

  @Input()
  noti! : INotification
  @Output() notiRead : EventEmitter<string> = new EventEmitter();

  
  constructor(private _service : NotificationService){
  }

  ngOnInit(): void {
  }

  readMe() {
    this._service.markAsRead(this.noti.id).subscribe(
       ()=> {this.notiRead.emit(  this.noti.id )}
    );
  }
}
