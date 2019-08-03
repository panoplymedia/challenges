import { Component, OnInit } from '@angular/core';
import { ApiService } from './api/api.service';

import { Observable, of } from 'rxjs';
import { map } from 'rxjs/operators';

import { Token } from './api/api.service';

import { DomSanitizer } from '@angular/platform-browser';
import { MatIconRegistry } from '@angular/material/icon';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'acme-cult-sales-portal';

  email: string = "";
  password: string = "";

  token: Observable<Token>;

  constructor(
    private apiService: ApiService,
    iconRegistry: MatIconRegistry,
    sanitizer: DomSanitizer) {
    iconRegistry.addSvgIcon('check', sanitizer.bypassSecurityTrustResourceUrl('assets/check.svg'));
  }

  ngOnInit() {
    this.token = this.apiService.GetToken();
  }

  login(email: string, password: string): void {
    this.token = this.apiService.Login(email, password).pipe(
      result => {
        this.email = "";
        this.password = "";
        return result;
      }
    );
  }
}
