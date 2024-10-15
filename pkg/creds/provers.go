package creds

func (s *CredService) GetAvailable() string {
	return "url"
}

func (s *CredService) SetAvailable(url string, val bool) error {
	return nil
}
