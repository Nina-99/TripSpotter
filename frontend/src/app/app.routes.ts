import { Routes } from '@angular/router';
import { LoginComponent } from './auth/components';
import { MapComponent } from './layers/components/map/map.component';
import { UploadComponent } from './layers/components/upload/upload.component';
import { RegisterComponent } from './auth/components/register/register.component';
import { AuthGuard } from './auth/core/guards';
import { NoAuthGuard } from './auth/core/guards/NoAuth.guard';
import { RoleGuard } from './auth/core/guards/role.guard';

export const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  // { path: 'login', component: LoginComponent, canActivate: [NoAuthGuard] },
  { path: 'login', component: LoginComponent },
  // { path: 'register', component: RegisterComponent, canActivate: [NoAuthGuard], },
  { path: 'register', component: RegisterComponent },
  { path: 'map', component: MapComponent, canActivate: [AuthGuard] },
  { path: 'uploads', component: UploadComponent, canActivate: [RoleGuard] },
  // { path: 'unauthorized', component: UnauthorizedComponent }
];
