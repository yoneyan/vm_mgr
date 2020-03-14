import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class AuthService {
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

    return this.http.post('http://localhost:8081/api/v1/token', body, httpOptions)
      .toPromise()
      .then((result: any) => {
        if (result.result === true) {
          this.r = true
          localStorage.setItem('id_token', result.token);
          localStorage.setItem('name', result.name)
          console.log("Auth OK")
          console.log("Token: " + result.token)
          alert("ok")
          alert("Token: " + result.token)
        } else {
          this.r = false
          console.log("Auth NG")
          alert("NG")
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

  logincheck(): Promise<boolean> {
    return new Promise(function(resolve, reject){
      reject(false)
      // resolve(true)

    })
  }


}
