import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EducationDialogComponent } from './education-dialog.component';

describe('EducationDialogComponent', () => {
  let component: EducationDialogComponent;
  let fixture: ComponentFixture<EducationDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EducationDialogComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EducationDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
