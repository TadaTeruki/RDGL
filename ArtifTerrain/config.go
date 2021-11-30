package artif_terrain

type GlobalConfig struct{
	Seed int64 // 地形のシード値
	NoizeScaleKm float64 // ノイズの生成単位の大きさ(KM)
	NoizeOctave int // ノイズの細分処理回数
	// ノイズの細分処理単位内の地形の粗さ : 0.7程度でリアス式海岸、0.4程度でなめらかな浜
	NoizeMinPersistence float64 // ノイズの細分処理単位内の地形の粗さ(最小) 
	NoizeMaxPersistence float64 // ノイズの細分処理単位内の地形の粗さ(最大)
	// 陸地率の範囲は計算量に反比例(広いほど小さい)
	MinLand float64 // 最小陸地率
	MaxLand float64 // 最大陸地率
	LocalTerrainSelectionQuality int // LocalTerrainのクオリティ O(N) 最低1
	OceanCheckIntervalKm float64 // 水たまり認識の基準点の間隔
	// O(LocalTerrain::NSKm*LocalTerrain::WEKm/(OceanCheckIntervalKm^2) )
	LiverCheckIntervalKm float64 // 川の間隔
	TerrainReverseScale float64 // 地形うらがえし処理を適用している地形の高さ調整(~倍)
	VirtualOceanElevation float64 // 水たまり認識の基準高度
}

func GetGlobalConfig() GlobalConfig {
	var conf GlobalConfig
	conf.Seed = 15
	conf.NoizeScaleKm = 1000.0
	conf.NoizeOctave = 20
	conf.NoizeMinPersistence = 0.4
	conf.NoizeMaxPersistence = 0.7
	conf.MinLand = 0.6
	conf.MaxLand = 0.95
	conf.LocalTerrainSelectionQuality = 100
	conf.OceanCheckIntervalKm = 1.0
	conf.LiverCheckIntervalKm = 1.0	
	conf.TerrainReverseScale = 0.01
	conf.VirtualOceanElevation = 10.0
	return conf
}