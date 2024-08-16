package org.skyeidos.power.toys.searcher.impl;

import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import org.apache.lucene.document.Document;
import org.apache.lucene.index.StoredFields;
import org.apache.lucene.search.Query;
import org.apache.lucene.search.ScoreDoc;
import org.apache.lucene.search.TopDocs;
import org.skyeidos.power.toys.entity.Product;
import org.skyeidos.power.toys.searcher.AbstractSearcher;

import java.util.ArrayList;
import java.util.List;

@Slf4j
public class ProductSearcher extends AbstractSearcher<Product> {

  @Override
  public String getFilePath() {
    return "products";
  }

  @SneakyThrows
  public List<Product> search(Query query, int nums) {
    log.debug("query:{}", query);
    TopDocs topDocs = searcher.search(query, nums);
    StoredFields storedFields = searcher.storedFields();
    List<Product> products = new ArrayList<>();
    for (ScoreDoc hit : topDocs.scoreDocs) {
      Document doc = storedFields.document(hit.doc);
      Product product = Product.builder().name(doc.get("name"))
          .description(doc.get("description"))
          .url(doc.get("url"))
          .build();
      products.add(product);
    }
    return products;
  }
}
