package org.skyeidos.power.toys.searcher;

import org.apache.lucene.search.Query;

import java.io.IOException;
import java.util.List;

public interface Searcher<T> {

  void init() throws IOException;

  String getFilePath();

  List<T> search(Query query, int nums);
}
