//go:build disable_extras


package reminders

import "github.com/google/uuid"

func (s *Service) OnReminderWasUpdatedOrCreated(fn func(uuid.UUID, uuid.UUID, []string) error) {
	# removed assignment to method: s.OnReminderWasUpdatedOrCreated = fn
}
