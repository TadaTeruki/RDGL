/*
Artograph/TerrainGeneration/nef.go
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

package terrain_generation

import "math"

// NEF
// hypsographic curve expression

func MakeNEFPoint(n,e float64) NEFPoint{
	var np NEFPoint
	np.NoiseLevel = n
	np.Elevation = e
	return np
}

func (obj *WorldTerrainObject) SetNEFPoint(){
	var NEFPointList = []NEFPoint{
		MakeNEFPoint(0.0, -obj.ElevationBaseM),
		MakeNEFPoint(obj.Config.StandardOceanProportion*0.7, -obj.ElevationBaseM*0.5),
		MakeNEFPoint(obj.Config.StandardOceanProportion*0.85, -obj.ElevationBaseM*0.1),
		MakeNEFPoint(obj.Config.StandardOceanProportion, 0.0),
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
	return -math.Pow((nlv-bx)/(cx-bx),2.3)*(by-cy)+by
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
	return math.Pow((1-nlv)/(1-dx),0.15)*(dy-ey)+ey
}


func (obj *WorldTerrainObject) GetElevationFromNoiseLevel(nlv float64)float64{
	if nlv < 0 { return -obj.ElevationBaseM }
	if nlv < obj.NEFPointList[1].NoiseLevel { return obj.NEF0(nlv) }
	if nlv < obj.NEFPointList[2].NoiseLevel { return obj.NEF1(nlv) }
	if nlv < obj.NEFPointList[3].NoiseLevel { return obj.NEF2(nlv) }
	if nlv < 1.0 { return obj.NEF3(nlv) }
	return obj.ElevationBaseM
}
