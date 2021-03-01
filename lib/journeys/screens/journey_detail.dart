import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:finall/journeys/journeys.dart';

class JourneyDetail extends StatelessWidget {
  static const routeName = 'journeyDetail';
  final Journey journey;

  JourneyDetail({@required this.journey});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('${this.journey.distance}'),
        actions: [
          IconButton(
            icon: Icon(Icons.edit),
            onPressed: () => Navigator.of(context).pushNamed(
              AddUpdateJourney.routeName,
              arguments: JourneyArgument(journey: this.journey, edit: true),
            ),
          ),
          SizedBox(
            width: 32,
          ),
          IconButton(
              icon: Icon(Icons.delete),
              onPressed: () {
                context.read<JourneyBloc>().add(JourneyDelete(this.journey));
                Navigator.of(context).pushNamedAndRemoveUntil(
                    JourneyList.routeName, (route) => false);
              }),
        ],
      ),
      body: Card(
        child: Column(
          children: [
            ListTile(
              title: Text('Source Adress: ${this.journey.source}'),
              subtitle: Text('Price: ${this.journey.price}'),
            ),
            Text(
              'Travel Details',
              style: TextStyle(
                fontSize: 18,
                fontWeight: FontWeight.bold,
              ),
            ),
            SizedBox(
              height: 10,
            ),
            Text(this.journey.destination),
          ],
        ),
      ),
    );
  }
}
