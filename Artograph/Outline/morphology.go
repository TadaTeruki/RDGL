package outline

import(
	"github.com/cnkei/gospline"
	terrain "../TerrainGeneration"
	"math"
	"fmt"
)

type MlPoint struct{
	x int
	y int
}

type ClassGroupUnit struct{
	parent int
	is_top bool
	highest_child int
}

func GetInterpolatedClass(obj *terrain.LocalTerrainObject, x, y, xp, yp int) float64{

	data_w := len(obj.ElevationTable[0])
	data_h := len(obj.ElevationTable)

	var bx, by int
	base := 0.0
	for{
		x -= xp
		y -= yp
		if x < 0 || y < 0 || x >= data_w || y >= data_h {
			x = bx
			y = by
			break
		}
		base++
		bx = x
		by = y
	}


	var dist []float64
	var class []float64

	for i := 0; ; i++{
		x += xp
		y += yp
		if x < 0 || y < 0 || x >= data_w || y >= data_h { break }
		if obj.ElevationTable[y][x] != obj.ElevationTable[by][bx]{
			dist = append(dist, float64(i))
			class = append(class, math.Max(obj.ElevationTable[y][x],obj.ElevationTable[by][bx]))
		}
		bx = x
		by = y
	}

	spline := gospline.NewCubicSpline(dist, class)
	return spline.At(base)
	
}


func ClassGroupId(obj *terrain.LocalTerrainObject, x, y int) int{
	data_w := len(obj.ElevationTable[0])
	return y*data_w+x
}

