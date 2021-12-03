/*
Artograph/TerrainGeneration/leveling.go
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
	utility "../Utility"
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

func (ocl *LevelingLayer) MarkLeveling(obj *LocalTerrainObject, x, y int, elevation_level float64) bool {
	leveling_table := ocl.LevelingTable
	
	if obj.GetElevationByKmPoint(leveling_table[y][x].XKm, leveling_table[y][x].YKm) > elevation_level {
		return false
	}

	if  x == 0 || x == len(leveling_table[0])-1 ||
		y == 0 || y == len(leveling_table)-1 {
		leveling_table[y][x].IsLeveling = true
		leveling_table[y][x].ElevationLevel = elevation_level
		return true
	}

	if  leveling_table[y-1][x].IsLeveling == true ||
		leveling_table[y+1][x].IsLeveling == true ||
		leveling_table[y][x-1].IsLeveling == true ||
		leveling_table[y][x+1].IsLeveling == true {

		leveling_table[y][x].IsLeveling = true
		leveling_table[y][x].ElevationLevel = elevation_level
		return true
	}

	return false

}

func (obj *LocalTerrainObject) MakeLevelingLayer(){

	ocl := &obj.LevelingLayerObj

	pond_interval_km := obj.WorldTerrain.Config.LevelingIntervalKm
	ocl.LevelingTable = make([][]LevelingPoint, int(math.Floor(obj.NSKm/pond_interval_km)))
	checked := make([][]bool, len(ocl.LevelingTable))
	for y := 0; y<len(ocl.LevelingTable); y++{
	 	ocl.LevelingTable[y] = make([]LevelingPoint, int(math.Floor(obj.WEKm/pond_interval_km)))
		checked[y] = make([]bool, len(ocl.LevelingTable[0]))
		for x := 0; x<len(ocl.LevelingTable[y]); x++{
			ocl.LevelingTable[y][x].XKm = float64(x)*pond_interval_km
			ocl.LevelingTable[y][x].YKm = float64(y)*pond_interval_km
			ocl.LevelingTable[y][x].IsLeveling = false
			checked[y][x] = false
		}
	}

	var open []Point
	
	xl := int(obj.WEKm/obj.WorldTerrain.Config.LevelingStartPointIntervalKm)
	for x := 0; x<xl; x++ {
		open = append(open, MakePoint(len(ocl.LevelingTable[0])*x/xl,0))
		open = append(open, MakePoint(len(ocl.LevelingTable[0])*x/xl,len(ocl.LevelingTable)-1))
	}
	
	yl := int(obj.NSKm/obj.WorldTerrain.Config.LevelingStartPointIntervalKm)
	for y := 1; y<yl-1; y ++ {
		open = append(open, MakePoint(0,len(ocl.LevelingTable)*y/yl))
		open = append(open, MakePoint(len(ocl.LevelingTable[0])-1,len(ocl.LevelingTable)*y/yl))
	}

	utility.EchoProcessPercentage("Leveling", 0)
	checked_sum := 0.0
	checked_all := float64(len(ocl.LevelingTable[0])*len(ocl.LevelingTable))

	for elv := -obj.WorldTerrain.ElevationBaseM; elv <= obj.WorldTerrain.ElevationBaseM; elv += obj.WorldTerrain.Config.LevelingHeightM {
		
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
					mo := ocl.MarkLeveling(obj, open[i].X,open[i].Y-1, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X,open[i].Y-1)] = struct{}{}
						mos = true
					}
				}
				if open[i].Y+1 < len(ocl.LevelingTable) && checked[open[i].Y+1][open[i].X] == false {
					mo := ocl.MarkLeveling(obj, open[i].X,open[i].Y+1, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X,open[i].Y+1)] = struct{}{}
						mos = true
					}
				}
	
				if open[i].X-1 >= 0 && checked[open[i].Y][open[i].X-1] == false {
					mo := ocl.MarkLeveling(obj, open[i].X-1,open[i].Y, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X-1,open[i].Y)] = struct{}{}
						mos = true
					}
				}
				if open[i].X+1 < len(ocl.LevelingTable[0]) && checked[open[i].Y][open[i].X+1] == false {
					mo := ocl.MarkLeveling(obj, open[i].X+1,open[i].Y, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X+1,open[i].Y)] = struct{}{} 
						mos = true
					}
				}

				if mos == false {
					nxopen[MakePoint(open[i].X,open[i].Y)] = struct{}{}

				} else {
					checked[open[i].Y][open[i].X] = true

					checked_before_sum := checked_sum
					checked_sum++
					if math.Floor(checked_before_sum/checked_all*10) != math.Floor(checked_sum/checked_all*10) {
						utility.EchoProcessPercentage("Leveling", checked_sum/checked_all)
					}

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

	utility.EchoProcessEnd("Leveling")

	obj.LevelingCheckIsAvailable = true
}

func (ocl LevelingLayer) GetLevelingPointByKmPoint(obj *LocalTerrainObject, xKm, yKm float64) LevelingPoint{
	pond_interval_km := obj.WorldTerrain.Config.LevelingIntervalKm
	x := int(math.Min(math.Round(xKm/pond_interval_km),float64(len(ocl.LevelingTable[0])-1)))
	y := int(math.Min(math.Round(yKm/pond_interval_km),float64(len(ocl.LevelingTable)-1)))
	return ocl.LevelingTable[y][x]
}