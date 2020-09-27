package com.jrciii.logprocessing

/**
  * Represents a merged set of byte ranges
  * @param ipAddress
  * @param userAgent
  * @param request
  * @param byteRanges If only one, its complete (if the end of the range is the file size.
  *                   Incomplete if more than one.
  */
case class MergedByteRanges(ipAddress: String,
                            userAgent: String,
                            request: String,
                            byteRanges: Seq[ByteRange])