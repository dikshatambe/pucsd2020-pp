import { Component } from '@angular/core';
import { environment } from '../environments/environment';
import { Router } from '@angular/router';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'project';
  loginSuccess = false;
  firstLogin = true;
  username = "";
  password = "";

  getEnable = false;
  putEnable = false;
  delEnable = false;
  postEnable = false;
  mainEnable = true;

  operationEnable = false;


  GetMethod() {
    console.log("GET");
    this.getEnable = true;
    this.putEnable = false;
    this.postEnable = false;
    this.delEnable = false;
  }

  postMethod() {
    console.log("POST");
    this.getEnable = false;
    this.putEnable = false;
    this.postEnable = true;
    this.delEnable = false;
  }

  putMethod() {
    console.log("PUT");
    this.getEnable = false;
    this.putEnable = true;
    this.postEnable = false;
    this.delEnable = false;
  }

  deleteMethod() {
    console.log("DELETE");
    this.getEnable = false;
    this.putEnable = false;
    this.postEnable = false;
    this.delEnable = true;
  }

  goPostComponent() {
    this.mainEnable = false;
  }

  performOperation() {
    this.operationEnable = true;
  }

  login() {
    this.loginSuccess = true;
    this.firstLogin = false;
  }

}
