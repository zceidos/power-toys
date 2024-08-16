package org.skyeidos.power.toys.indexer;

import java.io.IOException;

public interface Indexer<T> {

  void init() throws IOException;

  void index(T entity);

  void merge(int maxSegments);

  void flush() throws IOException;

  void commit() throws IOException;

}
