package org.skyeidos.power.toys.indexer;

import org.apache.lucene.index.IndexWriter;
import org.apache.lucene.index.IndexWriterConfig;
import org.apache.lucene.store.Directory;
import org.apache.lucene.store.MMapDirectory;

import java.io.File;
import java.io.IOException;

public abstract class AbstractIndexer<T> implements Indexer<T> {

  protected IndexWriter writer;


  public void init() throws IOException {
    File file = new File("./products");
    if (!(file.exists() || file.mkdirs())) {
      throw new IllegalStateException("dictionary not exists and create failed");
    }
    Directory directory = new MMapDirectory(file.toPath());
    writer = new IndexWriter(directory, new IndexWriterConfig());
  }

  public void flush() throws IOException {
    writer.flush();
  }
}
