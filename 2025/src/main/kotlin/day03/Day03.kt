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


            intervals.add(0 to bulbs.size)
            while (lit.size != 12) {
                val currentInterval = intervals.removeLast()

                val max = if ((currentInterval.first - currentInterval.second) == 0) {
                    bulbs.withIndex().elementAt(currentInterval.first)
                } else {
                    bulbs.subList(currentInterval.first, currentInterval.second).withIndex().maxBy { it.value }
                }

                val maxFixedIndex = if ((currentInterval.first - currentInterval.second) == 0) max else IndexedValue(currentInterval.first + max.index, max.value)
                lit.add(maxFixedIndex)

                if (maxFixedIndex.index != currentInterval.first) {
                    val leftInterval = Pair(currentInterval.first, maxFixedIndex.index)
                    intervals.addLast(leftInterval)
                }

                if ( currentInterval.second - (maxFixedIndex.index + 1) > 0) {
                    val rightInterval = Pair(maxFixedIndex.index + 1, currentInterval.second)
                    intervals.addLast(rightInterval)
                }
            }

            lit.sortBy { it.index }

            val partialSum = lit.fold("") {str, num -> str + num.value.toString()}.toLong()
            println(partialSum)
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