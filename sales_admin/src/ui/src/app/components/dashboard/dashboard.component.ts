import { Component, OnInit, Input } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';

import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

import { ApiService } from '../../api/api.service';
import { DialogComponent } from './dialog/dialog.component';

@Component({
    selector: 'dashboard-component',
    templateUrl: './dashboard.component.html'
})

export class DashboardComponent implements OnInit {
    salesRevenue: Observable<number>;

    constructor(
        private apiService: ApiService,
        public dialog: MatDialog
    ) { }

    ngOnInit() {
        this.salesRevenue = this.getSalesRevenue();
    }

    public openUploadDialog() {
        let dialogRef = this.dialog.open(DialogComponent, { width: '50%', height: '50%' });

        dialogRef.afterClosed().subscribe(() => {
            this.salesRevenue = this.getSalesRevenue();
        });
    }

    getSalesRevenue(): Observable<number> {
        return this.apiService.GetSalesSummary().pipe(
            map(result => result.totalSalesRevenue)
        );
    }
}