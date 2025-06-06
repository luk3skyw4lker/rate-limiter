package store

type Store interface {
	GetCurrentCount(clientId string) (int, error)
	IncrementCount(clientId string) (int, error)
	ResetCount(clientId string) error
	Close() error
}

// In-memory store implementation
type InMemoryStore struct {
	counts map[string]int
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		counts: make(map[string]int),
	}
}

func (s *InMemoryStore) GetCurrentCount(clientId string) (int, error) {
	count, exists := s.counts[clientId]
	if !exists {
		return 0, nil // No count found for clientId
	}
	return count, nil
}

func (s *InMemoryStore) IncrementCount(clientId string) (int, error) {
	count, exists := s.counts[clientId]
	if !exists {
		count = 0
	}
	count++
	s.counts[clientId] = count
	return count, nil
}

func (s *InMemoryStore) ResetCount(clientId string) error {
	if _, exists := s.counts[clientId]; !exists {
		return nil // No count to reset for clientId
	}
	delete(s.counts, clientId)
	return nil
}

func (s *InMemoryStore) Close() error {
	// In-memory store does not require any cleanup
	return nil
}
