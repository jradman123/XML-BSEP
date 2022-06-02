import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CompanyRequestsComponent } from './company-requests.component';

describe('CompanyRequestsComponent', () => {
  let component: CompanyRequestsComponent;
  let fixture: ComponentFixture<CompanyRequestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CompanyRequestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CompanyRequestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
