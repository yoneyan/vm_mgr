import { Injectable } from '@angular/core';
import {environment} from "../../environments/environment";
import {HttpClient, HttpHeaders} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class SpecService {

  constructor(
    private http: HttpClient

  ) { }

  public getCPU(): Promise<any> {
    return this.http.get<any>("../assets/cpu.json")
      .toPromise()
      .then((r) => {
        return r
      })
  }

  public getMemory(): Promise<any> {
    return this.http.get<any>("../assets/memory.json")
      .toPromise()
      .then((r) => {
        return r
      })
  }

  public getStorage(): Promise<any> {
    return this.http.get<any>("../assets/storage.json")
      .toPromise()
      .then((r) => {
        return r
      })
  }

  public getStorageSize(): Promise<any> {
    return this.http.get<any>("../assets/storage_size.json")
      .toPromise()
      .then((r) => {
        return r
      })
  }

}
