import {Component, OnInit} from '@angular/core';
import {FormBuilder} from "@angular/forms";

import {LoginService} from "../login.service";
import {AuthService, Login} from "../auth.service";

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
    private loginService: LoginService,
    // private loginservice: AuthService,
  ) {


    // this.loginForm = this.formBuilder.group({
    //   name: '',
    //   pass: '',
    // });
  }

  ngOnInit() {
  }

  onClickSubmit(data) {
    let result: any
    result = this.loginService.verifyUser(data)
    console.log(result)

    if (result.__zone_symbol__value === false) {

      console.log(result)
      // alert("Entered UserName : " + data.pass);
      // alert("Result : " + result);
    }
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
}
