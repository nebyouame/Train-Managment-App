import 'package:finall/journeys/repository/journey_repository.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:finall/journeys/bloc/bloc.dart';

class JourneyBloc extends Bloc<JourneyEvent, JourneyState> {
  final JourneyRepository journeyRepository;

  JourneyBloc({@required this.journeyRepository})
      : assert(journeyRepository != null),
        super(JourneyLoading());

  @override
  Stream<JourneyState> mapEventToState(JourneyEvent event) async* {
    if (event is JourneyLoad) {
      yield JourneyLoading();
      try {
        final journeys = await journeyRepository.getJourneys();
        yield JourneysLoadSuccess(journeys);
      } catch (e) {
        print(e);
        yield JourneyOperationFailure();
      }
    }

    if (event is JourneyCreate) {
      try {
        await journeyRepository.createJourney(event.journey);
        final journeys = await journeyRepository.getJourneys();
        yield JourneysLoadSuccess(journeys);
      } catch (e) {
        print(e);
        yield JourneyOperationFailure();
      }
    }

    if (event is JourneyUpdate) {
      try {
        await journeyRepository.updateJourney(event.journey);
        final journeys = await journeyRepository.getJourneys();
        yield JourneysLoadSuccess(journeys);
      } catch (e) {
        print(e);
        yield JourneyOperationFailure();
      }
    }

    if (event is JourneyDelete) {
      try {
        await journeyRepository.deleteJourney(event.journey.id);
        final journeys = await journeyRepository.getJourneys();
        yield JourneysLoadSuccess(journeys);
      } catch (e) {
        print(e);
        yield JourneyOperationFailure();
      }
    }
  }
}
