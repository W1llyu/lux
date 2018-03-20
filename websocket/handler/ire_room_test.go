package handler

import "testing"

func TestPushSidekiqTask (t *testing.T) {
	pushSidekiqTask(1, "1HPxfEV6greQ4UCronTs", LEAVETYPE)
	pushSidekiqTask(1, "1HPxfEV6greQ4UCronTs", JOINTYPE)
}