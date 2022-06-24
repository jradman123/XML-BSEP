import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { UserService } from 'src/app/services/UserService/user.service';
import {Clipboard} from '@angular/cdk/clipboard';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ConfirmDialogComponent, ConfirmDialogModel } from '../confirm-dialog/confirm-dialog.component';

@Component({
  selector: 'app-two-factor-auth',
  templateUrl: './two-factor-auth.component.html',
  styleUrls: ['./two-factor-auth.component.css']
})
export class TwoFactorAuthComponent implements OnInit {
  isChecked = false;
  code = '';
  getCodeVisible = false;
  isCodeVisible = false;
  email = localStorage.getItem('email')!

  constructor(private service : UserService, private dialog: MatDialog,
     private clipboard: Clipboard, private _snackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.service.check2FAStatus(this.email).subscribe((res) =>
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
        this.service.enable2FA(this.email, false).subscribe(
          res => {
            this.code = res.secret
          }
        )
      });
    } else {
      this.service.enable2FA(this.email, true).subscribe(
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