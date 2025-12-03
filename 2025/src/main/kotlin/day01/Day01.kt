package day01

import java.io.File
import kotlin.collections.fold

/**
 * Advent of Code 2025 - Day 01
 */
object Day01 {

    /**
     * Solve Part 1 of Day 01.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 1.
     */
    fun part1(input: List<Pair<Char, Int>>): Any {
        val target = 0
        val start = 50
        val max = 99

        val res = input.fold(Pair(0, start)) {
            acc, pair ->
            val magnitude = pair.second % (max + 1)
            val current = acc.second
            val newStart = when(pair.first){
                'L' -> {
                    if (current >= magnitude) {
                        current - magnitude
                    } else {
                        (max + 1) - (magnitude - current)
                    }
                }

                else -> {
                    (current + magnitude) % (max + 1)
                }
            }

            if (newStart == target) {
                Pair(acc.first + 1, newStart)
            } else {
                Pair(acc.first, newStart)
            }
        }
        return res.first
    }

    /**
     * Solve Part 2 of Day 02.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 2.
     */
    fun part2(input: List<Pair<Char, Int>>): Any {
        val target = 0
        val start = 50
        val max = 99

        val res = input.fold(Pair(0, start)) {
                acc, pair ->
            var zeroes = Math.floorDiv(pair.second, (max + 1))
            val magnitude = pair.second % (max + 1)
            val current = acc.second
            val newStart = when(pair.first){
                'L' -> {
                    if (current >= magnitude) {
                        current - magnitude
                    } else {
                        if (current != 0) {
                            zeroes++
                        }
                        (max + 1) - (magnitude - current)
                    }
                }

                else -> {
                    if (current + magnitude > (max + 1)) {
                        zeroes++
                    }
                    (current + magnitude) % (max + 1)
                }
            }



            if (newStart == target) {
                Pair(acc.first + 1 + zeroes, newStart)

            } else {
                Pair(acc.first + zeroes, newStart)
            }
        }
        return res.first
    }
}

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Usage: ./gradle <inputfile>")
        return
    }

    val input: List<Pair<Char, Int>> = File(args[0]).readLines()
        .map { it[0] to it.drop(1).toInt() }

    println(Day01.part2(input))
}