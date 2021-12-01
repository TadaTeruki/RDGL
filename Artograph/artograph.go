/*
Artograph/artograph.go
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


package artograph

import(
	terrain "./TerrainGeneration"
	"fmt"
	"math"
)
type ArtoTerrainSurface struct{
	Seed 				int64
	ElevationAbsM		float64
	ElevationIntervalM	float64
	UnitKm				float64
	VerticalKm			float64
	HorizontalKm		float64


	side_width		float64
	w_obj terrain.WorldTerrainObject
	l_obj terrain.LocalTerrainObject
}

func (ats *ArtoTerrainSurface) default_ats(){
	ats.Seed = 0
	ats.ElevationAbsM = 8000
	ats.UnitKm = 1
	ats.VerticalKm = 500
	ats.HorizontalKm = 1000
	ats.ElevationIntervalM = 5
	ats.side_width = 3
}

const ARTO_ERROR_1 = "<Artograph> Error : The coodinate assigned to 'GetElevationByKmPoint' is outside of the TerrainSurface"

func NewTerrainSurface(seed int64) ArtoTerrainSurface{
	var ats ArtoTerrainSurface
	ats.default_ats()
	ats.Seed = seed
	return ats
}

func (ats *ArtoTerrainSurface) Generate(){

	ats.l_obj.NSKm = ats.VerticalKm + ats.UnitKm * ats.side_width * 2
	ats.l_obj.WEKm = ats.HorizontalKm + ats.UnitKm * ats.side_width * 2

	ats.w_obj.NSKm = ats.l_obj.NSKm * ats.side_width
	ats.w_obj.WEKm = ats.l_obj.WEKm * ats.side_width

	ats.w_obj.ElevationBaseM = ats.ElevationAbsM
	ats.w_obj.Config = terrain.GetGlobalConfig()
	ats.w_obj.Config.Seed = ats.Seed
	ats.w_obj.Config.LiverCheckIntervalKm = ats.UnitKm
	ats.w_obj.Config.OceanCheckIntervalKm = ats.UnitKm
	ats.w_obj.Config.TerrainLevelingHeightM = ats.ElevationIntervalM
	ats.w_obj.Config.TerrainLevelingIntervalKm = math.Max(ats.l_obj.NSKm, ats.l_obj.WEKm)/100

	ats.w_obj.SetNEFPoint()
	ats.w_obj.MakeWorldTerrain()


	ats.l_obj.WorldTerrain = &ats.w_obj
	ats.l_obj.MakeLocalTerrain()
	
}

func (ats *ArtoTerrainSurface) GetElevationByKmPoint(xKm, yKm float64) (float64, error){
	if xKm < 0 || yKm < 0 || xKm > ats.HorizontalKm || yKm > ats.VerticalKm {
		err := fmt.Errorf(ARTO_ERROR_1)
		return 0, err
	} 
	dxKm := xKm + ats.UnitKm * ats.side_width
	dyKm := yKm + ats.UnitKm * ats.side_width

	return ats.l_obj.GetElevationByKmPoint(dxKm, dyKm), nil
}