func Morphology(obj *terrain.LocalTerrainObject, first bool, base_class *[][]float64, class_group *[]ClassGroupUnit){

	data_w := len(obj.ElevationTable[0])
	data_h := len(obj.ElevationTable)

	var class_border_list []MlPoint
	var morphology_table []MlPoint

	for y := 0; y < data_h; y++ {
		for x := 0; x < data_w; x++ {
			(*base_class)[y][x] = -1
			if obj.ElevationTable[y][x] == 0 {
				(*class_group)[ClassGroupId(obj,x,y)].parent = -1
			} else {
				(*class_group)[ClassGroupId(obj,x,y)].parent = ClassGroupId(obj,x,y)
			}
			(*class_group)[ClassGroupId(obj,x,y)].is_top = true
			(*class_group)[ClassGroupId(obj,x,y)].highest_child = (*class_group)[ClassGroupId(obj,x,y)].parent

		}
	}

	class_group_root := func(x, y int) int{
		tid := ClassGroupId(obj,x,y)
		if tid == -1 { return -1 }
		for ; tid != (*class_group)[tid].parent; {
			tid = (*class_group)[tid].parent
			if tid == -1 { return -1 }
		}
		return tid
	}

	
	for y := 0; y < data_h; y++ {
		for x := 0; x < data_w; x++ {
			check_class_border := func(sx, sy, ex, ey int){
				if ex < 0 || ey < 0 || ex >= data_w || ey >= data_h { return }

				if obj.ElevationTable[sy][sx] == 0 {
					return
				}

				if obj.ElevationTable[sy][sx] != obj.ElevationTable[ey][ex]{
					var target MlPoint
					target.x = sx
					target.y = sy
					class_border_list = append(class_border_list, target)
					(*base_class)[sy][sx] = math.Max(obj.ElevationTable[sy][sx], obj.ElevationTable[ey][ex])
					
					if obj.ElevationTable[sy][sx] < obj.ElevationTable[ey][ex]{
						sroot := class_group_root(sx,sy)
						(*class_group)[sroot].is_top = false
					}

				} else {
					var id int
					sroot, eroot := class_group_root(sx,sy), class_group_root(ex,ey)
					if ClassGroupId(obj,sx,sy) < ClassGroupId(obj,ex,ey) {
						id = sroot
					} else {
						id = eroot
					}
					(*class_group)[sroot].parent = id
					(*class_group)[sroot].is_top = (*class_group)[sroot].is_top && (*class_group)[eroot].is_top
					(*class_group)[eroot].parent = id
					(*class_group)[eroot].is_top = (*class_group)[sroot].is_top && (*class_group)[eroot].is_top
				}

				
			}

			check_class_border(x, y, x-1, y)
			check_class_border(x, y, x+1, y)
			check_class_border(x, y, x, y-1)
			check_class_border(x, y, x, y+1)

		}
	}

	morphology_table = class_border_list


	for ;len(morphology_table) > 0; {

		var new_morphology_table []MlPoint

		morphology_dilation := func(sx, sy, ex, ey int){
			if ex < 0 || ey < 0 || ex >= data_w || ey >= data_h { return }
		
			var target MlPoint
			target.x = ex
			target.y = ey
		
			if (*base_class)[ey][ex] == -1 && obj.ElevationTable[ey][ex] == obj.ElevationTable[sy][sx]{
		
				new_morphology_table = append(new_morphology_table, target)
				(*base_class)[ey][ex] = (*base_class)[sy][sx]
		
				if first == true {
					root := class_group_root(ex,ey)
					if root != -1 && (*class_group)[root].is_top == true {
						(*class_group)[root].highest_child = ClassGroupId(obj,ex,ey)
					}
				}
			}
		
			if (*base_class)[ey][ex] != -1 && (*base_class)[ey][ex] != (*base_class)[sy][sx]{
				class_border_list = append(class_border_list, target)
			}
		}
		
		
		for i := 0; i < len(morphology_table); i++{
			morphology_dilation(morphology_table[i].x, morphology_table[i].y, morphology_table[i].x-1, morphology_table[i].y)
			morphology_dilation(morphology_table[i].x, morphology_table[i].y, morphology_table[i].x+1, morphology_table[i].y)
			morphology_dilation(morphology_table[i].x, morphology_table[i].y, morphology_table[i].x, morphology_table[i].y-1)
			morphology_dilation(morphology_table[i].x, morphology_table[i].y, morphology_table[i].x, morphology_table[i].y+1)
			
		}
		morphology_table = new_morphology_table
		
	}

	top_list := make(map[MlPoint]struct{})
	
	for y := 0; y < data_h; y++ {
		for x := 0; x < data_w; x++ {
			if obj.ElevationTable[y][x] != 0.0 {
				obj.ElevationTable[y][x] = (obj.ElevationTable[y][x]+(*base_class)[y][x])/2
			}
			root := class_group_root(x, y)
			if first == true && root != -1 && (*class_group)[root].is_top == true {
				child := (*class_group)[root].highest_child
				var target MlPoint
				target.x = child%data_w
				target.y = child/data_w
				top_list[target] = struct{}{}
			}
		}
	}
	
	if first == true {
		for top_point := range top_list {

			a := GetInterpolatedClass(obj,top_point.x, top_point.y, 1, 0)
			b := GetInterpolatedClass(obj,top_point.x, top_point.y, 1, 1)
			c := GetInterpolatedClass(obj,top_point.x, top_point.y, 0, 1)
			d := GetInterpolatedClass(obj,top_point.x, top_point.y, 1, -1)

			obj.ElevationTable[top_point.y][top_point.x] = (a+b+c+d)/4
			
		}
	}

}


func Interpolate(obj *terrain.LocalTerrainObject, config *terrain.GlobalConfig){

	data_w := len(obj.ElevationTable[0])
	data_h := len(obj.ElevationTable)

	base_class := make([][]float64, data_h)
	class_group := make([]ClassGroupUnit, data_w*data_h)
	for y := 0; y < data_h; y++ {
		base_class[y] = make([]float64, data_w)
	}

	for i := 0; i < config.OutlineInterpolationQuality; i++{
		fmt.Println("processing...(",i,"/",config.OutlineInterpolationQuality,")")
		Morphology(obj, i == 0, &base_class, &class_group)
	}

}