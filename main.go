package main

import (
	"fmt"
	"strconv"
)

func HitungJarakLastFasility(data []map[string]bool, index int) int {
	hasil := 0

	return hasil
}

func CekJarakFromLast(data []map[string]bool, index int, dataCari string) int {

	if index == 0 {
		return 0
	}
	count := 0

	fmt.Println("len data ", len(data))
	for i := len(data) - 1; i >= 0; {
		fmt.Println("len data ", data[i])
		fmt.Println("data cari ", data[i][dataCari])
		if data[i][dataCari] == false {
			count += 1
		}
		i--
	}
	return count
}

func HitungBobot() {

}

func CekIfFasilityIsTrue(req string, fasility map[string]bool) bool {
	r := false
	if fasility[req] {
		r = true
	}
	return r
}

type DataResult struct {
	Warung  int
	Counter int
	laundry int
}

func main() {
	block := []map[string]bool{
		{ //0
			"warung":  true,  // 0
			"counter": false, // 1
			"laundry": false, // 2
		},
		{ //1
			"warung":  false, // 1
			"counter": true,  // 0
			"laundry": false, // 1
		},
		{ //2
			"warung":  false, // 2
			"counter": false, // 1
			"laundry": true,  // 0
		},
	}

	req := []string{"laundry"}

	result := []map[string]DataResult{}
	for _, v := range req {
		t := map[string]DataResult{}
		for j, w := range block {
			// cek bobot berdasarkan params request
			fmt.Println("cek fasilitas => ", v, j, w, CekIfFasilityIsTrue(v, w))
			if CekIfFasilityIsTrue(v, w) {
				k := strconv.Itoa(j)
				t[k] = DataResult{
					// Warung:  CekJarakFromLast(block, j, v),
					// Counter: CekJarakFromLast(block, j, v),
					laundry: CekJarakFromLast(block, j, v),
				}
				result = append(result, t)
			}
		}
	}
	fmt.Println(result)
}
