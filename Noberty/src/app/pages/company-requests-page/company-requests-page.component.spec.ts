import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CompanyRequestsPageComponent } from './company-requests-page.component';

describe('CompanyRequestsPageComponent', () => {
  let component: CompanyRequestsPageComponent;
  let fixture: ComponentFixture<CompanyRequestsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CompanyRequestsPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CompanyRequestsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
