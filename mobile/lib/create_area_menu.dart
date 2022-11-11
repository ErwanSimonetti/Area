import 'package:flutter/material.dart';
import 'package:workspace/menu_list_auth_plateform.dart';
import './drop_down_menu.dart';
import './create_area_menu.dart';

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
  List<DropdownMenuItem<String>> dropDownActions = [
    DropdownMenuItem(child: Text("Appeler Erwan"), value: "Appeler Erwan"),
    DropdownMenuItem(child: Text("Ratio Pixelle"), value: "Ratio Pixelle"),
  ];
  final _images = [
    "assets/Deezer.png",
    "assets/Discord.png",
    "assets/Spotify.png",
  ];

  List<DropdownMenuItem<String>> get dropdownItems {
    List<DropdownMenuItem<String>> menuItems = [
      DropdownMenuItem(child: Text("Discord"), value: "Discord"),
      DropdownMenuItem(child: Text("Spotify"), value: "Spotify"),
      DropdownMenuItem(child: Text("Deezer"), value: "Deezer"),
      DropdownMenuItem(child: Text("Meteo"), value: "Meteo"),
    ];
    return menuItems;
  }

  Map<String, List<String>> get listOfActions {
    Map<String, List<String>> menuList = {
      "Discord": ["Appeler Erwan", "Ratio Pixelle", "Jaj ?"],
      "Spotify": [
        "Nouvel album",
        "Lancer Chante"
      ],
      "Deezer": ["Jouter Despacito"],
      "Meteo": ["Il pleut", "Bo", "Neige"],
    };
    return menuList;
  }

  void listDropDown(List<String> list) {
    list.forEach((element) {
      DropdownMenuItem<String> temp =
          DropdownMenuItem(child: Text(element), value: element);
      dropDownActions.add(temp);
    });
  }

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
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                          return AlertDialog(
                            content: Stack(
                              children: <Widget>[
                                Positioned(
                                  right: -40.0,
                                  top: -40.0,
                                  child: InkResponse(
                                    onTap: () {
                                      Navigator.of(context).pop();
                                    },
                                    child: CircleAvatar(
                                      child: Icon(Icons.close),
                                      backgroundColor: Colors.red,
                                    ),
                                  ),
                                ),
                                Form(
                                  key: _formKey,
                                  child: Column(
                                    mainAxisSize: MainAxisSize.min,
                                    children: <Widget>[
                                      Padding(
                                        padding: EdgeInsets.all(8.0),
                                        child: DropdownButton(
                                          value: selectedValue,
                                          onChanged: ((value) {
                                            setState(() {
                                              selectedValue = value.toString();
                                              dropDownActions.clear();
                                              listDropDown(listOfActions[selectedValue] ?? []);
                                            });
                                          }),
                                          items: dropdownItems,
                                        ),
                                      ),
                                      Padding(
                                        padding: EdgeInsets.all(8.0),
                                        child: DropdownButton(
                                          value: selectedAction,
                                          onChanged: ((value) {
                                            selectedAction = value.toString();
                                          }),
                                          items: dropDownActions,
                                        ),
                                      ),
                                      Padding(
                                        padding: const EdgeInsets.all(8.0),
                                        child: MaterialButton(
                                          child: Text("Submit"),
                                          onPressed: () {
                                            print(selectedValue);
                                            print(selectedAction);
                                            selectedValue = "Discord";
                                            selectedAction = "Appeler Erwan";
                                            listDropDown(listOfActions["Discord"] ?? []);
                                            Navigator.of(context).pop();
                                          },
                                        ),
                                      )
                                    ],
                                  ),
                                ),
                              ],
                            ),
                          );
                        });
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
