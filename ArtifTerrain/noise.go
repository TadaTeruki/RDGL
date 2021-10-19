package artif_terrain


func (obj *WorldTerrainObject) MakeWorldTerrain(){
	obj.NoiseSrc.SetSeed(obj.Config.Seed)
	obj.AdjNoiseSrc.SetSeed(obj.Config.Seed)	
}