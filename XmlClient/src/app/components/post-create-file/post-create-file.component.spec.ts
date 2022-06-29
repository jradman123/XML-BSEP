import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PostCreateFileComponent } from './post-create-file.component';

describe('PostCreateFileComponent', () => {
  let component: PostCreateFileComponent;
  let fixture: ComponentFixture<PostCreateFileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PostCreateFileComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PostCreateFileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
