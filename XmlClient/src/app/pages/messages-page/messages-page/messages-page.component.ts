import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { auto } from '@popperjs/core';
import { MessageCreateComponent } from 'src/app/components/message-create/message-create/message-create.component';
import { IMesssages } from 'src/app/interfaces/message';
import { MessageService } from 'src/app/services/messsage-service/messaage.service';

@Component({
  selector: 'app-messages-page',
  templateUrl: './messages-page.component.html',
  styleUrls: ['./messages-page.component.css']
})
export class MessagesPageComponent implements OnInit {
  Sent: IMesssages
  Received:IMesssages
  constructor(
    public _matDialog: MatDialog,
    private _service: MessageService,
    private _snackBar: MatSnackBar,

  ) {
    this.Sent = {} as IMesssages
    this.Received = {} as IMesssages
    this.refresh()
   }
   

  ngOnInit(): void {
  }
  refresh() {
    const getSentMessages = {
      next: (res: IMesssages) => {
        console.log(res);
  
        if (res.Messages.length == 0) {
          return
        }
  
        this.Sent = res
       
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open("Error happend" + "!", '', {duration : 3000,panelClass: ['snack-bar']});
      },
    }
    const getReceivedMessages = {
      next: (res: IMesssages) => {
        console.log(res);
  
        if (res.Messages.length == 0) {
          return
        }
  
        this.Received = res
       
      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open("Error happend" + "!", '', {duration : 3000,panelClass: ['snack-bar']});
      },
    }
    this._service.GetSentMessages().subscribe(getSentMessages)
    this._service.GetReceivedMessages().subscribe(getReceivedMessages)
  }
  newMessage() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = '500px';
    dialogConfig.width = auto;

    this._matDialog.open(MessageCreateComponent, dialogConfig);
  }
}
