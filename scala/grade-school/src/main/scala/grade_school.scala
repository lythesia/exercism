class School {
  private val _db = scala.collection.mutable.Map[Int, List[String]]()

  def add(n: String, g: Int) = _db(g) = _db.getOrElseUpdate(g, List()) :+ n

  def grade(g: Int) = _db.getOrElse(g, Nil)

  def db = _db.toMap

  def sorted = scala.collection.immutable.SortedMap(_db.mapValues(_.sorted).toSeq :_*)
}
