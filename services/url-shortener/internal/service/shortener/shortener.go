package shortener

type Service struct{}

func (s *Service) Get(shortenURL string) (string, error) {
	return "", nil
}

func (s *Service) Set(totalURL string) (string, error) {
	return "", nil
}

func (s *Service) Delete(shortenURL string) error {
	return nil
}
