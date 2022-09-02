import { Injectable } from '@angular/core';
import { from, Observable, of } from 'rxjs';
import { catchError, switchMap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class ApiClientService {
  private readonly BASE_API_PATH = '/api';

  get<T>(url: string): Observable<T | null> {
    const req = fetch(`${this.BASE_API_PATH}/${url}`);

    return from(req).pipe(
      switchMap((req) => req.json()),
      catchError((e) => {
        // TODO(marchenkov): replace with shared logging service
        console.log(e);
        return of(null);
      })
    );
  }
}
