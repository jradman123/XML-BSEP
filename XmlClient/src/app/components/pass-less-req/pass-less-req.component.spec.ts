import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PassLessReqComponent } from './pass-less-req.component';

describe('PassLessReqComponent', () => {
  let component: PassLessReqComponent;
  let fixture: ComponentFixture<PassLessReqComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PassLessReqComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PassLessReqComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
