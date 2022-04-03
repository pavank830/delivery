package bestroute

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pavank830/delivery/utils"
)

// GetBestRoute -- endpoint that gives best route to the delivery executive
func GetBestRoute(w http.ResponseWriter, r *http.Request) {

	var req GetBestRouteReq
	resp := GetBestRouteResp{
		APIResp: APIResp{
			ResponseCode: utils.ResponseOK,
		},
	}
	var err error
	if err := utils.ParseRequest(w, r, &req); err != nil {
		return
	}
	defer func() {
		if err != nil {
			resp.ResponseCode = utils.ResponseFailed
			resp.ResponseDescription = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		out, _ := json.Marshal(resp)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Write(out)
	}()
	resp.Path, resp.TotalTime, err = findBestRoute(req.DeliveryDetail)
	if len(resp.Path) == 0 {
		err = fmt.Errorf("couldn't find the best route")
	}
}
