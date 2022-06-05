import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewJobOfferComponent } from './new-job-offer.component';

describe('NewJobOfferComponent', () => {
  let component: NewJobOfferComponent;
  let fixture: ComponentFixture<NewJobOfferComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewJobOfferComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewJobOfferComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
