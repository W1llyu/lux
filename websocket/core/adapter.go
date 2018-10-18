package core

import (
	"sync"
	"github.com/W1llyu/go-socket.io"
	"github.com/W1llyu/lux/websocket/cache"
)

/*
 重新实现原socketio包的broadcastAdapter
 赋予其他一些特性
 */
type LuxBroadcast struct {
	// 缓存
	c cache.CacheAdapter
	// 房间与socket map
	m map[string]map[string]socketio.Socket
	sync.RWMutex
}

func NewLuxBroadcast(cacheAdapter cache.CacheAdapter) socketio.BroadcastAdaptor {
	return &LuxBroadcast{
		c: cacheAdapter,
		m: make(map[string]map[string]socketio.Socket),
	}
}


func (b *LuxBroadcast) Join(room string, socket socketio.Socket) error {
	b.Lock()
	defer b.Unlock()

	sockets, ok := b.m[room]
	if !ok {
		sockets = make(map[string]socketio.Socket)
	}

	_, ok = sockets[socket.Id()]
	if ok {
		return nil
	}

	b.c.IncreaseRoomMemberCount(room, 1)
	sockets[socket.Id()] = socket
	b.m[room] = sockets

	return nil
}

func (b *LuxBroadcast) Leave(room string, socket socketio.Socket) error {
	b.Lock()
	defer b.Unlock()

	sockets, ok := b.m[room]
	if !ok {
		return nil
	}

	_, ok = sockets[socket.Id()]
	if !ok {
		return nil
	}

	b.c.DecreaseRoomMemberCount(room, 1)
	delete(sockets, socket.Id())
	if len(sockets) == 0 {
		delete(b.m, room)
		return nil
	}
	b.m[room] = sockets
	return nil
}

func (b *LuxBroadcast) Send(ignore socketio.Socket, room, event string, args ...interface{}) error {
	b.RLock()
	sockets := b.m[room]
	for id, s := range sockets {
		if ignore != nil && ignore.Id() == id {
			continue
		}
		s.Emit(event, args...)
	}
	b.RUnlock()
	return nil
}

func (b *LuxBroadcast) Sockets() map[string]map[string]socketio.Socket {
	return b.m
}