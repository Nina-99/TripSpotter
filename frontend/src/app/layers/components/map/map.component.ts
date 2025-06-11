import {
  Component,
  OnInit,
  Inject,
  NgZone,
  ChangeDetectorRef,
  PLATFORM_ID,
} from '@angular/core';
import { isPlatformBrowser } from '@angular/common';
import { CommonModule } from '@angular/common';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { WeatherService } from '../../services/weather.service';
import { LayerService } from '../../services/layer.service';
import { EmailService } from '../../services/email.service';
import { getPopupService } from '../../services/popup.service';
import { environment } from '../../../../environments/environment';
import { ReviewService } from '../../services/review.service';
import { HttpClient } from '@angular/common/http';
import { AuthService } from '../../../auth/core/services/auth.service';

@Component({
  selector: 'app-map',
  standalone: true,
  imports: [CommonModule, FormsModule, ReactiveFormsModule],
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.scss'],
})
export class MapComponent implements OnInit {
  private map!: any;
  private layerControl!: any;
  private pruneClusterGroup: any;

  private isBrowser: boolean;
  private L!: any;
  private apiKey = environment.LOCATIONIQ_API_KEY;

  forecastList: any[] = [];
  locationLabel = '';
  selectedFeature: any = null;
  selectedFeatureContent = '';
  usernameLabel: string | null = null;
  featurePopup = '';
  private leyendaControl: any;
  reviewStars: number = 0;
  hoverStars: number = 0;
  reviewText: string = '';
  selectedSiteId: number = 0;
  markerId: number = 0;
  role: string | null = null;

  constructor(
    private http: HttpClient,
    private weatherService: WeatherService,
    private layerService: LayerService,
    private popupService: getPopupService,
    private reviewService: ReviewService,
    private ngZone: NgZone,
    private cdr: ChangeDetectorRef,
    @Inject(PLATFORM_ID) platformId: Object,
  ) {
    this.isBrowser = isPlatformBrowser(platformId);
  }

  ngOnInit(): void {
    if (!this.isBrowser) return;
    if (typeof window !== 'undefined') {
      this.role = localStorage.getItem('role');
    }

    if (typeof window !== 'undefined') {
    // this.role = localStorage.getItem('role');
      this.usernameLabel = localStorage.getItem('username')
    }
    // this.emailService.emailLabel$.subscribe(
    //   (label) => (this.emailLabel = label),
    // );

    // Importar Leaflet dinámicamente solo en navegador
    import('leaflet').then((mod) => {
      this.L = mod;
      (window as any).L = mod;
      Promise.all([
        this.loadScript('assets/locationiq/leaflet-geocoder-locationiq.min.js'),
        this.loadStylesheet(
          'assets/locationiq/leaflet-geocoder-locationiq.min.css',
        ),
        this.loadStylesheet('assets/libs/leaflet-simple-locate.css'),
        this.loadScript('assets/libs/leaflet-simple-locate.min.js'),
        this.loadScript('assets/libs/leaflet.legend.js'),
        this.loadStylesheet('assets/libs/leaflet.legend.css'),
        this.loadScript('assets/libs/PruneCluster.js'),
        this.loadStylesheet('assets/libs/LeafletStyleSheet.css'),
        this.loadScript(
          'assets/libs/leaflet-sidebar-v2/js/leaflet-sidebar.min.js',
        ),
        this.loadStylesheet(
          'assets/libs/leaflet-sidebar-v2/css/leaflet-sidebar.min.css',
        ),
        this.loadScript(
          'assets/libs/Leaflet.SidePanel/leaflet-sidepanel.min.js',
        ),
        this.loadStylesheet(
          'assets/libs/Leaflet.SidePanel/leaflet-sidepanel.css',
        ),
      ]).then(() => this.setupMap());
    });
  }

  private loadScript(src: string): Promise<void> {
    return new Promise((resolve) => {
      const script = document.createElement('script');
      script.src = src;
      script.onload = () => resolve();
      document.body.appendChild(script);
    });
  }

  private loadStylesheet(href: string): void {
    const link = document.createElement('link');
    link.rel = 'stylesheet';
    link.href = href;
    document.head.appendChild(link);
  }

