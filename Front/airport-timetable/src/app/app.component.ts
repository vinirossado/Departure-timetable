import { Component, OnInit } from '@angular/core';
import { BaseModal } from 'src/models/base.model';
import { TimeTableModelObject } from 'src/models/time-table.model';
import { TimeTableService } from 'src/services/time-table.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'airport-timetable';

  infs!: TimeTableModelObject[];
  constructor(private _timeTableService: TimeTableService) {
  }

  ngOnInit(): void {
    this._timeTableService.getTimetable().subscribe((x: BaseModal<number,TimeTableModelObject[]> | null) => {
      console.log(x);
      this.infs = x?.data;
    });
  }

  identify = (index: number, item: any): string => item.index;

}
