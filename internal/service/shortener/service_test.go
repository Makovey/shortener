package shortener

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger/stdout"
	repoMock "github.com/Makovey/shortener/internal/repository/mock"
	repoModel "github.com/Makovey/shortener/internal/repository/model"
	common "github.com/Makovey/shortener/internal/service/model"
	"github.com/Makovey/shortener/internal/transport/model"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateShortURL(t *testing.T) {
	type fields struct {
		repoError   error
		originalURL string
	}
	type want struct {
		err bool
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "successfully saved short url",
			fields: fields{
				originalURL: fmt.Sprintf("https://%s.com", randomString(10)),
			},
			want: want{},
		},
		{
			name: "cannot save short url: repo error",
			fields: fields{
				originalURL: fmt.Sprintf("https://%s.com", randomString(10)),
				repoError:   errors.New("repo error"),
			},
			want: want{
				err: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := repoMock.NewMockRepository(ctrl)
			repo.
				EXPECT().
				SaveUserURL(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.fields.repoError).
				Times(1)

			s := &Service{
				repo:   repo,
				pinger: repoMock.NewPingerMock(nil),
				cfg:    config.NewConfigDummy(),
				log:    stdout.NewLoggerDummy(),
			}

			got, err := s.CreateShortURL(context.Background(), tt.fields.originalURL, uuid.NewString())

			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			}
		})
	}
}

func TestService_GetFullURL(t *testing.T) {
	type fields struct {
		repoError error
		repoRes   *repoModel.UserURL
	}
	type want struct {
		err bool
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "successfully get full url",
			fields: fields{
				repoRes: &repoModel.UserURL{OriginalURL: randomString(10), IsDeleted: false},
			},
			want: want{},
		},
		{
			name: "cannot get full url: short url not found",
			fields: fields{
				repoError: errors.New("repo error"),
			},
			want: want{
				err: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := repoMock.NewMockRepository(ctrl)
			repo.
				EXPECT().
				GetFullURL(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.fields.repoRes, tt.fields.repoError).
				Times(1)

			s := &Service{
				repo:   repo,
				pinger: repoMock.NewPingerMock(nil),
				cfg:    config.NewConfigDummy(),
				log:    stdout.NewLoggerDummy(),
			}

			got, err := s.GetFullURL(context.Background(), randomString(5), uuid.NewString())

			if tt.want.err {
				assert.Error(t, err)
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			}
		})
	}
}

func TestService_ShortBatch(t *testing.T) {
	type fields struct {
		repoError error
		argument  []model.ShortenBatchRequest
	}
	type want struct {
		err bool
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "successfully short all batch",
			fields: fields{
				argument: []model.ShortenBatchRequest{
					{
						CorrelationID: "1",
						OriginalURL:   randomString(10),
					},
				},
			},
			want: want{},
		},
		{
			name: "cannot short batch: not found",
			fields: fields{
				repoError: errors.New("repo error"),
			},
			want: want{
				err: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := repoMock.NewMockRepository(ctrl)
			repo.
				EXPECT().
				SaveUserURLs(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.fields.repoError).
				Times(1)

			s := &Service{
				repo:   repo,
				pinger: repoMock.NewPingerMock(nil),
				cfg:    config.NewConfigDummy(),
				log:    stdout.NewLoggerDummy(),
			}

			got, err := s.ShortBatch(context.Background(), tt.fields.argument, uuid.NewString())

			if tt.want.err {
				assert.Error(t, err)
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, len(tt.fields.argument), len(got))
			}
		})
	}
}

func TestService_GetAllURLs(t *testing.T) {
	type fields struct {
		repoError error
		repoRes   []common.ShortenBatch
	}
	type want struct {
		err bool
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "successfully get all users url",
			fields: fields{
				repoRes: []common.ShortenBatch{
					{
						CorrelationID: "1",
						OriginalURL:   randomString(10),
					},
					{
						CorrelationID: "2",
						OriginalURL:   randomString(10),
					},
				},
			},
			want: want{},
		},
		{
			name: "cannot get all users url: repo error",
			fields: fields{
				repoError: errors.New("repo error"),
			},
			want: want{
				err: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := repoMock.NewMockRepository(ctrl)
			repo.
				EXPECT().
				GetUserURLs(gomock.Any(), gomock.Any()).
				Return(tt.fields.repoRes, tt.fields.repoError).
				Times(1)

			s := &Service{
				repo:   repo,
				pinger: repoMock.NewPingerMock(nil),
				cfg:    config.NewConfigDummy(),
				log:    stdout.NewLoggerDummy(),
			}

			got, err := s.GetAllURLs(context.Background(), uuid.NewString())

			if tt.want.err {
				assert.Error(t, err)
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			}
		})
	}
}

func TestService_DeleteUsersURLs(t *testing.T) {
	type fields struct {
		repoError error
		shortURLS []string
	}
	type want struct {
		err bool
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "successfully mark all urls as deleted",
			fields: fields{
				shortURLS: []string{randomString(5), randomString(5)},
			},
			want: want{},
		},
		{
			name: "cannot mark all urls as deleted: repo error",
			fields: fields{
				shortURLS: []string{randomString(5), randomString(5)},
				repoError: errors.New("repo error"),
			},
			want: want{
				err: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := repoMock.NewMockRepository(ctrl)
			repo.
				EXPECT().
				MarkURLAsDeleted(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.fields.repoError).
				Times(len(tt.fields.shortURLS))

			s := &Service{
				repo:   repo,
				pinger: repoMock.NewPingerMock(nil),
				cfg:    config.NewConfigDummy(),
				log:    stdout.NewLoggerDummy(),
			}

			errs := s.DeleteUsersURLs(context.Background(), uuid.NewString(), tt.fields.shortURLS)

			if tt.want.err {
				assert.NotEmpty(t, errs)
			} else {
				assert.Contains(t, errs, nil)
			}
		})
	}
}

func TestService_CheckPing(t *testing.T) {
	type field struct {
		pingErr error
	}
	type want struct {
		err bool
	}
	tests := []struct {
		name  string
		field field
		want  want
	}{
		{
			name:  "successfully ping",
			field: field{},
			want:  want{},
		},
		{
			name: "ping error",
			field: field{
				pingErr: errors.New("ping error"),
			},
			want: want{
				err: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := repoMock.NewMockRepository(ctrl)

			s := &Service{
				repo:   repo,
				pinger: repoMock.NewPingerMock(tt.field.pingErr),
				cfg:    config.NewConfigDummy(),
				log:    stdout.NewLoggerDummy(),
			}

			err := s.CheckPing(context.Background())

			if tt.want.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func randomString(length int) string {
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
