package com.jrciii.logprocessing

import com.jrciii.logprocessing.ByteRangeStreamManager._

import scala.io.StdIn

/**
  * This object represents the entry point to the program. It starts the demo stream from the CSV
  * and allows the user to query the resulting aggregations, which are stored in the 'delivered' table in memory
  *
  * Example queries:
  *   To get incomplete range requests:
  *   > select * from delivered where size(ranges) > 1
  *
  *   To see if a request was completely delivered, select for an ipAddress, userAgent, and request and check if
  *   byteRanges is a single range:
  *
  *   > select * from delivered where ipAddress='183.3.129.45' AND userAgent='Mozilla/5.0 (iPhone; CPU iPhone OS 13_6
  *   like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1' AND
  *   request='/32668757-95c0-4ded-9b81-71607a644e92' AND byteRanges=array(named_struct('_1', 0, '_2', 1741))
  */
object ByteRangeMergerApp extends App {
  // Start the stream, reading from the CSV path provided by the command line arguments
  val stream = startStream(createDataset(args(0)))

  // Wait for the CSV processing to finish
  stream.processAllAvailable()

  // Execute user supplied Spark SQL on the stream until the user enters 'quit'
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
