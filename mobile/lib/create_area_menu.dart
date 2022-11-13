import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:workspace/menu_list_auth_plateform.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:grouped_buttons/grouped_buttons.dart';
import 'package:http/http.dart' as http;
import 'globals.dart' as globals;
import './drop_down_menu.dart';
import './create_area_menu.dart';
import 'dart:async';
import 'package:flutter/foundation.dart';
import 'dart:convert';
import 'dart:io';
import 'actions.dart';
import './send_area.dart';

class CreationAreaMenu extends StatefulWidget {
  const CreationAreaMenu();

  @override
  State<CreationAreaMenu> createState() => _CreationAreaMenuState();
}

class _CreationAreaMenuState extends State<CreationAreaMenu> {
  String actionPlateforme = "";
  String reactionPlateforme = "";
  String actionChoosen = "";
  String reactionChoosen = "";
  // String selectedAction = "temperature_over_N_degrees";
  int currentStep = 1;
  List<Parser> listElement = [];
  bool haveResult = false;
  List<String> actions = [];
  CheckboxGroup? plateformes;
  CheckboxGroup? plateformesActions;
  CheckboxGroup? plateformes2;
  CheckboxGroup? plateformesreactions;

  @protected
  @mustCallSuper
  void initState() {
    if (listElement.isEmpty) {
      fetchAreaction();
    }
  }

  @override
  Widget build(context) {
    return Padding(
        padding: const EdgeInsets.all(0),
        child: ListView(
          scrollDirection: Axis.vertical,
          shrinkWrap: true,
          children: <Widget>[
            Container(
              child: plateformes,
              decoration: BoxDecoration(border: Border.all(color: Colors.blue)),
            ),
            Container(
              child: plateformesActions,
              decoration: BoxDecoration(border: Border.all(color: Colors.red)),
            ),
            Container(
              child: plateformes2,
              decoration:
                  BoxDecoration(border: Border.all(color: Colors.green)),
            ),
            Container(
              child: plateformesreactions,
              decoration:
                  BoxDecoration(border: Border.all(color: Colors.orangeAccent)),
            ),
            MaterialButton(
              minWidth: double.infinity,
              height: 60,
              onPressed: () {
                createArgArea();
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
        ));
  }

  Parser? getAction() {
    for (var element in listElement) {
      if (element.plateforme == actionPlateforme) {
        for (var value in element.action) {
          if (value.description == actionChoosen) {
            return element;
          }
        }
      }
    }
    return null;
  }

  Parser? getReaction() {
    for (var element in listElement) {
      if (element.plateforme == reactionPlateforme) {
        for (var value in element.reaction) {
          if (value.description == reactionChoosen) {
            return element;
          }
        }
      }
    }
    return null;
  }

  void createArgArea() {
    Parser tempAction = getAction()!;
    Parser tempReaction = getReaction()!;

    print(tempAction.plateforme);
    print(tempAction.action);
    print(tempReaction.plateforme);
    print(tempReaction.reaction);

    Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => CreationAreaForm(tempAction, actionChoosen, tempReaction, reactionChoosen)),
    );
  }

  void callListActionChanged(String label) {
    setState(() {
      actionPlateforme = label;
      plateformesActions = setPlateformesActions();
    });
  }

  void callListReactionChanged(String label) {
    setState(() {
      reactionPlateforme = label;
      plateformesreactions = setPlateformesReactions();
    });
  }

  CheckboxGroup setPlateformesListAction() {
    return CheckboxGroup(
      labels: getPlateformeAction(),
      disabled: [],
      checked: [],
      onChange: (bool isChecked, String label, int index) {
        callListActionChanged(label);
      },
      onSelected: (List<String> checked) {

      },
    );
  }

  CheckboxGroup setPlateformesListReaction() {
    return CheckboxGroup(
      labels: getPlateformeReaction(),
      disabled: [],
      checked: [],
      onChange: (bool isChecked, String label, int index) {
        callListReactionChanged(label);
      },
      onSelected: (List<String> checked) {

      },
    );
  }

  CheckboxGroup setPlateformesActions() {
    return CheckboxGroup(
      labels: getActionWithPlateforme(actionPlateforme),
      disabled: [],
      checked: [],
      onChange: (bool isChecked, String label, int index) {
        actionChoosen = label;
      },
      onSelected: (List<String> checked) {

      },
    );
  }

  CheckboxGroup setPlateformesReactions() {
    print(getReactionWithPlateforme(reactionPlateforme));
    return CheckboxGroup(
      labels: getReactionWithPlateforme(reactionPlateforme),
      disabled: [],
      checked: [],
      onChange: (bool isChecked, String label, int index) {
        reactionChoosen = label;
      },
      onSelected: (List<String> checked) {

      },
    );
  }

  // void updateContent(CheckboxGroup box, List<String> list) {
  //   setState(() {
  //     box. = list;
  //   });
  // }

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
      listElement.add(Parser(
          action: listAction,
          reaction: listReaction,
          plateforme: element['name']));
    }
  }

  Future<void> fetchAreaction() async {
    final response = await http.get(
        Uri.parse("http://10.0.2.2:8080/area/user/propositions"),
        headers: <String, String>{
          'cookie': globals.tokenUser
        }).timeout(Duration(seconds: 30));

    if (response.statusCode == 200) {
      haveResult = true;
      // Map<String, dynamic> myMap = json.decode(response.body);
      iterateJson(response.body);
      setState(() {
        plateformes = setPlateformesListAction();
        plateformesActions = setPlateformesActions();
        plateformes2 = setPlateformesListReaction();
        plateformesreactions = setPlateformesReactions();
      });
      // print(responseData);
    } else {
      Fluttertoast.showToast(msg: 'Error: Failed request');
    }
  }

  List<String> getPlateformeAction() {
    List<String> listPlateforme = [];
    listElement.forEach((element) {
      if (element.action.length != 0) {
        listPlateforme.add(element.plateforme);
      }
    });
    return listPlateforme;
  }

  List<String> getPlateformeReaction() {
    List<String> listPlateforme = [];
    listElement.forEach((element) {
      if (element.reaction.length != 0) {
        listPlateforme.add(element.plateforme);
      }
    });
    return listPlateforme;
  }

  List<String> getActionWithPlateforme(String str) {
    List<String> listPlateforme = [];
    listPlateforme.clear();

    listElement.forEach((element) {
      if (element.plateforme == str) {
        element.action.forEach((value) {
          listPlateforme.add(value.description.toString());
        });
      }
    });
    return listPlateforme;
  }

  List<String> getReactionWithPlateforme(String str) {
    List<String> listPlateforme = [];
    listPlateforme.clear();

    print(str);
    listElement.forEach((element) {
      if (element.plateforme == str) {
        element.reaction.forEach((value) {
          listPlateforme.add(value.description.toString());
        });
      }
    });
    return listPlateforme;
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
