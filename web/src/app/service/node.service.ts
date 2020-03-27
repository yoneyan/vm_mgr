import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {AuthService} from "./auth/auth.service";
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class NodeService {

  constructor(
    private http: HttpClient,
    private auth: AuthService
  ) {
  }

  public getNode(): Promise<any> {
    let url: string = "http://" + environment.APIHostIP + ":8080/api/v1/node"
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.auth.getAuthHeader(),
        // 'Accept': '*/*',
      })
    }
    return this.http.get<any>(url, httpOptions)
      .toPromise()
      .then((r) => {
        return r
      })
  }
}
