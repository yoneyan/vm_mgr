import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {AuthService} from "./auth/auth.service";
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class GroupService {

  constructor(
    private http: HttpClient,
    private auth: AuthService
  ) {
  }

  public getGroup(): Promise<any> {
    let url: string = environment.http + "://" + environment.APIHostIP + "/api/v1/group"
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
