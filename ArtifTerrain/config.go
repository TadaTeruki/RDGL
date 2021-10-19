package artif_terrain

type GlobalConfig struct{
	Seed int64 // 地形のシード値
	NoizeScaleKm float64 // ノイズの生成単位の大きさ(KM)
	NoizeOctave int // ノイズの細分処理回数
	NoizeMinPersistence float64 // ノイズの細分処理単位内の地形の粗さ(最小) 0.7程度でリアス式海岸、0.4程度でなめらかな浜
	NoizeMaxPersistence float64 // ノイズの細分処理単位内の地形の粗さ(最大)
}

func GetGlobalConfig() GlobalConfig {
	var conf GlobalConfig
	conf.Seed = 20
	conf.NoizeScaleKm = 10000.0
	conf.NoizeOctave = 10
	conf.NoizeMinPersistence = 0.4
	conf.NoizeMaxPersistence = 0.7
	return conf
}