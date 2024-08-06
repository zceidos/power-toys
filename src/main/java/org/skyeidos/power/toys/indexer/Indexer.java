package org.skyeidos.power.toys.indexer;

import java.io.IOException;

public interface Indexer<T> {

  void init() throws IOException;

  void index(T entity);

  void merge();

  void flush() throws IOException;

}
