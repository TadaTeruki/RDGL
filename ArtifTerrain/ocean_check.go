package artif_terrain
import(
	"math"
)
type OceanPoint struct{
	XKm float64
	YKm float64
	IsOcean bool
}

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

func (obj *LocalTerrainObject) MarkOcean(x, y int) bool {


	if obj.GetElevationByKmPoint(obj.OceanTable[y][x].XKm, obj.OceanTable[y][x].YKm) >= 0.0 {
		return false
	}

	if  x == 0 || x == len(obj.OceanTable[0])-1 ||
		y == 0 || y == len(obj.OceanTable)-1 {
		obj.OceanTable[y][x].IsOcean = true
		return true
	}

	if  obj.OceanTable[y-1][x].IsOcean == true ||
		obj.OceanTable[y+1][x].IsOcean == true ||
		obj.OceanTable[y][x-1].IsOcean == true ||
		obj.OceanTable[y][x+1].IsOcean == true {

		obj.OceanTable[y][x].IsOcean = true
		return true
	}

	return false

}

func (obj *LocalTerrainObject) MakeOceanTable(){
	pond_interval_km := obj.WorldTerrain.Config.OceanCheckIntervalKm
	obj.OceanTable = make([][]OceanPoint, int(math.Floor(obj.NSKm/pond_interval_km)))
	checked := make([][]bool, len(obj.OceanTable))
	for y := 0; y<len(obj.OceanTable); y++{
	 	obj.OceanTable[y] = make([]OceanPoint, int(math.Floor(obj.WEKm/pond_interval_km)))
		checked[y] = make([]bool, len(obj.OceanTable[0]))
		for x := 0; x<len(obj.OceanTable[y]); x++{
			obj.OceanTable[y][x].XKm = float64(x)*pond_interval_km
			obj.OceanTable[y][x].YKm = float64(y)*pond_interval_km
			obj.OceanTable[y][x].IsOcean = false
			checked[y][x] = false
		}
	}

	var open, nwopen []Point
	open = append(open, MakePoint(0,0))

	for ;len(open) > 0;{
		nwopen = []Point{}
		for i := 0; i<len(open); i++ {
			if checked[open[i].Y][open[i].X] == true {
				continue
			}
			if open[i].Y-1 >= 0 && checked[open[i].Y-1][open[i].X] == false {
				mo := obj.MarkOcean(open[i].X,open[i].Y-1)
				if mo == true {nwopen = append(nwopen, MakePoint(open[i].X,open[i].Y-1)) }
			}
			if open[i].Y+1 < len(obj.OceanTable) && checked[open[i].Y+1][open[i].X] == false {
				mo := obj.MarkOcean(open[i].X,open[i].Y+1)
				if mo == true { nwopen = append(nwopen, MakePoint(open[i].X,open[i].Y+1)) }
			}

			if open[i].X-1 >= 0 && checked[open[i].Y][open[i].X-1] == false {
				mo := obj.MarkOcean(open[i].X-1,open[i].Y)
				if mo == true {nwopen = append(nwopen, MakePoint(open[i].X-1,open[i].Y)) }
			}
			if open[i].X+1 < len(obj.OceanTable[0]) && checked[open[i].Y][open[i].X+1] == false {
				mo := obj.MarkOcean(open[i].X+1,open[i].Y)
				if mo == true { nwopen = append(nwopen, MakePoint(open[i].X+1,open[i].Y)) }
			}
			checked[open[i].Y][open[i].X] = true
		}
		open = nwopen
	}

	
	/*
	for y := 0; y<len(obj.OceanTable); y++{
		for x := 0; x<len(obj.OceanTable[y]); x++{
			obj.MarkOcean(x, y);
		}
	}

	for y := len(obj.OceanTable)-1; y>=0; y--{
		for x := len(obj.OceanTable[y])-1; x>=0; x--{
			obj.MarkOcean(x, y);
		}
	}
	*/


	obj.OceanCheckIsAvailable = true
}

func (obj *LocalTerrainObject) CheckOceanByKmPoint(xKm, yKm float64) bool{
	pond_interval_km := obj.WorldTerrain.Config.OceanCheckIntervalKm
	x := int(math.Min(math.Round(xKm/pond_interval_km),float64(len(obj.OceanTable[0])-1)))
	y := int(math.Min(math.Round(yKm/pond_interval_km),float64(len(obj.OceanTable)-1)))
	return obj.OceanTable[y][x].IsOcean
}