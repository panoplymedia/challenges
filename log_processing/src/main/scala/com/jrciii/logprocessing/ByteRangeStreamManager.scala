package com.jrciii.logprocessing

import org.apache.spark.sql.{Dataset, SparkSession}
import org.apache.spark.sql.catalyst.ScalaReflection
import org.apache.spark.sql.functions.{array_sort, collect_list, struct}
import org.apache.spark.sql.streaming.StreamingQuery
import org.apache.spark.sql.types.StructType

/**
  * This object contains functions to manage the Spark Structured Streaming layer of the application.
  */
object ByteRangeStreamManager {
  /**
    * Gets or creates the Spark Session
    * @return The current Spark Session
    */
  def createSparkSession: SparkSession = {
    val spark = SparkSession.builder().appName("Byte Range Delivery Merger").master("local[2]").getOrCreate()
    spark.sparkContext.setLogLevel("ERROR")
    spark
  }

  /**
    * Creates the Dataset of merged byte ranges per ipAddress, userAgent and request.
    * A single merged byte range of 0 to size in bytes, e.g. Seq((0-1024)) means all bytes were delivered.
    * Multiple ranges in the sequence mean some bytes were not delivered, e.g. Seq((0-200), (400-500))
    * This can easily be changed to use a different input, such as Kafka. CSV is being used for demo purposes.
    * @param csvPath The directory containing the CSV files to stream.
    * @return
    */
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

  /**
    * Defines the sink and starts the stream. The sink can easily be changed to something else, like ElasticSearch.
    * Using the Memory sink for demo purposes.
    * @param dataset
    * @return
    */
  def startStream(dataset: Dataset[MergedByteRanges]): StreamingQuery = {
    dataset
      .writeStream
      .outputMode("complete")
      .format("memory")
      .queryName("delivered")
      .start()
  }
}
