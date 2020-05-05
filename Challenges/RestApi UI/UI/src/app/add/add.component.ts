import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';

@Component({
  selector: 'app-add',
  templateUrl: './add.component.html',
  styleUrls: ['./add.component.css']
})
export class AddComponent implements OnInit {
  fname = "";
  lname = "";
  email = "";


  constructor(public rest: RestService) { }

  ngOnInit(): void {
  }

  getValue() {
    console.log("GET");
    let postString = { "first_name": this.fname, "last_name": this.lname, "email": this.email }

    this.rest.postData(JSON.stringify(postString)).subscribe((data: {}) => {
      console.log(data);
    });
  }
}
