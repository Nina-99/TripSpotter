import { ChangeDetectionStrategy, Component } from '@angular/core';
import { AuthService } from '../../core/services/auth.service';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { EmailService } from '../../../layers/services/email.service';

@Component({
  selector: 'app-login',
  imports: [FormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginComponent {
  email = '';
  password = '';

  constructor(
    private auth: AuthService,
    private router: Router,
    private emailService: EmailService,
  ) {}

  onLogin() {
    this.emailService.setEmail(this.email);
    this.auth.initializeUserRole(this.email);
    this.auth.login({ email: this.email, password: this.password }).subscribe({
      next: () => this.router.navigate(['/map']),
      // this.auth.saveToken(res.token);
      error: (err) => alert(err.error.message || 'Login Falied'),
    });
  }
}
