package terrain_generation

import(
	"math"
)

func TerrainAdjustmentFade(px, strength float64) float64{
	x := px
	if x == 0 { return x }
	return math.Atan((2*x-1.0)*math.Tan(strength*math.Pi*0.5))/(strength*math.Pi*0.5)*0.5+0.5
}

func (obj *WorldTerrainObject) GetElevationByKmPoint(xKm, yKm float64) float64{

	xfix := xKm/obj.Config.NoizeScaleKm
	yfix := yKm/obj.Config.NoizeScaleKm

	noise_adj := obj.NoiseSrc.OctaveNoise(1, 0.5, xfix, yfix, 0.0)

	noise := obj.NoiseSrc.OctaveNoiseFixed(obj.Config.NoizeOctave,
		obj.Config.NoizeMinPersistence+noise_adj*(obj.Config.NoizeMaxPersistence-obj.Config.NoizeMinPersistence),
		xfix, yfix, obj.Z)

	return obj.GetElevationFromNoiseLevel(TerrainAdjustmentFade(noise, noise_adj))
	
}

func (obj *LocalTerrainObject) GetElevationByKmPoint(xKm, yKm float64) float64{
	relv := obj.WorldTerrain.GetElevationByKmPoint(xKm+obj.xKm, yKm+obj.yKm)


	if obj.OceanCheckIsAvailable == true {
		oc := obj.OceanLayerObj.GetOceanPointByKmPoint(obj, xKm, yKm)

		if oc.IsOcean == true{
			diff := (oc.ElevationLevel-relv)
			relv = oc.ElevationLevel-diff*obj.WorldTerrain.Config.PondDepthProportion
			if oc.ElevationLevel >= 0.0 && relv < 0.0 {
				relv = 0.0
			}
		}
	}

	
	if relv > 0.0 && obj.LiverCheckIsAvailable == true {
		relv = relv * obj.CheckLiverCavityByKmPoint(xKm, yKm)
	}
	
	
	return relv
}