package day09

import java.io.File
import kotlin.math.abs
import kotlin.math.max

/**
 * Advent of Code 2025 - Day 09
 */
object Day09 {

    // Calculate cross products of two vectors between three points
    fun crossProduct(points: List<Pair<Long, Long>>): Long {
        val p0 = points[0]
        val p1 = points[1]
        val p2 = points[2]

        val v1x = p1.first - p0.first
        val v1y = p1.second - p0.second
        val v2x = p2.first - p0.first
        val v2y = p2.second - p0.second

        return v1x * v2y - v1y * v2x
    }

    // Helper function to check if point (px, py) lies on segment p1-p2
    fun isPointOnSegment(px: Long, py: Long, p1: Pair<Long, Long>, p2: Pair<Long, Long>): Boolean {
        val x1 = p1.first
        val y1 = p1.second
        val x2 = p2.first
        val y2 = p2.second

        // 1. Check Bounding Box (Is it within the range of the segment?)
        // If outside the rectangle formed by p1 and p2, it can't be on the segment.
        if (px < minOf(x1, x2) || px > maxOf(x1, x2) ||
            py < minOf(y1, y2) || py > maxOf(y1, y2)) {
            return false
        }

        val crossProduct = (px - x1) * (y2 - y1) - (py - y1) * (x2 - x1)

        return crossProduct == 0L
    }

    fun strictIntersection(p1: Pair<Long, Long>, p2: Pair<Long, Long>, p3: Pair<Long, Long>, p4: Pair<Long, Long>): Boolean {
        val d1 = crossProduct(listOf(p1, p2, p3))
        val d2 = crossProduct(listOf(p1, p2, p4))
        val d3 = crossProduct(listOf(p3, p4, p1))
        val d4 = crossProduct(listOf(p3, p4, p2))

        return ((d1 > 0 && d2 < 0) or (d1 < 0 && d2 > 0)) &&
                ((d3 > 0 && d4 < 0) or (d3 < 0 && d4 > 0))
    }

    fun pointInPolygon(point: Pair<Long, Long>, polygon: List<Pair<Long, Long>>): Boolean {
        var inside = false
        val n = polygon.size
        val x = point.first
        val y = point.second

        for (i in 0 until n) {
            val p1 = polygon[i]
            val p2 = polygon[(i + 1) % n]

            if (isPointOnSegment(x, y, p1, p2)) {
                return true
            }

            // 1. Check Y-straddle:
            // One point must be strictly above, the other below or equal.
            // This implicitly handles vertices by counting them only for the "upper" edge.
            val straddlesY = (p1.second > y) != (p2.second > y)

            if (straddlesY) {
                // 2. Check X-position:
                // Calculate the X-coordinate where the edge crosses the ray's Y-line.
                // We only care if this crossing is to the RIGHT of our point.

                // Formula derived from line equation:
                val intersectX = p1.first + (y - p1.second) * (p2.first - p1.first).toDouble() / (p2.second - p1.second)

                if (x < intersectX) {
                    inside = !inside
                }
            }
        }
        return inside
    }

    fun isConvexPolygon(points: List<Pair<Long, Long>>): Boolean {
        if (points.size < 4) return true // Triangles are always convex

        var sign = 0L
        val n = points.size

        for (i in 0 until n) {
            val cp = crossProduct(
                listOf(
                    points[i],
                    points[(i + 1) % n],
                    points[(i + 2) % n]
                )
            )
            if (cp != 0L) {
                if (sign == 0L) {
                    sign = if (cp > 0) 1 else -1
                } else if ((cp > 0 && sign < 0) || (cp < 0 && sign > 0)) {
                    return false
                }
            }
        }
        return true
    }

    fun intersects(square: List<Pair<Long, Long>>, polygon: List<Pair<Long, Long>>): Boolean {
        // Check if any edge of the square intersects with any edge of the polygon
        for (i in square.indices) {
            val p1 = square[i]
            val p2 = square[(i + 1) % square.size]
            for (j in polygon.indices) {
                val p3 = polygon[j]
                val p4 = polygon[(j + 1) % polygon.size]
                if (strictIntersection(p1, p2, p3, p4)) {
                    return true
                }
            }
        }
        return false
    }
    /**
     * Solve Part 1 of Day 09.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 1.
     */
    fun part1(input: List<Pair<Long, Long>>): Any {

        var max = 0L
        for (i in 0..<input.size) {
            for(j in i + 1..<input.size) {
                val p1 = input[i]
                val p2 = input[j]
                val size = (abs(p1.first - p2.first) + 1) * (abs(p1.second - p2.second) + 1)
                max = max(max, size)

            }
        }
        return max
    }

    /**
     * Solve Part 2 of Day 09.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 2.
     */
    fun part2(input: List<Pair<Long, Long>>): Any {
        /**
         *  I have a Concave Polygon with N vertices.
         *  I need the find the largest area of a rectangle that can fit inside the polygon.
         *  Where two of the opposite corners of the rectangle are vertices of the polygon.
         */

        var max = 0L

        for (i in 0..<input.size) {
            for(j in i + 1..<input.size) {
                val p1 = input[i]
                val p2 = input[j]
                val size = (abs(p1.first - p2.first) + 1) * (abs(p1.second - p2.second) + 1)

                if (size > max) {
                    val rect = listOf(
                        Pair(minOf(p1.first, p2.first), minOf(p1.second, p2.second)),
                        Pair(maxOf(p1.first, p2.first), minOf(p1.second, p2.second)),
                        Pair(maxOf(p1.first, p2.first), maxOf(p1.second, p2.second)),
                        Pair(minOf(p1.first, p2.first), maxOf(p1.second, p2.second))
                    )

                    if (intersects(rect, input)) {
                        continue
                    }

                    if (rect.any { !pointInPolygon(it, input) }) {
                        continue
                    }
                    max = size
                }
            }
        }
        return max
    }
}

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Usage: ./gradlew day09 -Pargs=<inputfile>")
        return
    }

    val input = File(args[0]).useLines { lines ->
        lines.map { line ->
            val numbers = line.split(',').map { it.trim().toLong() }
            if (numbers.size != 2) {
                throw IllegalArgumentException("Each line must contain exactly two numbers separated by a comma.")
            }
            Pair(numbers[0], numbers[1])
        }.toList()
    }


    println(Day09.part2(input))
}