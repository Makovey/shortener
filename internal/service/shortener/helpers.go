package shortener

import (
	"fmt"

	repo "github.com/Makovey/shortener/internal/service/model"
	"github.com/Makovey/shortener/internal/transport/model"
)

func fromTransportToRepoShortenBatch(batch []model.ShortenBatchRequest) []repo.ShortenBatch {
	var res []repo.ShortenBatch
	for _, req := range batch {
		tmp := repo.ShortenBatch{
			CorrelationID: req.CorrelationID,
			ShortURL:      generateShortURL(req.OriginalURL),
			OriginalURL:   req.OriginalURL,
		}

		res = append(res, tmp)
	}

	return res
}

func fromRepoToShortenBatchResponse(batch []repo.ShortenBatch, baseReturnedURL string) []model.ShortenBatchResponse {
	var res []model.ShortenBatchResponse
	for _, req := range batch {
		tmp := model.ShortenBatchResponse{
			CorrelationID: req.CorrelationID,
			ShortURL:      fmt.Sprintf("%s/%s", baseReturnedURL, req.ShortURL),
		}

		res = append(res, tmp)
	}

	return res
}
