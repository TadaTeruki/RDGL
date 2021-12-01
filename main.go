
package main

import(
	output "./Artograph/Output"
	artograph "./Artograph"
)


func main(){
	ats := artograph.NewTerrainSurface(700)
	ats.ElevationAbsM = 8000
	ats.UnitKm = 1
	ats.VerticalKm = 1000
	ats.HorizontalKm = 1000

	ats.Generate()
	output.WriteArtoTerrainSurfaceToPNG("data.png", &ats, 300, -1)

}