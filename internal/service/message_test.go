package service

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rogalni/cng-hello-backend/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockId uuid.UUID = uuid.New()
)

type MRMock struct {
	mock.Mock
}

func (mr *MRMock) FindAll(ctx context.Context) ([]*model.Message, error) {
	return []*model.Message{}, nil

}

func (mr *MRMock) FindById(ctx context.Context, id uuid.UUID) (*model.Message, error) {
	if id == mockId {
		return &model.Message{Id: id, Code: "code", Text: "text"}, nil
	} else {
		return nil, errors.NewErrorNotFound("Message not found")
	}

}

func (mr *MRMock) Create(ctx context.Context, m *model.Message) error {
	return nil
}
func (mr *MRMock) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func TestMessageService_GetMessage(t *testing.T) {
	mms := MessageService{
		messageRepository: &MRMock{},
	}

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		ms      MessageService
		args    args
		want    *model.Message
		wantErr bool
		err     error
	}{
		{
			name: "happy case",
			ms:   mms,
			args: args{
				ctx: context.Background(),
				id:  mockId,
			},
			want: &model.Message{
				Id:   mockId,
				Code: "code",
				Text: "text",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "not existing",
			ms:   mms,
			args: args{
				ctx: context.Background(),
				id:  uuid.New(),
			},
			want:    nil,
			wantErr: true,
			err:     errors.NewErrorNotFound("Message not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ms.GetMessage(tt.args.ctx, tt.args.id)
			assert.ErrorIs(t, err, tt.err, "Not found error should be thrown")
			assert.Equal(t, tt.want, got, "Should be equal")
		})
	}
}
