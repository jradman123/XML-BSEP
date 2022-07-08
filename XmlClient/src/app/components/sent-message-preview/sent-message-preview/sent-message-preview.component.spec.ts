import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SentMessagePreviewComponent } from './sent-message-preview.component';

describe('SentMessagePreviewComponent', () => {
  let component: SentMessagePreviewComponent;
  let fixture: ComponentFixture<SentMessagePreviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SentMessagePreviewComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SentMessagePreviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
