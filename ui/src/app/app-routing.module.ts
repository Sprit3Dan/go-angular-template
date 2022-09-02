import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './routes/home/home.component';
import { CanActivateHome } from './routes/home/can-activate-home';
import { HomeModule } from './routes/home/home.module';
import { AuthComponent } from './routes/auth/auth.component';

const routes: Routes = [
  { path: 'home', component: HomeComponent, canActivate: [CanActivateHome] },
  { path: 'auth', component: AuthComponent },
  { path: '**', redirectTo: 'home' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes), HomeModule],
  exports: [RouterModule],
})
export class AppRoutingModule {}
