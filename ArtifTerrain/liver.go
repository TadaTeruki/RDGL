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

func (obj *LocalTerrainObject) MakeLiverTable(){
	liver_interval_km := obj.WorldTerrain.Config.LiverCheckIntervalKm
	obj.LiverTable = make([][]LiverPoint, int(math.Ceil(obj.NSKm/liver_interval_km)))
	//checked := make([][]bool, len(obj.LiverTable))
	var lv_order []Point

	for y := 0; y<len(obj.LiverTable); y++{
		obj.LiverTable[y] = make([]LiverPoint, int(math.Ceil(obj.WEKm/liver_interval_km)))
		//checked[y] = make([]bool, len(obj.LiverTable[0]))
		for x := 0; x<len(obj.LiverTable[y]); x++{
			obj.LiverTable[y][x].XKm = float64(x)*liver_interval_km
			obj.LiverTable[y][x].YKm = float64(y)*liver_interval_km
			obj.LiverTable[y][x].Direction = DIRECTION_NONE
			obj.LiverTable[y][x].Cavity = 1.0
			obj.LiverTable[y][x].BaseElevation = obj.GetElevationByKmPoint(obj.LiverTable[y][x].XKm, obj.LiverTable[y][x].YKm)
			//checked[y][x] = false
			if obj.LiverTable[y][x].BaseElevation >= 0.0 {
				lv_order = append(lv_order, MakePoint(x, y))
			}
			
		}
	}

	sort.Slice(lv_order, func(i, j int) bool {
		return obj.LiverTable[lv_order[i].Y][lv_order[i].X].BaseElevation >
			    obj.LiverTable[lv_order[j].Y][lv_order[j].X].BaseElevation
	})
	/*
	for i := 0; i<len(lv_order); i++{

	}
	*/

	get_liver_point_from_km_point := func(xKm, yKm float64) *LiverPoint{
		lt := &obj.LiverTable[int(math.Min(yKm/liver_interval_km,float64(len(obj.LiverTable)-1)))][int(math.Min(xKm/liver_interval_km,float64(len(obj.LiverTable[0])-1)))]
		return lt
	}

	path_score := func(xKm, yKm, sxKm, syKm, elevation float64) float64{
		return elevation
	}

	loop_out_condition := func(xKm, yKm, sxKm, syKm, elevation float64) bool{

		lt := get_liver_point_from_km_point(xKm, yKm)
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
			lt := get_liver_point_from_km_point(path[j].XKm,path[j].YKm)
			xd := path[j].XKm - path[j-1].XKm
			yd := path[j].YKm - path[j-1].YKm

			if xd == 0 && yd > 0 { lt.Direction = DIRECTION_NORTH }
			if xd > 0 && yd == 0 { lt.Direction = DIRECTION_WEST }
			if xd == 0 && yd < 0 { lt.Direction = DIRECTION_SOUTH }
			if xd < 0 && yd == 0 { lt.Direction = DIRECTION_EAST }

			lt2 := get_liver_point_from_km_point(path[j-1].XKm,path[j-1].YKm)

			if lt.BaseElevation*lt.Cavity > lt2.BaseElevation*lt2.Cavity {
				lt.Cavity = lt2.BaseElevation*lt2.Cavity/lt.BaseElevation
			}

		}
	}

	obj.LiverCheckIsAvailable = true
}
