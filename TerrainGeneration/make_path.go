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


