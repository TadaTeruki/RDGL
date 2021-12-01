/*
Artograph/TerrainGeneration/ocean.go
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
	//"fmt"
)

type Point struct{
	X int
	Y int
}

func MakePoint(x, y int) Point{
	var p Point
	p.X = x
	p.Y = y
	return p
}

func (ocl *OceanLayer) MarkOcean(obj *LocalTerrainObject, x, y int, elevation_level float64) bool {
	ocean_table := ocl.OceanTable
	
	if obj.GetElevationByKmPoint(ocean_table[y][x].XKm, ocean_table[y][x].YKm) > elevation_level {
		return false
	}

	if  x == 0 || x == len(ocean_table[0])-1 ||
		y == 0 || y == len(ocean_table)-1 {
		ocean_table[y][x].IsOcean = true
		ocean_table[y][x].ElevationLevel = elevation_level
		return true
	}

	if  ocean_table[y-1][x].IsOcean == true ||
		ocean_table[y+1][x].IsOcean == true ||
		ocean_table[y][x-1].IsOcean == true ||
		ocean_table[y][x+1].IsOcean == true {

		ocean_table[y][x].IsOcean = true
		ocean_table[y][x].ElevationLevel = elevation_level
		return true
	}

	return false

}

func (obj *LocalTerrainObject) MakeOceanLayer(){

	ocl := &obj.OceanLayerObj

	pond_interval_km := obj.WorldTerrain.Config.OceanCheckIntervalKm
	ocl.OceanTable = make([][]OceanPoint, int(math.Floor(obj.NSKm/pond_interval_km)))
	checked := make([][]bool, len(ocl.OceanTable))
	for y := 0; y<len(ocl.OceanTable); y++{
	 	ocl.OceanTable[y] = make([]OceanPoint, int(math.Floor(obj.WEKm/pond_interval_km)))
		checked[y] = make([]bool, len(ocl.OceanTable[0]))
		for x := 0; x<len(ocl.OceanTable[y]); x++{
			ocl.OceanTable[y][x].XKm = float64(x)*pond_interval_km
			ocl.OceanTable[y][x].YKm = float64(y)*pond_interval_km
			ocl.OceanTable[y][x].IsOcean = false
			checked[y][x] = false
		}
	}

	var open []Point
	
	xl := int(obj.WEKm/obj.WorldTerrain.Config.TerrainLevelingIntervalKm)
	for x := 0; x<xl; x++ {
		open = append(open, MakePoint(len(ocl.OceanTable[0])*x/xl,0))
		open = append(open, MakePoint(len(ocl.OceanTable[0])*x/xl,len(ocl.OceanTable)-1))
	}
	
	yl := int(obj.NSKm/obj.WorldTerrain.Config.TerrainLevelingIntervalKm)
	for y := 1; y<yl-1; y ++ {
		open = append(open, MakePoint(0,len(ocl.OceanTable)*y/yl))
		open = append(open, MakePoint(len(ocl.OceanTable[0])-1,len(ocl.OceanTable)*y/yl))
	}

	

	for elv := -obj.WorldTerrain.ElevationBaseM; elv <= obj.WorldTerrain.ElevationBaseM; elv += obj.WorldTerrain.Config.TerrainLevelingHeightM {
		
		nxopen := make(map[Point]struct{})
		for ;len(open) > 0;{
			nwopen := make(map[Point]struct{})

			for i := 0; i<len(open); i++ {

				if len(checked) < open[i].Y || len(checked[0]) < open[i].X {
					continue
				}
				if checked[open[i].Y][open[i].X] == true {
					continue
				}

				mos := false
				
				if open[i].Y-1 >= 0 && checked[open[i].Y-1][open[i].X] == false {
					mo := ocl.MarkOcean(obj, open[i].X,open[i].Y-1, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X,open[i].Y-1)] = struct{}{}
						mos = true
					}
				}
				if open[i].Y+1 < len(ocl.OceanTable) && checked[open[i].Y+1][open[i].X] == false {
					mo := ocl.MarkOcean(obj, open[i].X,open[i].Y+1, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X,open[i].Y+1)] = struct{}{}
						mos = true
					}
				}
	
				if open[i].X-1 >= 0 && checked[open[i].Y][open[i].X-1] == false {
					mo := ocl.MarkOcean(obj, open[i].X-1,open[i].Y, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X-1,open[i].Y)] = struct{}{}
						mos = true
					}
				}
				if open[i].X+1 < len(ocl.OceanTable[0]) && checked[open[i].Y][open[i].X+1] == false {
					mo := ocl.MarkOcean(obj, open[i].X+1,open[i].Y, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X+1,open[i].Y)] = struct{}{} 
						mos = true
					}
				}

				if mos == false {
					nxopen[MakePoint(open[i].X,open[i].Y)] = struct{}{}

				} else {
					checked[open[i].Y][open[i].X] = true
				}
				
			}

			open = []Point{}

			for point, _ := range nwopen{
				open = append(open, point)

			}
			
		}

		open = []Point{}
		
		for point, _ := range nxopen{
			open = append(open, point)
		}
	}

	obj.OceanCheckIsAvailable = true
}

func (ocl OceanLayer) GetOceanPointByKmPoint(obj *LocalTerrainObject, xKm, yKm float64) OceanPoint{
	pond_interval_km := obj.WorldTerrain.Config.OceanCheckIntervalKm
	x := int(math.Min(math.Round(xKm/pond_interval_km),float64(len(ocl.OceanTable[0])-1)))
	y := int(math.Min(math.Round(yKm/pond_interval_km),float64(len(ocl.OceanTable)-1)))
	return ocl.OceanTable[y][x]
}