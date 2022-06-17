import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ConfirmDialogComponent, ConfirmDialogModel } from '../confirm-dialog/confirm-dialog.component';
import { Clipboard } from '@angular/cdk/clipboard';
import { MatSnackBar } from '@angular/material/snack-bar';
import { UserService } from 'src/app/services/user-service/user.service';
import { IUsername } from 'src/app/interfaces/username';


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
  usernameJson!: IUsername;

  constructor(private service: UserService, private dialog: MatDialog,
    private clipboard: Clipboard, private _snackBar: MatSnackBar) {

  }

  ngOnInit(): void {
    this.service.check2FAStatus(this.username).subscribe((res) =>
      this.isChecked = res
    )
  }

  enable2fa(event: any) {

    if (this.isChecked) {
    
      this.service.disable2FA(this.username).subscribe(
        res => {        
          this.code = ""       
          this.isChecked = false
        }
      )
    } else {
      this.service.enable2FA(this.username).subscribe(
        res => {        
          this.code = res.uri       
          this.isChecked = res.res
        }
      )
      this.getCodeVisible = true
    }

  }

  showCode() {
    this.isCodeVisible = true;
  }

  copyMe() {
    this._snackBar.open('Code copied to clipboard.', '', {
      duration: 3000
    });
    this.clipboard.copy(this.code);
  }

}
