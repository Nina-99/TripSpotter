import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';

@Injectable({ providedIn: 'root' })
export class LayerService {
  constructor(private http: HttpClient) {}

  getLayer(): Observable<any> {
    return this.http.get(`${environment.API_URL}/layers/`);
  }
}
