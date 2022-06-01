import { ComponentFixture, TestBed } from '@angular/core/testing';

import { JobOfferListViewComponent } from './job-offer-list-view.component';

describe('JobOfferListViewComponent', () => {
  let component: JobOfferListViewComponent;
  let fixture: ComponentFixture<JobOfferListViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ JobOfferListViewComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(JobOfferListViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
