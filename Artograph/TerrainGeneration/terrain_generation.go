/*
Artograph/TerrainGeneration/terrain_generation.go
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
	perlin "../PerlinNoise"
)

// WorldTerrain ... Whole area to generate terrains
type WorldTerrainObject struct{
	NoiseSrc, AdjNoiseSrc perlin.PerlinObject
	NSKm float64
	WEKm float64
	Z float64
	ElevationBaseM float64
	NEFPointList []NEFPoint
	Config GlobalConfig
}

// LocalTerrain ... A part of terrains generated in WorldTerrain
// (Effects are only applied for LocalTerrain)
type LocalTerrainObject struct{
	WorldTerrain *WorldTerrainObject
	xKm float64
	yKm float64
	NSKm float64
	WEKm float64
	LevelingLayerObj LevelingLayer
	LevelingCheckIsAvailable bool
	LiverTable [][]LiverPoint
	LiverCheckIsAvailable bool
}

type GlobalConfig struct{ // details -> <globalconfig.go>
	Seed int64
	NoizeScaleKm float64
	NoizeOctave int
	NoizeMinPersistence float64
	NoizeMaxPersistence float64
	MinLand float64
	MaxLand float64
	LocalTerrainSelectionQuality int
	LevelingCheckIntervalKm float64
	LiverCheckIntervalKm float64
	TerrainReverseScale float64
	TerrainLevelingHeightM float64
	TerrainLevelingIntervalKm float64
	PondDepthProportion float64

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

type LiverPoint struct{
	XKm float64
	YKm float64
	Direction int
	Cavity float64
	BaseElevation float64
}

type NEFPoint struct {
	NoiseLevel float64
	Elevation float64
}

type LevelingPoint struct{
	XKm float64
	YKm float64
	IsLeveling bool
	ElevationLevel float64
}

type LevelingLayer struct{
	LevelingTable [][]LevelingPoint
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