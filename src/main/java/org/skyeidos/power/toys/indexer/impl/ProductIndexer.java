package org.skyeidos.power.toys.indexer.impl;

import lombok.SneakyThrows;
import org.apache.lucene.document.Document;
import org.apache.lucene.document.Field;
import org.apache.lucene.document.StringField;
import org.apache.lucene.document.TextField;
import org.skyeidos.power.toys.entity.Product;
import org.skyeidos.power.toys.indexer.AbstractIndexer;

public class ProductIndexer extends AbstractIndexer<Product> {


  @Override
  @SneakyThrows
  public void index(Product product) {
    Document document = new Document();
    document.add(new StringField("name", product.getName(), Field.Store.YES));
    document.add(new TextField("description", product.getDescription(), Field.Store.YES));
    document.add(new StringField("url",product.getUrl(), Field.Store.YES));
    writer.addDocument(document);
  }

  @Override
  @SneakyThrows
  public void merge(int maxSegments) {
    writer.forceMerge(maxSegments);
  }
}
