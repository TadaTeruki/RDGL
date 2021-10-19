package artif_terrain

import "math"

type NEFPoint struct {
	NoiseLevel float64
	Elevation float64
}

func MakeNEFPoint(n,e float64) NEFPoint{
	var np NEFPoint
	np.NoiseLevel = n
	np.Elevation = e
	return np
}

func (obj *WorldTerrainObject) SetNEFPoint(){
	var NEFPointList = []NEFPoint{
		MakeNEFPoint(0.0, -obj.ElevationBaseM),
		MakeNEFPoint(0.4, -4000),
		MakeNEFPoint(0.65, -1000),
		MakeNEFPoint(0.7, 0),
		MakeNEFPoint(1.0, obj.ElevationBaseM),
	}
	obj.NEFPointList = NEFPointList
}

func (obj *WorldTerrainObject) NEF0(nlv float64) float64{
	//ax := obj.NEFPointList[0].NoiseLevel
	ay := obj.NEFPointList[0].Elevation
	bx := obj.NEFPointList[1].NoiseLevel
	by := obj.NEFPointList[1].Elevation
	return (1-math.Pow(nlv/bx,0.3))*(ay-by)+by
}

func (obj *WorldTerrainObject) NEF1(nlv float64) float64{
	bx := obj.NEFPointList[1].NoiseLevel
	by := obj.NEFPointList[1].Elevation
	cx := obj.NEFPointList[2].NoiseLevel
	cy := obj.NEFPointList[2].Elevation
	return -math.Pow((nlv-bx)/(cx-bx),2.7)*(by-cy)+by
}

func (obj *WorldTerrainObject) NEF2(nlv float64) float64{

	cx := obj.NEFPointList[2].NoiseLevel
	cy := obj.NEFPointList[2].Elevation
	dx := obj.NEFPointList[3].NoiseLevel
	dy := obj.NEFPointList[3].Elevation
	return math.Pow((nlv-dx)/(cx-dx),2.1)*(cy-dy)+dy
}

func (obj *WorldTerrainObject) NEF3(nlv float64) float64{
	dx := obj.NEFPointList[3].NoiseLevel
	dy := obj.NEFPointList[3].Elevation
	//ex := obj.NEFPointList[4].NoiseLevel
	ey := obj.NEFPointList[4].Elevation
	return math.Pow((1-nlv)/(1-dx),0.12)*(dy-ey)+ey
}


func (obj *WorldTerrainObject) GetElevationFromNoiseLevel(nlv float64)float64{
	if nlv < 0 { return -obj.ElevationBaseM }
	if nlv < obj.NEFPointList[1].NoiseLevel { return obj.NEF0(nlv) }
	if nlv < obj.NEFPointList[2].NoiseLevel { return obj.NEF1(nlv) }
	if nlv < obj.NEFPointList[3].NoiseLevel { return obj.NEF2(nlv) }
	if nlv < 1.0 { return obj.NEF3(nlv) }
	return obj.ElevationBaseM
}
