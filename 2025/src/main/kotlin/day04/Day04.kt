package day04

import java.io.File

/**
 * Advent of Code 2025 - Day 04
 */
object Day04 {

    private val moves = listOf(0 to 1, 1 to 0, -1 to 0 , 0 to -1, 1 to 1, 1 to -1, -1 to 1, -1 to -1)

    fun inBounds(row: Int, col: Int, rows: Int, cols : Int): Boolean {
        return (row > -1) && (row < rows)  && (col > -1) && (col < cols)
    }

    /**
     * Solve Part 1 of Day 04.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 1.
     */
    fun part1(input: List<List<Char>>): Any {

        val rows = input.size
        val cols = input[0].size
        var accessible = 0
        for (row in 0..rows - 1) {
            for (col in 0..cols - 1) {
                if (input[row][col] == '.') continue
                var neighborCount = 0
                for (move in moves) {
                    val newRow = row + move.first
                    val newCol = col + move.second

                    if (inBounds(newRow, newCol, rows, cols) && input[newRow][newCol] == '@') {
                        neighborCount++
                        if (neighborCount > 3) break
                    }
                }
                if (neighborCount <= 3) accessible++
            }
        }
        return accessible
    }

    /**
     * Solve Part 2 of Day 04.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 2.
     */
    fun part2(input: MutableList<MutableList<Char>>): Any {

        val rolls = mutableSetOf<Pair<Int, Int>>()

        val rows = input.size
        val cols = input[0].size

        for (row in 0..rows-1) {
            for (col in 0..cols-1) {
                if (input[row][col] == '@') {
                    rolls.add(row to col)
                }
            }
        }

        val start = rolls.size
        while (true) {
            val toRemove = mutableListOf<Pair<Int, Int>>()

            for (roll in rolls) {
                var neighborCount = 0
                for (move in moves) {
                    val newRow = roll.first + move.first
                    val newCol = roll.second + move.second

                    if (inBounds(newRow, newCol, rows, cols) && input[newRow][newCol] == '@') {
                        neighborCount++
                        if (neighborCount > 3) break
                    }
                }
                if (neighborCount <= 3) {
                    toRemove.add(roll)
                    input[roll.first][roll.second] = '.'
                }
            }

            if (toRemove.isEmpty()) break
            rolls.removeAll(toRemove)


        }
        return start - rolls.size
    }
}

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Usage: ./gradlew day04 -Pargs=<inputfile>")
        return
    }

    val input: List<List<Char>> = File(args[0]).readLines().map { it.toCharArray().toList() }
    val inputMutable = input.map { it.toMutableList() }.toMutableList()
    println(Day04.part2(inputMutable))

}