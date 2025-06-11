import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { HttpClient } from '@angular/common/http';

@Injectable({ providedIn: 'root' })
export class getPopupService {
  private img: any;
  constructor() {}
  getPopupContent(feature: any, typeAtribute: any): string {
    const props = feature.properties;
    let imagen: string = '';
    switch (typeAtribute) {
      case 'Aeropuerto':
        imagen = `${environment.STATIC_URL}/${props.ruta_img}`;
        return `<div style="text-align:center"><b>${props.nombre_aer}</b><br><b>Ciudad: ${props.ciudad}</b></div>`;
        break;
      case 'Terminal':
        imagen = `${environment.STATIC_URL}/${props.ruta_img}`;
        return `<div style="text-align:center"><b>${props.nombre_com}</b><br><b>Ciudad: ${props.ciudad}</b></div>`;
        break;
      case 'Teleférico':
        return `<div style="text-align:center"><b>Ciudad: ${props.ciudad}</b><br><b>Estacion: ${props.nombre}</b><br><b>Linea: ${props.linea}</b>`;
        break;
      case 'Turistico':
        return `<div style="text-align:center"><b>${props.nombre}</b><br><b class"pIcon"><img src="assets/icons/default.png" width="20"> ${props.municipio}, ${props.provincia}, ${props.departamen}</b></div>`;
        break;
      case 'Club nocturno':
        return `<div style="text-align:center"><b>${props.nombre}</b><br><b>Ciudad: ${props.dep}</b></div>`;
        break;
      case 'Casino':
        return `<div style="text-align:center"><b>${props.nombre}</b><br><b>Ciudad: ${props.dep}</b></div>`;
        break;
      case 'Centros comerciales':
        return `<div style="text-align:center"><b>${props.tipo}</b><br><b>Ciudad: ${props.dep}</b></div>`;
        break;
      case 'Plazas':
        return `<div style="text-align:center"><b>${props.tipo}</b><br><b>Ciudad: ${props.dep}</b></div>`;
        break;
      case 'Parques':
        return `<div style="text-align:center"><b>${props.nombre}</b><br><b>Ciudad: ${props.dep}</b></div>`;
        break;
      case 'Parques infantiles':
        return `<div style="text-align:center"><b>${props.nombre}</b><br><b>Ciudad: ${props.dep}</b></div>`;
        break;
      case 'Parque  de perros':
        return `<div style="text-align:center"><b>${props.nombre}</b><br><b>Ciudad: ${props.dep}</b></div>`;
        break;
      case 'Parque de atracciones':
        return `<div style="text-align:center"><b>${props.nombre}</b><br><b>Ciudad: ${props.dep}</b></div>`;
        break;
      case 'Parque acuatico':
        return `<div style="text-align:center"><b>${props.nombre}</b><br><b>Ciudad: ${props.dep}</b></div>`;
        break;
      default:
        return `<strong>nada</strong><br>`;
    }
  }

  checkImage(url: string): Promise<boolean> {
    return new Promise((resolve) => {
      const img = new Image();
      img.onload = () => resolve(true);
      img.onerror = () => resolve(false);
      img.src = url;
    });
  }

  private imagenes: { [nombre: string]: string } = {};

  async precargarImagen(
    ruta: string,
    feature: any,
    actualizarCallback?: (feature: any) => void,
  ) {
    if (this.imagenes[ruta]) return; // ya está verificada

    const exists = await this.checkImage(ruta);
    this.imagenes[ruta] = exists
      ? ruta
      : `${environment.STATIC_URL}/img-default.png`;

    // Opcional: actualizar el contenido del panel si se proporciona un callback
    if (actualizarCallback) {
      actualizarCallback(feature);
    }
  }
  // async precargarImagen(nombre: string) {
  //   // const img = nombre;
  //   const exists = await this.checkImage(nombre);
  //
  //   if (exists) {
  //     // Solo asigna si la imagen existe
  //     this.imagenes[nombre] = nombre;
  //   } else if (!this.imagenes[nombre]) {
  //     // Solo asigna imagen por defecto si aún no existe ninguna
  //     this.imagenes[nombre] = `${environment.STATIC_URL}/img-default.png`;
  //   }
  // }

