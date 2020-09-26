package com.jrciii.logprocessing

import org.apache.spark.sql.{Dataset, SparkSession}
import org.apache.spark.sql.catalyst.ScalaReflection
import org.apache.spark.sql.functions.{array_sort, collect_list, struct}
import org.apache.spark.sql.streaming.StreamingQuery
import org.apache.spark.sql.types.StructType

object ByteRangeStreamManager {
  def createSparkSession: SparkSession = {
    val spark = SparkSession.builder().appName("Byte Range Delivery Merger").master("local[2]").getOrCreate()
    spark.sparkContext.setLogLevel("ERROR")
    spark
  }

  def createDataset(csvPath: String): Dataset[MergedByteRanges] = {
    val spark = SparkSession.getActiveSession.getOrElse(createSparkSession)
    import spark.implicits._
    spark
      .readStream
      .option("sep", "\t")
      .option("mode", "DROPMALFORMED")
      .schema(ScalaReflection.schemaFor[ByteRangeRequest].dataType.asInstanceOf[StructType])
      .csv(csvPath)
      .filter("status IN (200, 206)")
      .selectExpr("ipAddress", "userAgent", "request", "cast(split(byteRange, '-')[0] as int) as startByte",
        "cast(split(byteRange, '-')[1] as int) as endByte")
      .groupBy("ipAddress", "userAgent", "request")
      .agg(array_sort(collect_list(struct("startByte", "endByte")).as("ranges")))
      .as[(String, String, String, Seq[(Int, Int)])]
      .flatMap(ByteRangeMerger.mergeByteRanges)
  }

  def startStream(dataset: Dataset[MergedByteRanges]): StreamingQuery = {
    dataset
      .writeStream
      .outputMode("complete")
      .format("memory")
      .queryName("delivered")
      .start()
  }
}
