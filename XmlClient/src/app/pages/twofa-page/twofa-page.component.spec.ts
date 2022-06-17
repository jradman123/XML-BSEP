import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TwofaPageComponent } from './twofa-page.component';

describe('TwofaPageComponent', () => {
  let component: TwofaPageComponent;
  let fixture: ComponentFixture<TwofaPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TwofaPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TwofaPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
