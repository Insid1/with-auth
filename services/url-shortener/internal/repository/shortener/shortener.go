package shortener

type Repository struct{}

func (r *Repository) Get(shortenURL string) (string, error) {
	return "", nil
}

func (r *Repository) Set(URL string) (string, error) {
	return "", nil
}

func (r *Repository) Delete(URL string) error {
	return nil
}
