package artif_terrain

import(
	"math"
	"math/rand"
)

func (obj *WorldTerrainObject) MakeWorldTerrain(){
	obj.NoiseSrc.SetSeed(obj.Config.Seed)
	obj.AdjNoiseSrc.SetSeed(obj.Config.Seed)	
}

// LocalTerrainを評価。count:地形評価の精度 N(count^2)
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
			score += elevation*elevation
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
	obj.OceanCheckIsAvailable = false

	for i:=0; i<submit_model_num; i++{
		cobj[i] = *obj
		cobj[i].xKm = rand.Float64()*(obj.WorldTerrain.WEKm-obj.WEKm)
		cobj[i].yKm = rand.Float64()*(obj.WorldTerrain.NSKm-obj.NSKm)
		
		score, available := cobj[i].SubmitLocalTerrain(10)
		if available == false{
			i--
			continue
		}

		if score > max_score {
			max_score = score
			select_ad = i
		}
	}

	obj.xKm = cobj[select_ad].xKm
	obj.yKm = cobj[select_ad].yKm
	obj.MakeOceanTable()


}
