
package main

import(
	terrain "./ArtifTerrain"
	utility "./Utility"
)

func main(){
	utility.Enable()

	var w_obj terrain.WorldTerrainObject
	var l_obj terrain.LocalTerrainObject
	
	w_obj.NSKm = 20000.0
	w_obj.WEKm = 40000.0
	w_obj.ElevationBaseM = 8000.0
	w_obj.Config = terrain.GetGlobalConfig()
	w_obj.SetNEFPoint()
	w_obj.MakeWorldTerrain()

	l_obj.NSKm = 500.0
	l_obj.WEKm = 1000.0
	l_obj.WorldTerrain = &w_obj
	l_obj.MakeLocalTerrain()
	l_obj.WriteLocalToPNG(1000, -1)
	//w_obj.WriteWorldToPNG(2000, -1)

}