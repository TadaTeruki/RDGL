/*
github.com/TadaTeruki/RDGL/TerrainGeneration/hcf.go
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

// HCF
// hypsographic curve function

func MakeHCFPoint(n,e float64) HCFPoint{
	var np HCFPoint
	np.NoiseLevel = n
	np.Elevation = e
	return np
}

func (obj *WorldTerrainObject) SetHCFPoint(){
	ocean_proportion := 1.0-obj.Config.StandardLandProportion
	var HCFPointList = []HCFPoint{
		MakeHCFPoint(0.0, -obj.ElevationAbsM),
		MakeHCFPoint(ocean_proportion*0.7, -obj.ElevationAbsM*0.5),
		MakeHCFPoint(ocean_proportion*0.85, -obj.ElevationAbsM*0.1),
		MakeHCFPoint(ocean_proportion, 0.0),
		MakeHCFPoint(1.0, obj.ElevationAbsM),
	}
	obj.HCFPointList = HCFPointList
}

func (obj *WorldTerrainObject) HCF0(nlv float64) float64{
	//ax := obj.HCFPointList[0].NoiseLevel
	ay := obj.HCFPointList[0].Elevation
	bx := obj.HCFPointList[1].NoiseLevel
	by := obj.HCFPointList[1].Elevation
	return (1-math.Pow(nlv/bx,0.3))*(ay-by)+by
}

func (obj *WorldTerrainObject) HCF1(nlv float64) float64{
	bx := obj.HCFPointList[1].NoiseLevel
	by := obj.HCFPointList[1].Elevation
	cx := obj.HCFPointList[2].NoiseLevel
	cy := obj.HCFPointList[2].Elevation
	return -math.Pow((nlv-bx)/(cx-bx),2.3)*(by-cy)+by
}

func (obj *WorldTerrainObject) HCF2(nlv float64) float64{

	cx := obj.HCFPointList[2].NoiseLevel
	cy := obj.HCFPointList[2].Elevation
	dx := obj.HCFPointList[3].NoiseLevel
	dy := obj.HCFPointList[3].Elevation
	return math.Pow((nlv-dx)/(cx-dx),2.1)*(cy-dy)+dy
}

func (obj *WorldTerrainObject) HCF3(nlv float64) float64{
	dx := obj.HCFPointList[3].NoiseLevel
	dy := obj.HCFPointList[3].Elevation
	//ex := obj.HCFPointList[4].NoiseLevel
	ey := obj.HCFPointList[4].Elevation
	return math.Pow((1-nlv)/(1-dx),0.15)*(dy-ey)+ey
}


func (obj *WorldTerrainObject) GetElevationFromNoiseLevel(nlv float64)float64{
	if nlv < 0 { return -obj.ElevationAbsM }
	if nlv < obj.HCFPointList[1].NoiseLevel { return obj.HCF0(nlv) }
	if nlv < obj.HCFPointList[2].NoiseLevel { return obj.HCF1(nlv) }
	if nlv < obj.HCFPointList[3].NoiseLevel { return obj.HCF2(nlv) }
	if nlv < 1.0 { return obj.HCF3(nlv) }
	return obj.ElevationAbsM
}
