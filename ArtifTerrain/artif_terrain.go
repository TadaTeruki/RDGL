package artif_terrain

import(
	perlin "../PerlinNoise"
)

type WorldTerrainObject struct{
	NoiseSrc, AdjNoiseSrc perlin.PerlinObject
	NSKm float64
	WEKm float64
	ElevationBaseM float64
	NEFPointList []NEFPoint
	Config GlobalConfig
}

type LocalTerrainObject struct{
	WorldTerrain *WorldTerrainObject
	xKm float64
	yKm float64
	NSKm float64
	WEKm float64
	OceanTable [][]OceanPoint
	OceanCheckIsAvailable bool
}

