package org.skyeidos.power.toys;

import lombok.extern.slf4j.Slf4j;
import org.skyeidos.power.toys.entity.Product;
import org.skyeidos.power.toys.indexer.Indexer;
import org.skyeidos.power.toys.indexer.impl.ProductIndexer;

import java.io.IOException;

@Slf4j
public class PowerToysApplication {

  public static void main(String[] args) throws IOException {
    Indexer<Product> indexer = new ProductIndexer();
    indexer.init();
    indexer.index(Product.builder().name("test").description("test").build());
    indexer.merge();
  }

}
