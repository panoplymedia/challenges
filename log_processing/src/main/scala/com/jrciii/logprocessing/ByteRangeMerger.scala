package com.jrciii.logprocessing

/**
  * This object contains the function to merge byte ranges.
  */
object ByteRangeMerger {
  /**
    * Merges delivered byte ranges. Assumes incoming ranges are sorted by start byte and end byte, ascending.
    * @param ranges The byte ranges for each ipAddress, userAgent and request to merge.
    * @return Merged ranges. If there is more than one range, gaps have occurred.
    */
  def mergeByteRanges(ranges: MergedByteRanges): Option[MergedByteRanges] = {
    ranges.byteRanges.headOption.map(first => {
      val (accumulatedMerges, finalRange) =
        ranges.byteRanges.tail.foldLeft((Vector[ByteRange](), first))((accumulated, newRange) => {
          val previouslyMerged = accumulated._1
          val currentRange = accumulated._2
          // If the new range starts before or on the previous range's end, merge the range
          if (newRange.startByte <= currentRange.endByte)
            (previouslyMerged, ByteRange(currentRange.startByte, Math.max(currentRange.endByte, newRange.endByte)))
          // Otherwise add that range to the list of merged ranges and start a new round of merging
          // with the latest range
          else
            (previouslyMerged :+ currentRange, newRange)
        })
      ranges.copy(byteRanges = accumulatedMerges :+ finalRange)
    })
  }
}
