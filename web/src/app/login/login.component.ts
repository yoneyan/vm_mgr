import {Component, OnInit} from '@angular/core';
import {FormBuilder} from "@angular/forms";
import { Router } from '@angular/router';

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

  test() {
    // this.loginService.test();
    window.alert('Your product has been added to the cart!');
  }

  title = 'test'

  constructor(
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private router: Router,
    // private loginservice: Test,
  ) { }

  ngOnInit() {
  }

  onClickSubmit(data) {
    let result: any
    result = this.authService.verifyUser(data)
    console.log(result)

    this.router.navigate(['/'])

    //
    //
    // onClickSubmit(value: any) {
    //   this.submitted = true;
    //   this.errormsg = undefined;
    //
    //   // ユーザ認証する
    //   this.loginservice.login(this.model)
    //     .then((token: string) => {
    //     })
    //     .catch((err: any) => {
    //       this.errormsg = 'Eメールアドレスまたはパスワードが違います。';
    //       console.log(err);
    //     });
    // }
  }
  // Login(): void {
  //   let result: any
  //   console.log(this.loginForm);
  //   // this.router.navigate(['/products']);
  //   result = this.authService.verifyUser(this.)
  //   console.log(result)
  //
  // }

}
