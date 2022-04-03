package bestroute

import (
	"errors"
	"fmt"
	"math"
)

func findBestRoute(data DeliveryData) (bestPath []string, time float64, err error) {

	locPosMap := make(map[string]generalDetails)     // location/Position map
	restaurantMap := make(map[string]restaurantInfo) // restaurant Map  holds list of all restaurant
	dep := make(map[string]string)                   // dependncy map  holds the relation of the customer and his restaurant
	disMap := make(map[string]map[string]float64)    //disMap holds the distance between any two locations

	dev := data.DeliveryExec.ID //delivery executive id
	locPosMap[dev] = generalDetails{
		GeoPos: GeoPos{
			Lat: data.DeliveryExec.Loc.Lat,
			Lon: data.DeliveryExec.Loc.Lon,
		},
		ID: data.DeliveryExec.ID,
	}
	var dst string
	numVerticies := 1 //total num locations
	graph := make(map[string][]string)
	for _, val := range data.DeliveryList {
		cus := val.Consumer.ID
		if _, ok := locPosMap[cus]; ok {
			err = errors.New("id's must be unique")
			return
		}
		locPosMap[cus] = generalDetails{
			GeoPos: GeoPos{
				Lat: val.Consumer.Loc.Lat,
				Lon: val.Consumer.Loc.Lon,
			},
			ID: cus,
		}

		res := val.Restaurant.ID
		if _, ok := locPosMap[res]; ok {
			err = errors.New("id's must be unique")
			return
		}
		locPosMap[res] = generalDetails{
			GeoPos: GeoPos{
				Lat: val.Restaurant.Loc.Lat,
				Lon: val.Restaurant.Loc.Lon,
			},
			ID: res,
		}
		restaurantMap[res] = val.Restaurant
		if dst == "" {
			dst = cus
		}
		dep[cus] = res
		addEdge(graph, dev, res)
		if _, ok := disMap[dev]; !ok {
			disMap[dev] = make(map[string]float64)
		}
		disMap[dev][res] = haversineDist(locPosMap[dev].Lat, locPosMap[dev].Lon,
			locPosMap[res].Lat, locPosMap[res].Lon)
		numVerticies += 2
	}

	for key1 := range locPosMap {
		for key2 := range locPosMap {
			if key1 != key2 && key1 != dev && key2 != dev {
				addEdge(graph, key1, key2)
				if _, ok := disMap[key1]; !ok {
					disMap[key1] = make(map[string]float64)
				}
				if val, ok := disMap[key2][key1]; ok {
					disMap[key1][key2] = val
				} else {
					disMap[key1][key2] = haversineDist(locPosMap[key1].Lat,
						locPosMap[key1].Lon, locPosMap[key2].Lat, locPosMap[key2].Lon)
				}
			}
		}
	}

	fmt.Printf("Following are all different paths from delivery exec: %s\n", dev)
	finalPaths := findAllValidPaths(graph, dev, dst, numVerticies, dep)
	fmt.Println(finalPaths)
	bestPath, time = findBestPath(locPosMap, restaurantMap, data.DeliveryExec.AvgSpeed, disMap, finalPaths)
	return
}

func findAllValidPaths(graph map[string][]string, start, dst string, n int, dep map[string]string) [][]string {
	visited := make(map[string]bool, n)
	var finalPaths [][]string // finalPaths holds all possible valid paths
	var path []string
	finalPaths = constructAllValidPaths(graph, start, dst, visited, path, dep, finalPaths)
	return finalPaths
}

func constructAllValidPaths(graph map[string][]string, start, dst string,
	visited map[string]bool, path []string, dep map[string]string, finalPaths [][]string) [][]string {
	visited[start] = true
	path = append(path, start)
	totalPath := 2*len(dep) + 1
	if start == dst {
		if len(dep) == 2 && len(path) < totalPath {
			for key, val := range dep {
				if key != dst {
					if len(path) == totalPath-1 {
						temp := append(path, key)
						finalPaths = append(finalPaths, temp)
					}
					if len(path) == totalPath-2 {
						temp := append(path, val, key)
						finalPaths = append(finalPaths, temp)
					}
				}
			}
		} else {
			finalPaths = append(finalPaths, path)
		}
	} else {
		for _, v := range graph[start] {
			if visited[v] == false {
				if val, ok := dep[v]; ok {
					if visited[val] == true {
						finalPaths = constructAllValidPaths(graph, v, dst, visited, path, dep, finalPaths)
					}
				} else {
					finalPaths = constructAllValidPaths(graph, v, dst, visited, path, dep, finalPaths)
				}
			}
		}
	}
	path = path[:len(path)-1]
	visited[start] = false
	return finalPaths
}

func findBestPath(locPosMap map[string]generalDetails,
	restaurantMap map[string]restaurantInfo, speed float64,
	disMap map[string]map[string]float64, finalPaths [][]string) ([]string, float64) {
	var bestPath []string
	var minTime float64 = math.MaxFloat64
	for _, path := range finalPaths {
		var totalTime float64
		for j := 0; j < len(path)-1; j++ {
			posA := path[j]
			posB := path[j+1]
			var dis float64
			if val, ok := disMap[posA][posB]; ok {
				dis = val
			} else {
				dis = haversineDist(locPosMap[posA].Lat, locPosMap[posA].Lon, locPosMap[posB].Lat, locPosMap[posB].Lon)
			}
			time := dis / speed
			if val, ok := restaurantMap[posA]; ok {
				mealPrepTimeHrs := val.MealPrepTime / 60
				if totalTime < mealPrepTimeHrs {
					totalTime += (mealPrepTimeHrs - totalTime)
				}
			}
			totalTime += time
		}
		if totalTime <= minTime {
			minTime = totalTime
			bestPath = path
		}
	}
	return bestPath, minTime
}

func addEdge(graph map[string][]string, u, v string) {
	graph[u] = append(graph[u], v)
}

func haversineDist(lat1, lon1, lat2, lon2 float64) float64 {

	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0
	lat1 = (lat1) * math.Pi / 180.0
	lat2 = (lat2) * math.Pi / 180.0
	a := math.Pow(math.Sin(dLat/2), 2) + math.Pow(math.Sin(dLon/2), 2)*math.Cos(lat1)*math.Cos(lat2)
	var rad float64 = 6371
	c := 2 * math.Asin(math.Sqrt(a))
	return rad * c
}
