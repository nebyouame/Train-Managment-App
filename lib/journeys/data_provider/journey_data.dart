import 'dart:convert';
import 'package:meta/meta.dart';
import 'package:finall/journeys/models/journey.dart';
import 'package:http/http.dart' as http;

class JourneyDataProvider {
  final _baseUrl = 'http://10.4.107.185:8181';
  final http.Client httpClient;

  JourneyDataProvider({@required this.httpClient}) : assert(httpClient != null);

  Future<Journey> createJourney(Journey journey) async {
    final response = await httpClient.post(
      Uri.http('10.4.107.185:8181', '/journeys'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'source': journey.source,
        'price': journey.price,
        'destination': journey.destination,
        'distance': journey.distance,
      }),
    );

    if (response.statusCode == 201) {
      return Journey.fromJson(jsonDecode(response.body));
    } else {
      throw Exception('Failed to create course.');
    }
  }

  Future<List<Journey>> getJourney() async {
    final response = await httpClient.get('10.4.107.185:8181/journeys');

    if (response.statusCode == 200) {
      final journeys = jsonDecode(response.body) as List;
      return journeys.map((journey) => Journey.fromJson(journey)).toList();
    } else {
      throw Exception('Failed to load journeys');
    }
  }

  Future<List<Journey>> getJourneys() async {
    final response = await httpClient.get('$_baseUrl/journeys');

    if (response.statusCode == 200) {
      final journeys = jsonDecode(response.body) as List;
      return journeys.map((journey) => Journey.fromJson(journey)).toList();
    } else {
      throw Exception('Failed to load Journey');
    }
  }

  Future<void> deleteJourney(int id) async {
    final http.Response response = await httpClient.delete(
      '$_baseUrl/journeys/$id',
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
    );

    if (response.statusCode != 204) {
      throw Exception('Failed to delete journey.');
    }
  }

  Future<void> updateJourney(Journey journey) async {
    final http.Response response = await httpClient.put(
      '$_baseUrl/journeys/${journey.id}',
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'id': journey.id,
        'source': journey.source,
        'price': journey.price,
        'destination': journey.destination,
        'distance': journey.distance,
      }),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to update journey.');
    }
  }
}
