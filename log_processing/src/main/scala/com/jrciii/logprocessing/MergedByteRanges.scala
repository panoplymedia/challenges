package com.jrciii.logprocessing

case class MergedByteRanges(ipAddress: String, userAgent: String, request: String, ranges: Seq[(Int, Int)])