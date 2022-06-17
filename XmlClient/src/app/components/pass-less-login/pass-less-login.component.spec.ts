import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PassLessLoginComponent } from './pass-less-login.component';

describe('PassLessLoginComponent', () => {
  let component: PassLessLoginComponent;
  let fixture: ComponentFixture<PassLessLoginComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PassLessLoginComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PassLessLoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
