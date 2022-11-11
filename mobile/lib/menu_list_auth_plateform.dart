import 'package:flutter/material.dart';

class ListPlateform extends StatefulWidget {
  ListPlateform(this.title, this.map);

  final String title;
  final Map<String, void Function()> map;

  @override
  _ListPlateformState createState() => new _ListPlateformState();
}

class _ListPlateformState extends State<ListPlateform> {
  Widget build(BuildContext context) {
    return new Scaffold(
      body: ListView(
        scrollDirection: Axis.vertical,
        children: buildListImageCard(widget.map),
      ),
    );
  }

  List<Widget> buildListImageCard(Map<String, void Function()> map) {
    List<Widget> list = [];

    map.forEach((key, value) {
      list.add(Container(
          margin: const EdgeInsets.symmetric(vertical: 20.0),
          height: 200.0,
          child: Ink.image(
              image: NetworkImage(key),
              child: InkWell(
                onTap: () {
                  value();
                },
              ))));
    });

    return list;
  }

  Widget buildImageCard(String link) => Container(
      margin: const EdgeInsets.symmetric(vertical: 20.0),
      height: 200.0,
      child: Ink.image(
          image: NetworkImage(link),
          child: InkWell(
            onTap: () {
              print("Hello");
            },
          )));
}
