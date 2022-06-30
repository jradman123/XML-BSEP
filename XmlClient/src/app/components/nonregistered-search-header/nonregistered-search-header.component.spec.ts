import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NonregisteredSearchHeaderComponent } from './nonregistered-search-header.component';

describe('NonregisteredSearchHeaderComponent', () => {
  let component: NonregisteredSearchHeaderComponent;
  let fixture: ComponentFixture<NonregisteredSearchHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NonregisteredSearchHeaderComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NonregisteredSearchHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
