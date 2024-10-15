package creds

func (s *CredService) CreateExchange(address string, mode string) (string, error) {

	return "ex", nil
}

func (s *CredService) DeleteExchange(address string) error {

	return nil
}

func (s *CredService) GetExchange(address string, mode string) (string, error) {

	return "link/qr", nil
}

func (s *CredService) GetExchangeOwner(id string) (string, error) {

	return "address", nil
}
