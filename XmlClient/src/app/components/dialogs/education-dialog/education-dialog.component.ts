import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-education-dialog',
  templateUrl: './education-dialog.component.html',
  styleUrls: ['./education-dialog.component.css']
})
export class EducationDialogComponent implements OnInit {

  data! : any;
  constructor(public dialogRef: MatDialogRef<EducationDialogComponent>) { }

  ngOnInit(): void {
  }

  close(): void {
    this.dialogRef.close();
  }


}
