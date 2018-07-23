package main

import "fmt"

var labyrinth = [10][]string{}

func initLab()  {
	labyrinth[0] = []string{"#","o","#","#","#","#","#","#","#","#"}
	labyrinth[1] = []string{"#","o","#","#","#","#","#","#","#","#"}
	labyrinth[2] = []string{"#","o","o","o","o","#","#","#","#","#"}
	labyrinth[3] = []string{"#","#","o","#","o","#","#","#","#","#"}
	labyrinth[4] = []string{"#","#","#","#","o","#","#","#","#","#"}
	labyrinth[5] = []string{"#","#","o","#","o","o","#","#","#","#"}
	labyrinth[6] = []string{"#","#","o","#","#","o","#","#","#","#"}
	labyrinth[7] = []string{"#","#","o","o","o","o","#","#","#","#"}
	labyrinth[8] = []string{"#","#","o","#","#","#","#","#","#","#"}
	labyrinth[9] = []string{"#","#","o","#","#","#","#","#","#","#"}
}

var startPosition  = Position{0,1}
var endPosition = Position{8,2}

type Position struct {
	y,x int
}

type CurrentPosition struct {
	position Position
	parent* CurrentPosition
}

func (currentPosition CurrentPosition) canMoveLeft() bool {
	if currentPosition.position.x > 0 && canMove(currentPosition.position.y, currentPosition.position.x-1) {
		return true
	} else {
		return false
	}
}

func (currentPosition CurrentPosition) canMoveRight() bool {
	if currentPosition.position.x < len(labyrinth[currentPosition.position.y])-1 && canMove(currentPosition.position.y, currentPosition.position.x+1) {
		return true
	} else {
		return false
	}
}

func (currentPosition CurrentPosition) canMoveUp() bool {
	if currentPosition.position.y > 0 && canMove(currentPosition.position.y-1, currentPosition.position.x) {
		return true
	} else {
		return false
	}
}

func (currentPosition CurrentPosition) canMoveDown() bool {
	if currentPosition.position.y < len(labyrinth)-1 && canMove(currentPosition.position.y+1, currentPosition.position.x) {
		return true
	} else {
		return false
	}
}

func (currentPosition CurrentPosition) moveLeft() CurrentPosition {
	return CurrentPosition{Position{currentPosition.position.y, currentPosition.position.x-1},&currentPosition}
}

func (currentPosition CurrentPosition) moveRight() CurrentPosition {
	return CurrentPosition{Position{currentPosition.position.y, currentPosition.position.x+1},&currentPosition}
}

func (currentPosition CurrentPosition) moveUp() CurrentPosition {
	return CurrentPosition{Position{currentPosition.position.y-1, currentPosition.position.x},&currentPosition}
}

func (currentPosition CurrentPosition) moveDown() CurrentPosition {
	return CurrentPosition{Position{currentPosition.position.y+1, currentPosition.position.x},&currentPosition}
}

func canMove(y int,x int) bool {
	return labyrinth[y][x] == "o"
}

func main(){
	initLab()
	startPosition := CurrentPosition{startPosition,nil}
	path := step(&startPosition,make([]CurrentPosition,0),make([]CurrentPosition,0))
	exitPosition := find(path,CurrentPosition{endPosition,nil})
	if exitPosition == nil {
		fmt.Println("Exit has not been found")
	} else {
		positions := make([]CurrentPosition,0)
		fmt.Println(traverse(exitPosition,positions))
	}

}

func traverse(position* CurrentPosition, positions []CurrentPosition) *[]CurrentPosition {
	positions = append(positions,*position)
	if position.parent != nil {
		return traverse(position.parent,positions)
	} else {
		return &positions
	}
}

func step(currentPosition* CurrentPosition, positionsToVisit []CurrentPosition,visited []CurrentPosition) []CurrentPosition {
	if currentPosition == nil{
		return visited
	} else {
		if currentPosition.canMoveUp() && !contains(visited,currentPosition.moveUp()){
			positionsToVisit = append(positionsToVisit,currentPosition.moveUp())
		}
		if currentPosition.canMoveRight() && !contains(visited,currentPosition.moveRight()){
			positionsToVisit = append(positionsToVisit,currentPosition.moveRight())
		}
		if currentPosition.canMoveLeft() && !contains(visited,currentPosition.moveLeft()){
			positionsToVisit = append(positionsToVisit,currentPosition.moveLeft())
		}
		if currentPosition.canMoveDown() && !contains(visited,currentPosition.moveDown()){
			positionsToVisit = append(positionsToVisit,currentPosition.moveDown())
		}
	}
	if len(positionsToVisit) == 0 {
		return step(nil,positionsToVisit,visited)
	} else {
		return step(&positionsToVisit[0],positionsToVisit[1:],append(visited,*currentPosition))
	}
}

func find(positions []CurrentPosition,position CurrentPosition) *CurrentPosition {
	for _,val := range positions {
		if val.position.x == position.position.x && val.position.y == position.position.y {
			return &val
		}
	}
	return nil
}

func contains(positions []CurrentPosition, position CurrentPosition) bool{
	for _,val := range positions {
		if val.position.x == position.position.x && val.position.y == position.position.y{
			return true
		}
	}
	return false
}


