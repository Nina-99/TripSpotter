import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class EmailService {
  private emailLabelSubject = new BehaviorSubject<string>('');

  emailLabel$ = this.emailLabelSubject.asObservable();

  setEmail(emailLabel: string) {
    this.emailLabelSubject.next(emailLabel);
  }
}
