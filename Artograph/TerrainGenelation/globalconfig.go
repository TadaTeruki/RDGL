package terrain_generation

// sets parameters for configuration
func GetGlobalConfig() GlobalConfig {

	var conf GlobalConfig

	// The seed value for terrain-generation
	conf.Seed = 12
	
	// Scale of a plate of a continent (Square)
	conf.NoizeScaleKm = 1000.0

	// Fineness of terrain
	conf.NoizeOctave = 20

	// Minimum-complicatedness of terrain (Example : 0.4 -> genelates plains or beaches)
	conf.NoizeMinPersistence = 0.4

	// Maximum-complicatedness of terrain (Example : 0.7 -> genelates steep valleys or sawtooth shaped coasts)
	conf.NoizeMaxPersistence = 0.7

	// Minimum-proportion of land (Example : 0.7 -> 70% (of generated terrain) will covered with land)
	conf.MinLand = 0.6					
	
	// Maximum-proportion of land 
	conf.MaxLand = 0.95
	
	// Quality of terrain-generation
	conf.LocalTerrainSelectionQuality = 100

	// Effects configuration
	conf.OceanCheckIntervalKm = 1.0 
	conf.LiverCheckIntervalKm = 1.0	

	// Elevation adjustment of cavity areas (Example : 0.01 -> previous_elevation*(-0.01) (m) )
	conf.TerrainReverseScale = 0.01

	// Elevation of cavity areas to adjust
	conf.VirtualOceanElevation = 10.0
	return conf
}