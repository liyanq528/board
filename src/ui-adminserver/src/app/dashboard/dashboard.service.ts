import { Injectable } from '@angular/core';
import { HttpClient, HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { ComponentStatus } from './component-status.model';

const BASE_URL = '/v1/admin';

@Injectable()
export class DashboardService {

  constructor(private http: HttpClient) {
  }

  monitorContainer(): Observable<any> {
    const token = window.sessionStorage.getItem('token');
    return this.http.get(
      `${BASE_URL}/monitor?token=${token}`,
      { observe: 'response', })
      .pipe(map((res: HttpResponse<Array<ComponentStatus>>) => {
        return res.body;
      }));
  }
}
