import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { IUsername } from 'src/app/interfaces/username';
import { UserService } from 'src/app/services/user-service/user.service';
import { Clipboard } from '@angular/cdk/clipboard';
import { ApiKeyService } from 'src/app/services/api-key-service/api-key.service';
import { UserDetails } from 'src/app/interfaces/user-details';
import { NotificationService } from 'src/app/services/notification-service/notification.service';
import { NotificationSettings } from 'src/app/interfaces/notification-settings';
import { GetSettingsResponse } from 'src/app/interfaces/get-settings-response';
import { ChangeSettingsRequest } from 'src/app/interfaces/change-settings-request';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent implements OnInit {

  @Input()
  user! : UserDetails;
  userDetails! : UserDetails;
  isChecked = false;
  code = '';
  getCodeVisible = false;
  isCodeVisible = false;
  username = localStorage.getItem('username')!
  usernameJson!: IUsername;
  apiKey!: string;
  clicked : boolean = false;
  privateChecked! : boolean;
  messagesNotificationChecked! : boolean;
  postsNotificationChecked! : boolean;
  connectionsNotificationChecked! : boolean;
  notificationSettings! : GetSettingsResponse;
  changeSettingsRequest! : ChangeSettingsRequest;

  constructor(private service: UserService, private dialog: MatDialog,
    private clipboard: Clipboard, private _snackBar: MatSnackBar, private apiKeyService: ApiKeyService,
    private notificationService : NotificationService, private userService : UserService) {
      this.notificationSettings = {} as GetSettingsResponse
      this.changeSettingsRequest = {} as ChangeSettingsRequest
      this.changeSettingsRequest.settings = {} as NotificationSettings
      this.userDetails = {} as UserDetails
      this.setSettings();
      this.setOnInit();
     }

  ngOnInit(): void {
   /* this.service.check2FAStatus(this.username).subscribe((res) =>
      this.isChecked = res
    );
    this.privateChecked = this.user.profileStatus === "PRIVATE"
    this.setSettings();*/
  
  }

  setOnInit() {
    this.service.check2FAStatus(this.username).subscribe((res) =>
      this.isChecked = res
    );
    this.userService.getUserDetails(localStorage.getItem('username')).subscribe({
      next: (data: UserDetails) => {
        this.userDetails = data
        this.privateChecked = this.userDetails.profileStatus === "PRIVATE"
      },
    });
  }

  setSettings() {
    this.notificationService.getUsersNotificationsSettings(this.username).subscribe({
      next: (data: GetSettingsResponse) => {
        this.messagesNotificationChecked = data.settings.messages;
        this.connectionsNotificationChecked = data.settings.connections;
        this.postsNotificationChecked = data.settings.posts;

      },
    })
    }  


  changeStatus() {
    console.log(this.privateChecked);
    
    this.privateChecked = !this.privateChecked
    this.service.changePrivacyStatus(this.user.username, this.privateChecked ? "PRIVATE" : "PUBLIC").subscribe(
      res => {
        this._snackBar.open('Privacy settings updated.', '', {
          duration: 3000
        });
      },
    );
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

  changeConnectionNotifications()  {
    this.connectionsNotificationChecked = !this.connectionsNotificationChecked
    this.createChangeSettingsRequest();
    this.notificationService.changeNotificationSettings(this.changeSettingsRequest).subscribe(
      res => {
        console.log("usao")
        this._snackBar.open('Connection notifications updated.', '', {
          duration: 3000
        });
      },
    );
    

  }

  createChangeSettingsRequest() {
    this.changeSettingsRequest.username = this.username;
    this.changeSettingsRequest.settings.connections = this.connectionsNotificationChecked
    this.changeSettingsRequest.settings.messages = this.messagesNotificationChecked
    this.changeSettingsRequest.settings.posts = this.postsNotificationChecked
  }

  changePostNotifications() {
    this.postsNotificationChecked = !this.postsNotificationChecked
    this.createChangeSettingsRequest();
    this.notificationService.changeNotificationSettings(this.changeSettingsRequest).subscribe(
      res => {
        console.log("usao")
        this._snackBar.open('Post notifications updated.', '', {
          duration: 3000
        });
      },
    );

  }

  changeMessageNotifications() {
    this.messagesNotificationChecked = !this.messagesNotificationChecked
    this.createChangeSettingsRequest();
    this.notificationService.changeNotificationSettings(this.changeSettingsRequest).subscribe(
      res => {
        console.log("usao")
        this._snackBar.open('Message notifications updated.', '', {
          duration: 3000
        });
      },
    );

  }

}
