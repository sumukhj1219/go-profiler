package stats

import (
	"log"
	"net/http"

	"github.com/arl/statsviz"
	"github.com/sumukhj1219/go-profiler/internals/config"
)

func NewStats(cfg config.Config) {
	mux := http.NewServeMux()
	statsviz.Register(mux)

	go func() {
		addr := cfg.Host + ":" + cfg.StatsVizPort
		log.Printf("Statsviz running at http://%s/debug/statsviz", addr)
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.Fatal("Statsviz server failed:", err)
		}
	}()
}
