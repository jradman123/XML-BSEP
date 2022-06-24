import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { TfaComponent } from 'src/app/components/tfa/tfa.component';

@Component({
  selector: 'app-edit-user',
  templateUrl: './edit-user.component.html',
  styleUrls: ['./edit-user.component.css']
})
export class EditUserComponent implements OnInit {

  constructor(public matDialog: MatDialog) { }

  ngOnInit(): void {
  }
  opet2fa() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = false;
    dialogConfig.id = 'modal-component';
    dialogConfig.height = 'fit-content';
    dialogConfig.width = '800px';
    this.matDialog.open(TfaComponent, dialogConfig);
  }

}
