package sjmq

import "gorm.io/gorm"

//type GormDBContextFactory func() *GormDBContext

type GormDBContext struct {
	source *gorm.DB
	tx     *gorm.DB
}

func NewGormDBContext(db *gorm.DB) *GormDBContext {
	return &GormDBContext{source: db}
}

func (g *GormDBContext) DB() *gorm.DB {
	if g.tx != nil {
		return g.tx
	} else {
		return g.source
	}
}

func (g *GormDBContext) Begin() error {
	g.tx = g.source.Begin()
	return nil
}

func (g *GormDBContext) Commit() error {
	err := g.tx.Commit().Error
	if err != nil {
		return err
	}
	g.tx = nil
	return nil
}

func (g *GormDBContext) IsInTransaction() bool {
	return g.tx != nil
}

func (g *GormDBContext) Rollback() {
	g.tx.Rollback()
	g.tx = nil
}

type SenderContext struct {
	*GormDBContext
	notifyHandler func()
}

func NewSenderContext(db *gorm.DB, notifyHandler func()) *SenderContext {
	s := new(SenderContext)
	s.GormDBContext = NewGormDBContext(db)
	s.notifyHandler = notifyHandler
	return s
}

func (s *SenderContext) Commit() error {
	err := s.GormDBContext.Commit()
	if err != nil {
		return err
	}
	s.notifyHandler()
	return nil
}
