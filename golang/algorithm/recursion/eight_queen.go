package main

import (
    "fmt"
    "math"
)

// 八皇后
// 在 8×8 格的国际象棋上摆放 8 个皇后，使其不能互相攻击，即
// 任意两个皇后都不能处于同一行、同一列或同一斜线上，问有多少种摆法。
// 思路分析
// 先将第一个皇后摆放在第一行第一列，接着将第二个皇后先摆在第二行第
// 一列，如果不行再摆放在第二行第二列，还不行继续在摆放在下一列，依
// 次类推，整个过程递归进行，当最后一个皇后摆放完成后，回溯到上一个皇后，继续摆放。
// 使用一个一维数组 positions 存储象棋摆放的位置，下标表示第几个象棋，数组的值
// 表示第几列。那么摆放第 n 个皇后时，**positions[i] == positions[n]** 表示在同一列
// **|n - i| == |positions[n] - positions[i]|** 表示在同一斜线。

type EightQueen struct {
    positions []int // 存储每 8 个皇后，一种摆放的位置
}

func NewEightQueen() *EightQueen {
    positions := make([]int, 8)
    return &EightQueen{positions: positions}
}

func (queen *EightQueen) putQueen(n int) {
    // 最后一个皇后已经放置完成
    if n == len(queen.positions) {
        // 打印当前摆放的位置
        fmt.Println(queen.positions)
        return
    }
    for i := 0; i < len(queen.positions); i++ {
        // i=0 时，假设当前皇后可以放在第一列
        // 如果不能放，将进行下一次循环，当前皇后放在下一个位置
        queen.positions[n] = i
        // 判断是否可以放
        if queen.isCanPut(n) {
            // 放置下一个皇后
            queen.putQueen(n + 1)
        }
    }
}


// 判断当前皇后是否和已经摆放过的皇后冲突
func (queen *EightQueen) isCanPut(n int) bool {
    positions := queen.positions
    for i := 0; i < n; i++ {
        // positions[n] == positions[i] 表示在同一列
        // math.Abs(float64(n - i)) == math.Abs(float64(positions[n] - positions[i]) 表示同一斜线
        if positions[n] == positions[i] ||
            math.Abs(float64(n - i)) == math.Abs(float64(positions[n] - positions[i])){
            return false
        }
    }
    return true
}

func main() {
    eightQueen := NewEightQueen()
    eightQueen.putQueen(0)
}
