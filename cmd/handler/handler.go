package handler

import (
	timezoneProto "github.com/justIGreK/Reminders-Timezone/pkg/go/timezone"
	"google.golang.org/grpc"
)

type Handler struct {
	server      grpc.ServiceRegistrar
	timezone TimezoneService
}

func NewHandler(grpcServer grpc.ServiceRegistrar, tzSRV TimezoneService) *Handler {
	return &Handler{server: grpcServer, timezone: tzSRV}
}
func (h *Handler) RegisterServices() {
	h.registerTxService(h.server, h.timezone)
}

func (h *Handler) registerTxService(server grpc.ServiceRegistrar, tz TimezoneService) {
	timezoneProto.RegisterTimezoneServiceServer(server, &TimezoneServiceServer{TzSRV: tz})
}
