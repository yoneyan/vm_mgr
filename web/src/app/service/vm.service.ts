import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {AuthService} from "./auth/auth.service";
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class VmService {

  constructor(private http: HttpClient,
              private auth: AuthService) {
  }

  public getUserVM(): Promise<any> {
    let url: string = environment.http + "://" + environment.APIHostIP + "/api/v1/vm"
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

  public getVM(id): Promise<any> {
    let url: string = environment.http + "://" + environment.APIHostIP + "/api/v1/vm/" + id
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

  public createVM(nodeid, vmname, group, cpu, mem, storage, storagetype, imagename, imagetag): Promise<any> {
    let url: string = environment.http + "://" + environment.APIHostIP + "/api/v1/vm"
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.auth.getAuthHeader(),
        // 'Accept': '*/*',
      })
    }
    console.log(cpu)
    console.log(mem)
    console.log(storage)
    let body: any = {
      nodeid: parseInt(nodeid),
      vmname: vmname,
      group: group,
      cpu: parseInt(cpu),
      mem: parseInt(mem),
      storage: parseInt(storage),
      storagetype: parseInt(storagetype),
      autostart: 1,
      imagename: imagename,
      imagetype: imagetag
    };

    return this.http.put<any>(url, body, httpOptions)
      .toPromise()
      .then((r) => {
        return r
      })
  }

  public startVM(id): Promise<any> {
    let url: string = environment.http + "://" + environment.APIHostIP + "/api/v1/vm/" + id + "/power"
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.auth.getAuthHeader(),
        // 'Accept': '*/*',
      })
    }
    const body: any = {};

    return this.http.put<any>(url, body, httpOptions)
      .toPromise()
      .then((r) => {
        return r
      })
  }

  public stopVM(id): Promise<any> {
    let url: string = environment.http + "://" + environment.APIHostIP + "/api/v1/vm/" + id + "/power"
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.auth.getAuthHeader(),
      }),
      body: {
        force: true
      }
    }

    return this.http.delete<any>(url, httpOptions)
      .toPromise()
      .then((r) => {
        return r
      })
  }

  public shutdownVM(id): Promise<any> {
    let url: string = environment.http + "://" + environment.APIHostIP + "/api/v1/vm/" + id + "/power"
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.auth.getAuthHeader(),
      }),
      body: {
        force: false
      }
    }

    return this.http.delete<any>(url, httpOptions)
      .toPromise()
      .then((r) => {
        return r
      })
  }

  public resetVM(id): Promise<any> {
    let url: string = environment.http + "://" + environment.APIHostIP + "/api/v1/vm/" + id + "/reset"
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.auth.getAuthHeader(),
        // 'Accept': '*/*',
      })
    }
    const body: any = {};


    return this.http.put<any>(url, body, httpOptions)
      .toPromise()
      .then((r) => {
        return r
      })
  }

  public pauseVM(id): Promise<any> {
    let url: string = environment.http + "://" + environment.APIHostIP + "/api/v1/vm/" + id + "/pause"
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.auth.getAuthHeader(),
        // 'Accept': '*/*',
      })
    }
    const body: any = {};

    return this.http.put<any>(url, body, httpOptions)
      .toPromise()
      .then((r) => {
        return r
      })
  }

  public resumeVM(id): Promise<any> {
    let url: string = environment.http + "://" + environment.APIHostIP + "/api/v1/vm/" + id + "/pause"
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
        'Authorization': this.auth.getAuthHeader(),
        // 'Accept': '*/*',
      })
    }

    return this.http.delete<any>(url, httpOptions)
      .toPromise()
      .then((r) => {
        return r
      })
  }

}
