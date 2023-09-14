package pack_handler

import (
	"bytes"
	"errors"
	"fmt"
	pack_usecase "gimshark-test/server/internal/usecase/pack"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestHandlerGetPacksNumber(t *testing.T) {
	var packSizes = []uint64{250, 500, 1000, 2000, 5000}

	tests := []struct {
		name     string
		items    interface{}
		wantCode int
		wantResp string
	}{
		{
			name:     "Valid request",
			items:    10,
			wantCode: http.StatusOK,
			wantResp: "{\"250\":1}\n",
		},
		{
			name:     "Zero items",
			items:    0,
			wantCode: http.StatusBadRequest,
			wantResp: "Bad request: items should be greater than 0\n",
		},
		{
			name:     "Invalid JSON. Items -1",
			items:    -1,
			wantCode: http.StatusBadRequest,
			wantResp: "Bad request: invalid data\n",
		},
		{
			name:     "Invalid JSON. Items is not a number",
			items:    "s",
			wantCode: http.StatusBadRequest,
			wantResp: "Bad request: invalid data\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPost,
				"/packs",
				bytes.NewBuffer([]byte(
					[]byte(fmt.Sprintf(`{"items": %v}`, tt.items)),
				)),
			)
			assert.NoError(t, err)

			req.Header.Set("Content-Type", "Application/Json")

			log := zap.NewExample()
			useCase := pack_usecase.New(packSizes)

			handler := New(log, useCase)
			rr := httptest.NewRecorder()
			handler.GetPacksNumber().ServeHTTP(rr, req)

			assert.Equal(t, tt.wantCode, rr.Code)
			assert.Equal(t, tt.wantResp, rr.Body.String())
		})
	}
}

func TestHandlerValidate(t *testing.T) {
	var packSizes = []uint64{}

	tests := []struct {
		name    string
		in      *GetPacksNumberIn
		wantErr error
	}{
		{
			name: "Valid request",
			in:   &GetPacksNumberIn{Items: 10},
		},
		{
			name:    "Invalidrequest. Items is 0",
			in:      &GetPacksNumberIn{Items: 0},
			wantErr: errors.New("items should be greater than 0"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := zap.NewExample()
			useCase := pack_usecase.New(packSizes)

			handler := New(log, useCase)
			if err := handler.validateReq(tt.in); err != nil {
				assert.Equal(t, tt.wantErr, err)
			}

		})
	}
}
