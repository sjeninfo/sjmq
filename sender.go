package sjmq

type ISender interface {
	SendEvent(event interface{}) error
}

type Sender struct {
	outboxRepo *OutboxRepository
	ctx        *SenderContext
}

func NewSender(ctx *SenderContext) *Sender {
	sender := new(Sender)
	sender.outboxRepo = NewOutboxRepository(ctx.GormDBContext)
	sender.ctx = ctx
	return sender
}

func (s *Sender) SendEvent(event interface{}) error {
	err := s.outboxRepo.AddEvent(event)
	if err != nil {
		return err
	}
	if !s.ctx.IsInTransaction() {
		s.ctx.notifyHandler()
	}
	return nil
}
