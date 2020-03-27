import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateVmComponent } from './create-vm.component';

describe('CreateVmComponent', () => {
  let component: CreateVmComponent;
  let fixture: ComponentFixture<CreateVmComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CreateVmComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateVmComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
