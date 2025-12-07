package day03

import java.io.File
import kotlin.collections.ArrayDeque
import kotlin.collections.mutableListOf


/**
 * Advent of Code 2025 - Day 03
 */
object Day03 {

    /**
     * Solve Part 1 of Day 03.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 1.
     */
    fun part1(input: List<List<Long>>): Long{

        return input.fold(0L) { cumSum, it ->
            val firstBat = it.subList(0, it.size - 1).withIndex().maxBy { it.value }
            val secondBat = it.subList(firstBat.index + 1, it.size).withIndex().maxBy { it.value }
            cumSum + (firstBat.value.toString() + secondBat.value.toString()).toLong()
        }
    }

    /**
     * Solve Part 2 of Day 03.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 2.
     */
    fun part2(input: List<List<Long>>): Any {


        // Treat it like a tree
        // Max on an interval
        // then Max on the right interval
        // then Max on the left interval
        return input.fold(0L) { sum, bulbs ->
            val intervals = ArrayDeque<Pair<Int, Int>>()
            val lit = mutableListOf<IndexedValue<Long>>()


            intervals += 0 to bulbs.size

            while (lit.size < 12) {
                val (start, end) = intervals.removeLast()

                val maxIndexed = if (end - start == 1) {
                    bulbs.withIndex().elementAt(start)
                } else {
                    bulbs.subList(start, end)
                        .withIndex()
                        .maxBy { it.value }
                        .let { IndexedValue(start + it.index, it.value) }
                }

                lit += maxIndexed

                if (maxIndexed.index > start) {
                    intervals += start to maxIndexed.index
                }
                if (maxIndexed.index + 1 < end) {
                    intervals += (maxIndexed.index + 1) to end
                }
            }

            lit.sortBy { it.index }
            val partialSum =
                lit.joinToString(separator = "") { it.value.toString() }.toLong()
            sum + partialSum
        }
    }
}

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Usage: ./gradlew day03 -Pargs=<inputfile>")
        return
    }

    val input: List<List<Long>> = File(args[0]).readLines().map { it -> it.chunked(1) }.map { it -> it.map { it -> it.toLong() } }
    println(Day03.part2(input))
}