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

// class MyDropDownMenu extends StatefulWidget {
//   MyDropDownMenu(this._content, this._element);

//   List<String> _content;
//   String _element;
//   List<String> _checked = [];

//   List<String> get content {
//     return _content;
//   }

//   void set content(List<String> newContent) {
//     _content = newContent;
//   }

//   String get element {
//     return _element;
//   }

//   void set element(String newElement) {
//     _element = newElement;
//   }

//   // DropdownButton get button {
//   //   return _button!;
//   // }

//   @override
//   Widget build(context) {
//     print(_content);
//     print(_element);
//     return CheckboxGroup(
//       labels: content,
//       disabled: [],
//       checked: _checked,
//       onChange: (bool isChecked, String label, int index) =>
//           print("isChecked: $isChecked   label: $label  index: $index"),
//       onSelected: (List<String> checked) =>
//           print("checked: ${checked.toString()}"),
//     );
//   }
// }

class CreationAreaMenu extends StatefulWidget {
  const CreationAreaMenu();

  @override
  State<CreationAreaMenu> createState() => _CreationAreaMenuState();
}

class _CreationAreaMenuState extends State<CreationAreaMenu> {
  String actionPlateforme = "weather";
  String reactionPlateforme = "weather";
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
  void initState() {}

  @override
  Widget build(context) {
    if (listElement.isEmpty) {
      fetchAreaction();
    }
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
        ));
  }

  void callListAction(String label) {
    setState(() {
      print("set state");
      actionPlateforme = label;
      plateformesActions = setPlateformesActions();
    });
  }

  void callListReaction(String label) {
    setState(() {
      print("call list reaction");
      reactionPlateforme = label;
      plateformesreactions = setPlateformesReactions();
    });
  }

  CheckboxGroup setPlateformesListAction() {
    return CheckboxGroup(
      labels: getPlateforme(),
      disabled: [],
      checked: [],
      onChange: (bool isChecked, String label, int index) {
        print("isChecked: $isChecked   label: $label  index: $index");
        callListAction(label);
      },
      onSelected: (List<String> checked) =>
          print("checked: ${checked.toString()}"),
    );
  }

  CheckboxGroup setPlateformesListReaction() {
    return CheckboxGroup(
      labels: getPlateforme(),
      disabled: [],
      checked: [],
      onChange: (bool isChecked, String label, int index) {
        print("isChecked: $isChecked   label: $label  index: $index");
        callListReaction(label);
      },
      onSelected: (List<String> checked) =>
          print("checked: ${checked.toString()}"),
    );
  }

  CheckboxGroup setPlateformesActions() {
    return CheckboxGroup(
      labels: getActionWithPlateforme(actionPlateforme),
      disabled: [],
      checked: [],
      onChange: (bool isChecked, String label, int index) =>
          print("isChecked: $isChecked   label: $label  index: $index"),
      onSelected: (List<String> checked) =>
          print("checked: ${checked.toString()}"),
    );
  }

  CheckboxGroup setPlateformesReactions() {
    print(getReactionWithPlateforme(reactionPlateforme));
    return CheckboxGroup(
      labels: getReactionWithPlateforme(reactionPlateforme),
      disabled: [],
      checked: [],
      onChange: (bool isChecked, String label, int index) =>
          print("isChecked: $isChecked   label: $label  index: $index"),
      onSelected: (List<String> checked) =>
          print("checked: ${checked.toString()}"),
    );
  }

  // void updateContent(CheckboxGroup box, List<String> list) {
  //   setState(() {
  //     box. = list;
  //   });
  // }

  Future<void> postArea() async {
    var response =
        await http.post(Uri.parse("http://10.0.2.2:8080/area/create"),
            headers: {},
            body: json.encode({
              'action_service': '',
              'action_func': '',
              'action_func_params': '',
              'reaction_service': '',
              'reaction_func': '',
              'reaction_func_params': '',
            }));

    if (response.statusCode == 200) {
      Fluttertoast.showToast(msg: 'Area successfuly created');
    } else {
      Fluttertoast.showToast(msg: 'Error cannot created');
    }

    ///area/create
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

  List<String> getPlateforme() {
    List<String> listPlateforme = [];
    listElement.forEach((element) {
      listPlateforme.add(element.plateforme);
    });
    print(listPlateforme.length);
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
        print("oui");
        element.reaction.forEach((value) {
          print(value.description);
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
