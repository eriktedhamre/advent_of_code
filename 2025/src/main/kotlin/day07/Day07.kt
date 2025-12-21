package day07

import java.io.File

/**
 * Advent of Code 2025 - Day 07
 */
object Day07 {

    fun inBounds(row: Int, col: Int, rows: Int, cols : Int): Boolean {
        return (row > -1) && (row < rows)  && (col > -1) && (col < cols)
    }

    /**
     * Solve Part 1 of Day 07.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 1.
     */
    fun part1(grid: MutableList<MutableList<Char>>, start: Pair<Long, Long>): Long {
        val lines = ArrayDeque<Pair<Long, Long>>()
        lines.addLast(start.copy(first = start.first + 1L))

        var splits = 0L
        while (lines.isNotEmpty()) {
            val pos = lines.removeFirst()
            val row = pos.first.toInt()
            val col = pos.second.toInt()

            if (!inBounds(row, col, grid.size, grid[0].size)) {
                continue
            }

            val cell = grid[row][col]
            if (cell == '|') {
                continue
            } else if (cell == '.') {
                grid[row][col] = '|'
                lines.addLast((row + 1).toLong() to col.toLong())
            } else {
                splits++
                lines.addLast((row + 1).toLong() to (col + 1).toLong())
                lines.addLast((row + 1).toLong() to (col - 1).toLong())
            }
        }
        return splits
    }

    /**
     * Solve Part 2 of Day 07.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 2.
     */
    fun part2(grid: List<List<Char>>, start: Pair<Long, Long>): Any {
        // Memoization

        val memTable = mutableMapOf<Pair<Long, Long>, Long>()
        fun countPaths(pos: Pair<Long, Long>): Long {
            if (memTable.containsKey(pos)) {
                return memTable[pos]!!
            }

            val row = pos.first.toInt()
            val col = pos.second.toInt()

            if (!inBounds(row, col, grid.size, grid[0].size)) {
                return 1L
            }

            val cell = grid[row][col]
            if (cell == '.') {
                val paths = countPaths((row + 1).toLong() to col.toLong())
                memTable[pos] = paths
                return paths
            } else {
                val leftPaths = countPaths((row + 1).toLong() to (col - 1).toLong())
                val rightPaths = countPaths((row + 1).toLong() to (col + 1).toLong())
                val totalPaths = leftPaths + rightPaths
                memTable[pos] = totalPaths
                return totalPaths
            }
        }
        return countPaths(start)
    }
}

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Usage: ./gradlew day07 -Pargs=<inputfile>")
        return
    }

    val grid = mutableListOf<MutableList<Char>>()
    var startPos : Pair<Long, Long>? = null

    File(args[0]).useLines { lines ->
        lines.withIndex().forEach { line ->
            grid.add(line.value.toMutableList())
            if (line.value.contains("S")) {
                startPos = line.index.toLong() to line.value.indexOf("S").toLong()
            }
        }
    }

    if (startPos == null) {
        println("Start position not found!")
        return
    }

    println(Day07.part2(grid, startPos))
}