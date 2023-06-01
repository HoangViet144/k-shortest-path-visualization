package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"

	"googlemaps.github.io/maps"
)

const googleMapApiKey = "API_KEY"

//var location = []string{
//	"268 Lý Thường Kiệt, Phường 14, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
//	"45 Hàn Thuyên, Bến Nghé, Quận 11, Thành phố Hồ Chí Minh, Vietnam",
//	"144 Lý Thường Kiệt, Phường 14, Quận 10, Thành phố Hồ Chí Minh 700000, Vietnam",
//	"495 Đ. Lý Thái Tổ, Phường 9, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
//	"10.767660,106.674153",
//	"10.767894,106.674544",
//	"241Bis Đ. Cách Mạng Tháng 8, Phường 4, Quận 3, Thành phố Hồ Chí Minh 700000, Vietnam",
//	"264A-264B-264C, 264B - 264C Đ. Nguyễn Thị Minh Khai, Phường 6, Quận 3, Thành phố Hồ Chí Minh 700000, Vietnam",
//	"126 Đ. Nguyễn Thị Minh Khai, Phường 6, Quận 3, Thành phố Hồ Chí Minh 70000, Vietnam",
//	"1 Tô Hiến Thành, Phường 14, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
//	"10.776405,106.663603",
//	"10.782877,106.672105",
//	"Binh, Đ. Võ Thị Sáu Vòng Xoay Công Trường Dân Chủ, & Bùng, Quận 3, Thành phố Hồ Chí Minh, Vietnam",
//	"173/14 Đ. Nguyễn Thị Minh Khai, Phường 2, Quận 3, Thành phố Hồ Chí Minh, Vietnam",
//	"QM98+2R7, Đ. Thành Thái, Phường 8, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
//	"231 Đ. 3 Tháng 2, Phường 10, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
//	"10.766338,106.678627",
//	"10.775019,106.686788",
//	"QMCV+G7Q, Ngã sáu Phù Đổng, Phường Phạm Ngũ Lão, Quận 1, Thành phố Hồ Chí Minh, Vietnam",
//	"200A Lý Tự Trọng, Phường Bến Thành, Quận 1, Thành phố Hồ Chí Minh 700000, Vietnam",
//	"10.776016,106.691627",
//	"160 Đ. Nam Kỳ Khởi Nghĩa, Phường 6, Quận 1, Thành phố Hồ Chí Minh, Vietnam",
//	"10.776785,106.699882",
//	"103 Pasteur, Bến Nghé, Quận 1, Thành phố Hồ Chí Minh, Vietnam",
//}

var location = []string{
	"268 Lý Thường Kiệt, Phường 14, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
	"45 Hàn Thuyên, Bến Nghé, Quận 11, Thành phố Hồ Chí Minh, Vietnam",
	"270 B, Lý Thường Kiệt, Phường 14, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
	"10.780678,106.658833",
	"175 Đ. Thành Thái, Phường 14, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
	"10.770488,106.658230",
	"10.782881,106.672134",
	"144 Lý Thường Kiệt, Phường 14, Quận 10, Thành phố Hồ Chí Minh 700000, Vietnam",
	"10.767604,106.667086",
	"Binh, Đ. Võ Thị Sáu Vòng Xoay Công Trường Dân Chủ, & Bùng, Quận 3, Thành phố Hồ Chí Minh, Vietnam",
	"231 Đ. 3 Tháng 2, Phường 10, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
	"QM9F+3QC, Phường 2, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
	"241Bis Đ. Cách Mạng Tháng 8, Phường 4, Quận 3, Thành phố Hồ Chí Minh 700000, Vietnam",
	"173/14 Đ. Nguyễn Thị Minh Khai, Phường 2, Quận 3, Thành phố Hồ Chí Minh, Vietnam",
	"10.773615,106.689442",
	"10.779552,106.694980",
	"QMCV+G7Q, Ngã sáu Phù Đổng, Phường Phạm Ngũ Lão, Quận 1, Thành phố Hồ Chí Minh, Vietnam",
	"664 Đ. Cách Mạng Tháng 8, Phường 15, Quận 3, Thành phố Hồ Chí Minh, Vietnam",
	"10.783382,106.690759",
	"217 Đ. Nam Kỳ Khởi Nghĩa, Phường 7, Quận 3, Thành phố Hồ Chí Minh, Vietnam",
	"88 Đ. Thành Thái, Phường 12, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
	"Chago, 816 Đ. Sư Vạn Hạnh, Phường 12, Quận 10, Thành phố Hồ Chí Minh, Vietnam",
	"10.769610,106.670836",
	"10.773795,106.677803",
	"307A, Số Cũ/155A Tô Hiến Thành, Quận 10, Thành phố Hồ Chí Minh 700000, Vietnam",
	"90 Đ. Cách Mạng Tháng 8, Phường 6, Quận 3, Thành phố Hồ Chí Minh, Vietnam",
	"10.782457,106.691810",
	"10.776786,106.699887",
	"Đ. Nguyễn Du, Bến Nghé, Quận 1, Thành phố Hồ Chí Minh, Vietnam",
	"10.778663,106.697897",
	"01 Đ. Cao Thắng, Phường 2, Quận 3, Thành phố Hồ Chí Minh 70000, Vietnam",
}

