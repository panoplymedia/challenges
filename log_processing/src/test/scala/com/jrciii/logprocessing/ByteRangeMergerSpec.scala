package com.jrciii.logprocessing

import org.scalatest.{FlatSpec, Matchers}

class ByteRangeMergerSpec extends FlatSpec with Matchers {
  behavior of "Byte Range Merger"

  it should "merge overlapping ranges into one" in {
    val merged =
      ByteRangeMerger
        .mergeByteRanges(MergedByteRanges("127.0.0.1", "Test User Agent", "/request", Seq(ByteRange(0, 200),
          ByteRange(200, 400), ByteRange(100, 101)).sorted))
    merged shouldBe MergedByteRanges("127.0.0.1", "Test User Agent", "/request", Seq(ByteRange(0, 400)))
  }

  it should "merge sets of overlapping ranges into multiple ranges" in {
    val merged =
      ByteRangeMerger
        .mergeByteRanges(MergedByteRanges("127.0.0.1", "Test User Agent", "/request",
          Seq(ByteRange(777, 778), ByteRange(0, 200), ByteRange(201, 400), ByteRange(100, 101)).sorted))
    merged shouldBe MergedByteRanges("127.0.0.1", "Test User Agent", "/request",
      Seq(ByteRange(0, 200), ByteRange(201, 400), ByteRange(777, 778)))
  }
}
