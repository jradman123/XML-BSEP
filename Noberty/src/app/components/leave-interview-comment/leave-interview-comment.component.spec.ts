import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LeaveInterviewCommentComponent } from './leave-interview-comment.component';

describe('LeaveInterviewCommentComponent', () => {
  let component: LeaveInterviewCommentComponent;
  let fixture: ComponentFixture<LeaveInterviewCommentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LeaveInterviewCommentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LeaveInterviewCommentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
