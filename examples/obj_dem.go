package main

import(
	output "../Artograph/Output"
	artograph "../Artograph"
)

func main(){
	dem := artograph.NewDEM(25)
	dem.ElevationAbsM = 8000
	dem.UnitKm = 2
	dem.VerticalKm = 1000
	dem.HorizontalKm = 1000
	dem.LevelingIntervalM = 5

	artograph.EnableProcessLog()

	dem.Generate()

	// (filename, pointer of ArtoDEM object, width of PNG image, height of PNG image, 
	//        scale of elevation, whether Z-axis is vertical or not ['true' -> Z-axis is vertical, 'false' -> Y-axis is vertical ] )
	output.WriteDEMtoOBJ("output.obj", &dem, 100, -1, 5.0, false)

}