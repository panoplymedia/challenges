package com.jrciii.logprocessing

import com.google.common.collect.ComparisonChain

case class ByteRange(startByte: Int, endByte: Int) extends Ordered[ByteRange] {
  override def compare(that: ByteRange): Int =
    ComparisonChain.start().compare(startByte, that.startByte).compare(endByte, that.endByte).result()
}
