
/*
examples/hello_dem.go
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
	rdg "github.com/TadaTeruki/RDGL"
	"fmt"
)

func main(){
	// Use 'NewDEM' to make a new DEM (assigning the seed value).
	dem := rdg.NewDEM(14)
	// Start DEM generation.
	dem.Generate()
	// DEM generation will be completed automatically. Easy!

	// ---

	// Enabling 'ProcessLog', you can check the stage of process during DEM generation.  
	// If you think it unnecessary, comment out the below row to disable ProcessLog.
	rdg.EnableProcessLog()

	// Now let's output the DEM data onto your terminal.
	// 'HorizontalKm' means the horizontal width of the DEM data (unit:Km)
	
	for yKm := 0.0; yKm < dem.HorizontalKm; yKm += dem.HorizontalKm/50.0{
		str := ""
		// 'VerticalKm' means the vertical width of the DEM data (unit:Km)
		for xKm := 0.0; xKm < dem.VerticalKm; xKm += dem.HorizontalKm/50.0{

			// Get elevation (unit:Meter) by Km coodinate
			elevation, _ := dem.GetElevationByKmPoint(xKm, yKm)

			// The maximum elevation is 8000M, and the minimum elvation is -8000M by default.
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

	// Great!
}
