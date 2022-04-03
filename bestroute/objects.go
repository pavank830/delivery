package bestroute

// api paths
const (
	BestRoutePath = "/api/bestroute"
)

//DeliveryData -- struct contains delivery executive details, list of deliveries
type DeliveryData struct {
	DeliveryExec DeliveryExecDetails `json:"delivery_exec_info"`
	DeliveryList []DeliveryInfo      `json:"delivery_info"`
}

// DeliveryExecDetails -- struct contains delivery executive id, location,avg speed
type DeliveryExecDetails struct {
	ID       string  `json:"id"`
	Loc      GeoPos  `json:"location"`
	AvgSpeed float64 `json:"speed"`
}

//GeoPos -- struct contains latitude and longitude
type GeoPos struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

//generalDetails --
type generalDetails struct {
	ID string `json:"id"`
	GeoPos
}

//consumerInfo -- struct contains the consumer id,location
type consumerInfo struct {
	ID  string `json:"id"`
	Loc GeoPos `json:"location"`
}

//restaurantInfo -- struct contains the restaurant id,location and meal prep time
type restaurantInfo struct {
	ID           string  `json:"id"`
	Loc          GeoPos  `json:"location"`
	MealPrepTime float64 `json:"prep_time"` //prep time in mins
}

//DeliveryInfo -- struct contains the consumer info and his ordered restaurant info
type DeliveryInfo struct {
	Consumer   consumerInfo   `json:"consumer"`
	Restaurant restaurantInfo `json:"restaurant"`
}

// GetBestRouteReq -- request for /api/bestroute endpoint
type GetBestRouteReq struct {
	DeliveryDetail DeliveryData `json:"deliver_data"`
}

// APIResp - basic any api response struct
type APIResp struct {
	ResponseCode        int    `json:"response_code"`
	ResponseDescription string `json:"response_description"`
}

// GetBestRouteResp -- response struct for /api/bestroute endpoint
type GetBestRouteResp struct {
	Path      []string `json:"path"`
	TotalTime float64  `json:"total_time"`
	APIResp
}

// ----------------------------------