  private setupMap(): void {
    if (!this.L) return;
    const L = this.L;
    const key = this.apiKey;

    // const streets = L.tileLayer.Unwired({
    //   key,
    //   scheme: 'streets',
    //   crossOrigin: false,
    // });
    const satellite = L.tileLayer(
      'https://tiles.stadiamaps.com/tiles/alidade_satellite/{z}/{x}/{y}{r}.{ext}',
      {
        maxZoom: 19,
        minZoom: 0,
        attribution:
          '&copy; CNES, Airbus DS | &copy; Stadia Maps | &copy; OpenStreetMap contributors',
        ext: 'jpg',
        crossOrigin: false,
      },
    );

    this.map = L.map('map', {
      center: [-17.38, -66.16],
      zoom: 11,
      layers: [satellite],
      zoomControl: false,
    });

    L.control.scale().addTo(this.map);
    const control = new L.Control.SimpleLocate({
      position: 'bottomright',
      className: 'button-locate',
      afterMarkerAdd: () => {
        console.log('afterMarkerAdded');
        const elem = document.getElementById('leaflet-simple-locate-icon-spot');
        if (elem) {
          elem.addEventListener('click', (event) => {
            const latlng = control.getLatLng();
            const latlng_str = `geolocation: [${Math.round(latlng.lat * 100000) / 100000}, ${Math.round(latlng.lng * 100000) / 100000}]`;

            const accuracy = control.getAccuracy();
            const accuracy_str = `accuracy: ${Math.round(accuracy)} meter`;

            const angle = control.getAngle();
            const angle_str = `orientation: ${Math.round(angle)} degree`;

            L.popup()
              .setLatLng(latlng)
              .setContent(
                `<p style="margin: 0.25rem 0 0.25rem 0">${latlng_str}</p><p style="margin: 0.25rem 0 0.25rem 0">${accuracy_str}</p><p style="margin: 0.25rem 0 0.25rem 0">${angle_str}</p>`,
              )
              .openOn(this.map);

            event.stopPropagation();
            event.preventDefault();
          });
        }
      },
    }).addTo(this.map);

    L.control
      .geocoder(key, {
        url: 'https://api.locationiq.com/v1',
        expanded: true,
        panToPoint: true,
        focus: true,
        position: 'topright',
      })
      .addTo(this.map);

    this.layerControl = L.control
      .layers({ Satélite: satellite }, undefined, {
        position: 'bottomleft',
      })
      .addTo(this.map);

    this.map.on('click', (e: L.LeafletMouseEvent) => {
      const { lat, lng } = e.latlng;
      this.ngZone.run(() => {
        this.getForecast(lat, lng);
      });
    });
    var sidepanelLeft = L.control
      .sidepanel('mySidepanelLeft', {
        tabsPosition: 'left',
        startTab: 'tab-2',
      })
      .addTo(this.map);

    this.fetchGeojsonLayers();
  }

  tiposUnicos: string[] = [];
  tipoSeleccionado: string = 'TODOS';

  openSelect(selectElement: HTMLSelectElement): void {
    // Simula un clic para abrir el menú
    selectElement.focus();
  }

  onTipoChange(event: Event): void {
    const selectElement = event.target as HTMLSelectElement;
    const tipo = selectElement.value;
    this.tipoSeleccionado = tipo;
    this.clearAllLayers();
    this.fetchGeojsonLayers(tipo);
  }

  private clearAllLayers(): void {
    this.map.eachLayer((layer: any) => {
      if (!(layer instanceof this.L.TileLayer)) {
        this.map.removeLayer(layer);
      }
    });
    this.layerControl = this.L.control.layers().addTo(this.map); // Reinicia controles
  }

