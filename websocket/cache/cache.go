package cache

type CacheAdapter interface {
	// 清空数据
	ClearAll() error

	// 增加房间成员数量
	IncreaseRoomMemberCount(room string, count int) error

	DecreaseRoomMemberCount(room string, count int) error
}
