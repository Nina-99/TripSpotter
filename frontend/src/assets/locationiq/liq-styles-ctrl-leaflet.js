
// class TileLayer.Unwired
// Tile layer for Unwired Map tiles
L.TileLayer.Unwired = L.TileLayer.extend({

	options: {
		subdomains: 'abc',
		minZoom: 2,
		maxZoom: 18,
		scheme: 'streets',
		format: 'png',
		key: '', //access token, not the main API token!
		attributions: [] //depends on map layer
	},


	initialize: function initialize(options) {
		options = L.setOptions(this, options);

		// If it's not Retina, don't add anything here, it'll be 256px by default
		options.tileResolution = '';

		options.attributions = ['&copy; <a href="https://locationiq.com/?ref=maps" target="_blank">LocationIQ Maps</a>', '<a href="https://www.openstreetmap.org/copyright" target="_blank">OpenStreetMap</a> Contributors'];

		if(options.scheme == "streets") {
			options.scheme = "streets";
		} else if(options.scheme == "earth") {
			options.format = "jpg";
			options.maxZoom = 18;
			options.attributions = ['&copy; <a href="https://locationiq.com/?ref=maps" target="_blank">LocationIQ Maps</a>', '&copy; <a href="http://bit.ly/mapsapiderive">DigitalGlobe</a>'];
		} else if(options.scheme == "hybrid") {
			options.format = "jpg";
			options.maxZoom = 18;
			options.attributions = ['&copy; <a href="https://locationiq.com/?ref=maps" target="_blank">LocationIQ Maps</a>', '&copy; <a href="http://bit.ly/mapsapiderive">DigitalGlobe</a>', '<a href="https://www.openstreetmap.org/copyright" target="_blank">OpenStreetMap</a> Contributors'];
		}
				
		//If it's Retina, we need to serve 512px tiles
		if (L.Browser.retina) {
			options.tileResolution = '@2x';
		}

		var path = '/v3/{scheme}/r/{z}/{x}/{y}{tileResolution}.{format}?key={key}';
		var tileServer = 'https://{s}-tiles.locationiq.com';


		var tileUrl = tileServer + path;
		
		L.TileLayer.prototype.initialize.call(this, tileUrl, options);


	},

	onAdd: function onAdd(map) {
		L.TileLayer.prototype.onAdd.call(this, map);

		if (!this._attributionBBoxes) {
			this._findCopyrightBBox(this.options);
		}


	},

	onRemove: function onRemove(map) {
		L.TileLayer.prototype.onRemove.call(this, map);

		this._map.attributionControl.removeAttribution(this._attributionText);

		this._map.off('moveend zoomend resetview', this._findCopyrightBBox, this);
	},

	_findCopyrightBBox: function _findCopyrightBBox(options) {		
		var attributions = options.attributions;
		var attributionText = attributions.join(', ');
		this._map.attributionControl.addAttribution(this._attributionText = attributionText);
	}

});



L.tileLayer.Unwired = function(opts){

	return new L.TileLayer.Unwired(opts);

};


