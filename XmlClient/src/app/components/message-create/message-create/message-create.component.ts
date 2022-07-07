import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IComment } from 'src/app/interfaces/comment';
import { IMesssage } from 'src/app/interfaces/message';
import { MessageService } from 'src/app/services/messsage-service/messaage.service';
import { CommentCreateComponent } from '../../comment-create/comment-create.component';

@Component({
  selector: 'app-message-create',
  templateUrl: './message-create.component.html',
  styleUrls: ['./message-create.component.css']
})
export class MessageCreateComponent implements OnInit {
  createForm!: FormGroup;
  newMessaage: IMesssage;
  username!: string
  constructor(
    private _formBuilder: FormBuilder,
    public dialogRef: MatDialogRef<CommentCreateComponent>,
    private _snackBar: MatSnackBar,
    private _service: MessageService,
  ) {
    this.newMessaage = {} as IMesssage
    this.username = localStorage.getItem('username')!
    this.createForm = this._formBuilder.group({
      message: new FormControl('', Validators.required)
    })
   }

  ngOnInit(): void {
  }
  submitRequest(): void {
    if (this.createForm.invalid){
      return;
    }
    this.newMessaage.MessageText = this.createForm.value.message
    this.newMessaage.SenderUsername = this.username
    this.newMessaage.ReceiverUsername = this.username
    console.log(this.newMessaage);
    
    const sendMessage = {
      next: (res:IMesssage) => {
        console.log(res);

      },
      error: (err: HttpErrorResponse) => {
        this._snackBar.open(err.error.message + "!", 'Dismiss', { duration: 3000 });
      },
    }
    this._service.SendMessage(this.newMessaage).subscribe(sendMessage)
  }
  
  clearForm() {
    this.createForm.reset()
  }
}
