package main

import "fmt"

type pos struct {
    x, y int
}

type figureOnBoard struct {
	f string
	p pos
}

var dim pos
var inc = func(i int) int { return i + 1 }
var dec = func(i int) int { return i - 1 }

func main() {
    // entry point: board dimension and quantity of each type
    // of figures (Kings, Queens, Bishops, Rooks and Knights)
	start(pos{3, 4}, 0, 0, 1, 2, 1)
}

func start(d pos, kings, queens, bishops, rooks, knights int) {
    if d.x < 1 || d.y < 1 {
    	fmt.Println("The input dimensions should be positive values, exit program")
    } else if kings < 0 || queens < 0 || bishops < 0 || rooks < 0 || knights < 0 {
    	fmt.Println("The quantity of each of figure should be positive, exit program")
    } else if kings == 0 && queens == 0 && bishops == 0 && rooks == 0 && knights == 0 {
	    fmt.Println("At least one figure should be present, exit program")
    } else {
    	fmt.Println(internalStart(d, kings, queens, bishops, rooks, knights))
    }
}

func internalStart(d pos, kings, queens, bishops, rooks, knights int) int {
	dim = d
    figures := make([]string, 0, kings + queens + bishops + rooks + knights)
    figures = append(figures, makeFigureSlice("king", kings)...)
    figures = append(figures, makeFigureSlice("queen", queens)...)
    figures = append(figures, makeFigureSlice("bishop", bishops)...)
    figures = append(figures, makeFigureSlice("rook", rooks)...)
    figures = append(figures, makeFigureSlice("knight", knights)...)
	return sum(permutations(figures))
}

func sum(l *[][]string) int {
	if len(*l) == 0 {
		return 0
	} else {
		temp := (*l)[1:]
		return calc2(&pos{1, 1}, &[]figureOnBoard {}, &(*l)[0]) + sum(&temp)
	}
}

func calc2(p *pos, path *[]figureOnBoard, f *[]string) int {
	if !isInBounds(p) {
		return 0
	} else {
		fi := figureOnBoard{(*f)[0], *p}
		newPath := append(*path, fi)
		newPos := nextPos(p)
		x := 0
		if check(path, &fi) {
			if len((*f)[1:]) == 0 {
				x = 1
			} else {
				temp := (*f)[1:]
				x = calc2(newPos, &newPath, &temp)
			}
		}
		return x + calc2(newPos, path, f)
	}
}

func check(path *[]figureOnBoard, n *figureOnBoard) bool {
	for _, value := range *path {
		if !isCheck(&value, n) { return false }
	}
	return true
}

func isCheck(a, b *figureOnBoard) bool {
	return internalCheck(a, b) && internalCheck(b, a)
}

func internalCheck(a, b *figureOnBoard) bool {
	switch a.f {
		case "king": return b.p != pos{a.p.x - 1, a.p.y} && b.p != pos{a.p.x - 1, a.p.y - 1} && b.p != pos{a.p.x, a.p.y - 1} && b.p != pos{a.p.x + 1, a.p.y - 1} && b.p != pos{a.p.x + 1, a.p.y} && b.p != pos{a.p.x + 1, a.p.y + 1} && b.p != pos{a.p.x, a.p.y + 1} && b.p != pos{a.p.x - 1, a.p.y + 1}
		case "queen": return a.p.x != b.p.x && a.p.y != b.p.y && calcBishop(&a.p, &b.p, dec, dec) && calcBishop(&a.p, &b.p, dec, inc) && calcBishop(&a.p, &b.p, inc, dec) && calcBishop(&a.p, &b.p, inc, inc) 
		case "bishop": return calcBishop(&a.p, &b.p, dec, dec) && calcBishop(&a.p, &b.p, dec, inc) && calcBishop(&a.p, &b.p, inc, dec) && calcBishop(&a.p, &b.p, inc, inc)
		case "rook": return a.p.x != b.p.x && a.p.y != b.p.y
		case "knight": return b.p != pos{a.p.x - 1, a.p.y - 2} && b.p != pos{a.p.x - 2, a.p.y - 1} && b.p != pos{a.p.x - 1, a.p.y + 2} && b.p != pos{a.p.x - 2, a.p.y + 1} && b.p != pos{a.p.x + 1, a.p.y + 2} && b.p != pos{a.p.x + 2, a.p.y + 1} && b.p != pos{a.p.x + 1, a.p.y - 2} && b.p != pos{a.p.x + 2, a.p.y - 1}
	}
	panic("Undefined figure " + a.f)
}

func calcBishop(bishop, p *pos, opx, opy func(int) int) bool {
	if !isInBounds(bishop) {
		return true
	} else {
		b := pos{opx(bishop.x), opy(bishop.y)}
		if (b == *p) { 
			return false
		} else {
			return calcBishop(&b, p, opx, opy)
		}
	}
}

func isInBounds(p *pos) bool {
	return p.x > 0 && p.y > 0 && p.x <= dim.x && p.y <= dim.y
}

func nextPos(p *pos) *pos {
	if p.x + 1 > dim.x {
		return &pos{1, p.y + 1}
	} else {
		return &pos{p.x + 1, p.y}
	}
}

func permutations(figures []string) *[][]string {
	r := make([][]string, 0)
	internalPerm([]string {}, figures, figures, len(figures), &r)	
	return &r
}

func internalPerm(res, f, all []string, z int, result *[][]string) {
	if len(f) == 0 {
		if len(res) == z { *result = append(*result, res) }
	} else {
		q := filter(f[0], all)
		internalPerm(append(res, f[0]), q, q, z, result)
		internalPerm(res, skipSameFigures(f[0], f[1:]), all, z, result)
	}
}

func filter(f string, list []string) []string {
	if len(list) == 0 {
		return []string {}
	} else if list[0] != f {
		return append([]string {list[0]}, filter(f, list[1:])...)
	} else {
		return list[1:]
	}
}
  
func skipSameFigures(f string, l []string) []string {
	if len(l) == 0 || f != l[0] {
		return l
	} else {
		return skipSameFigures(f, l[1:])
	}
}

func makeFigureSlice(figure string, qnt int) []string {
	r := make([]string, qnt)
	for i, _ := range r {
		r[i] = figure
	}
	return r
}
