import {Component, OnInit} from '@angular/core';
import {FormBuilder} from "@angular/forms";

import {LoginService} from "../login.service";

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent implements OnInit {
  loginForm;
  private result: boolean;

  test() {
    // this.loginService.test();
    window.alert('Your product has been added to the cart!');
  }

  title = 'test'

  constructor(
    private formBuilder: FormBuilder,
    private loginService: LoginService,
  ) {


    // this.loginForm = this.formBuilder.group({
    //   name: '',
    //   pass: '',
    // });
  }

  ngOnInit() {
  }

  onClickSubmit(data) {
    this.result = this.loginService.verifyUser(data)
    alert("Entered UserName : " + data.pass);
    alert("Result : " + this.result);

  }

}


export class LoginUserModel {
  constructor(
    public name: string,
    public pass: string,
  ) {
  }
}
