package com.jrciii.logprocessing

import com.jrciii.logprocessing.ByteRangeStreamManager._

import scala.io.StdIn

object ByteRangeMergerApp extends App {
  val stream = startStream(createDataset(args(0)))

  Iterator
    .continually({
      print("> ")
      StdIn.readLine
    })
    .takeWhile(_ != "quit")
    .foreach(x => try {
      stream.sparkSession.sql(x).show(false)
    } catch {
      case e: Exception =>
        println("Error processing query. Exception:")
        println(e.getMessage)
      case t: Throwable => throw t
    })

  stream.stop()
}
