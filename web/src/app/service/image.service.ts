import {Injectable} from '@angular/core';
import {environment} from "../../environments/environment";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {AuthService} from "./auth/auth.service";

@Injectable({
  providedIn: 'root'
})
export class ImageService {

  constructor(
    private http: HttpClient,
    private auth: AuthService
  ) {
  }

  public getImage(): Promise<any> {
    let url: string = "http://" + environment.APIHostIP + ":8080/api/v1/image"
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
