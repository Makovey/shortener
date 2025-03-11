package mapper

import (
	"github.com/Makovey/shortener/internal/generated/shortener"
	"github.com/Makovey/shortener/internal/service/model"
	transportModel "github.com/Makovey/shortener/internal/transport/model"
)

func FromBatchProtoToBatchRequest(proto *shortener.PostBatchURLRequest) []transportModel.ShortenBatchRequest {
	var batch []transportModel.ShortenBatchRequest

	for _, val := range proto.GetBatch() {
		batch = append(batch, transportModel.ShortenBatchRequest{
			CorrelationID: val.CorrelationID,
			OriginalURL:   val.OriginalURL,
		})
	}

	return batch
}

func FromBatchRequestToPostBatchProto(models *[]transportModel.ShortenBatchResponse) *shortener.PostBatchURLResponse {
	var batch []*shortener.BatchURLResponse

	for _, val := range *models {
		batch = append(batch, &shortener.BatchURLResponse{
			CorrelationID: val.CorrelationID,
			ShortURL:      val.ShortURL,
		})
	}

	return &shortener.PostBatchURLResponse{Batch: batch}
}

func FromBatchToGetUserURLsResponse(models []model.ShortenBatch) *shortener.GetUserURLsResponse {
	var urls []*shortener.UserURL

	for _, val := range models {
		urls = append(urls, &shortener.UserURL{
			CorrelationID: val.CorrelationID,
			OriginalURL:   val.OriginalURL,
			ShortURL:      val.ShortURL,
		})
	}

	return &shortener.GetUserURLsResponse{UserURLs: urls}
}
