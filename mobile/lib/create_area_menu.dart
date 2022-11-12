import 'package:flutter/material.dart';
import 'package:workspace/menu_list_auth_plateform.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:http/http.dart' as http;
import 'globals.dart' as globals;
import './drop_down_menu.dart';
import './create_area_menu.dart';
import 'dart:async';
import 'package:flutter/foundation.dart';
import 'dart:convert';
import 'dart:io';
import 'actions.dart';

class SelectableImage extends StatelessWidget {
  const SelectableImage({
    Key? key,
    required this.imageAsset,
  }) : super(key: key);
  final List<String> imageAsset;
  @override
  Widget build(BuildContext context) {
    return new Scaffold(
        body: Padding(
      padding: const EdgeInsets.all(10),
      child: ListView(
        children: printTextList(imageAsset),
      ),
    ));
  }
}

class CreationAreaMenu extends StatefulWidget {
  CreationAreaMenu({Key? key, this.title = ""}) : super(key: key);

  final String title;

  @override
  _CreationAreaMenuState createState() => new _CreationAreaMenuState();
}

class _CreationAreaMenuState extends State<CreationAreaMenu> {
  final _formKey = GlobalKey<FormState>();
  String selectedValue = "Discord";
  String selectedAction = "Appeler Erwan";
  int currentStep = 0;
  List<Parser> listElement = [];


  Widget build(BuildContext context) {
    return new Scaffold(
        body: Padding(
            padding: const EdgeInsets.all(10),
            child: ListView(
              children: <Widget>[
                MaterialButton(
                  minWidth: double.infinity,
                  height: 60,
                  onPressed: () {
                    fetchAreaction();
                  },
                  color: Colors.redAccent,
                  shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(40)),
                  child: const Text(
                    "Add an action !",
                    style: TextStyle(
                      fontWeight: FontWeight.w600,
                      fontSize: 16,
                    ),
                  ),
                ),
                Text(
                  '+',
                  style: TextStyle(fontWeight: FontWeight.bold, fontSize: 15),
                  textAlign: TextAlign.center,
                ),
                MaterialButton(
                  minWidth: double.infinity,
                  height: 60,
                  onPressed: () {},
                  color: Colors.blueAccent,
                  shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(40)),
                  child: const Text(
                    "Add an reaction !",
                    style: TextStyle(
                      fontWeight: FontWeight.w600,
                      fontSize: 16,
                    ),
                  ),
                ),
              ],
            )));
  }

  void iterateJson(String jsonStr) {
    var myMap = jsonDecode(jsonStr);

    for (var element in myMap) {
      List<Reac> listAction = [];
      List<Reac> listReaction = [];

      for (var word in element['actions']) {
        listAction.add(Reac.fromJson(word));
      }
      for (var word in element['reactions']) {
        listReaction.add(Reac.fromJson(word));
      }
      listElement.add(Parser(action: listAction, reaction: listReaction, plateforme: element['name']));
    }
  }

  Future<void> fetchAreaction() async {
    print(globals.tokenUser);
    final response = await http.get(
        Uri.parse("http://10.0.2.2:8080/area/user/propositions"),
        headers: <String, String>{
          'cookie': globals.tokenUser
        }).timeout(Duration(seconds: 30));

    print('${response.statusCode}');
    if (response.statusCode == 200) {
      // Map<String, dynamic> myMap = json.decode(response.body);
      iterateJson(response.body);
      // print(responseData);
    } else {
      Fluttertoast.showToast(msg: 'Error: Failed request');
    }
  }
}



List<Widget> printTextList(List<String> map) {
  List<Widget> list = [];

  map.forEach((element) {
    list.add(
      GestureDetector(
        onTap: () {},
        child: Image.asset(
          element,
          fit: BoxFit.cover,
        ),
      ),
    );
  });
  return (list);
}
