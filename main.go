
package main

import(
	terrain "./ArtifTerrain"
	utility "./Utility"
)

func main(){
	utility.Enable()

	var w_obj terrain.WorldTerrainObject
	
	w_obj.NSKm = 20000.0
	w_obj.WEKm = 40000.0
	w_obj.ElevationBaseM = 8000.0
	w_obj.Config = terrain.GetGlobalConfig()
	w_obj.SetNEFPoint()
	w_obj.MakeWorldTerrain()
	w_obj.WriteWorldToPNG(2000, -1)

}