import {Component, OnInit} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Router} from "@angular/router";
import {AuthService} from "../../service/auth/auth.service";
import {Observable} from "rxjs";
import {retry} from "rxjs/operators";
import {VmService} from "../../service/vm.service";


@Component({
  selector: 'app-vm',
  templateUrl: './vm.component.html',
  styleUrls: ['./vm.component.css']
})
export class VmComponent implements OnInit {

  constructor(private VmService: VmService) {
  }

  public test: VMData[] ;
  public a : string[];
  public b : string;

  ngOnInit(): void {
    this.getUserVM()
  }
    getUserVM():void{
      this.VmService.getUserVM()
        .then(d => {
          // let data:
          this.test = d
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
