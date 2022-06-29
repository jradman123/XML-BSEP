import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PostsCommentsViewComponent } from './posts-comments-view.component';

describe('PostsCommentsViewComponent', () => {
  let component: PostsCommentsViewComponent;
  let fixture: ComponentFixture<PostsCommentsViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PostsCommentsViewComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PostsCommentsViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
