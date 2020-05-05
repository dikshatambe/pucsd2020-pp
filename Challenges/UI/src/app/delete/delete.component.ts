import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';

@Component({
  selector: 'app-delete',
  templateUrl: './delete.component.html',
  styleUrls: ['./delete.component.css']
})
export class DeleteComponent implements OnInit {

  delByIdEnable = false;
  delAllEnable = false;
  id = 0;

  constructor(public rest: RestService) { }

  ngOnInit(): void {
  }

  deleteAll() {
    this.delAllEnable = true;
    this.delByIdEnable = false;
    this.rest.extractData().subscribe((data: {}) => {
      this.createDatasource(data);
    });
  }

  deleteById() {
    this.delByIdEnable = true;
    this.delAllEnable = false;
  }

  createDatasource(data) {

  }

  getValue() {

  }


}

