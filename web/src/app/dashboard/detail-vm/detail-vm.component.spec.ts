import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DetailVMComponent } from './detail-vm.component';

describe('DetailVMComponent', () => {
  let component: DetailVMComponent;
  let fixture: ComponentFixture<DetailVMComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DetailVMComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DetailVMComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
