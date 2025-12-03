plugins {
    kotlin("jvm") version "2.2.21"
    application
}

application {
    mainClass.set("day01.Day01Kt")
}

repositories {
    mavenCentral()
}

dependencies {
    testImplementation(kotlin("test"))
}

tasks.test {
    useJUnitPlatform()
}

(1..25).forEach { day ->
    val dayStr = day.toString().padStart(2, '0')

    tasks.register<JavaExec>("day$dayStr") {
        group = "aoc"
        description = "Runs Advent of Code Day $dayStr"

        classpath = sourceSets["main"].runtimeClasspath
        mainClass.set("day$dayStr.Day${dayStr}Kt")

        // Forward arguments if provided using -Pargs="..."
        val argsProp = project.findProperty("args") as String?
        if (argsProp != null) {
            args = argsProp.split(" ")
        }
    }

    // ---------- Test Task ----------
    tasks.register<Test>("testDay$dayStr") {
        group = "aoc"
        description = "Run tests for Day $dayStr only"

        // run only tests in package dayXX.*
        useJUnitPlatform()
        filter {
            includeTestsMatching("day$dayStr.*")
        }

        testClassesDirs = sourceSets["test"].output.classesDirs
        classpath = sourceSets["test"].runtimeClasspath
    }
}