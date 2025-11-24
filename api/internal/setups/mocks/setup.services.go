package mocks

import "github.com/stretchr/testify/mock"

type ISetupServices struct {
	mock.Mock
}

func (m *ISetupServices) InitSetup() error {
	args := m.Called()
	return args.Error(0)
}
