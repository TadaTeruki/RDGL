/*
examples/outline_dem.go
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
	output "../Artograph/Output"
	artograph "../Artograph"
	
)

func main(){

	dem := artograph.NewDEM(18)

	dem.ElevationAbsM = 8000
	dem.UnitKm = 2
	dem.VerticalKm = 1000
	// When use OutlineInterpolation, you can substitute -1 for either VerticalKm or HorizontalKm
	//  to apply the aspect ratio of outline image to the DEM.
	dem.HorizontalKm = -1
	dem.LevelingIntervalM = 5

	artograph.EnableProcessLog()

	// assign the filename of outline image.
	dem.Process("./example.png")

	// (filename, pointer of ArtoDEM object, width of PNG image, height of PNG image, shadow)
	output.WriteDEMtoPNGwithShadow("output.png", &dem, 300, -1, output.DefaultShadow(&dem))
}