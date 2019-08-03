import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders, HttpEventType, HttpResponse } from '@angular/common/http';

import { Observable, of, Subject } from 'rxjs';
import { catchError, map } from 'rxjs/operators'

import { environment } from '../../environments/environment';

export interface Token {
    token: string;
}

export interface SalesSummary {
    totalSalesRevenue: number;
}

@Injectable({
    providedIn: 'root'
})
export class ApiService {
    private token: Token;

    constructor(private client: HttpClient) { }

    Login(email: string, password: string): Observable<Token> {
        return this.client.post<Token>(environment.apiBaseUrl + '/login', null, {
            params: {
                "email": email,
                "password": password,
            }
        }).pipe(
            map(resp => {
                this.token = resp;
                localStorage.setItem('currentToken', JSON.stringify(resp));
                return resp;
            }),
            catchError(err => {
                console.log(err);
                return of(undefined);
            })
        );
    }


    GetToken(): Observable<Token> {
        let token = localStorage.getItem('currentToken');
        if (!!token) {
            this.token = JSON.parse(token);
            return of(this.token);
        }
        return of(undefined);
    }

    UploadCsv(file: File): Observable<number> {
        let status: Observable<number>;
        let formData: FormData = new FormData();
        formData.append('file', file, file.name);

        let progress = new Subject<number>();

        let headers = new HttpHeaders();
        headers = headers.set("Authorization", "Bearer " + this.token.token);

        let req = new HttpRequest('POST', environment.apiBaseUrl + '/sale/upload', formData, {
            reportProgress: true,
            headers: headers
        });

        this.client.request(req).subscribe(event => {
            if (event.type === HttpEventType.UploadProgress) {
                const percentDone = Math.round(100 * event.loaded / event.total);
                progress.next(percentDone);
            } else if (event instanceof HttpResponse) {
                progress.complete();
            }
        });

        status = progress.asObservable();

        return status;
    }

    GetSalesSummary(): Observable<SalesSummary> {
        return this.client.get<SalesSummary>(environment.apiBaseUrl + '/sale/summary', {
            headers: {
                "Authorization": 'Bearer ' + this.token.token
            },
            reportProgress: true
        }).pipe(
            map(resp => {
                return resp;
            }),
            catchError(err => {
                console.log(err);
                return of({ totalSalesRevenue: 0 });
            })
        )
    }
}