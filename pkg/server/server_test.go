package server

import (
	"net/http"
	"testing"

	"bitbucket.org/btcid/startrack/internal/config"
)

func SampleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("test"))
	}
}

func TestServer_AddRoutes(t *testing.T) {
	type fields struct {
		config *config.Config
	}
	type args struct {
		routes Routes
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Register Routes",
			fields: fields{},
			args: args{
				routes: Routes{
					{
						Name:    "test",
						Path:    "/test",
						Handler: SampleHandler(),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				config: tt.fields.config,
			}
			s.AddRoutes(tt.args.routes)
		})
	}
}
