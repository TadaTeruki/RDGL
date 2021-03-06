/*
github.com/TadaTeruki/RDGL/TerrainGeneration/terrain_generation.go
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

import(
	perlin "github.com/TadaTeruki/RDGL/PerlinNoise"
)


type InternalConfig struct{ // details -> <globalconfig.go>
	Seed int64
	PlateSizeKm float64
	NoizeOctave int
	NoizeMinPersistence float64
	NoizeMaxPersistence float64
	LocalTerrainSelectionQuality int
	LevelingHeightM float64
	PlainElevationProportion float64
	ContShelfElevationProportion float64
	LakeDepthProportion float64
	LiverEndPointElevationProportion float64
	OutlineInterpolationQuality int
	OutlineNoiseMinStrength float64
	OutlineNoiseMaxStrength float64
	StandardLandProportion float64
	OutlineNoizeMinPersistence float64
	OutlineNoizeMaxPersistence float64
	MapSideWidthKm float64
	RootIntervalKm float64
	UnitKm float64
	

}


const DIRECTION_NONE  = 0
const DIRECTION_NORTH = 1
const DIRECTION_WEST  = 2
const DIRECTION_SOUTH = 3
const DIRECTION_EAST  = 4

type KmPoint struct{
	XKm float64
	YKm float64
}

type HCFPoint struct {
	NoiseLevel float64
	Elevation float64
}


type UnitPoint struct{
	XKm float64
	YKm float64
	
	// leveling
	IsLeveling bool
	ElevationLevel float64

	// liver
	Direction int
	Cavity float64
	BaseElevation float64
	Root *KmPoint
	RootDistKm float64
}

/*
type UnitPoint struct{
	XKm float64
	YKm float64
	IsLeveling bool
	ElevationLevel float64
}
*/

type UnitLayer struct{
	Table [][]UnitPoint
}

type PathPoint struct{
	point KmPoint
	parent KmPoint
	score float64
	flag int // 0:untreated 1:opened 2:closed
}

func MakeKmPoint(xKm, yKm float64) KmPoint{
	var r KmPoint
	r.XKm = xKm
	r.YKm = yKm
	return r
} 

// WorldTerrain ... Whole area to generate terrains
type WorldTerrainObject struct{
	NoiseSrc, AdjNoiseSrc perlin.PerlinObject
	NSKm float64
	WEKm float64
	Z float64
	ElevationAbsM float64
	HCFPointList []HCFPoint
	Config InternalConfig
}

// LocalTerrain ... A part of terrains generated in WorldTerrain
// (Effects are only applied for LocalTerrain)
type LocalTerrainObject struct{
	WorldTerrain *WorldTerrainObject
	xKm float64
	yKm float64
	NSKm float64
	WEKm float64

	RootList []KmPoint

	UnitLayerObj UnitLayer
	LevelingCheckIsAvailable bool
	LiverCheckIsAvailable bool

	ElevationTableIsAvailable bool
	ElevationTable [][]float64
}
