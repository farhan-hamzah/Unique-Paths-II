package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
    if m == 0 {
        return 0
    }
    n := len(obstacleGrid[0])

    // Jika titik awal atau akhir adalah rintangan, tidak ada jalur
    if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
        return 0
    }

    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    dp[0][0] = 1

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                dp[i][j] = 0 // Tidak bisa lewat rintangan
            } else {
                if i > 0 {
                    dp[i][j] += dp[i-1][j]
                }
                if j > 0 {
                    dp[i][j] += dp[i][j-1]
                }
            }
        }
    }

    return dp[m-1][n-1]
}

func main() {
    // Contoh tanpa rintangan
    grid1 := [][]int{
        {0, 0, 0},
        {0, 1, 0},
        {0, 0, 0},
    }
    fmt.Println("Unique paths with obstacles:", uniquePathsWithObstacles(grid1)) // Output: 2

    // Contoh dengan rintangan awal
    grid2 := [][]int{
        {1, 0},
        {0, 0},
    }
    fmt.Println("Unique paths with obstacles:", uniquePathsWithObstacles(grid2)) // Output: 0
}
