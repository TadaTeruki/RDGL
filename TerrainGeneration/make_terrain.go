
/*
github.com/TadaTeruki/RDGL/TerrainGeneration/make_terrain.go
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
	utility "github.com/TadaTeruki/RDGL/Utility"
	"sort"
)

func (obj *WorldTerrainObject) MakeWorldTerrain(){
	obj.NoiseSrc.SetSeed(obj.Config.Seed)
	obj.AdjNoiseSrc.SetSeed(obj.Config.Seed)
}

// N(count^2)
func (obj *LocalTerrainObject) SubmitLocalTerrain(count int) (float64, bool){
	
	var score float64
	fcount := float64(count)

	var minWHkm float64

	if obj.NSKm < obj.WEKm { 
		minWHkm = obj.NSKm
	} else {
		minWHkm = obj.WEKm
	}

	cxKm := obj.xKm + obj.WEKm/2
	cyKm := obj.yKm + obj.NSKm/2
	count_land := 0
	count_all := 0

	for yKm := obj.yKm; yKm < obj.yKm + obj.NSKm; yKm += minWHkm/fcount{
		for xKm := obj.xKm; xKm < obj.xKm + obj.WEKm; xKm += minWHkm/fcount{
			distKm := math.Sqrt((cxKm-xKm)*(cxKm-xKm)+(cyKm-yKm)*(cyKm-yKm))
			rDistKm := 0.5-distKm/(minWHkm/2)
			_,_ = distKm, rDistKm
			elevation := obj.WorldTerrain.GetElevationByKmPoint(xKm, yKm)
			if elevation >= 0 {
				count_land++
			}
			count_all++
			//score += elevation
		}
	}
	
	land := float64(count_land)/float64(count_all)
	if land > 1.0 { land = 1.0 }
	if land < 0.0 { land = 0.0 }

	score = 1.0 - math.Abs(obj.WorldTerrain.Config.StandardLandProportion-land)
	
	return score, true
}

func (obj *LocalTerrainObject) MakeLocalTerrain(){
	
	rand.Seed(0)

	submit_model_num := obj.WorldTerrain.Config.LocalTerrainSelectionQuality
	cobj := make([]LocalTerrainObject, submit_model_num)
	var max_score float64
	var select_ad int
	var select_z float64


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

}

func (obj *LocalTerrainObject) SetRoot(){
	root_interval_km := obj.WorldTerrain.Config.RootIntervalKm
	unit_km := obj.WorldTerrain.Config.UnitKm
	//shelf_elevation_m := obj.WorldTerrain.Config.ContShelfElevationProportion*obj.WorldTerrain.ElevationAbsM
	//start_x_km := (obj.WEKm-math.Floor(obj.WEKm/root_interval_km)*root_interval_km)*0.5
	//start_y_km := (obj.NSKm-math.Floor(obj.NSKm/root_interval_km)*root_interval_km)*0.5
	/*
	for syKm := 0.0; syKm < obj.NSKm; syKm += root_interval_km  {
		for sxKm := 0.0; sxKm < obj.WEKm; sxKm += root_interval_km  {
			//obj.RootList = append(obj.RootList, MakeKmPoint(xKm, yKm))

			if sxKm != 0.0 && syKm != 0.0 && sxKm+root_interval_km <= obj.NSKm && syKm+root_interval_km <= obj.WEKm {
				continue
			}
			
			min_point := MakeKmPoint(sxKm, syKm);
			max_elevation := -obj.WorldTerrain.ElevationAbsM
			for yKm := syKm; yKm < syKm+root_interval_km; yKm += unit_km{
				for xKm := sxKm; xKm < sxKm+root_interval_km; xKm += unit_km{
					elevation := obj.GetElevationByKmPoint(xKm,yKm)
					if elevation > max_elevation {
						max_elevation = elevation
						min_point = MakeKmPoint(xKm, yKm);
					}
				}
			}
			obj.RootList = append(obj.RootList, min_point)
			
		}
	}*/

	
	for syKm := 0.0; syKm < obj.NSKm; syKm += root_interval_km  {
		for sxKm := 0.0; sxKm < obj.WEKm; sxKm += root_interval_km  {
			//obj.RootList = append(obj.RootList, MakeKmPoint(xKm, yKm))
			
			if sxKm != 0.0 && syKm != 0.0 && sxKm+root_interval_km <= obj.NSKm && syKm+root_interval_km <= obj.WEKm {
				continue
			}
			
			
			min_point := MakeKmPoint(sxKm, syKm);
			min_elevation := -obj.WorldTerrain.ElevationAbsM
			for yKm := syKm; yKm < syKm+root_interval_km; yKm += unit_km{
				for xKm := sxKm; xKm < sxKm+root_interval_km; xKm += unit_km{
					elevation := obj.GetElevationByKmPoint(xKm,yKm)
					if elevation < min_elevation {
						min_elevation = elevation
						min_point = MakeKmPoint(xKm, yKm);
					}
				}
			}
			obj.RootList = append(obj.RootList, min_point)
			
		}
	}
	

	sort.Slice(obj.RootList, func(i,j int) bool{
		i_elevation := obj.GetElevationByKmPoint(obj.RootList[i].XKm,obj.RootList[i].YKm)
		j_elevation := obj.GetElevationByKmPoint(obj.RootList[j].XKm,obj.RootList[j].YKm)
		return i_elevation < j_elevation
	})
}

func (obj *LocalTerrainObject) TransformProcess(leveling bool, liver bool){

	obj.SetRoot()

	obj.MakeUnitLayer()
	

	if leveling == true {
		obj.Leveling()
	}


	if liver == true {
		obj.MakeLiver()
	}

	


	utility.EchoProcessEnd("DEM generation")
}