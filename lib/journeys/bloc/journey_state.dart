import 'package:equatable/equatable.dart';
import 'package:finall/journeys/models/journey.dart';

class JourneyState extends Equatable {
  const JourneyState();

  @override
  List<Object> get props => [];
}

class JourneyLoading extends JourneyState {}

class JourneysLoadSuccess extends JourneyState {
  final List<Journey> journeys;

  JourneysLoadSuccess([this.journeys = const []]);

  @override
  List<Object> get props => [journeys];
}

class JourneyOperationFailure extends JourneyState {}
