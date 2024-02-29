package healthcheck

import (
	"net/http"
)

func HealthCheckController(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("health OK!"))
	
}
