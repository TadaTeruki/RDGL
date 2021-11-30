
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