import { Component, OnInit } from '@angular/core';
import { RestService } from '../rest.service';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit {
  getAllEnable = false;
  getByIdEnable = false;
  myJSON = "";

  displayedColumns: string[] = ['id', 'first_name', 'last_name', 'email'];

  dataSource = [
  ];

  constructor(public rest: RestService) { }

  ngOnInit(): void {
  }

  getAll() {
    this.getAllEnable = true;
    this.getByIdEnable = false;
    this.dataSource = [
      { "id": "19", "first_name": "Ankita", "last_name": "Mishra", "email": "anky@gmail.com" },
      { "id": "20", "first_name": "Chinmay", "last_name": "Kale", "email": "chkale11@gmail.com" },
      { "id": "21", "first_name": "Saavi", "last_name": "Sirsat", "email": "saavi20@gmail.com" },
      { "id": "22", "first_name": "Tilak", "last_name": "Pawar", "email": "tp@gmail.com" },
    ];

  }

  getById() {
    this.getAllEnable = false;
    this.getByIdEnable = true;
    this.rest.extractData().subscribe((data: {}) => {
      this.createDatasource(data);
    });
  }

  createDatasource(data) {
    console.log(data.data[0].id);

    this.myJSON = JSON.stringify(data.data[0]);
    console.log(this.myJSON);
    let key = Object.keys(data.data[0]);
    let value = Object.values(data.data[0]);
    var d = {}
    for (var i = 0; i < key.length; i++) {
      d[key[i]] = value[i];
    }
    this.dataSource = [d];
    console.log(typeof (this.dataSource));
  }

  //applyFilter(event: Event) {
  //var dataSource = new MatTableDataSource(this.dataSource);
  //var filterValue = (event.target as HTMLInputElement).value;
  //dataSource.filter = filterValue.trim().toLowerCase();
  //}

}

