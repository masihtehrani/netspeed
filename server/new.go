package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/masihtehrani/netspeed/providers"
)

const PORT = 6010

func New() error {
	r := mux.NewRouter()

	provider, err := providers.New()
	if err != nil {
		return err
	}

	r.HandleFunc("/{provider}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		providerName := vars["provider"]

		p := provider.Builder(providerName)

		download, upload, err := p.GetSpeed()
		if err != nil {
			_ = json.NewEncoder(w).Encode(struct {
				Error string `json:"error"`
			}{
				Error: err.Error(),
			})
		} else {
			_ = json.NewEncoder(w).Encode(struct {
				Download string `json:"download"`
				Upload   string `json:"upload"`
			}{
				Download: fmt.Sprintf("%f Mbps", download),
				Upload:   fmt.Sprintf("%f Mbps", upload),
			})
		}
		
		w.Header().Set("Content-Type", "application/json")
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", PORT), r)
	if err != nil {
		return err
	}

	log.Printf("server run at port: %d\n", PORT)

	return nil
}