var edgeList = [][]int{
	{0, 2}, {2, 0},
	{2, 3}, {3, 2},
	{3, 4}, {4, 3},
	{0, 5}, {5, 0},
	{5, 4}, {4, 5},
	{5, 7}, {7, 5},
	{4, 20}, {20, 4},
	{20, 8}, {8, 20},
	{7, 8}, {8, 7},
	{4, 24}, {24, 4},
	{20, 21}, {21, 20},
	{24, 21}, {21, 24},
	{8, 22}, {22, 8},
	{21, 22}, {22, 21},
	{22, 10}, {10, 22},
	{10, 23}, {23, 10},
	{23, 9}, {9, 23},
	{24, 6}, {6, 24},
	{3, 17}, {17, 3},
	{6, 17}, {17, 6},
	{6, 9}, {9, 6},
	{8, 11}, {11, 8},
	{10, 11}, {11, 10},
	{11, 13}, {13, 11},
	{11, 12},
	{9, 12}, {12, 9},
	{12, 25}, {25, 12},
	{14, 25}, {25, 14},
	{23, 30}, {30, 23},
	{13, 30}, {30, 13},
	{30, 14}, {14, 30},
	{9, 19},
	{12, 18},
	{19, 18}, {18, 19},
	{25, 26}, {26, 25},
	{18, 26},
	{14, 15},
	{15, 1},
	{14, 16}, {16, 14},
	{16, 27},
	{27, 28},
	{28, 29},
	{29, 1},
}

func getMatrixDistanceFromGoogleMap(location []string) *GraphData {
	client, err := maps.NewClient(maps.WithAPIKey(googleMapApiKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	matrixSize := len(location)
	fmt.Printf("Total location: %d\n", matrixSize)

	distanceMatrix := make([][]int, matrixSize)
	for i := 0; i < matrixSize; i++ {
		distanceMatrix[i] = make([]int, matrixSize)
	}

	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			if i == j {
				distanceMatrix[i][j] = 0
			} else {
				distanceMatrix[i][j] = math.MaxInt
			}
		}
	}

	for _, edge := range edgeList {
		u := edge[0]
		v := edge[1]
		rsp, err := client.DistanceMatrix(
			context.Background(), &maps.DistanceMatrixRequest{
				Origins:      []string{location[u]},
				Destinations: []string{location[v]},
			},
		)
		if err != nil {
			log.Fatalf("Directions request failed: %v", err)
		}

		distanceMatrix[u][v] = rsp.Rows[0].Elements[0].Distance.Meters
	}

	return &GraphData{
		Location:       location,
		DistanceMatrix: distanceMatrix,
	}
}

func writeGraphDataToFile(data *GraphData) {
	jsonData, _ := json.Marshal(data)
	os.WriteFile("data.json", jsonData, 0644)
}

func loadGraphDataFromFile() *GraphData {
	jsonData, _ := os.ReadFile("data.json")
	data := &GraphData{}
	json.Unmarshal(jsonData, data)
	return data
}

func generateSampleData() *GraphData {
	return &GraphData{
		Location: []string{"C", "D", "E", "F", "G", "H"},
		DistanceMatrix: [][]int{
			{
				0, 3, 2, math.MaxInt, math.MaxInt, math.MaxInt,
			},
			{
				math.MaxInt, 0, math.MaxInt, 4, math.MaxInt, math.MaxInt,
			},
			{
				math.MaxInt, 1, 0, 2, 3, math.MaxInt,
			},
			{
				math.MaxInt, math.MaxInt, math.MaxInt, 0, 2, 1,
			},
			{
				math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 0, 2,
			},
			{
				math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, 0,
			},
		},
	}
}
