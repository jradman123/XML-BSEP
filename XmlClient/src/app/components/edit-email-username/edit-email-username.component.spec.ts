import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditEmailUsernameComponent } from './edit-email-username.component';

describe('EditEmailUsernameComponent', () => {
  let component: EditEmailUsernameComponent;
  let fixture: ComponentFixture<EditEmailUsernameComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EditEmailUsernameComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EditEmailUsernameComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
