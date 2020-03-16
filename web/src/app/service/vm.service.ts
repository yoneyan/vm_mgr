import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {AuthService} from "./auth/auth.service";

@Injectable({
  providedIn: 'root'
})
export class VmService {

  constructor(private http: HttpClient,
              private auth: AuthService) {
  }

  public getUserVM(): Promise<any> {
    let url: string = "http://localhost:8080/api/v1/vm"
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.auth.getAuthHeader(),
        // 'Accept': '*/*',
      })
    }
    console.log(url)


    return this.http.get<any>(url, httpOptions)
      .toPromise()
      .then((r) => {
        return r
      })

  }
}
