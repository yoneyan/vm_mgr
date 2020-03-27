import { Component, OnInit } from '@angular/core';
import {AuthService} from "../service/auth/auth.service";

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  private Authentication: boolean;

  constructor(private auth: AuthService) { }

  ngOnInit(): void {
    setTimeout(()=>{
      this.Authentication = true

    },3000)
  }

}
