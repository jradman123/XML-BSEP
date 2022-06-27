import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-post-create',
  templateUrl: './post-create.component.html',
  styleUrls: ['./post-create.component.css']
})
export class PostCreateComponent implements OnInit {
  imageSrc!: string;
  defaultImg = true
  createForm!: FormGroup;
  links!: string[];
  myForm = new FormGroup({

    file: new FormControl('', [Validators.required]),
    fileSource: new FormControl('', [Validators.required])
  });


  constructor(private _formBuilder: FormBuilder,) {
    this.createForm = this._formBuilder.group({
      Links: new FormControl('', [
        Validators.required,
      ]),
    })
    this.links = []
  }

  ngOnInit(): void {
  }
  get f() {
    return this.myForm.controls;
  }
  addItem() {
    if (this.createForm.value.Links == null) {
      return
    }
    console.log(this.createForm.value.Links);
    
    this.links.push(this.createForm.value.Links)
  }
  onFileChange(event: any) {
    this.defaultImg = false
    const reader = new FileReader();

    if (event.target.files && event.target.files.length) {
      const [file] = event.target.files;
      reader.readAsDataURL(file);

      reader.onload = () => {

        this.imageSrc = reader.result as string;

        this.myForm.patchValue({
          fileSource: reader.result
        });

      };
    }
  }

  submit() {
    console.log(this.myForm.value);
  }
}
