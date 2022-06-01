import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CompanyListViewComponent } from './company-list-view.component';

describe('CompanyListViewComponent', () => {
  let component: CompanyListViewComponent;
  let fixture: ComponentFixture<CompanyListViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CompanyListViewComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CompanyListViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
