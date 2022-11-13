import 'package:http/http.dart' as http;
import 'globals.dart' as globals;
import 'package:flutter/material.dart';
import 'globals.dart';
import 'actions.dart';
import 'dart:async';
import 'dart:convert';
import 'dart:io';

class CreationAreaForm extends StatefulWidget {
  CreationAreaForm(Parser this.actionService, String this.actionValue,
      Parser this.reactionService, String this.reactionValue);
  Parser actionService;
  String actionValue;
  Parser reactionService;
  String reactionValue;

  @override
  State<CreationAreaForm> createState() => _CreationAreaFormState();
}

class _CreationAreaFormState extends State<CreationAreaForm> {
  @override
  List<Widget> listArgumentsAction = [];
  List<Widget> listArgumentsReaction = [];
  List<TextEditingController> listTextEditingControlAction = [];
  List<TextEditingController> listTextEditingControlReaction = [];

  Widget build(context) {
    listArgumentsAction = getListTextAction(getArgAction());
    listArgumentsReaction = getListTextReaction(getArgReaction());
    return new Scaffold(
      body: Padding(
        padding: const EdgeInsets.all(10),
        child: ListView(
          children: [
            Column(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: listArgumentsAction,
            ),
            Column(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: listArgumentsReaction,
            ),
            MaterialButton(
              minWidth: double.infinity,
              height: 60,
              onPressed: () {
                postArea();
              },
              color: Colors.redAccent,
              shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(40)),
              child: const Text(
                "Create Area",
                style: TextStyle(
                  fontWeight: FontWeight.w600,
                  fontSize: 16,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  String getActionParams() {
    String result = "";

    for (var i in listTextEditingControlAction) {
      if (result.length == 0) {
        result = i.text;
      } else {
        result = result + "@@@" + i.text;
      }
    }
    print(result);
    return result;
  }

  String getReactionsParams() {
    String result = "";

    for (var i in listTextEditingControlReaction) {
      if (result.length == 0) {
        result = i.text;
      } else {
        result = result + "@@@" + i.text;
      }
    }
    print(result);
    return result;
  }

  Future<void> postArea() async {
    print(getArgAction());
    var response =
        await http.post(Uri.parse("http://10.0.2.2:8080/area/create"),
            headers: {'cookie': globals.tokenUser},
            body: json.encode({
              'action_service': widget.actionService.plateforme,
              'action_func': widget.actionValue,
              'action_func_params': getActionParams(),
              'reaction_service': widget.reactionService.plateforme,
              'reaction_func': widget.reactionValue,
              'reaction_func_params': getReactionsParams(),
            }));

    if (response.statusCode == 200) {
      print('Area successfuly created');
    } else {
      print('Error cannot created');
    }

    ///area/create
  }

  List<Widget> getListTextAction(List<String> list) {
    List<Widget> listWidget = [];

    for (var i in list) {
      TextEditingController temp = TextEditingController();
      listWidget.add(makeInput(temp, label: i));
      listTextEditingControlAction.add(temp);
    }
    return listWidget;
  }

  List<Widget> getListTextReaction(List<String> list) {
    List<Widget> listWidget = [];

    for (var i in list) {
      TextEditingController temp = TextEditingController();
      listWidget.add(makeInput(temp, label: i));
      listTextEditingControlReaction.add(temp);
    }
    return listWidget;
  }

  List<String> getArgAction() {
    for (var val in widget.actionService.action) {
      if (val.description == widget.actionValue) {
        final List<String> strs =
            val.field_names.map((e) => e.toString()).toList();
        return strs;
      }
    }
    return [];
  }

  List<String> getArgReaction() {
    for (var val in widget.reactionService.reaction) {
      if (val.description == widget.reactionValue) {
        final List<String> strs =
            val.field_names.map((e) => e.toString()).toList();
        return strs;
      }
    }
    return [];
  }

  Widget makeInput(TextEditingController text, {label, obsureText = false}) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          label,
          style: TextStyle(
              fontSize: 10, fontWeight: FontWeight.w400, color: Colors.black87),
        ),
        SizedBox(
          height: 5,
        ),
        TextField(
          obscureText: obsureText,
          controller: text,
          decoration: InputDecoration(
            contentPadding: EdgeInsets.symmetric(vertical: 0, horizontal: 4),
            enabledBorder: OutlineInputBorder(
              borderSide: BorderSide(
                color: Colors.grey.shade400,
              ),
            ),
            border: OutlineInputBorder(
                borderSide: BorderSide(color: Colors.grey.shade400)),
          ),
        ),
        SizedBox(
          height: 30,
        )
      ],
    );
  }
}
