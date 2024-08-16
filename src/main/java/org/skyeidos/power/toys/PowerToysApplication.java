package org.skyeidos.power.toys;

import com.apifan.common.random.RandomSource;
import com.apifan.common.random.entity.Poem;
import lombok.extern.slf4j.Slf4j;
import org.apache.lucene.index.Term;
import org.apache.lucene.search.*;
import org.skyeidos.power.toys.entity.Product;
import org.skyeidos.power.toys.indexer.Indexer;
import org.skyeidos.power.toys.indexer.impl.ProductIndexer;
import org.skyeidos.power.toys.searcher.Searcher;
import org.skyeidos.power.toys.searcher.impl.ProductSearcher;
import org.springframework.util.StopWatch;

import java.io.IOException;
import java.util.List;

@Slf4j
public class PowerToysApplication {

  public static void main(String[] args) throws IOException {
//     index();
    Searcher<Product> searcher = new ProductSearcher();
    searcher.init();
    Query query =new BooleanQuery.Builder()
        .add(new WildcardQuery(new Term("name","vivo*")), BooleanClause.Occur.MUST)
        .add(new WildcardQuery(new Term("url","*top")), BooleanClause.Occur.MUST)
        .build();
    StopWatch stopWatch = new StopWatch();
    stopWatch.start("query");
    List<Product> products = searcher.search(query, 10);
    stopWatch.stop();
    System.out.println(products);
    System.out.println(stopWatch.prettyPrint());
  }

  private static void index() throws IOException {
    Indexer<Product> indexer = new ProductIndexer();
    indexer.init();
    for (int i = 0; i < 10; i++) {
      for (int j = 0; j < 1000; j++) {
        String mobileModel = RandomSource.otherSource()
            .randomMobileModel();
        Poem poem = RandomSource.languageSource().randomTangPoem();
        Product product = Product.builder()
            .name(mobileModel).description(poem.getContent()[0])
            .url(RandomSource.internetSource().randomDomain(16))
            .build();
        indexer.index(product);
      }
      indexer.flush();
    }
    indexer.merge(3);
    indexer.commit();
  }

}
