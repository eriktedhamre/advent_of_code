package day08

import java.io.File
import java.util.PriorityQueue
import kotlin.math.abs
import kotlin.math.max
import kotlin.math.min
import kotlin.math.pow
import kotlin.math.sqrt

/**
 * Advent of Code 2025 - Day 08
 */
object Day08 {

    data class Point(val x: Double, val y: Double, val z: Double)

    data class Edge(val from: Point, val to: Point)


    /**
     * Solve Part 1 of Day 08.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 1.
     */
    fun part1(input: List<Point>): Any {
        val edgeList = mutableListOf<Pair<Long, Edge>>()
        val clusterMap = mutableMapOf<Point, Int>()
        val freqTable = mutableMapOf<Int, Int>()
        var clusterId = 0
        val newConnectionsLimit = 1000L
        var newConnections = 0

        for (i in 0..<input.size) {
            for(j in i + 1..<input.size) {
                val p1 = input[i]
                val p2 = input[j]
                val dist = sqrt((p1.x - p2.x).pow(2)  + (p1.y - p2.y).pow(2) + (p1.z - p2.z).pow(2)).toLong()
                edgeList.add(Pair(dist, Edge(p1, p2)))
            }
        }

        edgeList.sortBy { it.first }

        for ( i in 0 until edgeList.size ) {
            val edge = edgeList[i]
            val fromCluster = clusterMap[edge.second.from]
            val toCluster = clusterMap[edge.second.to]
            newConnections++
            if (fromCluster == null && toCluster == null) {
                clusterMap[edge.second.from] = clusterId
                clusterMap[edge.second.to] = clusterId
                clusterId++

            } else if (fromCluster != null && toCluster == null) {
                clusterMap[edge.second.to] = fromCluster

            } else if (fromCluster == null && toCluster != null) {
                clusterMap[edge.second.from] = toCluster

            } else if (fromCluster != toCluster) {
                val oldCluster = toCluster
                val newCluster = fromCluster!!

                for ((point, cluster) in clusterMap) {
                    if (cluster == oldCluster) {
                        clusterMap[point] = newCluster
                    }
                }
            }

            if (newConnections >= newConnectionsLimit) break
        }

        clusterMap.forEach { (_, cluster) ->
            freqTable[cluster] = freqTable.getOrDefault(cluster, 0) + 1
        }
        return freqTable.values.sortedByDescending { it }.subList(0, 3).reduce { acc, i -> acc * i }
    }

    /**
     * Solve Part 2 of Day 08.
     *
     * @param input Raw puzzle input as a single string.
     * @return The answer for Part 2.
     */
    fun part2(input: List<Point>): Long {
        val edgeList = mutableListOf<Pair<Long, Edge>>()

        var sets = mutableListOf<Set<Point>>()
        var pointToSetIndex = mutableMapOf<Point, Int>()
        for (i in 0..<input.size) {
            for(j in i + 1..<input.size) {
                val p1 = input[i]
                val p2 = input[j]
                val dist = sqrt((p1.x - p2.x).pow(2)  + (p1.y - p2.y).pow(2) + (p1.z - p2.z).pow(2)).toLong()
                edgeList.add(Pair(dist, Edge(p1, p2)))
            }
        }

        edgeList.sortBy { it.first }
        var edge  = Edge(Point(0.0,0.0,0.0), Point(0.0,0.0,0.0))

        for ( i in 0 until edgeList.size ) {
            edge = edgeList[i].second
            val fromSetIndex = pointToSetIndex[edge.from]
            val toSetIndex = pointToSetIndex[edge.to]


            if (fromSetIndex == null && toSetIndex == null) {
                val newSet = mutableSetOf<Point>()
                newSet.add(edge.from)
                newSet.add(edge.to)
                sets.add(newSet)
                val newIndex = sets.size - 1
                pointToSetIndex[edge.from] = newIndex
                pointToSetIndex[edge.to] = newIndex

            } else if (fromSetIndex != null && toSetIndex == null) {
                sets[fromSetIndex] = sets[fromSetIndex].plus(edge.to)
                pointToSetIndex[edge.to] = fromSetIndex

            } else if (fromSetIndex == null && toSetIndex != null) {
                sets[toSetIndex] = sets[toSetIndex].plus(edge.from)
                pointToSetIndex[edge.from] = toSetIndex


            } else if (fromSetIndex != toSetIndex) {
                val oldSetIndex = toSetIndex!!
                val newSetIndex = fromSetIndex!!
                sets[newSetIndex] = sets[newSetIndex].plus(sets[oldSetIndex])
                for (point in sets[oldSetIndex]) {
                    pointToSetIndex[point] = newSetIndex
                }
            }

            if (sets[pointToSetIndex[edge.from]!!].size == input.size) {
                break
            }
        }
        return (edge.from.x * edge.to.x).toLong()
    }
}

fun main(args: Array<String>) {
    if (args.isEmpty()) {
        println("Usage: ./gradlew day08 -Pargs=<inputfile>")
        return
    }

    val points = mutableListOf<Day08.Point>()
    File(args[0]).useLines { lines ->
        lines.forEach { line ->
            val numbers = line.split(',').map { it.trim().toDouble() }
            points.add(Day08.Point(numbers[0], numbers[1], numbers[2]))
        }
    }
    println(Day08.part2(points))
}