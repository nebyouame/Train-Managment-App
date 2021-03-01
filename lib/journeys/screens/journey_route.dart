import 'package:flutter/material.dart';
import 'package:finall/journeys/journeys.dart';

class JourneyAppRoute {
  static Route generateRoute(RouteSettings settings) {
    if (settings.name == '/') {
      return MaterialPageRoute(builder: (context) => JourneyList());
    }

    if (settings.name == AddUpdateJourney.routeName) {
      JourneyArgument args = settings.arguments;
      return MaterialPageRoute(
          builder: (context) => AddUpdateJourney(
                args: args,
              ));
    }

    if (settings.name == JourneyDetail.routeName) {
      Journey journey = settings.arguments;
      return MaterialPageRoute(
          builder: (context) => JourneyDetail(
                journey: journey,
              ));
    }

    return MaterialPageRoute(builder: (context) => JourneyList());
  }
}

class JourneyArgument {
  final Journey journey;
  final bool edit;
  JourneyArgument({this.journey, this.edit});
}
