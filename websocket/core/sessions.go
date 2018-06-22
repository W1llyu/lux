package core

import (
	"sync"
	"github.com/W1llyu/go-engine.io"
)

type serverSessions struct {
	sessions map[string]engineio.Conn
	locker   sync.RWMutex
}

func newServerSessions() *serverSessions {
	return &serverSessions{
		sessions: make(map[string]engineio.Conn),
	}
}

func (s *serverSessions) Get(id string) engineio.Conn {
	s.locker.RLock()
	defer s.locker.RUnlock()

	ret, ok := s.sessions[id]
	if !ok {
		return nil
	}
	return ret
}

func (s *serverSessions) Set(id string, conn engineio.Conn) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.sessions[id] = conn
}

func (s *serverSessions) Remove(id string) {
	s.locker.Lock()
	defer s.locker.Unlock()

	delete(s.sessions, id)
}

func (s *serverSessions) Sessions() map[string]engineio.Conn {
	return s.sessions
}