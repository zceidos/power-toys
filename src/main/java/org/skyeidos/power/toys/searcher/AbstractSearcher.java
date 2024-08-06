package org.skyeidos.power.toys.searcher;

import org.apache.lucene.index.DirectoryReader;
import org.apache.lucene.index.IndexReader;
import org.apache.lucene.index.IndexWriter;
import org.apache.lucene.index.IndexWriterConfig;
import org.apache.lucene.search.IndexSearcher;
import org.apache.lucene.store.Directory;
import org.apache.lucene.store.MMapDirectory;

import java.io.File;
import java.io.IOException;

public abstract class AbstractSearcher<T> implements Searcher<T> {
  protected IndexSearcher searcher;

  @Override
  public void init() throws IOException {
    File file = new File(getFilePath());
    if (!(file.exists() || file.mkdirs())) {
      throw new IllegalStateException("dictionary not exists and create failed");
    }
    Directory directory = new MMapDirectory(file.toPath());
    IndexReader reader = DirectoryReader.open(directory);
    searcher = new IndexSearcher(reader);
  }

  @Override
  public String getFilePath() {
    return null;
  }
}
