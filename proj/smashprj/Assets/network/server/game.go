package main

import(
	"fmt"
//	"os"
	"net"
)

type GameObject struct{
	x float32
	y float32
	width float32
	height float32
	name string	
}

type GamePlayer struct{
	GameObject
	netconn net.Conn
}

type GameScene struct{
	GameObject	
	children map[string]*GamePlayer
}

func (m *GameScene) Init(){
	m.children = make(map[string]*GamePlayer)
}

func (m *GameScene) AddObj(newobj *GamePlayer){
	m.children[newobj.name] = newobj
}

func (m *GameScene) RemoveObj(name *string){
	delete(m.children, *name)
}

func (m *GameScene) GetObj(name string) *GamePlayer{
	return m.children[name]
}

func (m *GameScene) Shoot(name string){
	player := m.children[name]
	player.Shoot()
}

func (m *GamePlayer) Shoot(){
	fmt.Println("i shoot.")
}