package pack_usecase

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsecase_GetPacksNumber(t *testing.T) {
	var packSizes = []uint64{250, 500, 1000, 2000, 5000}

	type args struct {
		items uint64
	}

	tests := []struct {
		name string
		args args
		want map[uint64]uint64
	}{
		{
			name: "items = 1",
			args: args{1},
			want: map[uint64]uint64{250: 1},
		},
		{
			name: "items = 250",
			args: args{250},
			want: map[uint64]uint64{250: 1},
		},
		{
			name: "items = 251",
			args: args{251},
			want: map[uint64]uint64{500: 1},
		},
		{
			name: "items = 501",
			args: args{501},
			want: map[uint64]uint64{500: 1, 250: 1},
		},
		{
			name: "items = 12001",
			args: args{12001},
			want: map[uint64]uint64{5000: 2, 2000: 1, 250: 1},
		},
		{
			name: "items = 8751",
			args: args{8751},
			want: map[uint64]uint64{5000: 1, 2000: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New(packSizes)
			got := uc.GetPacksNumber(tt.args.items)
			assert.True(t, reflect.DeepEqual(got, tt.want))
		})
	}
}