  private fetchGeojsonLayers(tipoFiltro: string = 'TODOS'): void {
    const L = this.L;
    const legendItems: any[] = [];

    this.layerService.getLayer().subscribe((layersData: any) => {
      const tiposSet = new Set<string>(); // Para tipos únicos

      layersData.forEach((featureItem: any) => {
        const geojson = featureItem.geojson;

        const leafletLayer = L.geoJSON(geojson, {
          filter: (feature: any) => {
            const vocacion = feature.properties.vocacion;
            const tipo = feature.properties.tipo || 'sin tipo';
            tiposSet.add(vocacion);
            const incluir = tipoFiltro === 'TODOS' || vocacion === tipoFiltro;

            // Solo agrega a la leyenda si se incluye y no está ya
            if (incluir && !legendItems.some((item) => item.label === tipo)) {
              legendItems.push({
                label: tipo,
                type: 'image',
                url: `assets/icons/${tipo}.png`,
              });
            }
            return incluir;
          },
          onEachFeature: (feature: any, layerRef: any) => {
            const tipo = feature.properties.tipo || 'sin tipo';
            const img =
              feature.properties.ruta_img || `${feature.properties.nombre}.jpg`;
            const content = this.popupService.getPopupContent(feature, tipo);
            this.popupService.precargarImagen(
              `${environment.STATIC_URL}/${img}`,
              feature,
            );
            layerRef.bindPopup(content);

            layerRef.on('click', () => {
              const panel = document.getElementById('mySidepanelLeft');
              if (!panel?.classList.contains('opened')) {
                panel?.classList.add('opened');
                panel?.classList.remove('closed');
              }
              this.ngZone.run(() => {
                this.selectedFeature = feature;
                this.selectedFeatureContent =
                  this.popupService.getPanelContent(feature);
                this.selectedSiteId = feature.properties.ogc_fid;
                this.markerId = feature.properties.id
                console.log('este es el Id: ', this.selectedSiteId);
              });
              this.cdr.detectChanges();
            });
          },
          pointToLayer: (feature: any, latlng: any) => {
            const tipo = feature.properties.tipo;
            return L.marker(latlng, {
              icon: L.icon({
                iconUrl: `assets/icons/${tipo}.png`,
                iconSize: [25, 25],
                iconAnchor: [12, 41],
                popupAnchor: [0, 30],
              }),
            });
          },
        });

        const tipo = geojson?.features?.[0]?.properties?.tipo || 'Capa';
        this.layerControl.addOverlay(leafletLayer, tipo);
        this.map.addLayer(leafletLayer);
      });

      this.tiposUnicos = Array.from(tiposSet);
      console.log('Tipos únicos:', this.tiposUnicos);
      this.cdr.detectChanges();

      if (this.tiposUnicos.length === 0) {
        this.tiposUnicos = Array.from(tiposSet);
      }
      if (this.leyendaControl) {
        this.map.removeControl(this.leyendaControl);
        this.map.removeControl(this.layerControl);
      }
      if (legendItems.length > 0) {
        this.leyendaControl = L.control
          .legend({
            position: 'bottomright',
            title: 'Leyenda',
            legends: legendItems,
            collapsed: true,
            opacity: 1,
            column: 1,
          })
          .addTo(this.map);
      }
    });
  }

  private getForecast(lat: number, lon: number): void {
    const L = this.L;

    this.weatherService.getForecast(lat, lon).subscribe({
      next: (data) => {
        this.forecastList = data.list || [];
        const first = data.list?.[0];
        const city = data.city?.name || 'Ubicación';
        const desc = first?.weather?.[0]?.description || 'Sin datos';
        const temp = first?.main?.temp ?? 'N/A';
        const icon = first?.weather?.[0]?.icon;
        const iconUrl = `https://openweathermap.org/img/wn/${icon}@2x.png`;

        this.locationLabel = `${city} Temp: ${temp}°C`;
        L.popup()
          .setLatLng([lat, lon])
          .setContent(
            `<img src="${iconUrl}" width="40" height="40"/><br><b>${city}</b><br>${desc}<br>${temp} °C`,
          )
          .openOn(this.map);

        this.cdr.detectChanges();
      },
      error: () => alert('No se pudo obtener el pronóstico.'),
    });
  }

  setReviewStars(stars: number): void {
    this.reviewStars = stars;
  }

  submitReview(): void {
    const reviewData = {
      site_id: this.selectedSiteId,
      stars: this.reviewStars,
      text: this.reviewText,
    };

    this.reviewService.uploadReview(reviewData).subscribe({
      next: () => {
        alert('¡Gracias por tu reseña!');
        this.reviewText = '';
        this.reviewStars = 0;
      },
      error: (err) => {
        alert('Error al enviar reseña');
        console.error(err);
      },
    });
  }

  selectedFile: File | null = null;

  onFileSelected(event: Event) {
    const input = event.target as HTMLInputElement;
    this.selectedFile = input.files?.[0] || null;
  }

  onSubmit() {
    if (!this.selectedFile || !this.markerId) return;

    const formData = new FormData();
    formData.append('image', this.selectedFile);
    formData.append('marker_id', this.markerId.toString());

    this.http
      .post(`${environment.API_URL}/reviews/uploadImg/`, formData)
      .subscribe({
        next: (res) => console.log('Imagen subida:', res),
        error: (err) => console.error('Error al subir imagen:', err),
      });
  }
}
