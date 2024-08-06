package org.skyeidos.power.toys.entity;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Product {

  private Long id;
  private String name;
  private String description;
  private Double price;
  private String category;
  private String brand;
  private String color;
  private String size;
  private String material;
  private String style;
  private String pattern;
  private String image;
  private String url;
  private String status;
  private String createdBy;
  private String updatedBy;
  private String createdDate;
  private String updatedDate;

}
