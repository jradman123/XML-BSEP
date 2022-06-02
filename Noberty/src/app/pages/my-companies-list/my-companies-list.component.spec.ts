import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyCompaniesListComponent } from './my-companies-list.component';

describe('MyCompaniesListComponent', () => {
  let component: MyCompaniesListComponent;
  let fixture: ComponentFixture<MyCompaniesListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MyCompaniesListComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MyCompaniesListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
