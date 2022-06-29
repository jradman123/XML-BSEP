import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IUsername } from 'src/app/interfaces/username';
import { UserService } from 'src/app/services/user-service/user.service';
import { Clipboard } from '@angular/cdk/clipboard';
import { ApiKeyService } from 'src/app/services/api-key-service/api-key.service';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent implements OnInit {

  isChecked = false;
  code = '';
  getCodeVisible = false;
  isCodeVisible = false;
  username = localStorage.getItem('username')!
  usernameJson!: IUsername;
  apiKey!: string;
  clicked : boolean = false;

  constructor(private service: UserService, private dialog: MatDialog,
    private clipboard: Clipboard, private _snackBar: MatSnackBar,private apiKeyService: ApiKeyService) { }

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

  generate() {
    if(this.clicked){
      this.clicked = false;
    }else{
      this.clicked = true;
    }
    this.apiKeyService
      .generateApiKey(localStorage.getItem('username')!)
      .subscribe((res) => (this.apiKey = res.apiToken));
  }

}
