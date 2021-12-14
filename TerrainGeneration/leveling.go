/*
github.com/TadaTeruki/RDGL/TerrainGeneration/leveling.go
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
	utility "github.com/TadaTeruki/RDGL/Utility"
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

func (ocl *UnitLayer) MarkLeveling(obj *LocalTerrainObject, x, y int, elevation_level float64) bool {
	leveling_table := ocl.Table

	if obj.GetElevationByKmPoint(leveling_table[y][x].XKm, leveling_table[y][x].YKm) >= elevation_level {
		return false
	}

	if  ( y > 0 						&& leveling_table[y-1][x].IsLeveling == true ) ||
		( y < len(leveling_table)-1		&& leveling_table[y+1][x].IsLeveling == true ) ||
		( x > 0 						&& leveling_table[y][x-1].IsLeveling == true ) ||
		( x < len(leveling_table[0])-1	&& leveling_table[y][x+1].IsLeveling == true ) {

		leveling_table[y][x].IsLeveling = true
		if elevation_level > 0.0{
			leveling_table[y][x].ElevationLevel = elevation_level
		}
		
		return true
	}

	return false

}



func (obj *LocalTerrainObject) MakeUnitLayer(){

	ocl := &obj.UnitLayerObj

	unit_km := obj.WorldTerrain.Config.UnitKm
	ocl.Table = make([][]UnitPoint, int(math.Ceil(obj.NSKm/unit_km)))
	checked := make([][]bool, len(ocl.Table))
	for y := 0; y<len(ocl.Table); y++{
	 	ocl.Table[y] = make([]UnitPoint, int(math.Ceil(obj.WEKm/unit_km)))
		checked[y] = make([]bool, len(ocl.Table[0]))
		for x := 0; x<len(ocl.Table[y]); x++{
			ocl.Table[y][x].XKm = float64(x)*unit_km
			ocl.Table[y][x].YKm = float64(y)*unit_km

			// leveling
			ocl.Table[y][x].IsLeveling = false
			ocl.Table[y][x].ElevationLevel = obj.GetElevationByKmPoint(ocl.Table[y][x].XKm, ocl.Table[y][x].YKm)

			//liver
			ocl.Table[y][x].Direction = DIRECTION_NONE
			ocl.Table[y][x].Cavity = 1.0
			ocl.Table[y][x].RootDistKm = 0.0

			checked[y][x] = false
		}
	}

	//obj.LiverCheckIsAvailable = true
}
func (obj *LocalTerrainObject) Leveling(){

	ocl := &obj.UnitLayerObj

	unit_km := obj.WorldTerrain.Config.UnitKm
	checked := make([][]bool, len(ocl.Table))
	for y := 0; y<len(ocl.Table); y++{
		checked[y] = make([]bool, len(ocl.Table[0]))
		for x := 0; x<len(ocl.Table[y]); x++{
			checked[y][x] = false
		}
	}

	var open []Point

	shelf_elevation_m := obj.WorldTerrain.Config.ContShelfElevationProportion * obj.WorldTerrain.ElevationAbsM

	for i := 0 ; i< len(obj.RootList); i++{
		ix := int(math.Floor(obj.RootList[i].XKm/unit_km))
		iy := int(math.Floor(obj.RootList[i].YKm/unit_km))
		if ix >= len(ocl.Table[0]) {
			ix = len(ocl.Table[0])-1
		}
		if iy >= len(ocl.Table) {
			iy = len(ocl.Table)-1
		}

		open = append(open, MakePoint(ix, iy))
		ocl.Table[iy][ix].IsLeveling = true
	}

	utility.EchoProcessPercentage("Leveling", 0)
	checked_sum := 0.0
	checked_all := float64(len(ocl.Table[0])*len(ocl.Table))

	for elv := shelf_elevation_m; elv <= obj.WorldTerrain.ElevationAbsM; elv += obj.WorldTerrain.Config.LevelingHeightM {
		
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
				if open[i].Y+1 < len(ocl.Table) && checked[open[i].Y+1][open[i].X] == false {
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
				if open[i].X+1 < len(ocl.Table[0]) && checked[open[i].Y][open[i].X+1] == false {
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


func (ocl UnitLayer) GetUnitPointByKmPoint(obj *LocalTerrainObject, xKm, yKm float64) *UnitPoint{
	unit_km := obj.WorldTerrain.Config.UnitKm
	x := int(math.Min(math.Round(xKm/unit_km),float64(len(ocl.Table[0])-1)))
	y := int(math.Min(math.Round(yKm/unit_km),float64(len(ocl.Table)-1)))
	return &ocl.Table[y][x]
}

func (ocl UnitLayer) GetLevelingElevationByKmPoint(obj *LocalTerrainObject, xKm, yKm float64) float64{
	unit_km := obj.WorldTerrain.Config.UnitKm
	x := xKm/unit_km
	y := yKm/unit_km
	fx := math.Min(math.Max(math.Floor(x), 0),float64(len(ocl.Table[0])-1))
	fy := math.Min(math.Max(math.Floor(y), 0),float64(len(ocl.Table)-1))
	cx := math.Min(fx+1, float64(len(ocl.Table[0])-1))
	cy := math.Min(fy+1, float64(len(ocl.Table)-1))

	ffsc := (x-fx)*(y-fy)
	cfsc := (cx-x)*(y-fy)
	fcsc := (x-fx)*(cy-y)
	ccsc := (cx-x)*(cy-y)

	return  ccsc*ocl.Table[int(fy)][int(fx)].ElevationLevel +
			cfsc*ocl.Table[int(cy)][int(fx)].ElevationLevel +
			fcsc*ocl.Table[int(fy)][int(cx)].ElevationLevel +
			ffsc*ocl.Table[int(cy)][int(cx)].ElevationLevel

}
