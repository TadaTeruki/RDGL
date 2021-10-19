package artif_terrain

import(
	"math"
)

func TerrainAdjustmentFade(px, strength float64) float64{
	x := px
	if x == 0 { return x }
	return math.Atan((2*x-1.0)*math.Tan(strength*math.Pi*0.5))/(strength*math.Pi*0.5)*0.5+0.5
}


func (obj *WorldTerrainObject) GetElevation(x, y float64) float64{
/*
	var maxWHkm float64

	if obj.NSKm > obj.WEKm { 
		maxWHkm = obj.NSKm
	} else {
		maxWHkm = obj.WEKm
	}
	*/
	xfix := x/obj.Config.NoizeScaleKm
	yfix := y/obj.Config.NoizeScaleKm

	noise_adj := obj.NoiseSrc.OctaveNoise(1, 0.5, xfix, yfix, 0.0)

	noise := obj.NoiseSrc.OctaveNoiseFixed(obj.Config.NoizeOctave,
		obj.Config.NoizeMinPersistence+noise_adj*(obj.Config.NoizeMaxPersistence-obj.Config.NoizeMinPersistence),
		xfix, yfix, 0.0)
	
	return obj.GetElevationFromNoiseLevel(TerrainAdjustmentFade(noise, noise_adj))
	
}