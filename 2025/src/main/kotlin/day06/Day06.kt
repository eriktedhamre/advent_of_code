package day06

import java.io.File
import kotlin.math.max

/**
 * Advent of Code 2025 - Day 06
 */
object Day06 {

    /**
     * Solve Part 1 of Day 06.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 1.
     */
    fun part1(operandsNonRotated: List<List<Long>>, operators: List<Char> ): Any {

        var operands = mutableListOf<MutableList<Long>>()

        for (col in 0..(operandsNonRotated[0].size-1)) {
            var l = mutableListOf<Long>()
            for (row in 0..(operandsNonRotated.size-1)) {
                l.add(operandsNonRotated[row][col])
            }
            operands.add(l)
        }

        var sum = 0L
        for (item in operands.withIndex()) {
            var partialSum = if (operators.elementAt(item.index) == '+') {
                item.value.sum()
            } else {
                item.value.reduce { a, b -> a * b }
            }
            sum += partialSum
        }
        return sum
    }

    /**
     * Solve Part 2 of Day 06.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 2.
     */
    fun part2(operands: List<List<Long>>, operators: List<Char>): Any {

        var sum = 0L
        for (item in operands.withIndex()) {
            var partialSum = if (operators.elementAt(item.index) == '+') {
                item.value.sum()
            } else {
                item.value.reduce { a, b -> a * b }
            }
            sum += partialSum
        }
        return sum

    }
}

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Usage: ./gradlew day06 -Pargs=<inputfile>")
        return
    }

    val lines = File(args[0]).readLines()
    var operands = lines.take(lines.size-1)//.map { it.reversed() }
    val maxSize = operands.fold(0) {acc, string -> max(acc, string.length)}

    operands = operands.map { it.padEnd(maxSize) }

    var allFormatted = mutableListOf<List<Long>>()
    var formatted = mutableListOf<Long>()
    val emptyCol = operands.fold("") { acc, string -> acc + " " }
    println(emptyCol)
    for (i in operands[0].length-1 downTo 0) {

        var value = ""

        for (j in 0..operands.size-1){
            if (operands[j][i] != ' ') {
                value += operands[j][i]
            } else {
                value += ' '
            }
        }

        if (value == emptyCol){
            allFormatted.add(formatted.toList())
            formatted.clear()
        } else {
            formatted.add(value.trim().toLong())
            if (i == 0){
                allFormatted.add(formatted.toList())
            }
        }
    }

    //print(allFormatted)

    val operators = lines.last().split(Regex("\\s+")).map { it[0] }.reversed()

    println(Day06.part2(allFormatted, operators))


    /*
    File(args[0]).useLines { lines ->
        lines.forEach { line ->
            when {
                line.contains('*') || line.contains('+') -> {
                    val operatorLine = line.split(Regex("\\s+")).map { it[0] }
                    operators.add(operatorLine)
                }
                line.isNotEmpty() -> {
                    val numbers = line.split(Regex("\\s+")).map { it.toLong() }
                    operandsNonRotated.add(numbers)
                }
            }
        }
    }

    println(Day06.part1(operandsNonRotated, operators[0]))
    */
}