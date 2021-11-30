package artif_terrain

import(
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type KmPoint struct{
	XKm float64
	YKm float64
}

type PathPoint struct{
	point KmPoint
	parent KmPoint
	score float64
	flag int // 0:untreated 1:opened 2:closed
}

func MakeKmPoint(xKm, yKm float64) KmPoint{
	var r KmPoint
	r.XKm = xKm
	r.YKm = yKm
	return r
} 


func (obj *LocalTerrainObject) MakePath(
		start_x_km, start_y_km, interval_km float64,
		get_score func(x, y, sx, sy, elevation float64) float64,
		loop_out_condition func(x, y, sx, sy, elevation float64) bool)[]KmPoint{


	var path_point_list []PathPoint
	var path_point_address map[KmPoint]int

	path_point_address = make(map[KmPoint]int)

	var open_list *rbt.Tree
	open_list = rbt.NewWithIntComparator()

	var ptar KmPoint
	ptar.XKm = start_x_km
	ptar.YKm = start_y_km

	if ptar.XKm < 0 || ptar.YKm < 0 || ptar.XKm >= obj.WEKm || ptar.YKm >= obj.NSKm { return make([]KmPoint, 0) }

	get_score_id := func(x, y, sx, sy, elevation float64) int{
		return int(get_score(x, y, sx, sy, elevation))
	}

	open_path_point := func(point, parent KmPoint){

		if point.XKm < 0 || point.YKm < 0 || point.XKm >= obj.WEKm || point.YKm >= obj.NSKm { return }
		
		if _, ok := path_point_address[point]; ok {
			//exists
			return
		}

		//height := GetHeight(point.X, point.Y, data, property_list)
		elevation := obj.GetElevationByKmPoint(point.XKm, point.YKm)
		score := get_score(point.XKm, point.YKm, start_x_km, start_y_km, elevation)
		idscore := get_score_id(point.XKm, point.YKm, start_x_km, start_y_km, elevation)

		for {
			_, found := open_list.Get(idscore)
			if found == false {
				break
			}
			idscore++
		}
		var target PathPoint
		target.point = point
		target.parent = parent
		target.score = score
		target.flag = 1
		path_point_list = append(path_point_list, target)
		open_list.Put(idscore, len(path_point_list)-1)
		path_point_address[point] = len(path_point_list)-1

	}

	close_path_point := func(point KmPoint){
		elevation := obj.GetElevationByKmPoint(point.XKm, point.YKm)
		idscore := get_score_id(point.XKm, point.YKm, start_x_km, start_y_km, elevation)
		ad := path_point_address[point]
		path_point_list[ad].flag = 2
		for{
			cad, found := open_list.Get(idscore)
			if found == false || cad != ad {
				idscore++
				
				continue
			} else {
				open_list.Remove(idscore)
				break
			}
		}
	}

	open_path_point(ptar, ptar)

	for {
		
		if open_list.Size() == 0 { break }
		ad := open_list.Left().Value.(int)

		ptar = path_point_list[ad].point
		
		up := MakeKmPoint(ptar.XKm, ptar.YKm-interval_km)
		dw := MakeKmPoint(ptar.XKm, ptar.YKm+interval_km)
		lf := MakeKmPoint(ptar.XKm-interval_km, ptar.YKm)
		rg := MakeKmPoint(ptar.XKm+interval_km, ptar.YKm)

		open_path_point(up, ptar)
		open_path_point(dw, ptar)
		open_path_point(lf, ptar)
		open_path_point(rg, ptar)

		close_path_point(ptar)

		if loop_out_condition(ptar.XKm, ptar.YKm, start_x_km, start_y_km, obj.GetElevationByKmPoint(ptar.XKm, ptar.YKm)) == true{
			break
		}

	}

	var path []KmPoint

	btar := ptar
	
	for i := false;; i = true{
		if i == true && btar == ptar {
			break
		}
		path = append(path, ptar)
		btar = ptar
		ptar = path_point_list[path_point_address[ptar]].parent
	}

	return path
}