package artif_terrain

import(
	"math"
	"sort"
)

const DIRECTION_NONE  = 0
const DIRECTION_NORTH = 1
const DIRECTION_WEST  = 2
const DIRECTION_SOUTH = 3
const DIRECTION_EAST  = 4

type LiverPoint struct{
	XKm float64
	YKm float64
	Direction int
	Cavity float64
	BaseElevation float64
}

func (obj *LocalTerrainObject) GetLiverPointFromKmPoint(xKm, yKm float64) *LiverPoint{
	liver_interval_km := obj.WorldTerrain.Config.LiverCheckIntervalKm
	x := int(math.Min(xKm/liver_interval_km,float64(len(obj.LiverTable[0])-1)))
	y := int(math.Min(yKm/liver_interval_km,float64(len(obj.LiverTable)-1)))
	lt := &obj.LiverTable[y][x]
	return lt
}


func (obj *LocalTerrainObject) MakeLiverTable(){
	liver_interval_km := obj.WorldTerrain.Config.LiverCheckIntervalKm
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
			if obj.LiverTable[y][x].BaseElevation >= 0.0 {
				lv_order = append(lv_order, MakePoint(x, y))
			}
			
		}
	}

	sort.Slice(lv_order, func(i, j int) bool {
		return obj.LiverTable[lv_order[i].Y][lv_order[i].X].BaseElevation >
			    obj.LiverTable[lv_order[j].Y][lv_order[j].X].BaseElevation
	})

	path_score := func(xKm, yKm, sxKm, syKm, elevation float64) float64{
		return elevation
	}

	loop_out_condition := func(xKm, yKm, sxKm, syKm, elevation float64) bool{

		lt := obj.GetLiverPointFromKmPoint(xKm, yKm)
		if lt.Direction != DIRECTION_NONE {
			return true
		}		
		
		return elevation < 0
	}


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

			if j == 1 && lt2.Cavity < 0.99{
				/*
				p := (1.01-lt2.Cavity)/(1.01-lt.Cavity)
				for l := 1; l < len(path); l++{
					lt3 := obj.GetLiverPointFromKmPoint(path[l].XKm,path[l].YKm)
					lt3.Cavity = 1.0-(1.0-lt3.Cavity)*p
				}*/
				
				//p := (ltroot.BaseElevation-lt2.BaseElevation*lt2.Cavity)/(ltroot.BaseElevation-lt.BaseElevation*lt.Cavity)
				
				for l := 1; l < len(path); l++{
					p := float64(l)/float64(len(path))
					lt3 := obj.GetLiverPointFromKmPoint(path[l].XKm,path[l].YKm)
					lt3.Cavity = p*lt3.Cavity + (1.0-p)*lt2.Cavity
					//lt3 := obj.GetLiverPointFromKmPoint(path[l].XKm,path[l].YKm)
					//lt3.Cavity = ltroot.BaseElevation-(ltroot.BaseElevation-lt3.Cavity*lt3.BaseElevation)*p
				}
				

			} else {
				lt2.Cavity = lt.Cavity
			
				if lt.BaseElevation*lt.Cavity < lt2.BaseElevation*lt2.Cavity {
					lt2.Cavity = math.Max(lt.BaseElevation*lt.Cavity/lt2.BaseElevation,0)
				}
			}


		}
		/*
		for j := 0; j < len(path)-1; j++{
			lt := obj.GetLiverPointFromKmPoint(path[j+1].XKm,path[j+1].YKm)
			lt2 := obj.GetLiverPointFromKmPoint(path[j].XKm,path[j].YKm)

			if lt.BaseElevation*lt.Cavity < lt2.BaseElevation*lt2.Cavity {
				lt.Cavity = lt2.BaseElevation*lt2.Cavity/lt.BaseElevation
			}
		}
		*/
		
	}

	obj.LiverCheckIsAvailable = true
}

func (obj *LocalTerrainObject) CheckLiverCavityByKmPoint(xKm, yKm float64) float64{


	liver_interval_km := obj.WorldTerrain.Config.LiverCheckIntervalKm

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
