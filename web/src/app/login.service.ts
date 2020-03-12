import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LoginService {
  name = "a";
  pass = "";
  // test(){
  //   return 0;
  // }
  verifyUser(data){
    this.name = data.name
    this.pass = data.pass
    if (this.name == this.pass){
      return true
    }else{
      return false
    }
  }
}
