package artif_terrain


func (obj *WorldTerrainObject) GetElevation(x, y float64) float64{

	var maxWHkm float64

	if obj.NSKm > obj.WEKm { 
		maxWHkm = obj.NSKm
	} else {
		maxWHkm = obj.WEKm
	}
	x01 := x/maxWHkm
	y01 := y/maxWHkm
	noise := obj.NoiseSrc.OctaveNoiseFixed(10, 0.5, x01, y01, 0)
	return obj.GetElevationFromNoiseLevel(noise)
	
}