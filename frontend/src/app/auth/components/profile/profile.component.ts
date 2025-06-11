import { ChangeDetectionStrategy, Component } from '@angular/core';
import { AuthService } from '../../core/services/auth.service';

@Component({
  selector: 'app-profile',
  imports: [],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ProfileComponent {
  constructor(private auth: AuthService) {}

  logout() {
    this.auth.logout();
  }

  deleteAccount() {
    if (confirm('¿Seguro que deseas eliminar tu cuenta?')) {
      this.auth.deleteAccount().subscribe({
        next: () => alert('Cuenta eliminada'),
        error: (err) => alert(err.error.message || 'Error al eliminar cuenta'),
      });
    }
  }
}
