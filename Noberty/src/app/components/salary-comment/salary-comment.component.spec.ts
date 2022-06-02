import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SalaryCommentComponent } from './salary-comment.component';

describe('SalaryCommentComponent', () => {
  let component: SalaryCommentComponent;
  let fixture: ComponentFixture<SalaryCommentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SalaryCommentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SalaryCommentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
