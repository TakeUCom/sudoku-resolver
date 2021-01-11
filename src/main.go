package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	question := newQuestion()
	set(question)
	question.solve()
	question.print()
}

func set(question *Question) {
	qstr := make([]string, 9)
	qstr[0] = "?7???????"
	qstr[1] = "3195???72"
	qstr[2] = "?5???????"
	qstr[3] = "?35?????6"
	qstr[4] = "????4?823"
	qstr[5] = "9??23????"
	qstr[6] = "?9???84?7"
	qstr[7] = "?2????368"
	qstr[8] = "??64????1"
	for i := 0; i < 9; i++ {
		for j, s := range strings.Split(qstr[i], "") {
			if s != "?" {
				n, _ := strconv.Atoi(s)
				question.setNumber(i, j, n)
			}
		}
	}
}

type Cell struct {
	fixNumber    int
	availNumbers [9]bool
	packs        [][]*Cell
}

func newCell() *Cell {
	cell := new(Cell)
	cell.fixNumber = 0
	cell.availNumbers = [9]bool{true, true, true, true, true, true, true, true, true}
	cell.packs = make([][]*Cell, 0)
	return cell
}

func (cell *Cell) addPack(pack []*Cell) {
	cell.packs = append(cell.packs, pack)
}

func (cell *Cell) setNumber(number int) {
	cell.fixNumber = number
}

func (cell *Cell) removeAvailNumber(number int) {
	cell.availNumbers[number-1] = false
}

func (cell *Cell) isAvail(number int) bool {
	return cell.availNumbers[number-1]
}

func (cell *Cell) getAvailNumbers() []int {
	availNumbers := []int{}
	for i, avail := range cell.availNumbers {
		if avail {
			availNumbers = append(availNumbers, i+1)
		}
	}
	return availNumbers
}

type Question struct {
	cells [][]*Cell
	packs [][]*Cell
}

func newQuestion() *Question {
	q := new(Question)
	q.cells = make([][]*Cell, 9)
	for i := 0; i < 9; i++ {
		q.cells[i] = []*Cell{newCell(), newCell(), newCell(), newCell(), newCell(), newCell(), newCell(), newCell(), newCell()}
	}

	return q
}

func (question *Question) setNumber(x int, y int, number int) {
	fmt.Printf("x:%d.y:%d.number:%d.\n", x, y, number)
	question.cells[x][y].setNumber(number)
	r := x / 3
	c := y / 3

	for i := 0; i < 9; i++ {
		if i != x {
			question.cells[i][y].removeAvailNumber(number)
		}
		if i != y {
			question.cells[x][i].removeAvailNumber(number)
		}

		ri := i / 3
		ci := i % 3
		tx := r*3 + ri
		ty := c*3 + ci
		if tx != x && ty != y {
			question.cells[tx][ty].removeAvailNumber(number)
		}
	}
}

func (question *Question) getAvailNumbersOfRow(x, n int) []int {
	avails := []int{}
	for i := 0; i < 9; i++ {
		cell := question.cells[x][i]
		if cell.fixNumber == n {
			avails = []int{}
			break
		}

		if cell.isAvail(n) {
			avails = append(avails, i)
		}
	}
	return avails
}

func (question *Question) getAvailNumbersOfColumn(y, n int) []int {
	avails := []int{}
	for i := 0; i < 9; i++ {
		cell := question.cells[i][y]
		if cell.fixNumber == n {
			avails = []int{}
			break
		}

		if cell.isAvail(n) {
			avails = append(avails, i)
		}
	}
	return avails
}

func (question *Question) getAvailNumbersOfBlock(b, n int) []int {
	avails := []int{}
	r := b / 3
	c := b % 3

	for i := 0; i < 9; i++ {
		ri := i / 3
		ci := i % 3
		cell := question.cells[r*3+ri][c*3+ci]
		if cell.fixNumber == n {
			avails = []int{}
			break
		}

		if cell.isAvail(n) {
			avails = append(avails, i)
		}
	}

	return avails
}

func (question *Question) print() {
	for _, row := range question.cells {
		s := ""
		for _, cell := range row {
			if cell.fixNumber == 0 {
				s += "?"
			} else {
				s += strconv.Itoa(cell.fixNumber)
			}
		}
		fmt.Println(s)
	}
}

func (question *Question) solve() {
	for {
		count := question.solve1()
		count += question.solve2()
		if count == 0 {
			break
		}
	}
}

func (question *Question) solve1() int {
	count := 0
	for x, row := range question.cells {
		for y, cell := range row {
			availNumbers := cell.getAvailNumbers()
			if cell.fixNumber == 0 && len(availNumbers) == 1 {
				question.setNumber(x, y, availNumbers[0])
				count++
			}
		}
	}
	return count
}

func (question *Question) solve2() int {
	count := 0
	for i := 0; i < 9; i++ {
		for n := 1; n <= 9; n++ {
			availsRow := question.getAvailNumbersOfRow(i, n)
			if len(availsRow) == 1 {
				question.setNumber(i, availsRow[0], n)
				count++
			}

			availsColumn := question.getAvailNumbersOfColumn(i, n)
			if len(availsColumn) == 1 {
				question.setNumber(availsColumn[0], i, n)
				count++
			}

			availsBlock := question.getAvailNumbersOfBlock(i, n)
			if len(availsBlock) == 1 {
				r := i / 3
				c := i % 3
				ri := availsBlock[0] / 3
				ci := availsBlock[0] % 3
				question.setNumber(r*3+ri, c*3+ci, n)
				count++
			}
		}
	}
	return count
}
