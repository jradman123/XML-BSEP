import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RecoverPassRequestComponent } from './recover-pass-request.component';

describe('RecoverPassRequestComponent', () => {
  let component: RecoverPassRequestComponent;
  let fixture: ComponentFixture<RecoverPassRequestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RecoverPassRequestComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RecoverPassRequestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
