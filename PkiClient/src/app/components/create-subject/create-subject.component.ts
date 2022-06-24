import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { SubjectData } from 'src/app/interfaces/subject-data';
import { UserService } from 'src/app/services/UserService/user.service';

@Component({
  selector: 'app-create-subject',
  templateUrl: './create-subject.component.html',
  styleUrls: ['./create-subject.component.css'],
})
export class CreateSubjectComponent implements OnInit {
  @Output() newItemEvent = new EventEmitter<string>();
  hide = true;
  itsId!: number;
  subject: SubjectData | undefined;
  subjectForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl(),
    commonName: new FormControl(),
    organization: new FormControl(),
    organizationUnit: new FormControl(),
    locality: new FormControl(),
    country: new FormControl(),
    recoveryMail : new FormControl()
  });

  getErrorMessage() {
    if (this.subjectForm.controls['email'].errors?.required) {
      return 'You must enter a value';
    }
    return this.subjectForm.controls['email'].errors?.email
      ? 'Not a valid email'
      : '';
  }
  constructor(
    private userService: UserService,
    private _snackBar: MatSnackBar
  ) {}

  ngOnInit(): void {}

  onFormSubmit(): void {
    if(this.subjectForm.invalid){
      return;
    }
    this.subject = {
      id: 0,
      commonName: this.subjectForm.get('commonName')?.value,
      organization: this.subjectForm.get('organization')?.value,
      organizationUnit: this.subjectForm.get('organizationUnit')?.value,
      locality: this.subjectForm.get('locality')?.value,
      country: this.subjectForm.get('country')?.value,
      email: this.subjectForm.get('email')?.value,
      password: this.subjectForm.get('password')?.value,
      recoveryMail : this.subjectForm.get('recoveryMail')?.value
    };
    console.log(this.subject);
    this.userService.createSubject(this.subject).subscribe(
      (res) => {
        this._snackBar.open('Subject successfully created', 'Dismiss',{
          duration : 3000
         });
        console.log(res);
        this.itsId = res.id;
        console.log(this.itsId);
        this.newItemEvent.emit(this.itsId.toString());
      },
      (err) => {
        this._snackBar.open(
          'Subject could not be created! Please try again.',
          'Dismiss',
          {
            duration : 3000
           }
        );
      }
    );
  }
}
