package util

import "math/rand"

func GenerateRandomValue() (int, int) {
	water := rand.Intn(101)
	wind := rand.Intn(101)
	return water, wind
}

// func toJson(data Response) ([]byte, error) {
// 	json, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	return json, nil
// }

func GenerateStatus(water int, wind int) (string, string) {
	wtrStatus := ""
	wndStatus := ""

	if water <= 5 {
		wtrStatus = "aman"
	} else if water > 5 && water < 9 {
		wtrStatus = "siaga"
	} else {
		wtrStatus = "bahaya"
	}

	if wind <= 6 {
		wndStatus = "aman"
	} else if wind > 6 && wind < 16 {
		wndStatus = "siaga"
	} else {
		wndStatus = "bahaya"
	}

	return wtrStatus, wndStatus
}
