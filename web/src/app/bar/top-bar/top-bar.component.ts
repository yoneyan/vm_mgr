import { Component, OnInit } from '@angular/core';
import {AuthService} from "../../service/auth/auth.service";

@Component({
  selector: 'app-top-bar',
  templateUrl: './top-bar.component.html',
  styleUrls: ['./top-bar.component.css']
})
export class TopBarComponent implements OnInit {


  constructor(private Auth :AuthService) { }
  User = this.Auth.getAuthUser();

  ngOnInit(): void {
  }

}
