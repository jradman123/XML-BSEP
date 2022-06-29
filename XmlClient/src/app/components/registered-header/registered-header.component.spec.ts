import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisteredHeaderComponent } from './registered-header.component';

describe('RegisteredHeaderComponent', () => {
  let component: RegisteredHeaderComponent;
  let fixture: ComponentFixture<RegisteredHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisteredHeaderComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RegisteredHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
