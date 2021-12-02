
/*
Artograph/TerrainGeneration/make_terrain.go
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
	"math/rand"
	"fmt"
)

func (obj *WorldTerrainObject) MakeWorldTerrain(){
	obj.NoiseSrc.SetSeed(obj.Config.Seed)
	obj.AdjNoiseSrc.SetSeed(obj.Config.Seed)
}

// N(count^2)
func (obj *LocalTerrainObject) SubmitLocalTerrain(count int) (float64, bool){
	
	var score float64
	fcount := float64(count)

	var maxWHkm float64

	if obj.NSKm > obj.WEKm { 
		maxWHkm = obj.NSKm
	} else {
		maxWHkm = obj.WEKm
	}

	cxKm := obj.xKm + obj.WEKm/2
	cyKm := obj.yKm + obj.NSKm/2
	count_land := 0
	count_all := 0

	for yKm := obj.yKm; yKm < obj.yKm + obj.NSKm; yKm += maxWHkm/fcount{
		for xKm := obj.xKm; xKm < obj.xKm + obj.WEKm; xKm += maxWHkm/fcount{
			distKm := math.Sqrt((cxKm-xKm)*(cxKm-xKm)+(cyKm-yKm)*(cyKm-yKm))
			rDistKm := 0.5-distKm/(maxWHkm/2)
			_,_ = distKm, rDistKm
			elevation := obj.WorldTerrain.GetElevationByKmPoint(xKm, yKm)
			if elevation >= 0 {
				count_land++
			}
			count_all++
			score += elevation*elevation
		}
	}
	
	land := float64(count_land)/float64(count_all)

	if land < obj.WorldTerrain.Config.MinLand || land > obj.WorldTerrain.Config.MaxLand {
		return 0, false
	}
	

	return score, true
}

func (obj *LocalTerrainObject) MakeLocalTerrain(){
	
	rand.Seed(0)

	submit_model_num := obj.WorldTerrain.Config.LocalTerrainSelectionQuality
	cobj := make([]LocalTerrainObject, submit_model_num)
	var max_score float64
	var select_ad int
	var select_z float64

	obj.LevelingCheckIsAvailable = false
	obj.LiverCheckIsAvailable = false
	

	for i:=0; i<submit_model_num; i++{
		cobj[i] = *obj
		cobj[i].xKm = (obj.WorldTerrain.WEKm-obj.WEKm)*rand.Float64()
		cobj[i].yKm = (obj.WorldTerrain.NSKm-obj.NSKm)*rand.Float64()
		score, available := cobj[i].SubmitLocalTerrain(10)
		if available == false{
			i--
			obj.WorldTerrain.Z += 11
			continue
		}

		if score > max_score {
			max_score = score
			select_ad = i
			select_z = obj.WorldTerrain.Z
		}
	}
	
	obj.WorldTerrain.Z = select_z



	obj.xKm = cobj[select_ad].xKm
	obj.yKm = cobj[select_ad].yKm

	fmt.Println("Process : Leveling")
	obj.MakeLevelingLayer()
	fmt.Println("Process : Liver")
	obj.MakeLiverTable()
	fmt.Println("Terrain generation successfully finished")

}