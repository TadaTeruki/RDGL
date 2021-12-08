/*
github.com/TadaTeruki/RDGL/TerrainGeneration/liver.go
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
	"sort"
	utility "github.com/TadaTeruki/RDGL/Utility"
)

func (obj *LocalTerrainObject) GetLiverPointFromKmPoint(xKm, yKm float64) *LiverPoint{
	liver_interval_km := obj.WorldTerrain.Config.UnitKm
	x := int(math.Min(xKm/liver_interval_km,float64(len(obj.LiverTable[0])-1)))
	y := int(math.Min(yKm/liver_interval_km,float64(len(obj.LiverTable)-1)))
	lt := &obj.LiverTable[y][x]
	return lt
}

func (obj *LocalTerrainObject) MakeLiverTable(){
	
	liver_interval_km := obj.WorldTerrain.Config.UnitKm
	//side_width_km := obj.WorldTerrain.Config.MapSideWidthKm
	obj.LiverTable = make([][]LiverPoint, int(math.Ceil(obj.NSKm/liver_interval_km)))
	var lv_order []Point

	for y := 0; y<len(obj.LiverTable); y++{
		obj.LiverTable[y] = make([]LiverPoint, int(math.Ceil(obj.WEKm/liver_interval_km)))
		for x := 0; x<len(obj.LiverTable[y]); x++{
			obj.LiverTable[y][x].XKm = float64(x)*liver_interval_km
			obj.LiverTable[y][x].YKm = float64(y)*liver_interval_km
			obj.LiverTable[y][x].Direction = DIRECTION_NONE
			obj.LiverTable[y][x].Cavity = 1.0
			obj.LiverTable[y][x].BaseElevation = obj.GetElevationByKmPoint(obj.LiverTable[y][x].XKm, obj.LiverTable[y][x].YKm)
			if obj.LiverTable[y][x].BaseElevation >= obj.WorldTerrain.Config.LiverEndPointElevationProportion*obj.WorldTerrain.ElevationAbsM {
				lv_order = append(lv_order, MakePoint(x, y))
			}
			
		}
	}

	sort.Slice(lv_order, func(i, j int) bool {
		return obj.LiverTable[lv_order[i].Y][lv_order[i].X].BaseElevation >
			    obj.LiverTable[lv_order[j].Y][lv_order[j].X].BaseElevation
	})

	path_score := func(xKm, yKm, sxKm, syKm, elevation float64) float64{
		euc := math.Sqrt((xKm-sxKm)*(xKm-sxKm)+(yKm-syKm)*(yKm-syKm))
		lt := obj.GetLiverPointFromKmPoint(xKm, yKm)
		return elevation*lt.Cavity*(0.5+euc/(obj.NSKm+obj.WEKm/2))
	}

	loop_out_condition := func(xKm, yKm, sxKm, syKm, elevation float64) bool{

		//if xKm < side_width_km || yKm < side_width_km || xKm > obj.WEKm-side_width_km || yKm > obj.NSKm-side_width_km { return true }
		lt := obj.GetLiverPointFromKmPoint(xKm, yKm)
		if lt.Direction != DIRECTION_NONE { return true }		

		return elevation <= obj.WorldTerrain.Config.LiverEndPointElevationProportion*obj.WorldTerrain.ElevationAbsM
	}

	utility.EchoProcessPercentage("Liver simulation", 0)

	checked_sum := 0.0
	checked_all := float64(len(lv_order))

	// simulate the behavior of water
	for i := 0; i<len(lv_order); i++{
		ltroot := obj.LiverTable[lv_order[i].Y][lv_order[i].X]
		if ltroot.Direction != DIRECTION_NONE {
			continue
		}
		path := obj.MakePath(ltroot.XKm, ltroot.YKm, liver_interval_km, path_score, loop_out_condition)
		for j := len(path)-1; j >= 1; j--{
			lt := obj.GetLiverPointFromKmPoint(path[j].XKm,path[j].YKm)
			xd := path[j].XKm - path[j-1].XKm
			yd := path[j].YKm - path[j-1].YKm

			if xd == 0 && yd > 0 { lt.Direction = DIRECTION_NORTH }
			if xd > 0 && yd == 0 { lt.Direction = DIRECTION_WEST }
			if xd == 0 && yd < 0 { lt.Direction = DIRECTION_SOUTH }
			if xd < 0 && yd == 0 { lt.Direction = DIRECTION_EAST }
			
			lt2 := obj.GetLiverPointFromKmPoint(path[j-1].XKm,path[j-1].YKm)

			if j == 1 && lt2.Cavity < 1.0{
				for l := 1; l < len(path); l++{
					p := float64(l)/float64(len(path))
					lt3 := obj.GetLiverPointFromKmPoint(path[l].XKm,path[l].YKm)
					lt3.Cavity = p*lt3.Cavity + (1.0-p)*lt2.Cavity
				}

			} else {
				lt2.Cavity = lt.Cavity
				if lt.BaseElevation*lt.Cavity < lt2.BaseElevation*lt2.Cavity {
					lt2.Cavity = math.Max(lt.BaseElevation*lt.Cavity/lt2.BaseElevation,0)


				}
			}

			checked_before_sum := checked_sum
			checked_sum ++
			if math.Floor(checked_before_sum/checked_all*10) != math.Floor(checked_sum/checked_all*10) &&
			   checked_sum < checked_all{
				
				utility.EchoProcessPercentage("Liver simulation", checked_sum/checked_all)
			}
		}


	}

	utility.EchoProcessEnd("Liver simulation")

	obj.LiverCheckIsAvailable = true
	
}

func (obj *LocalTerrainObject) CheckLiverCavityByKmPoint(xKm, yKm float64) float64{


	liver_interval_km := obj.WorldTerrain.Config.UnitKm

	var nw, ne, sw, se KmPoint
	nw.XKm = math.Floor(xKm/liver_interval_km)*liver_interval_km
	ne.XKm = nw.XKm+liver_interval_km
	nw.YKm = math.Floor(yKm/liver_interval_km)*liver_interval_km
	sw.YKm = nw.YKm+liver_interval_km
	sw.XKm = nw.XKm
	se.XKm = ne.XKm
	ne.YKm = nw.YKm
	se.YKm = sw.YKm

	nwsc := (xKm-nw.XKm)*(yKm-nw.YKm)/(liver_interval_km*liver_interval_km)
	nesc := (ne.XKm-xKm)*(yKm-ne.YKm)/(liver_interval_km*liver_interval_km)
	swsc := (xKm-sw.XKm)*(sw.YKm-yKm)/(liver_interval_km*liver_interval_km)
	sesc := (se.XKm-xKm)*(se.YKm-yKm)/(liver_interval_km*liver_interval_km)

	nwcav := obj.GetLiverPointFromKmPoint(nw.XKm, nw.YKm).Cavity
	necav := obj.GetLiverPointFromKmPoint(ne.XKm, ne.YKm).Cavity
	swcav := obj.GetLiverPointFromKmPoint(sw.XKm, sw.YKm).Cavity
	secav := obj.GetLiverPointFromKmPoint(se.XKm, se.YKm).Cavity
	
	return nwcav*sesc + necav*swsc + swcav*nesc + secav*nwsc
}
