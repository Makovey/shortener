package mapper

import (
	"github.com/Makovey/shortener/internal/generated/service_info"
	"github.com/Makovey/shortener/internal/service/model"
)

// ToProtoFromStats из общей модели в прото формат
func ToProtoFromStats(stats *model.Stats) *service_info.StatsResponse {
	return &service_info.StatsResponse{
		Urls:  int64(stats.URLS),
		Users: int64(stats.Users),
	}
}
