import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { Observable } from 'rxjs';

export interface Review {
  site_id: number;
  stars: number;
  text: string;
}

@Injectable({
  providedIn: 'root',
})
export class ReviewService {
  constructor(private http: HttpClient) {}

  uploadReview(review: Review): Observable<any> {
    return this.http.post(`${environment.API_URL}/reviews/upload`, review);
  }
}
