package link_service

import "errors"

func (s *LinkService) RemoverLink(linkID string) error {
	link, err := s.repo.FindByID(linkID)
	if err != nil {
		return err
	}

	if link == nil {
		return errors.New("link não encontrado")
	}

	s.repo.Delete(link.ID.String())

	return nil
}