  // Luego puedes usarla así:
  getImagen(nombre: string): string {
    return this.imagenes[nombre] || `${environment.STATIC_URL}/img-default.png`;
  }
  getPanelContent(feature: any): string {
    const props = feature.properties;
    const typeAtribute = props.tipo;
    let imagen: string = '';
    switch (typeAtribute) {
      case 'Aeropuerto':
        this.img = `${environment.STATIC_URL}/${props.ruta_img}`;

        this.precargarImagen(this.img, feature).then(() => {
          imagen = this.getImagen(this.img);
        });
        imagen = this.getImagen(this.img);
        return `<img src="${imagen}" width="100%"/><br><p><b>${props.nombre_aer}</b><br><p><b>Ciudad: </b>${props.ciudad}</p><br><p><b>Vuelos </b>${props.tipo_aerod}</p><br><p><b>Altitud:</b>  ${props.elevacion_}</p>`;
        break;
      case 'Terminal':
        this.img = `${environment.STATIC_URL}/${props.ruta_img}`;

        this.precargarImagen(this.img, feature).then(() => {
          imagen = this.getImagen(this.img);
        });
        imagen = this.getImagen(this.img);
        return `<img src="${imagen}" width="100%"/><br><p><b>${props.nombre_com}</b><br><p><b>Ciudad: </b>${props.ciudad}</p><br><p><b><img src="assets/icons/default.png" width="20"> </b>${props.direccion}</p><br><p>  ${props.ubicacion}</p>`;
        break;
      case 'Teleférico':
        this.img = `${environment.STATIC_URL}/${props.ruta_img}`;

        this.precargarImagen(this.img, feature).then(() => {
          imagen = this.getImagen(this.img);
        });
        imagen = this.getImagen(this.img);
        console.log('imagen: ', imagen);
        return `<img src="${imagen}" width="100%"/><br><p><b>Linea </b>${props.linea}</p><br><p><b>Estacion </b>${props.nombre}</p><br><p><b>Referencia: </b>${props.referencia}</p>`;
        break;
      case 'Turistico':
        this.img = `${environment.STATIC_URL}/${props.nombre}.jpg`;

        this.precargarImagen(imagen, feature).then(() => {
          console.log('🖼️ Imagen verificada:', this.getImagen(imagen));
        });
        imagen = this.getImagen(this.img);
        console.log('no: ', this.getImagen(this.img));
        return `<img src="${imagen}" width="100%" /><br><h3>${props.nombre}</h3><br><b class"pIcon"><img src="assets/icons/default.png" width="20"> ${props.municipio}, ${props.provincia}, ${props.departamen}</b><br><b class"pIcon"><img src="assets/icons/reloj.png" width="20"> ${props.temporalid}</b><br><div class="container"><h5>ACTIVIDADES</h5></div><p>${props.actividade}</p><br><div class="container"><h5>INTERES</h5></div><p>${props.sitios_tur}</p><br><div class="container"><h5>HOSPEDAJE</h5></div><p>${props.alojamient}</p><br><div class="container"><h5>ALIMENTACION</h5></div><p>${props.alimentaci}</p><br><div class="container"><h5>TRANSPORTE</h5></div><p>${props.transporte}</p>`;
        break;
      case 'Club nocturno':
        return `<b>Ciudad: ${props.dep}</b><br><b>Nombre: ${props.nombre}</b>`;
        break;
      case 'Casino':
        return `<b>Ciudad: ${props.dep}</b><br><b>Nombre: ${props.nombre}</b>`;
        break;
      case 'Centros comerciales':
        return `<b>Ciudad: ${props.dep}</b><br><b>Nombre: ${props.nombre}</b>`;
        break;
      case 'Plazas':
        return `<b>Ciudad: ${props.dep}</b><br><b>Nombre: ${props.nombre}</b>`;
        break;
      case 'Parques':
        return `<b>Ciudad: ${props.dep}</b><br><b>Nombre: ${props.nombre}</b>`;
        break;
      case 'Parques infantiles':
        return `<b>Ciudad: ${props.dep}</b><br><b>Nombre: ${props.nombre}</b>`;
        break;
      case 'Parque  de perros':
        return `<b>Ciudad: ${props.dep}</b><br><b>Nombre: ${props.nombre}</b>`;
        break;
      case 'Parque de atracciones':
        return `<b>Ciudad: ${props.dep}</b><br><b>Nombre: ${props.nombre}</b>`;
        break;
      case 'Parque acuatico':
        return `<b>Ciudad: ${props.dep}</b><br><b>Nombre: ${props.nombre}</b>`;
        break;
      default:
        return `<strong>nada</strong><br>`;
    }
  }
}
