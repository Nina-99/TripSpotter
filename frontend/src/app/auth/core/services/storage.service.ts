import { Injectable } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class StorageService {
  isBrowser = typeof window !== 'undefined';

  get(key: string): string | null {
    return this.isBrowser ? localStorage.getItem(key) : null;
  }

  set(key: string, value: string) {
    if (this.isBrowser) localStorage.setItem(key, value);
  }

  remove(key: string) {
    if (this.isBrowser) localStorage.removeItem(key);
  }
}
