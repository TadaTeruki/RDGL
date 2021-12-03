/*
Artograph/Utility/utility.go
Copyright (C) 2021 Tada Teruki

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation; either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
package utility

import(
	"fmt"
	"math"
)

var Debug = fmt.Println

var ProcessLog bool = false

func EchoProcessPercentage(context string, process_proportion float64){
	if ProcessLog == false { return }
	fmt.Println("Process log : " + context + " (",math.Floor(process_proportion*100),"% )")
}

func EchoProcessEnd(context string){
	if ProcessLog == false { return }
	fmt.Println("Process log : " + context + " was successfully finished")
}