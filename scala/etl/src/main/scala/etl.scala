object ETL {
  def transform(old: Map[Int, Seq[String]]): Map[String, Int] = {
    old.flatMap { case (key, values) => values.map(_.toLowerCase -> key) }
  }
}
