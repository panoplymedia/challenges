package com.jrciii.logprocessing

import org.scalatest.{FlatSpec, Matchers}

class ByteRangeStreamManagerSpec extends FlatSpec with Matchers {
  behavior of "The Byte Range Merger Stream Manager"

  it should "process a directory of CSV files" in {
    val csvFile = this.getClass.getResource("/sample").getFile
    val dataset = ByteRangeStreamManager.createDataset(csvFile)
    val stream = ByteRangeStreamManager.startStream(dataset)
    stream.processAllAvailable()
    assert(stream.sparkSession.sql("select * from delivered where ").count() > 0)
  }
}
