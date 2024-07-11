package handlers

import (
	"net/http"
)

type TobbacoHandler interface {
	GetTobbacos(w http.ResponseWriter, r *http.Request)

	//PostTobaccos(w http.ResponseWriter, r *http.Request)
}
