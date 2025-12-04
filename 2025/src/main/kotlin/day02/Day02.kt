package day02

import kotlin.io.path.Path
import kotlin.collections.forEach
import kotlin.io.path.readText

/**
 * Advent of Code 2025 - Day 02
 */
object Day02 {

    /**
     * Solve Part 1 of Day 02.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 1.
     */
    fun part1(input: List<Pair<Long, Long>>): Long {

        val res = mutableListOf<Long>()

        input.forEach { (start, end) ->
            for (i in start..end) {
                val iString = i.toString()

                if (iString.length %2 != 0) continue

                if (iString.take(iString.length/2) == iString.takeLast(iString.length/2)){
                    res.add(i)
                }
            }
        }

        return res.sum()
    }

    /**
     * Solve Part 2 of Day 02.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 2.
     */
    fun part2(input: List<Pair<Long, Long>>): Long{
        val res = mutableListOf<Long>()

        input.forEach { (start, end) ->
            for (i in start..end) {
                val iString = i.toString()

                for (chunkSize in 1..(iString.length/2)) {
                    val list = iString.chunked(chunkSize)
                    val matches = list.all { it -> it == list[0] }
                    if (matches) {
                        res.add(i)
                        break
                    }
                }
            }
        }

        return res.sum()
    }
}

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Usage: ./gradlew day02 -Pargs=<inputfile>")
        return
    }

    val input: List<Pair<Long, Long>> = Path(args[0]).readText().split(',', '-').map { it.trim() }.chunked(2).map { (a, b) -> a.toLong() to b.toLong() }

    println(Day02.part2(input))

}