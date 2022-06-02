import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LeaveSallaryCommentComponent } from './leave-sallary-comment.component';

describe('LeaveSallaryCommentComponent', () => {
  let component: LeaveSallaryCommentComponent;
  let fixture: ComponentFixture<LeaveSallaryCommentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LeaveSallaryCommentComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LeaveSallaryCommentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
