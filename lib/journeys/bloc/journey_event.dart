import 'package:equatable/equatable.dart';
import 'package:finall/journeys/models/journey.dart';

abstract class JourneyEvent extends Equatable {
  const JourneyEvent();
}

class JourneyLoad extends JourneyEvent {
  const JourneyLoad();

  @override
  List<Object> get props => [];
}

class JourneyCreate extends JourneyEvent {
  final Journey journey;

  const JourneyCreate(this.journey);

  @override
  List<Object> get props => [journey];

  @override
  String toString() => 'Journey Created {journey: $journey}';
}

class JourneyUpdate extends JourneyEvent {
  final Journey journey;

  const JourneyUpdate(this.journey);

  @override
  List<Object> get props => [journey];

  @override
  String toString() => 'Journey Updated {journey: $journey}';
}

class JourneyDelete extends JourneyEvent {
  final Journey journey;

  const JourneyDelete(this.journey);

  @override
  List<Object> get props => [journey];

  @override
  toString() => 'Journey Deleted {journey: $journey}';
}
