package terrain_generation
import(
	"math"
	//"fmt"
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

func (ocl OceanLayer) MarkOcean(obj *LocalTerrainObject, x, y int, elevation_level float64) bool {
	ocean_table := ocl.OceanTable
	
	if obj.GetElevationByKmPoint(ocean_table[y][x].XKm, ocean_table[y][x].YKm) > elevation_level {
		return false
	}

	if  x == 0 || x == len(ocean_table[0])-1 ||
		y == 0 || y == len(ocean_table)-1 {
		ocean_table[y][x].IsOcean = true
		ocean_table[y][x].ElevationLevel = elevation_level
		return true
	}

	if  ocean_table[y-1][x].IsOcean == true ||
		ocean_table[y+1][x].IsOcean == true ||
		ocean_table[y][x-1].IsOcean == true ||
		ocean_table[y][x+1].IsOcean == true {

		ocean_table[y][x].IsOcean = true
		ocean_table[y][x].ElevationLevel = elevation_level
		return true
	}

	return false

}

func (obj *LocalTerrainObject) MakeOceanTable(elevation_level float64){

	ocl := obj.OceanLayers[elevation_level]

	pond_interval_km := obj.WorldTerrain.Config.OceanCheckIntervalKm
	ocl.OceanTable = make([][]OceanPoint, int(math.Floor(obj.NSKm/pond_interval_km)))
	checked := make([][]bool, len(ocl.OceanTable))
	for y := 0; y<len(ocl.OceanTable); y++{
	 	ocl.OceanTable[y] = make([]OceanPoint, int(math.Floor(obj.WEKm/pond_interval_km)))
		checked[y] = make([]bool, len(ocl.OceanTable[0]))
		for x := 0; x<len(ocl.OceanTable[y]); x++{
			ocl.OceanTable[y][x].XKm = float64(x)*pond_interval_km
			ocl.OceanTable[y][x].YKm = float64(y)*pond_interval_km
			ocl.OceanTable[y][x].IsOcean = false
			checked[y][x] = false
		}
	}

	var open []Point
	

	
	for x := 0; x<len(ocl.OceanTable[0]); x+=len(ocl.OceanTable[0])/10{
		open = append(open, MakePoint(x,0))
		open = append(open, MakePoint(x,len(ocl.OceanTable)-1))
	}
	for y := 1; y<len(ocl.OceanTable)-1; y+=(len(ocl.OceanTable[0])-1)/10{
		open = append(open, MakePoint(0,y))
		open = append(open, MakePoint(len(ocl.OceanTable[0])-1,y))
	}

	

	for elv := elevation_level; elv <= obj.WorldTerrain.ElevationBaseM; elv += 10 {
		
		nxopen := make(map[Point]struct{})
		//fmt.Println(elv)
		for ;len(open) > 0;{
			nwopen := make(map[Point]struct{})

			for i := 0; i<len(open); i++ {

				if checked[open[i].Y][open[i].X] == true {
					continue
				}

				mos := false
				
				if open[i].Y-1 >= 0 && checked[open[i].Y-1][open[i].X] == false {
					mo := ocl.MarkOcean(obj, open[i].X,open[i].Y-1, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X,open[i].Y-1)] = struct{}{}
						mos = true
					}
				}
				if open[i].Y+1 < len(ocl.OceanTable) && checked[open[i].Y+1][open[i].X] == false {
					mo := ocl.MarkOcean(obj, open[i].X,open[i].Y+1, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X,open[i].Y+1)] = struct{}{}
						mos = true
					}
				}
	
				if open[i].X-1 >= 0 && checked[open[i].Y][open[i].X-1] == false {
					mo := ocl.MarkOcean(obj, open[i].X-1,open[i].Y, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X-1,open[i].Y)] = struct{}{}
						mos = true
					}
				}
				if open[i].X+1 < len(ocl.OceanTable[0]) && checked[open[i].Y][open[i].X+1] == false {
					mo := ocl.MarkOcean(obj, open[i].X+1,open[i].Y, elv)
					if mo == true {
						nwopen[MakePoint(open[i].X+1,open[i].Y)] = struct{}{} 
						mos = true
					}
				}

				if mos == false {
					nxopen[MakePoint(open[i].X,open[i].Y)] = struct{}{}
					//nxopen = append(nxopen, MakePoint(open[i].X,open[i].Y)

				} else {
					checked[open[i].Y][open[i].X] = true
				}
				
			}

			open = []Point{}

			for point, _ := range nwopen{
				open = append(open, point)

			}
			
		}

		open = []Point{}
		
		for point, _ := range nxopen{
			/*
			if checked[nxopen[i].Y][nxopen[i].X] == false {
				open = append(open, nxopen[i])
			}
			*/
			open = append(open, point)
			//checked[nxopen[i].Y][nxopen[i].X] = false
		}
		
		//open = nxopen
	}


	ocl.Available = true

	obj.OceanLayers[elevation_level] = ocl
}

func (ocl OceanLayer) GetOceanPointByKmPoint(obj *LocalTerrainObject, xKm, yKm float64) OceanPoint{
	pond_interval_km := obj.WorldTerrain.Config.OceanCheckIntervalKm
	x := int(math.Min(math.Round(xKm/pond_interval_km),float64(len(ocl.OceanTable[0])-1)))
	y := int(math.Min(math.Round(yKm/pond_interval_km),float64(len(ocl.OceanTable)-1)))
	return ocl.OceanTable[y][x]
}