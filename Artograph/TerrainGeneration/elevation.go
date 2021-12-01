/*
Artograph/TerrainGeneration/elevation.go
Copyright (C) 2021 Tada Teruki

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation; either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
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

	if xKm < 0 { xKm = 0 }
	if yKm < 0 { yKm = 0 }
	if xKm > obj.WEKm { xKm = obj.WEKm }
	if yKm > obj.NSKm { yKm = obj.NSKm }

	relv := obj.WorldTerrain.GetElevationByKmPoint(xKm+obj.xKm, yKm+obj.yKm)

	if obj.LevelingCheckIsAvailable == true {
		oc := obj.LevelingLayerObj.GetLevelingPointByKmPoint(obj, xKm, yKm)

		if oc.IsLeveling == true{
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