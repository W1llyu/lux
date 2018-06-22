package handler

import (
	"fmt"
	"net/http"
	"github.com/irelia_socket/websocket/core"
	"encoding/json"
	"sort"
)

type clientStat struct {
	Client string
	Token string
	ConnCount int
	SocketIds []string
}

type By func(c1, c2 *clientStat) bool

func (by By) Sort(clientStats []clientStat) {
	ps := &clientStatSorter{
		clientStats: clientStats,
		by:        by,
	}
	sort.Sort(ps)
}

type clientStatSorter struct {
	clientStats []clientStat
	by func(c1, c2 *clientStat) bool
}

func (c *clientStatSorter) Len() int {
	return len(c.clientStats)
}

func (c *clientStatSorter) Swap(i, j int) {
	c.clientStats[i], c.clientStats[j] = c.clientStats[j], c.clientStats[i]
}

func (c *clientStatSorter) Less(i, j int) bool {
	return c.by(&c.clientStats[i], &c.clientStats[j])
}


func GetClientStat(w http.ResponseWriter, r *http.Request) {
	var clientSlice []clientStat
	for _, clientStat := range loadClientMap() {
		clientSlice = append(clientSlice, *clientStat)
	}
	sortFunc := func(c1, c2 *clientStat) bool {
		return c1.ConnCount > c2.ConnCount
	}
	By(sortFunc).Sort(clientSlice)
	res := make(map[string]interface{})
	res["clients"] = clientSlice
	body, _ := json.Marshal(res)
	fmt.Fprint(w, string(body))
}

func loadClientMap() map[string]*clientStat {
	sessions := core.Sessions.Sessions()
	clientMap := make(map[string]*clientStat)
	for sid, conn := range sessions {
		client := conn.Request().URL.Query().Get("client")
		token := conn.Request().URL.Query().Get("token")
		c, ok := clientMap[fmt.Sprintf("%s_%s", client, token)]
		if ok {
			c.ConnCount += 1
			c.SocketIds = append(c.SocketIds, sid)
		} else {
			clientMap[fmt.Sprintf("%s_%s", client, token)] = &clientStat{
				Client: client,
				Token: token,
				ConnCount: 1,
				SocketIds: []string{sid},
			}
		}
	}
	return clientMap
}