import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OverviewProfileComponent } from './overview-profile.component';

describe('OverviewProfileComponent', () => {
  let component: OverviewProfileComponent;
  let fixture: ComponentFixture<OverviewProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ OverviewProfileComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(OverviewProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
