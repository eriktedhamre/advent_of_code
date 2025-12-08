package day05

import java.io.File
import kotlin.math.max

/**
 * Advent of Code 2025 - Day 05
 */
object Day05 {

    /**
     * Solve Part 1 of Day 05.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 1.
     */
    fun part1(intervals : List<Pair<Long, Long>>, ids: List<Long>): Long {

        return ids.fold(0) {count, id ->
            val match = intervals.firstOrNull { limits ->
                id >= limits.first && id <= limits.second
            }

            if (match != null) {
                count + 1
            } else {
                count
            }
        }
    }

    /**
     * Solve Part 2 of Day 05.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 2.
     */
    fun part2(input: MutableList<Pair<Long, Long>>): Long {
        input.sortBy { it.first }
        val mergedIntervals = mutableListOf<Pair<Long, Long>>()

        var mergedInterval = input.first()
        for (interval in input) {
            // Interval start is inside the current interval
            if (interval.first >= mergedInterval.first && interval.first <= mergedInterval.second) {
                mergedInterval = mergedInterval.first to max(mergedInterval.second, interval.second)
            } else {
                mergedIntervals.add(mergedInterval)
                mergedInterval = interval
            }
        }
        mergedIntervals.add(mergedInterval)
        return mergedIntervals.fold(0L) {sum, interval -> sum + (interval.second - interval.first) + 1}
    }
}

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Usage: ./gradlew day05 -Pargs=<inputfile>")
        return
    }

    val intervals = mutableListOf<Pair<Long, Long>>()
    val ids = mutableListOf<Long>()

    File(args[0]).useLines { lines ->
        lines.forEach { line ->
            when {
                line.contains('-') -> {
                    val numbers = line.split('-').map { it.trim().toLong() }
                    intervals.add(numbers[0] to numbers[1])
                }
                line.isNotEmpty() -> ids.add(line.toLong())
            }
        }
    }


    println(Day05.part2(intervals))

}