import 'package:equatable/equatable.dart';
import 'package:flutter/material.dart';

@immutable
class Journey extends Equatable {
  Journey(
      {this.id,
      @required this.source,
      @required this.price,
      @required this.destination,
      @required this.distance});

  final int id;
  final String source;
  final int price;
  final String destination;
  final String distance;

  @override
  List<Object> get props => [id, source, price, destination, distance];

  factory Journey.fromJson(Map<String, dynamic> json) {
    return Journey(
      id: json['id'],
      source: json['source'],
      price: json['price'],
      destination: json['destination'],
      distance: json['distance'],
    );
  }

  @override
  String toString() =>
      'Course { id: $id, source: $source, price: $price, destination: $destination, distance: $distance  }';
}
