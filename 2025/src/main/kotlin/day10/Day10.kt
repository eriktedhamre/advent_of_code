package day10

import day10.Day10.intToBinaryString
import java.io.File

/**
 * Advent of Code 2025 - Day 10
 */
object Day10 {

    fun intToBinaryString(button: Int): String {
        return button.toString(2).padStart(32, '0')
    }

    fun stepsToGoal(goalState: Int, buttons: List<Int>): Int {
        // Queue holds: Pair(CurrentState, Depth)
        val queue = ArrayDeque<Pair<Int, Int>>()
        queue.add(0 to 0) // Start at state 0, depth 0

        val visited = mutableSetOf<Int>()
        visited.add(0)

        while (queue.isNotEmpty()) {
            val (current, depth) = queue.removeFirst()

            if (current == goalState) {
                return depth
            }

            for (button in buttons) {
                val next = current xor button
                if (next !in visited) {
                    visited.add(next)
                    queue.add(next to depth + 1)
                }
            }
        }
        return -1
    }

    /**
     * Solve Part 1 of Day 10.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 1.
     */
    fun part1(goalStates: List<Int>, buttons: List<List<Int>>): Any {

        return goalStates.withIndex().fold(0) { acc, state ->
            val steps = stepsToGoal(state.value, buttons[state.index])
            acc + steps
        }
    }

    /**
     * Solve Part 2 of Day 10.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 2.
     */
    fun part2(input: String): Any {
    return 0
    }
}

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Usage: ./gradlew day10 -Pargs=<inputfile>")
        return
    }

    var goalStates = mutableListOf<Int>()
    var buttons = mutableListOf<List<Int>>()
    var others = mutableListOf<List<Int>>()

    File(args[0]).useLines { lines ->
        lines.forEach { line ->
            val split = line.split(' ')
            var goal: Int = 0
            split.first().substring(1, split.first().length - 1).withIndex().forEach { it ->

                when (it.value) {
                    '.' -> goal = goal or (0 shl it.index)
                    '#' -> goal = goal or (1 shl it.index)
                    else -> throw IllegalArgumentException("Invalid character in goal state: $it")
                }
            }
            goalStates.add(goal)

            val other = split.last().substring(1, split.last().length - 1).split(',').map { it.toInt() }
            others.add(other)

            val currentButtons = split.subList(1, split.size - 1).map { buttonString ->
                var button: Int = 0
                buttonString.substring(1, buttonString.length - 1).split(',').map { it.toInt() }.withIndex().forEach {
                    button = button or (1 shl it.value)
                }
                button
            }
            buttons.add(currentButtons)
        }
    }

    println(Day10.part1(goalStates, buttons))
}