package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	Server string
	Port   string
}

func New(server, port string) *Config {
	return &Config{Server: server, Port: port}
}

func (c Config) delegate(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Server string `json:"m.server"`
	}{
		Server: fmt.Sprintf("%s:%s", c.Server, c.Port),
	}
	json.NewEncoder(w).Encode(status)
}
