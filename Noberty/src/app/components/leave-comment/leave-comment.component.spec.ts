import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LeaveCommentComponent } from './leave-comment.component';

describe('LeaveCommentComponent', () => {
  let component: LeaveCommentComponent;
  let fixture: ComponentFixture<LeaveCommentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LeaveCommentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LeaveCommentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
