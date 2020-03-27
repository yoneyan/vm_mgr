import {Injectable} from '@angular/core';
import {CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree, Router} from '@angular/router';
import {Observable} from 'rxjs';
import {AuthService} from "../service/auth/auth.service";
import {map} from "rxjs/operators";

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private auth: AuthService,
              private router: Router) {
  }

  // canActivate(
  //   next: ActivatedRouteSnapshot,
  //   state: RouterStateSnapshot): Observable<boolean> | Promise<boolean> | boolean {
  //   return true;

  async canActivate(next: ActivatedRouteSnapshot,
                    state: RouterStateSnapshot): Promise<boolean> {
    return this.auth.tokenCheck()
      .then((data)=>{
        if (data == true){
          console.log("TRUE")
          this.auth.isLogin = true
          return true
        }else{
          console.log("FALSE")
          this.auth.logout()
          return false
        }
      })
  }


  // canActivate(
  //   next: ActivatedRouteSnapshot,
  //   state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
  //   return true;
  // }

}
