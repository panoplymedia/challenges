import { NgModule } from '@angular/core';

import { OverlayModule } from '@angular/cdk/overlay';
import { PortalModule } from '@angular/cdk/portal';

import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatRadioModule } from '@angular/material/radio';
import { MatDialogModule } from '@angular/material/dialog';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatTabsModule } from '@angular/material/tabs';
import { MatCardModule } from '@angular/material/card';
import { MatTableModule } from '@angular/material/table';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { MatSelectModule } from '@angular/material/select';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatProgressBarModule } from '@angular/material/progress-bar';

@NgModule({
    imports: [
        FlexLayoutModule,
        PortalModule,
        OverlayModule,
        MatInputModule,
        MatButtonModule,
        MatDialogModule,
        MatRadioModule,
        MatListModule,
        MatIconModule,
        MatDividerModule,
        MatTabsModule,
        MatCardModule,
        MatTableModule,
        MatFormFieldModule,
        MatButtonToggleModule,
        MatSelectModule,
        MatTooltipModule,
        MatPaginatorModule,
        MatToolbarModule,
        MatProgressBarModule
    ],
    exports: [
        FlexLayoutModule,
        PortalModule,
        OverlayModule,
        MatInputModule,
        MatButtonModule,
        MatDialogModule,
        MatRadioModule,
        MatListModule,
        MatIconModule,
        MatDividerModule,
        MatTabsModule,
        MatCardModule,
        MatTableModule,
        MatFormFieldModule,
        MatButtonToggleModule,
        MatSelectModule,
        MatTooltipModule,
        MatPaginatorModule,
        MatToolbarModule,
        MatProgressBarModule
    ]
})
export class MatModule { }