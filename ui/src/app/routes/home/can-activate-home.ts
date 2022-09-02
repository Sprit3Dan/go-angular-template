import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivate,
  Router,
  RouterStateSnapshot,
  UrlTree,
} from '@angular/router';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { UserInfoService } from '../../services/user-info.service';

@Injectable({
  providedIn: 'root',
})
export class CanActivateHome implements CanActivate {
  constructor(
    private readonly router: Router,
    private readonly userInfoService: UserInfoService
  ) {}

  canActivate(
    _: ActivatedRouteSnapshot,
    __: RouterStateSnapshot
  ): Observable<UrlTree | boolean> {
    return this.userInfoService.isAuthenticated().pipe(
      map((isAuthenticated) => {
        if (isAuthenticated) {
          return true;
        }

        return this.router.createUrlTree(['/auth']);
      })
    );
  }
}
