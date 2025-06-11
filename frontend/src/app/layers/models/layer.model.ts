export interface LayerItem {
  id: number;
  name: string;
  feature: any;
  leafletLayer: L.GeoJSON;
  visible: boolean;
  color: string;
}
