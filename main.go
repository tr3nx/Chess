package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

type Board struct {
	grid     map[Coord]Piece
	Selected Coord
}

func NewBoard() *Board {
	b := &Board{
		grid: make(map[Coord]Piece),
	}

	// White
	b.Set(Piece{"Rook", 1}, Coord{1, 8})
	b.Set(Piece{"Knight", 1}, Coord{2, 8})
	b.Set(Piece{"Bishop", 1}, Coord{3, 8})
	b.Set(Piece{"Queen", 1}, Coord{4, 8})
	b.Set(Piece{"King", 1}, Coord{5, 8})
	b.Set(Piece{"Bishop", 1}, Coord{6, 8})
	b.Set(Piece{"Knight", 1}, Coord{7, 8})
	b.Set(Piece{"Rook", 1}, Coord{8, 8})

	// Black
	b.Set(Piece{"Rook", 2}, Coord{1, 1})
	b.Set(Piece{"Knight", 2}, Coord{2, 1})
	b.Set(Piece{"Bishop", 2}, Coord{3, 1})
	b.Set(Piece{"Queen", 2}, Coord{4, 1})
	b.Set(Piece{"King", 2}, Coord{5, 1})
	b.Set(Piece{"Bishop", 2}, Coord{6, 1})
	b.Set(Piece{"Knight", 2}, Coord{7, 1})
	b.Set(Piece{"Rook", 2}, Coord{8, 1})

	for i := 1; i <= 8; i++ {
		b.grid[Coord{i, 2}] = Piece{"Pawn", 2} // Black
		b.grid[Coord{i, 7}] = Piece{"Pawn", 1} // White
	}

	return b
}

func (b *Board) Render() string {
	var rend strings.Builder
	for j := 8; j >= 0; j-- {
		if j == 0 {
			rend.WriteString("    ")
		} else {
			rend.WriteString(strconv.Itoa(j) + "   ")
		}

		for i := 1; i <= 8; i++ {
			if j == 0 {
				rend.WriteString(string(alphamap[i-1]) + "   ")
				continue
			}

			p := b.Get(Coord{i, j})
			symbol := p.Symbol()

			if p.Side == 1 {
				symbol = color.HiBlackString(symbol)
			} else if p.Side == 2 {
				symbol = color.HiWhiteString(symbol)
			}

			if b.Selected.X == i && b.Selected.Y == j {
				highlight := color.New(color.FgHiBlack, color.BgHiYellow).SprintFunc()
				symbol = fmt.Sprintf("%v", highlight(symbol))
			}

			rend.WriteString(symbol + "   ")
		}
		rend.WriteString("\n\n")
	}
	return rend.String()
}

func (b *Board) Set(p Piece, c Coord) {
	b.grid[c] = p
}

func (b *Board) Get(c Coord) Piece {
	return b.grid[c]
}

func (b *Board) GetSymbol(c Coord) string {
	return b.grid[c].Symbol()
}

func (b *Board) Move(from, to Coord) {
	b.grid[to] = b.grid[from]
	b.grid[from] = Piece{}
	b.Unselect()
}

func (b *Board) Select(c Coord) Piece {
	b.Selected = c
	return b.grid[c]
}

func (b *Board) Unselect() {
	b.Selected = Coord{}
}

func (b *Board) ValidMove(from, to Coord) bool {
	p := b.Get(from)
	target := b.Get(to)

	if p.Side == target.Side {
		return false
	}

	switch p.Type {
	// case "King":
	// case "Queen":
	// case "Rook":
	// case "Knight":
	// case "Bishop":
	// case "Pawn":
	}
	return true
}

type Piece struct {
	Type string
	Side int
}

func (p Piece) Symbol() string {
	var symbol string
	switch p.Type {
	case "King":
		symbol = "K"
	case "Queen":
		symbol = "Q"
	case "Rook":
		symbol = "R"
	case "Knight":
		symbol = "N"
	case "Bishop":
		symbol = "B"
	case "Pawn":
		symbol = "P"
	default:
		symbol = "-"
	}
	return symbol
}

func min(check, min int) int {
	if check < min {
		return min
	}
	return check
}

func max(check, max int) int {
	if check > max {
		return max
	}
	return check
}

var alphamap = "abcdefgh"

func main() {
	b := NewBoard()

	var action int

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println(b.Render())
		fmt.Println(b.Select(Coord{b.Selected.X, b.Selected.Y}))
		switch action {
		case 0:
			fmt.Printf("Select a piece: ")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				input := scanner.Text()

				if input == "q" {
					goto Quit
				}

				x := strings.Index(alphamap, string(input[0])) + 1
				y, _ := strconv.Atoi(string(input[1]))

				b.Select(Coord{x, y})
				action = 1
				break
			}
		case 1:
			fmt.Printf("Select where to move: ")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				input := scanner.Text()

				if input == "q" {
					goto Quit
				}

				if input == "c" {
					b.Unselect()
					action = 0
					break
				}

				x := strings.Index(alphamap, string(input[0])) + 1
				y, _ := strconv.Atoi(string(input[1]))

				if b.ValidMove(b.Selected, Coord{x, y}) {
					b.Move(b.Selected, Coord{x, y})
					action = 0
				}
				break
			}
		}
	}

Quit:

	fmt.Println("done.")
}
