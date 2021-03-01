import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:finall/journeys/journeys.dart';

class AddUpdateJourney extends StatefulWidget {
  static const routeName = 'journeyAddUpdate';
  final JourneyArgument args;

  AddUpdateJourney({this.args});
  @override
  _AddUpdateJourneyState createState() => _AddUpdateJourneyState();
}

class _AddUpdateJourneyState extends State<AddUpdateJourney> {
  final _formKey = GlobalKey<FormState>();

  final Map<String, dynamic> _journey = {};

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('${widget.args.edit ? "Edit Journey" : "Add New Journey"}'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Form(
          key: _formKey,
          child: Column(
            children: [
              TextFormField(
                  initialValue:
                      widget.args.edit ? widget.args.journey.distance : '',
                  validator: (value) {
                    if (value.isEmpty) {
                      return 'Please enter Journey distance';
                    }
                    return null;
                  },
                  decoration: InputDecoration(labelText: 'Journey distance'),
                  onSaved: (value) {
                    setState(() {
                      this._journey["distance"] = value;
                    });
                  }),
              TextFormField(
                  initialValue:
                      widget.args.edit ? widget.args.journey.source : '',
                  validator: (value) {
                    if (value.isEmpty) {
                      return 'Please enter Journey Source Place';
                    }
                    return null;
                  },
                  decoration: InputDecoration(labelText: 'Journey Source'),
                  onSaved: (value) {
                    this._journey["source"] = value;
                  }),
              TextFormField(
                  initialValue: widget.args.edit
                      ? widget.args.journey.price.toString()
                      : '',
                  validator: (value) {
                    if (value.isEmpty) {
                      return 'Please enter Journey price';
                    }
                    return null;
                  },
                  decoration: InputDecoration(labelText: 'Journey Price'),
                  onSaved: (value) {
                    setState(() {
                      this._journey["price"] = int.parse(value);
                    });
                  }),
              TextFormField(
                  initialValue:
                      widget.args.edit ? widget.args.journey.destination : '',
                  validator: (value) {
                    if (value.isEmpty) {
                      return 'Please enter Journey destination';
                    }
                    return null;
                  },
                  decoration: InputDecoration(labelText: 'Journey destination'),
                  onSaved: (value) {
                    setState(() {
                      this._journey["destination"] = value;
                    });
                  }),
              Padding(
                padding: const EdgeInsets.symmetric(vertical: 16.0),
                child: ElevatedButton.icon(
                  onPressed: () {
                    final form = _formKey.currentState;
                    if (form.validate()) {
                      form.save();
                      final JourneyEvent event = widget.args.edit
                          ? JourneyUpdate(
                              Journey(
                                id: widget.args.journey.id,
                                distance: this._journey["distance"],
                                source: this._journey["source"],
                                price: this._journey["price"],
                                destination: this._journey["destination"],
                              ),
                            )
                          : JourneyCreate(
                              Journey(
                                distance: this._journey["distance"],
                                source: this._journey["source"],
                                price: this._journey["price"],
                                destination: this._journey["destination"],
                              ),
                            );
                      BlocProvider.of<JourneyBloc>(context).add(event);
                      Navigator.of(context).pushNamedAndRemoveUntil(
                          JourneyList.routeName, (route) => false);
                    }
                  },
                  label: Text('SAVE'),
                  icon: Icon(Icons.save),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
