import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'globals.dart' as globals;
import 'dart:async';
import 'package:flutter/foundation.dart';
import 'dart:convert';
import 'dart:io';
import 'actions.dart';
import './card.dart';

class WalletPage extends StatefulWidget {
  WalletPage({Key? key, this.title = ""}) : super(key: key);

  final String title;

  @override
  _WalletPageState createState() => new _WalletPageState();
}

class _WalletPageState extends State<WalletPage> {
  List<Card> _listCard = [];

  @protected
  @mustCallSuper
  void initState() {
    fetchAreaction();
  }

  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      body: Padding(
        padding: const EdgeInsets.all(10),
        child: ListView(
          children: <Widget>[
            Column(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: _listCard,
            ),
            MaterialButton(
              minWidth: double.infinity,
              height: 60,
              onPressed: () {
                setState() {
                  _listCard = getCard();
                }
              },
              color: Colors.redAccent,
              shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(40)),
              child: const Text(
                "Reload",
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

  List<Card> getCard() {
    return _listCard;
  }

  void iterateJson(String jsonStr) {
    var myMap = jsonDecode(jsonStr);

    for (var element in myMap) {
      setState(() {
        _listCard.add(buildCard(
            element['action_service'],
            element['action_func'],
            element['reaction_service'],
            element['reaction_func']));
      });
    }
  }

  Future<void> fetchAreaction() async {
    final response = await http.get(
        Uri.parse("http://10.0.2.2:8080/area/user/areas"),
        headers: <String, String>{
          'cookie': globals.tokenUser
        }).timeout(Duration(seconds: 30));

    if (response.statusCode == 200) {
      iterateJson(response.body);
    } else {}
  }
}
