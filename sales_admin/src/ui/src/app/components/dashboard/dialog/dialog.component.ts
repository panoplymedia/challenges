import { Component, ViewChild, ElementRef } from '@angular/core'
import { MatDialogRef } from '@angular/material/dialog';

import { Observable } from 'rxjs';

import { ApiService } from '../../../api/api.service';

@Component({
  selector: 'upload-dialog',
  templateUrl: './dialog.component.html',
  styleUrls: ['./dialog.component.scss'],
})
export class DialogComponent {
  @ViewChild('file', { static: false }) file: ElementRef;

  progress: Observable<number>;
  canBeClosed = true;
  primaryButtonText = 'Upload';
  showCancelButton = true;
  uploading = false;
  uploadSuccessful = false;

  public selectedFile: File;

  constructor(
    public dialogRef: MatDialogRef<DialogComponent>,
    private apiService: ApiService
  ) { }

  addFile() {
    this.file.nativeElement.click();
  }

  onFileAdded() {
    this.selectedFile = this.file.nativeElement.files[0];
  }

  closeDialog() {
    if (this.uploadSuccessful) {
      return this.dialogRef.close();
    }

    this.uploading = true;

    this.progress = this.apiService.UploadCsv(this.selectedFile);

    this.primaryButtonText = 'Finish';

    this.canBeClosed = false;

    this.dialogRef.disableClose = true;

    this.showCancelButton = false;

    this.progress.subscribe(() => {
      this.canBeClosed = true;
      this.dialogRef.disableClose = false;

      this.uploadSuccessful = true;

      this.uploading = false;
    });
  }
}