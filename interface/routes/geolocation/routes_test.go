package geolocation

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"go-ddd/interface/controllers/geolocation/mock"
)

func TestNewGeolocationRouter(t *testing.T) {
	t.Run("given a controller, create a new router", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		controllerMock := mock.NewMockController(ctrl)

		gr := NewGeolocationRouter(controllerMock)

		require.NotNil(t, gr)
	})
}

func TestRouter_Routes(t *testing.T) {
	t.Run("given a router, register routes successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		controllerMock := mock.NewMockController(ctrl)

		gr := NewGeolocationRouter(controllerMock)

		r := gr.Routes()
		require.NotNil(t, r)
	})
}
