import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';

@Component({
  selector: 'app-twofa-page',
  templateUrl: './twofa-page.component.html',
  styleUrls: ['./twofa-page.component.css']
})
export class TwofaPageComponent implements OnInit {
  createForm!: FormGroup;
  constructor(
    private _snackBar: MatSnackBar,
    private _router: Router,
    private formBuilder: FormBuilder) { }

  ngOnInit(): void {
    this.createForm = this.formBuilder.group({
      code: new FormControl('', [
        Validators.required,
        Validators.pattern('^[0-9]{6}$'),
      ]),
    });
  }
  onSbmit() {

    if (this.createForm.invalid) {
      return;
    }


  }
}
