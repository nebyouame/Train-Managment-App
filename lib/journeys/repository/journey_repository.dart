import 'package:finall/journeys/data_provider/journey_data.dart';
import 'package:meta/meta.dart';
import 'package:finall/journeys/models/journey.dart';

class JourneyRepository {
  final JourneyDataProvider dataProvider;

  JourneyRepository({@required this.dataProvider})
      : assert(dataProvider != null);

  Future<Journey> createJourney(Journey journey) async {
    print(journey.source);
    return await dataProvider.createJourney(journey);
  }

  Future<List<Journey>> getJourneys() async {
    return await dataProvider.getJourneys();
  }

  Future<void> updateJourney(Journey journey) async {
    await dataProvider.updateJourney(journey);
  }

  Future<void> deleteJourney(int id) async {
    await dataProvider.deleteJourney(id);
  }
}
