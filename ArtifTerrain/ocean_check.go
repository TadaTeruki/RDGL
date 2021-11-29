package artif_terrain
import(
	"math"
)
type OceanPoint struct{
	XKm float64
	YKm float64
	IsOcean bool
}

func (obj *LocalTerrainObject) MarkOcean(x, y int) {


	if obj.GetElevationByKmPoint(obj.OceanTable[y][x].XKm, obj.OceanTable[y][x].YKm) >= 0.0 {
		return
	}

	if  x == 0 || x == len(obj.OceanTable[0])-1 ||
		y == 0 || y == len(obj.OceanTable)-1 {
		obj.OceanTable[y][x].IsOcean = true
		return
	}

	if  obj.OceanTable[y-1][x].IsOcean == true ||
		obj.OceanTable[y+1][x].IsOcean == true ||
		obj.OceanTable[y][x-1].IsOcean == true ||
		obj.OceanTable[y][x+1].IsOcean == true {

		obj.OceanTable[y][x].IsOcean = true
		return
	}

}

func (obj *LocalTerrainObject) MakeOceanTable(){
	pond_interval_km := obj.WorldTerrain.Config.OceanCheckIntervalKm
	obj.OceanTable = make([][]OceanPoint, int(math.Floor(obj.NSKm/pond_interval_km)))
	for y := 0; y<len(obj.OceanTable); y++{
		obj.OceanTable[y] = make([]OceanPoint, int(math.Floor(obj.WEKm/pond_interval_km)))
		for x := 0; x<len(obj.OceanTable[y]); x++{
			obj.OceanTable[y][x].XKm = float64(x)*pond_interval_km
			obj.OceanTable[y][x].YKm = float64(y)*pond_interval_km
			obj.OceanTable[y][x].IsOcean = false
		}
	}

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

	obj.OceanCheckIsAvailable = true
}

func (obj *LocalTerrainObject) CheckOceanByKmPoint(xKm, yKm float64) bool{
	pond_interval_km := obj.WorldTerrain.Config.OceanCheckIntervalKm
	x := int(math.Min(math.Round(xKm/pond_interval_km),float64(len(obj.OceanTable[0])-1)))
	y := int(math.Min(math.Round(yKm/pond_interval_km),float64(len(obj.OceanTable)-1)))
	return obj.OceanTable[y][x].IsOcean
}