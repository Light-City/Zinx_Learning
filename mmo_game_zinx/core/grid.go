/*
 * @Author: 光城
 * @Date: 2020-10-29 16:56:43
 * @LastEditors: 光城
 * @LastEditTime: 2020-10-29 17:18:19
 * @Description: AOI 地图
 * @FilePath: /Zinx_Learning/mmo_game_zinx/core/grid.go
 */
package core

import (
	"fmt"
	"sync"
)

// AOI地图格子类型

type Grid struct {
	// 格子ID
	GID int
	// 格子的左边边界坐标
	MinX int
	// 格子的右边边界坐标
	MaxX int
	// 格子的上边边界坐标
	MinY int
	// 格子的下边边界坐标
	MaxY int
	// 当前格子内玩家/物体成员ID集合
	playerIDs map[int]bool
	// 保护当前集合的锁
	pIDLock sync.RWMutex
}

// 初始化当前格子的方法
func NewGrid(gID, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gID,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

// 给格子添加一个玩家
func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	g.playerIDs[playerID] = true
}

// 给格子删除一个玩家
func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerID)
}

// 得到当前格子的所有玩家
func (g *Grid) GetPlayerIDs() (playerIDs []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.Unlock()

	for k, _ := range g.playerIDs {
		playerIDs = append(playerIDs, k)
	}
	return
}

// fmt.Println(grid)=grid.String()
// 调试使用-打印出格子的基本信息
func (g *Grid) print() string {
	return fmt.Sprintf("Grid id:%d, minX:%d, maxX:%d, minY:%d, maxY:%d, playerIDs:%v",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MinY, g.playerIDs)
}
