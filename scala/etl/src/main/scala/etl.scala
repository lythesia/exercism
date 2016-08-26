object ETL {
  def transform(old: Map[Int, Seq[String]]): Map[String, Int] = 
    for {
      (key, values) <- old
      value <- values
    } yield (value.toLowerCase, key)
}
