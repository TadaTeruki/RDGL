
/*
examples/write_to_png_with_shadow.go
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
	output "../RDGL/Output"
	rdg "../RDGL"
)

func main(){
	dem := rdg.NewDEM(25)

	rdg.EnableProcessLog()

	dem.Generate()
	
	// (filename, pointer of ArtoDEM object, width of PNG image, height of PNG image, shadow)
	// width|height, when either of them is -1,
	//  will be applied proper value according to the height|width and the aspect ratio of DEM.
	output.WriteDEMtoPNGwithShadow("result.png", &dem, 300, -1, output.DefaultShadow(&dem))

}