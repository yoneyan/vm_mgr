import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable, throwError} from "rxjs";
import {timeout, catchError} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})

export class LoginService {

  // test(){
  //   return 0;
  // }
  private r: boolean;

  constructor(private http: HttpClient) {
  }

  private token: string

  public verifyUser(data: any): any {
    const body: any = {
      user: data.name,
      pass: data.pass
    };
    console.log("user: " + body.name)
    console.log("pass: " + body.pass)


    const httpOptions = {
      headers: new HttpHeaders("application/json")
    }

    this.http.post('http://localhost:8081/api/v1/token', body, httpOptions)
      .toPromise()
      .then((result: any) => {
        if (result.result === true) {
          this.r = true
          localStorage.setItem('id_token', result.token);
          console.log("Auth OK")
          console.log("Token: " + result.token)
        } else {
          this.r = false
          console.log("Auth NG")
          return "NG"
        }
      })
      .catch((err: any) => {
        return false
      })
  }

  logout() {
    localStorage.remoteItem('id_token');
  }

}
