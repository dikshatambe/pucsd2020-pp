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

  addEnable =false;
  delEnable =false;
  upEnable =false;
  serEnable =false;

  add()
  {
    this.addEnable =true;
    this.delEnable =false;
    this.upEnable =false;
    this.serEnable =false;
  }

  delete()
  {
    this.addEnable =false;
    this.delEnable =true;
    this.upEnable =false;
    this.serEnable =false;
  }

  update()
  {
    this.addEnable =false;
    this.delEnable =false;
    this.upEnable =true;
    this.serEnable =false;
  }

  search()
  {
    this.addEnable =false;
    this.delEnable =false;
    this.upEnable =false;
    this.serEnable =true;
  }

}
