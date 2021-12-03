/*
Artograph/TerrainGeneration/internalconfig.go
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

// sets parameters for configuration
func GetInternalConfig() InternalConfig {

	var conf InternalConfig

	// The seed value for terrain-generation
	conf.Seed = 0
	
	// Scale of a plate of a continent (Square)
	conf.NoizeScaleKm = 1000.0

	// Fineness of terrain
	conf.NoizeOctave = 30

	// Minimum-complicatedness of terrain (Example : 0.4 -> genelates plains or beaches)
	conf.NoizeMinPersistence = 0.4

	// Maximum-complicatedness of terrain (Example : 0.7 -> genelates steep valleys or sawtooth shaped coasts)
	conf.NoizeMaxPersistence = 0.7

	// Minimum-proportion of land (Example : 0.7 -> 70% (of generated terrain) will covered with land)
	conf.MinLandProportion = 0.5				
	
	// Maximum-proportion of land 
	conf.MaxLandProportion = 0.65
	
	// Quality of terrain-generation
	conf.LocalTerrainSelectionQuality = 100

	// Effects configuration
	conf.LevelingIntervalKm = 1.0 
	conf.LiverIntervalKm = 1.0	
	conf.LiverEndPointElevationM = -10.0

	// Elevation adjustment of cavity areas (Example : 0.01 -> previous_elevation*(-0.01) (m) )
	conf.TerrainReverseScale = 0.01

	conf.LevelingHeightM = 100.0
	conf.LevelingStartPointIntervalKm = 100.0
	conf.LevelingMinimumElevationM = -1500.0

	conf.PlainDepth = 0.3

	conf.OutlineInterpolationQuality = 10
	conf.OutlineNoiseStrength = 0.5

	conf.OutlineNoizeMinPersistence = 0.6
	conf.OutlineNoizeMaxPersistence = 0.8

	conf.StandardLandProportion = 0.5
	
	return conf
}