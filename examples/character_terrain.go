
/*
examples/character_terrain.go
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
package main

import(
	artograph "../Artograph"
	"fmt"
)


func main(){
	ats := artograph.NewTerrainSurface(700)
	ats.ElevationAbsM = 8000
	ats.UnitKm = 1
	ats.VerticalKm = 1000
	ats.HorizontalKm = 1000

	ats.Generate()

	// output
	for yKm := 0.0; yKm < ats.HorizontalKm; yKm+= 20.0{
		str := ""
		for xKm := 0.0; xKm < ats.VerticalKm; xKm+= 20.0{
			elevation, _ := ats.GetElevationByKmPoint(xKm, yKm)
			if elevation >= 1500{
				str += "@@"
			} else if elevation >= 500 {
				str += "[]"
			} else if elevation >= 0 {
				str += "__"
			} else {
				str += " "
			}
			
		}
		fmt.Println(str)
	}

}