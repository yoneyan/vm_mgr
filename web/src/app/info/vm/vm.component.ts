import {Component, OnInit} from '@angular/core';
import {VmService} from "../../service/vm.service";
import {RouterModule} from "@angular/router";


@Component({
  selector: 'app-vm',
  templateUrl: './vm.component.html',
  styleUrls: ['./vm.component.css']
})
export class VmComponent implements OnInit {

  constructor(private VmService: VmService,
              private router: RouterModule) {
  }

  public vms: VMData[];

  ngOnInit(): void {
    this.getUserVM()
  }

  getUserVM(): void {
    this.VmService.getUserVM()
      .then(d => {
        // let data:
        this.vms = d
        console.log(d)
      })
  }

}

interface VMData {
  nodeid: string
  id: string
  name: string
  cpu: string
  mem: string
  status: string
  autostart: boolean

}
