import {Component, OnInit} from '@angular/core';
import {Location} from '@angular/common';
import {ActivatedRoute} from "@angular/router";
import {VmService} from "../../service/vm.service"

@Component({
  selector: 'app-detail-vm',
  templateUrl: './detail-vm.component.html',
  styleUrls: ['./detail-vm.component.css']
})
export class DetailVMComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private location: Location,
    private vm: VmService,) {
  }

  public vmdata: VMData;

  ngOnInit(): void {
    this.getID()
  }

  getID(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    this.vm.getVM(id)
      .then(d => this.vmdata = d);
  }

  start(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    this.vm.startVM(id)
  }

  stop(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    this.vm.stopVM(id)
  }

  shutdown(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    this.vm.shutdownVM(id)
  }

  reset(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    this.vm.resetVM(id)
  }

  goBack(): void {
    this.location.back();
  }
}

interface VMData {
  nodeid: string
  id: string
  name: string
  cpu: string
  mem: string
  net: string
  vncurl: string
  status: string
  autostart: boolean

}

