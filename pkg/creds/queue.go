package creds

func (s *CredService) Enqueue(vp string, address string) error {
	return nil
}

func (s *CredService) GetNext() (string, string, error) {
	return "vp", "address", nil
}
