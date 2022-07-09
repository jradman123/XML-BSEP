import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { IMesssage } from 'src/app/interfaces/message';
import { MessageService } from 'src/app/services/messsage-service/messaage.service';

@Component({
  selector: 'app-message-create',
  templateUrl: './message-create.component.html',
  styleUrls: ['./message-create.component.css']
})
export class MessageCreateComponent implements OnInit {
  createForm!: FormGroup;
  newMessaage: IMesssage;
  username!: string
  receiver!: string

  constructor(
    private _formBuilder: FormBuilder,
    private _snackBar: MatSnackBar,
    private _service: MessageService,
    private _router : Router
  ) {
    this.newMessaage = {} as IMesssage
    this.username = localStorage.getItem('username')!
    this.createForm = this._formBuilder.group({
      message: new FormControl('', Validators.required)
    })
   }

  ngOnInit(): void {
    this.receiver = this._router.url.substring(14) ?? '';
    console.log(this.receiver);
  }
  submitRequest(): void {
    if (this.createForm.invalid){
      return;
    }
    this.newMessaage.MessageText = this.createForm.value.message
    this.newMessaage.SenderUsername = this.username
    this.newMessaage.ReceiverUsername = this.receiver
    console.log(this.newMessaage);
    
    const sendMessage = {
      next: (res:IMesssage) => {
        console.log(res);

      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open("Message cannot be sent" + "!", '', {duration : 3000,panelClass: ['snack-bar']});
      },
    }
    this._service.SendMessage(this.newMessaage).subscribe(sendMessage)
  }
  
  clearForm() {
    this.createForm.reset()
  }
}
