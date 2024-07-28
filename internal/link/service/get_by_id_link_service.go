package link_service

import "github.com/JohannBandelow/meus-links-go/internal/link"

func (s *LinkService) GetById(linkID string) (*link.Link, error) {
	link, err := s.repo.FindByID(linkID)
	if err != nil {
		return nil, err
	}

	return link, nil
}
