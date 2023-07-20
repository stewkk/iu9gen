package app

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AppSuite struct {
	suite.Suite
	app App
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}

func (s *AppSuite) SetupTest() {
	s.app = NewApp()
}
