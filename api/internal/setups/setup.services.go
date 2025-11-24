package setups

type SetupServices struct {
	SetupRepo ISetupRepository
}

func (service SetupServices) InitSetup() error {
	_, err := service.SetupRepo.CreateEntry()
	if err != nil {
		return err
	}

	return nil
}
