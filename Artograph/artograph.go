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
	outline "./Outline"
	utility "./Utility"
	"fmt"
	"math"
)

type ArtoDEM struct{
	Seed 				int64
	ElevationAbsM		float64
	LevelingIntervalM	float64
	UnitKm				float64
	VerticalKm			float64
	HorizontalKm		float64

	side_width		float64
	w_obj terrain.WorldTerrainObject
	l_obj terrain.LocalTerrainObject
}

func (dem *ArtoDEM) default_dem(){
	dem.Seed = 0
	dem.ElevationAbsM = 8000
	dem.UnitKm = 2
	dem.VerticalKm = 1000
	dem.HorizontalKm = 1000
	dem.LevelingIntervalM = 5
	dem.side_width = 3.0
}

const ARTO_ERROR_1 = "<Artograph> Error : The coodinate assigned to 'GetElevationByKmPoint' is outside of the DEM"

func NewDEM(seed int64) ArtoDEM{
	var dem ArtoDEM
	dem.default_dem()
	dem.Seed = seed
	return dem
}

func (dem *ArtoDEM) config(){

	dem.w_obj.Config = terrain.GetInternalConfig()

	dem.w_obj.Config.MapSideWidthKm = dem.UnitKm * dem.side_width

	dem.l_obj.NSKm = dem.VerticalKm + dem.w_obj.Config.MapSideWidthKm * 2
	dem.l_obj.WEKm = dem.HorizontalKm + dem.w_obj.Config.MapSideWidthKm * 2

	dem.w_obj.NSKm = dem.l_obj.NSKm * dem.side_width * 2
	dem.w_obj.WEKm = dem.l_obj.WEKm * dem.side_width * 2

	dem.w_obj.ElevationBaseM = dem.ElevationAbsM
	
	dem.w_obj.Config.Seed = dem.Seed
	dem.w_obj.Config.LiverIntervalKm = dem.UnitKm
	dem.w_obj.Config.LevelingIntervalKm = dem.UnitKm
	dem.w_obj.Config.LevelingHeightM = dem.LevelingIntervalM
	dem.w_obj.Config.LevelingStartPointIntervalKm = math.Max(dem.l_obj.NSKm, dem.l_obj.WEKm)/100

	dem.l_obj.WorldTerrain = &dem.w_obj

}

func (dem *ArtoDEM) Generate(){

	dem.config()
	dem.w_obj.SetNEFPoint()
	dem.w_obj.MakeWorldTerrain()

	dem.l_obj.MakeLocalTerrain()
	

	dem.l_obj.TransformProcess(true, true)
	
}

func (dem *ArtoDEM) Process(file string){
	data_fw, data_fh := outline.GetImageScale(file)

	if dem.HorizontalKm < 0 {
		dem.HorizontalKm = dem.VerticalKm/data_fh*data_fw
	}
	if dem.VerticalKm < 0 {
		dem.VerticalKm = dem.HorizontalKm/data_fw*data_fh
	}

	dem.config()
	dem.w_obj.SetNEFPoint()
	dem.w_obj.Config.NoizeMinPersistence = dem.w_obj.Config.OutlineNoizeMinPersistence
	dem.w_obj.Config.NoizeMaxPersistence = dem.w_obj.Config.OutlineNoizeMaxPersistence
	dem.w_obj.MakeWorldTerrain()
	
	outline.LoadTerrainData(&dem.l_obj, &dem.w_obj.Config, file)
	
	dem.l_obj.TransformProcess(true, true)
}

func (dem *ArtoDEM) GetElevationByKmPoint(xKm, yKm float64) (float64, error){
	if xKm < 0 || yKm < 0 || xKm > dem.HorizontalKm || yKm > dem.VerticalKm {
		err := fmt.Errorf(ARTO_ERROR_1)
		return 0, err
	} 
	dxKm := xKm + dem.UnitKm * dem.w_obj.Config.MapSideWidthKm
	dyKm := yKm + dem.UnitKm * dem.w_obj.Config.MapSideWidthKm

	return dem.l_obj.GetElevationByKmPoint(dxKm, dyKm), nil
}

func EnableProcessLog(){
	utility.ProcessLog = true
}