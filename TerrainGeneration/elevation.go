/*
github.com/TadaTeruki/RDGL/TerrainGeneration/elevation.go
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


func (obj *WorldTerrainObject) GetNoiseLevelByKmPoint(xKm, yKm float64) float64{

	xfix := xKm/obj.Config.PlateSizeKm
	yfix := yKm/obj.Config.PlateSizeKm

	noise_adj := obj.NoiseSrc.OctaveNoise(1, 0.5, xfix, yfix, 0.0)

	noise := obj.NoiseSrc.OctaveNoiseFixed(obj.Config.NoizeOctave,
		obj.Config.NoizeMinPersistence+noise_adj*(obj.Config.NoizeMaxPersistence-obj.Config.NoizeMinPersistence),
		xfix, yfix, obj.Z)

	return TerrainAdjustmentFade(noise, noise_adj)
	
}

func (obj *WorldTerrainObject) GetElevationByKmPoint(xKm, yKm float64) float64{
	
	noise := obj.GetNoiseLevelByKmPoint(xKm, yKm)
	return obj.GetElevationFromNoiseLevel(noise)
	
}

func (obj *LocalTerrainObject) GetElevationByKmPointFromElevationTable(xKm, yKm float64) float64{


	data_w := len(obj.ElevationTable[0])
	data_h := len(obj.ElevationTable)
	data_fw := float64(data_w)
	data_fh := float64(data_h)
	fx := xKm/obj.WEKm*data_fw
	fy := yKm/obj.NSKm*data_fh
	ffx := math.Floor(fx)
	ffy := math.Floor(fy)
	
	cfx := ffx+1
	cfy := ffy+1

	if cfx >= data_fw { cfx = data_fw-1 }
	if cfy >= data_fh { cfy = data_fh-1 }

	nw := (fx-ffx)*(fy-ffy)
	ne := (cfx-fx)*(fy-ffy)
	sw := (fx-ffx)*(cfy-fy)
	se := (cfx-fx)*(cfy-fy)

	return se*obj.ElevationTable[int(ffy)][int(ffx)]+sw*obj.ElevationTable[int(ffy)][int(cfx)]+
	       ne*obj.ElevationTable[int(cfy)][int(ffx)]+nw*obj.ElevationTable[int(cfy)][int(cfx)]
	
	//return obj.ElevationTable[int(ffy)][int(ffx)]

}

func (obj *LocalTerrainObject) GetElevationByKmPoint(xKm, yKm float64) float64{

	if xKm < 0 { xKm = 0 }
	if yKm < 0 { yKm = 0 }
	if xKm > obj.WEKm { xKm = obj.WEKm }
	if yKm > obj.NSKm { yKm = obj.NSKm }

	var relv float64

	if obj.ElevationTableIsAvailable == true {
		noise := obj.WorldTerrain.GetNoiseLevelByKmPoint(xKm+obj.xKm, yKm+obj.yKm)
		relv = obj.GetElevationByKmPointFromElevationTable(xKm, yKm)
		mins := obj.WorldTerrain.Config.OutlineNoiseMinStrength
		maxs := obj.WorldTerrain.Config.OutlineNoiseMaxStrength
		elevation_abs := obj.WorldTerrain.ElevationAbsM
		relv += math.Abs(noise-0.5)*2.0*(elevation_abs*(math.Abs(relv)/elevation_abs*(maxs-mins)+mins))
	} else {
		relv = obj.WorldTerrain.GetElevationByKmPoint(xKm+obj.xKm, yKm+obj.yKm)
	}

	leveling_elevation := obj.WorldTerrain.Config.LevelingMinimumElevationProportion*obj.WorldTerrain.ElevationAbsM
	liver_elevation := obj.WorldTerrain.Config.LiverEndPointElevationProportion*obj.WorldTerrain.ElevationAbsM

	if obj.LevelingCheckIsAvailable == true {
		oc := obj.LevelingLayerObj.GetLevelingPointByKmPoint(obj, xKm, yKm)

		if oc.IsLeveling == true &&
		   oc.ElevationLevel > leveling_elevation{
			diff := (oc.ElevationLevel-relv)
			relv = oc.ElevationLevel-diff*obj.WorldTerrain.Config.LakeDepthProportion
		}
	}

	
	if relv >= liver_elevation && obj.LiverCheckIsAvailable == true {
		
		arelv := (relv-liver_elevation) * obj.CheckLiverCavityByKmPoint(xKm, yKm)+liver_elevation
		relv = math.Min(relv,arelv)
	}
	
	return relv
}