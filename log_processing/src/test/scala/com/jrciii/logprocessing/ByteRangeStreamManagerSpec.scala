package com.jrciii.logprocessing

import org.scalatest.{FlatSpec, Matchers}

class ByteRangeStreamManagerSpec extends FlatSpec with Matchers {
  behavior of "The Byte Range Merger Stream Manager"

  private val csvFile = this.getClass.getResource("/sample").getFile
  private val dataset = ByteRangeStreamManager.createDataset(csvFile)
  private val stream = ByteRangeStreamManager.startStream(dataset)
  import dataset.sparkSession.implicits._
  stream.processAllAvailable()

  val expectedCompleteRange = MergedByteRanges(
    ipAddress = "183.3.129.45",
    userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1",
    request="/32668757-95c0-4ded-9b81-71607a644e92",
    byteRanges = Seq(ByteRange(0, 1741)))

  val expectedIncompleteRange = MergedByteRanges(
    ipAddress = "182.122.91.54",
    userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) " +
      "CriOS/83.0.4147.71 Mobile/15E148 Safari/604.1",
    request = "/ea65704e-6829-4871-a92a-5f0c3b80addf",
    byteRanges = Seq(ByteRange(0, 311), ByteRange(522, 1057)))

  private def testSql(expectedRange: MergedByteRanges) =
    s"select * from delivered where ipAddress='${expectedRange.ipAddress}' " +
      s"AND userAgent='${expectedRange.userAgent}' AND request='${expectedRange.request}'"

  it should "merge a complete set of overlapping requests into one" in {
    val actualCompletedRange: Array[MergedByteRanges] =
      stream
        .sparkSession
        .sql(testSql(expectedCompleteRange))
        .as[MergedByteRanges]
        .collect()

    actualCompletedRange shouldBe Array(expectedCompleteRange)
  }

  it should "merge an incomplete set of overlapping requests into multiple" in {
    val actualIncompleteRange: Array[MergedByteRanges] =
      stream
        .sparkSession
        .sql(testSql(expectedIncompleteRange))
        .as[MergedByteRanges]
        .collect()

    actualIncompleteRange shouldBe Array(expectedIncompleteRange)
  }
}
