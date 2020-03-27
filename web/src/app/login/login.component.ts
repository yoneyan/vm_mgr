import {Component, OnInit} from '@angular/core';
import {FormBuilder} from "@angular/forms";
import {Router} from '@angular/router';

import {AuthService} from "../service/auth/auth.service";

@Component({
  selector: 'app-sign-in',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm;
  // private result: boolean;
  //
  // model: Login = new Login('', '');
  // submitted: boolean = false;
  // errormsg: string = undefined;

  title = 'test'
  result = ""

  constructor(
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private router: Router,
    // private loginservice: Test,
  ) {
  }

  ngOnInit() {
  }

  onClickSubmit(data) {
    console.log(data)
    this.result = this.authService.verifyUser(data)

  }

}
