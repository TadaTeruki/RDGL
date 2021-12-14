/*
github.com/TadaTeruki/RDGL/TerrainGeneration/make_path.go
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
	prque "github.com/TadaTeruki/RDGL/PriorityQueue"
)

// optimal best path algorithm
func (obj *LocalTerrainObject) MakePath(
		start_x_km, start_y_km, interval_km float64,
		get_score func(x, y, sx, sy, elevation float64) float64,
		loop_out_condition func(x, y, sx, sy, elevation float64) bool) []KmPoint{

	var open prque.PriorityQueue

	var ptar KmPoint
	ptar.XKm = start_x_km
	ptar.YKm = start_y_km

	parent_index := make(map[KmPoint]KmPoint)

	if ptar.XKm < 0 || ptar.YKm < 0 || ptar.XKm >= obj.WEKm || ptar.YKm >= obj.NSKm { return make([]KmPoint, 0) }
	//-get_score(ptar.XKm, ptar.YKm, start_x_km, start_y_km, obj.GetElevationByKmPoint(ptar.XKm, ptar.YKm)))
	open.Push(prque.MakeObject(ptar, 0))
	parent_index[ptar] = ptar

	open_path_point := func(point, parent KmPoint){

		if point.XKm < 0 || point.YKm < 0 || point.XKm >= obj.WEKm || point.YKm >= obj.NSKm { return }
		if _, existence := parent_index[point]; existence == true { return }
		
		elevation := obj.GetElevationByKmPoint(point.XKm, point.YKm)
		score := -get_score(point.XKm, point.YKm, start_x_km, start_y_km, elevation)
		
		open.Push(prque.MakeObject(point, score))
		parent_index[point] = parent

	}

	for {
		psrc := open.GetFront()
		ptar = psrc.Value.(KmPoint)

		if loop_out_condition(ptar.XKm, ptar.YKm, start_x_km, start_y_km, obj.GetElevationByKmPoint(ptar.XKm, ptar.YKm)) == true{
			break
		}

		open.Pop()

		up := MakeKmPoint(ptar.XKm, ptar.YKm-interval_km)
		dw := MakeKmPoint(ptar.XKm, ptar.YKm+interval_km)
		lf := MakeKmPoint(ptar.XKm-interval_km, ptar.YKm)
		rg := MakeKmPoint(ptar.XKm+interval_km, ptar.YKm)

		open_path_point(up, ptar)
		open_path_point(dw, ptar)
		open_path_point(lf, ptar)
		open_path_point(rg, ptar)
	}

	var path []KmPoint

	path = append(path, ptar)

	btar := ptar
	ptar = parent_index[ptar]
	
	for ; btar != ptar ;{
		path = append(path, ptar)
		btar = ptar
		ptar = parent_index[ptar]
	}

	return path
}



/*
package terrain_generation

import(
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

// optimal best path algorithm
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
*/