
/*
examples/detailed_hello_dem.go
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
	ats := artograph.NewDEM(16)

	// [ats.ElevationAbsM] Absolute value of the maximum/minimum elevation(Meter)
	// (Example : ats.ElevationAbsM = 8000 -> the minimum/maximum elevation : -8000 ~ 8000)
	// (Recommendation : ats.ElevationAbsM >= 4000)
	ats.ElevationAbsM = 8000

	// [ats.UnitKm] The interval of datum point(Km) for generating liver & leveling terrain
	ats.UnitKm = 2

	// [ats.VerticalKm] The vertical width of the DEM data (Km)
	ats.VerticalKm = 1000

	// [ats.HorizontalKm] The horizontal width of the DEM data (Km)
	ats.HorizontalKm = 1000

	// [ats.LevelingIntervalM] The elevation unit(Meter)
	ats.LevelingIntervalM = 5

	ats.Generate()
	// ---
	
	for yKm := 0.0; yKm < ats.HorizontalKm; yKm += ats.HorizontalKm/50.0{
		str := ""
		for xKm := 0.0; xKm < ats.VerticalKm; xKm += ats.HorizontalKm/50.0{

			elevation, _ := ats.GetElevationByKmPoint(xKm, yKm)

			if elevation >= 800{
				str += "@@"
			} else if elevation >= 300 {
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
