import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { BaseModal } from 'src/models/base.model';
import { TimeTableModelObject } from 'src/models/time-table.model';

@Injectable({
  providedIn: 'root'
})
export class TimeTableService {

  constructor(private _httpClient: HttpClient) { }

  getTimetable(): Observable<BaseModal<number,TimeTableModelObject[]> | null> {
    return this._httpClient.get<BaseModal<number,TimeTableModelObject[]> | null>(`${environment.apiUrl}`)
  }
}
