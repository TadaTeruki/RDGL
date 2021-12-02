package main

import(
	output "../Artograph/Output"
	artograph "../Artograph"
	
)

func main(){

	ats := artograph.NewDEM(18)
	ats.VerticalKm = 1000
	ats.HorizontalKm = -1
	ats.UnitKm = 1.5
	ats.Process("./example.png")
	
	output.WriteArtoDEMToPNGWithShadow("output.png", &ats, 500, -1, output.DefaultShadow(&ats))


}