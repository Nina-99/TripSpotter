import { HttpClient } from '@angular/common/http';
import { ChangeDetectionStrategy, Component } from '@angular/core';
import { environment } from '../../../../environments/environment';

@Component({
  selector: 'app-upload',
  imports: [],
  templateUrl: './upload.component.html',
  styleUrl: './upload.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UploadComponent {
  constructor(private http: HttpClient) {}
  uploadShapefile(event: any) {
    const files = event.target.files;
    const formData = new FormData();
    for (let file of files) {
      formData.append('files', file);
    }
    formData.append('layerName', 'capa1');

    this.http
      .post(`${environment.API_URL}/layers/upload`, formData)
      .subscribe((res) => {
        alert('Cargado correctamente');
      });
  }
}
