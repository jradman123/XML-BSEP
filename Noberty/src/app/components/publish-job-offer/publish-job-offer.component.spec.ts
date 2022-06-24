import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PublishJobOfferComponent } from './publish-job-offer.component';

describe('PublishJobOfferComponent', () => {
  let component: PublishJobOfferComponent;
  let fixture: ComponentFixture<PublishJobOfferComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PublishJobOfferComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PublishJobOfferComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
