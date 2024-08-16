package org.skyeidos.power.toys.bench;

import org.apache.lucene.index.Term;
import org.apache.lucene.search.MatchAllDocsQuery;
import org.apache.lucene.search.Query;
import org.apache.lucene.search.TermQuery;
import org.openjdk.jmh.annotations.*;
import org.openjdk.jmh.results.format.ResultFormatType;
import org.openjdk.jmh.runner.Runner;
import org.openjdk.jmh.runner.RunnerException;
import org.openjdk.jmh.runner.options.Options;
import org.openjdk.jmh.runner.options.OptionsBuilder;
import org.skyeidos.power.toys.entity.Product;
import org.skyeidos.power.toys.searcher.Searcher;
import org.skyeidos.power.toys.searcher.impl.ProductSearcher;

import java.io.IOException;
import java.util.List;
import java.util.concurrent.TimeUnit;

@BenchmarkMode({Mode.AverageTime, Mode.Throughput})
@Warmup(iterations = 1)
@Measurement(iterations = 2, time = 1)
@OutputTimeUnit(TimeUnit.MICROSECONDS)
@Fork(value = 2)
@Threads(8)
@State(Scope.Benchmark)
@OperationsPerInvocation(1)
public class SearcherBeach {

  @Param({"10", "100","1000"})
  private int n;

  Searcher<Product> searcher;
  @Setup
  public void setup() throws IOException {
    searcher = new ProductSearcher();
    searcher.init();
    searcher.search(new MatchAllDocsQuery(),10);
  }

  @TearDown
  public void tearDown() {

  }


  @Benchmark
  public void simpleSearch(){
    Query query = new TermQuery(new Term("name","vivo X27 Pro"));
    List<Product> products = searcher.search(query, 10);
  }


  public static void main(String[] args) throws RunnerException {
    Options opt = new OptionsBuilder()
        .include(SearcherBeach.class.getSimpleName())
        .resultFormat(ResultFormatType.JSON)
        .result("result.json")
        .build();

    new Runner(opt).run();
  }
}
