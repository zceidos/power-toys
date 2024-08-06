package org.skyeidos.power.toys.searcher.impl;

import org.skyeidos.power.toys.entity.Product;
import org.skyeidos.power.toys.searcher.AbstractSearcher;

import javax.management.Query;
import java.util.Collections;
import java.util.List;

public class ProductSearcher extends AbstractSearcher<Product> {


  public List<Product> search(Query query, int nums) {
    return Collections.emptyList();
  }
}
