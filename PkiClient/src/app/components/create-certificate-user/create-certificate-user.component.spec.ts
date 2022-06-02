import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateCertificateUserComponent } from './create-certificate-user.component';

describe('CreateCertificateUserComponent', () => {
  let component: CreateCertificateUserComponent;
  let fixture: ComponentFixture<CreateCertificateUserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateCertificateUserComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateCertificateUserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
