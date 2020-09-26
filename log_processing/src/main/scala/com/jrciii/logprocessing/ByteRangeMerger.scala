package com.jrciii.logprocessing

object ByteRangeMerger {
  /**
    * Merges delivered byte ranges. Assumes ranges are sorted by start byte and end byte, ascending.
    * @param ranges The byte ranges for each ipAddress, userAgent and request to merge.
    * @return Merged ranges. If there is more than one range, gaps have occurred.
    */
  def mergeByteRanges(ranges: (String, String, String, Seq[(Int, Int)])): Option[MergedByteRanges] = {
    val (ipAddress, userAgent, request, byteRanges) = ranges
    byteRanges.headOption.map(first => {
      val (accumulatedMerges, finalRange) =
        byteRanges.tail.foldLeft((Vector[(Int, Int)](), first))((accumulated, newRange) => {
          val (start, end) = newRange
          val previouslyMerged = accumulated._1
          val currentRange = accumulated._2
          // If the new range starts before or on the previous range's end, merge the range
          if (start <= currentRange._2)
            (previouslyMerged, (currentRange._1, Math.max(currentRange._2, end)))
          // Otherwise add that range to the list of merged ranges and start a new round of merging
          // with the latest range
          else
            (previouslyMerged :+ currentRange, newRange)
        })
      MergedByteRanges(ipAddress, userAgent, request, accumulatedMerges :+ finalRange)
    })
  }
}
