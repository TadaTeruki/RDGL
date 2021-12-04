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
	
	// Size of a continental plate of a continent (Square)
	conf.PlateSizeKm = 1000.0

	// Fineness of terrain
	conf.NoizeOctave = 30

	// Minimum-complicatedness of terrain (Example : 0.4 -> genelates plains or beaches)
	conf.NoizeMinPersistence = 0.4

	// Maximum-complicatedness of terrain (Example : 0.7 -> genelates steep valleys or sawtooth shaped coasts)
	conf.NoizeMaxPersistence = 0.7

	conf.MapSideWidthKm = 0.0

	// Proportion of land of WorldTerrain
	conf.StandardLandProportion = 0.5
	
	// Quality of terrain-generation
	conf.LocalTerrainSelectionQuality = 100

	// The maximum depth of lake (LakeDepthProportion*ElevationAbsM)
	conf.LakeDepthProportion = 0.3

	// Liver effects configuration
	conf.LiverIntervalKm = 1.0	
	conf.LiverEndPointElevationProportion = -0.002

	// Leveling effects configuration
	conf.LevelingIntervalKm = 1.0 
	conf.LevelingHeightM = 100.0
	conf.LevelingStartPointIntervalKm = 10.0
	conf.LevelingMinimumElevationProportion = -0.5

	// Interpolation quality of outline data
	conf.OutlineInterpolationQuality = 10
	// Noise strength of outline data
	conf.OutlineNoiseMinStrength = 0.01
	conf.OutlineNoiseMaxStrength = 1.0

	// Minimum/Maximum-complicatedness of terrain
	conf.OutlineNoizeMinPersistence = 0.5
	conf.OutlineNoizeMaxPersistence = 0.8


	
	return conf
}