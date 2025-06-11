import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';

@Injectable({ providedIn: 'root' })
export class NoAuthGuard implements CanActivate {
  constructor(private router: Router) {}

  canActivate(): boolean {
    const token = localStorage.getItem('token');
    if (!token) {
      return true; // no está autenticado
    }
    this.router.navigate(['/map']); // si está autenticado, lo mando a mapa
    return false;
  }
}
