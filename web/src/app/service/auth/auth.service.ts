import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Router} from "@angular/router";
import {resetFakeAsyncZone} from "@angular/core/testing";

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private r: boolean;

  constructor(
    private http: HttpClient,
    private router: Router
  ) {
  }

  public isLogin = false

  private token: string
  const;
  defalutHttpOptions = {
    headers: new HttpHeaders("application/json")
  }

  private result:string


  public verifyUser(data: any): string {
    const body: any = {
      user: data.name,
      pass: data.pass
    };
    console.log("user: " + body.user)
    console.log("pass: " + body.pass)


    this.http.post('http://localhost:8081/api/v1/token', body, this.defalutHttpOptions)
      .toPromise()
      .then((result: any) => {
        if (result.result === true) {
          this.r = true
          localStorage.setItem('id_token', result.token);
          localStorage.setItem('name', result.name)
          console.log("Auth OK")
          console.log("Token: " + result.token)
          this.isLogin = true
          return location.href="/"
          // return this.router.navigate(['/'])
        } else {
          this.r = false
          console.log("Auth NG")
          this.result =  "Pass worng"
          // Promise.reject(false)
        }
      })
      .catch((err: any) => {
        return false
      })
    return this.result
  }

  logout() {
    localStorage.removeItem("name")
    localStorage.removeItem("id_token")
    this.isLogin = false
    this.router.navigate([ '/logout' ]);
  }

  public tokenCheck(): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.getAuthHeader(),
        // 'Accept': '*/*',
      })
    }
    const body: any = {
    };
    return this.http.post('http://localhost:8081/api/v1/token/check', body, httpOptions)
      .toPromise()
      .then((res) => {
        const response: any = res;
        console.log(response)
        return response.status
      })
      .catch((err)=> {
        console.log('Error occured.', err);
        // return Promise.reject(err.message || err);
        return false
      }
  )
  }

  loginCheck(): boolean {
    return this.isLogin
  }

  getAuthUser(): string{
    return localStorage.getItem("name")
  }

  getAuthHeader(): string {
    const token = localStorage.getItem('id_token');

    if (token) {
      return 'Bearer ' + token;
    } else {
      undefined;
    }
  }


}
