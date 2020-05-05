import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';

@Component({
  selector: 'app-update',
  templateUrl: './update.component.html',
  styleUrls: ['./update.component.css']
})
export class UpdateComponent implements OnInit {
  id = 0;
  fname = "";
  lname = "";
  email = "";

  constructor(public rest: RestService) { }

  ngOnInit(): void {
  }

  getValue() {
    console.log("pit");
    let postString = { "first_name": this.fname, "last_name": this.lname, "email": this.email };

    this.rest.update(this.id, JSON.stringify(postString)).subscribe((data: {}) => {
      console.log(data);
    });
  }

}

