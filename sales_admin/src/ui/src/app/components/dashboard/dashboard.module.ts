import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { MatModule } from '../../angular-material.module';

import { DashboardComponent } from './dashboard.component';
import { DialogComponent } from './dialog/dialog.component';

@NgModule({
    declarations: [
        DashboardComponent,
        DialogComponent
    ],
    imports: [
        MatModule,
        CommonModule
    ],
    exports: [DashboardComponent],
    entryComponents: [DialogComponent]
})

export class DashboardModule { }