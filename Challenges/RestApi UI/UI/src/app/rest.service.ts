import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { map, catchError, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class RestService {

  endpoint = 'http://localhost:9090/webapi/v1/user';
  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  };

  constructor(private http: HttpClient) { }

  postData(postString) {
    console.log(postString);
    return this.http.post(this.endpoint, postString);
  }


  extractData() {
    // return this.http.get(this.endpoint);
    return this.http.get(this.endpoint).pipe(map(res => res));
  }

  update(id, postString) {
    let url = this.endpoint + '/' + id.toString();
    return this.http.put(url, postString);
  }

}
