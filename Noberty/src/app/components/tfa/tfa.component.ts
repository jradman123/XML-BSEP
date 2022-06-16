import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { UserServiceService } from 'src/app/services/UserService/user-service.service';
import { ConfirmDialogComponent, ConfirmDialogModel } from '../confirm-dialog/confirm-dialog.component';
import {Clipboard} from '@angular/cdk/clipboard';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-tfa',
  templateUrl: './tfa.component.html',
  styleUrls: ['./tfa.component.css']
})
export class TfaComponent implements OnInit {
  isChecked = false;
  code = '';
  getCodeVisible = false;
  isCodeVisible = false;
  username = localStorage.getItem('username')!

  constructor(private service : UserServiceService, private dialog: MatDialog,
     private clipboard: Clipboard, private _snackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.service.check2FAStatus(this.username).subscribe((res) =>
      this.isChecked = res
    )
  }

  enable2fa(event:any ){

    if (this.isChecked) {
      const message = `Are you sure you want to disable 2 factor authentication?`;

      const dialogData = new ConfirmDialogModel("Confirm", message);
  
      const dialogRef = this.dialog.open(ConfirmDialogComponent, {
        maxWidth: "400px",
        data: dialogData
      });

      dialogRef.afterClosed().subscribe(dialogResult => {
        this.isChecked = !dialogResult;
        this.getCodeVisible =  !dialogResult;
        this.isCodeVisible =  !dialogResult;
        this.service.enable2FA(this.username, false).subscribe(
          res => {
            this.code = res.secret
          }
        )
      });
    } else {
      this.service.enable2FA(this.username, true).subscribe(
        res => {
          this.code = res.secret
        }
      )
      this.getCodeVisible = true
    }

  }

  showCode() {
    this.isCodeVisible = true;
  }

  copyMe() { 
    this._snackBar.open('Code copied to clipboard.','', {
    duration: 3000
  });
    this.clipboard.copy(this.code);
  }

}
