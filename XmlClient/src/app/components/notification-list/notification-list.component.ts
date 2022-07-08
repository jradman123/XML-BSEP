import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import Pusher from 'pusher-js';
import { INotification } from 'src/app/interfaces/notification';
import { PusherNotification } from 'src/app/interfaces/pusher-notification';
import { NotificationService } from 'src/app/services/notification-service/notification.service';

@Component({
  selector: 'app-notification-list',
  templateUrl: './notification-list.component.html',
  styleUrls: ['./notification-list.component.css']
})
export class NotificationListComponent implements OnInit {

  @Output() newNotifications : EventEmitter<number> = new EventEmitter();
  notifications : INotification[] = []
  sorted : INotification[] = []
  noOfNotReadNot = 0;

  constructor(private _notificationService: NotificationService) {
    this.notifications.forEach(n => {
      n = {} as INotification
    })
  }

  ngOnInit(): void {
    this._notificationService.getUsersNotifications(localStorage.getItem('username')!).subscribe(
      res => {

        console.log(res.notifications)
        
        res.notifications.forEach((not : INotification) => {
          if(!not.read) this.noOfNotReadNot++;
        })

        this.notifications = res.notifications.sort((a: INotification, b: INotification) => {
          let r = new Date(a.time)
          let q = new Date(b.time)

          if (r > q) {
            return -1;
          } else if (r < q) {
            return 1;
          } else {
            return 0;
          }

        })
      }
    )

    Pusher.logToConsole = true;

    const pusher = new Pusher('dd3ce2a9c4a58e3577a4', {
      cluster: 'eu'
    });

    const channel = pusher.subscribe('notifications');
    channel.bind('notification', (data: never) => {
      
      
      let recieved = data as PusherNotification

      let receivedCasted = {
        id : recieved.Id,
        content : recieved.Content,
        from :  recieved.NotificationFrom,
        to : recieved.NotificationTo,
        read : recieved.Read,
        redirectPath : recieved.RedirectPath,
        notificationType : recieved.Type,
        time : new Date(recieved.Timestamp)

      } as INotification
      
      if(receivedCasted.to === localStorage.getItem('username')){
        this.notifications.unshift(receivedCasted);

        this.newNotifications.emit(++this.noOfNotReadNot);
      }
    });
  }

  markNotiAsRead(notiId : string) {
    
    console.log(this.notifications)

    this.notifications.forEach( n => {
      if (n.id === notiId) {
        n.read = true;
      }
    })

    console.log(this.notifications)

    this.newNotifications.emit(--this.noOfNotReadNot);
  }
}
