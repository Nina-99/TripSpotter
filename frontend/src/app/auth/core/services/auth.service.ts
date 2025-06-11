import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Inject, Injectable, PLATFORM_ID } from '@angular/core';
import { Router } from '@angular/router';
import { BehaviorSubject, Observable, tap } from 'rxjs';
import { environment } from '../../../../environments/environment';
import { StorageService } from './storage.service';
import { isPlatformBrowser } from '@angular/common';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private tokenKey = 'auth_token';
  private isBrowser: boolean;
  private apiUrl = environment.API_URL;
  public isLoggedIn$ = new BehaviorSubject<boolean>(false);

  constructor(
    private http: HttpClient,
    private router: Router,
    @Inject(PLATFORM_ID) private platformId: Object,
  ) {
    this.isBrowser = isPlatformBrowser(platformId);
  }

  private hasToken(): boolean {
    return this.isBrowser && !!localStorage.getItem(this.tokenKey);
  }

  login(credentials: { email: string; password: string }): Observable<any> {
    return this.http.post(`${this.apiUrl}/login`, credentials).pipe(
      tap((res: any) => {
        if (this.isBrowser) {
          localStorage.setItem(this.tokenKey, res.token);
        }
        this.isLoggedIn$.next(true);
      }),
    );
  }
  register(data: {
    username: string;
    email: string;
    password: string;
  }): Observable<any> {
    return this.http.post(`${this.apiUrl}/register`, data);
  }

  getToken(): string | null {
    return this.isBrowser ? localStorage.getItem(this.tokenKey) : null;
  }

  getUserRole(email: string) {
    const token = localStorage.getItem('token');
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);

    return this.http.get<any>(`${environment.API_URL}/role/${email}`, {
      headers,
    });
  }

  // Método de ejemplo para llamar después del login
  initializeUserRole(email: string) {
    this.getUserRole(email).subscribe((user) => {
      localStorage.setItem('role', user.role);
      localStorage.setItem('username', user.username);
    });
  }

  getRole(): string | null {
    if (typeof window !== 'undefined') {
      return localStorage.getItem('role');
    }
    return null;
  }
  getUsername(): string | null {
    return localStorage.getItem('username');
  }

  isAdmin(): boolean {
    return this.getRole() === 'admin';
  }

  isUser(): boolean {
    return this.getRole() === 'user';
  }

  logout() {
    if (this.isBrowser) {
      localStorage.removeItem(this.tokenKey);
    }
    this.isLoggedIn$.next(false);
    this.router.navigate(['/login']);
  }

  deleteAccount(): Observable<any> {
    return this.http
      .delete(`${this.apiUrl}/delete`, {
        headers: { Authorization: `Bearer ${this.getToken()}` },
      })
      .pipe(tap(() => this.logout()));
  }
}
