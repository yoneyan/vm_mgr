import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Router} from "@angular/router";
import {environment} from '../../../environments/environment';

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

  private result: string


  public verifyUser(data: any): string {
    const body: any = {
      user: data.user,
      pass: data.pass
    };
    console.log("user: " + body.user)
    console.log("pass: " + body.pass)

    this.http.post(environment.http + '://' + environment.APIHostIP + '/api/v1/token', body, this.defalutHttpOptions)
      .toPromise()
      .then((result: any) => {
        if (result.result === true) {
          this.r = true
          localStorage.setItem('id_token', result.token);
          localStorage.setItem('user', result.username)
          localStorage.setItem('userid', result.userid)
          console.log("Auth OK")
          console.log("Token: " + result.token)
          this.isLogin = true
          return location.href = "/dashboard"
          // return this.router.navigate(['/'])
        } else {
          this.r = false
          console.log("Auth NG")
          this.result = "Wrong username or password !!"
          // Promise.reject(false)
        }
      })
      .catch((err: any) => {
        return false
      })
    return this.result
  }

  logout() {
    localStorage.removeItem("user")
    localStorage.removeItem("userid")
    localStorage.removeItem("id_token")
    this.isLogin = false
    this.router.navigate(['/logout']);
  }

  public tokenCheck(): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.getAuthHeader(),
        // 'Accept': '*/*',
      })
    }
    const body: any = {};

    return this.http.post(environment.http+'://' + environment.APIHostIP + '/api/v1/token/check', body, httpOptions)
      .toPromise()
      .then((res) => {
        const response: any = res;
        console.log(response)
        return response.result
      })
      .catch((err) => {
          console.log('Error occured.', err);
          // return Promise.reject(err.message || err);
          return false
        }
      )
  }

  loginCheck(): boolean {
    return this.isLogin
  }

  getAuthUser(): string {
    return localStorage.getItem("user")
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
