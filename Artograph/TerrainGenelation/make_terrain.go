package terrain_generation

import(
	"math"
	"math/rand"
)

func (obj *WorldTerrainObject) MakeWorldTerrain(){
	obj.NoiseSrc.SetSeed(obj.Config.Seed)
	obj.AdjNoiseSrc.SetSeed(obj.Config.Seed)
}

// N(count^2)
func (obj *LocalTerrainObject) SubmitLocalTerrain(count int) (float64, bool){
	
	var score float64
	fcount := float64(count)

	var maxWHkm float64

	if obj.NSKm > obj.WEKm { 
		maxWHkm = obj.NSKm
	} else {
		maxWHkm = obj.WEKm
	}

	cxKm := obj.xKm + obj.WEKm/2
	cyKm := obj.yKm + obj.NSKm/2
	count_land := 0
	count_all := 0

	for yKm := obj.yKm; yKm < obj.yKm + obj.NSKm; yKm += maxWHkm/fcount{
		for xKm := obj.xKm; xKm < obj.xKm + obj.WEKm; xKm += maxWHkm/fcount{
			distKm := math.Sqrt((cxKm-xKm)*(cxKm-xKm)+(cyKm-yKm)*(cyKm-yKm))
			rDistKm := 0.5-distKm/(maxWHkm/2)
			_,_ = distKm, rDistKm
			elevation := obj.WorldTerrain.GetElevationByKmPoint(xKm, yKm)
			if elevation >= 0 {
				count_land++
			}
			count_all++
			score += elevation//*elevation
		}
	}
	
	land := float64(count_land)/float64(count_all)

	
	
	if land < obj.WorldTerrain.Config.MinLand || land > obj.WorldTerrain.Config.MaxLand {
		return 0, false
	}
	

	return score, true
}


func (obj *LocalTerrainObject) MakeLocalTerrain(){
	
	rand.Seed(0)

	submit_model_num := obj.WorldTerrain.Config.LocalTerrainSelectionQuality
	cobj := make([]LocalTerrainObject, submit_model_num)
	var max_score float64
	var select_ad int
	var select_z float64
	obj.OceanCheckIsAvailable = false
	obj.LiverCheckIsAvailable = false
	

	for i:=0; i<submit_model_num; i++{
		cobj[i] = *obj
		cobj[i].xKm = (obj.WorldTerrain.WEKm-obj.WEKm)*rand.Float64()
		cobj[i].yKm = (obj.WorldTerrain.NSKm-obj.NSKm)*rand.Float64()
		score, available := cobj[i].SubmitLocalTerrain(5)
		if available == false{
			i--
			obj.WorldTerrain.Z += 1
			continue
		}

		if score > max_score {
			max_score = score
			select_ad = i
			select_z = obj.WorldTerrain.Z
		}
	}
	
	obj.WorldTerrain.Z = select_z



	obj.xKm = cobj[select_ad].xKm
	obj.yKm = cobj[select_ad].yKm

	obj.MakeOceanTable()
	obj.MakeLiverTable()

}
