package router
import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux{
	r:=chi.NewRouter()
	r.get("/health",func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOk)
		w.write([]byte("CodeBridge backend is running"))
	})
	return r
}