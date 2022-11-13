import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:flutter/material.dart';
import './drop_down_button.dart';
import './menu_list_auth_plateform.dart';
import './create_area_menu.dart';
import './wallet.dart';

class HomePage extends StatefulWidget {
  HomePage({Key? key, this.title = ""}) : super(key: key);

  final String title;

  @override
  _HomePageState createState() => new _HomePageState();
}

class _HomePageState extends State<HomePage> {
  PageController? _pageController;
  int _page = 0;
  String _title = "MyApp";
  Color _appBarColor = Colors.pink;

  @override
  Widget build(BuildContext context) {
    return new Scaffold(
      appBar: new AppBar(
        title: new Text(_title),
        backgroundColor: _appBarColor,
        leading: new CustomButtonTest(),
      ),
      body: PageView(
        children: <Widget>[
          Container(
            child: new WalletPage(),
          ),
          Container(
            child: new CreationAreaMenu(),
          ),
          Container(
            child: new ListPlateform(),
          ),
        ],
        controller: _pageController,
        onPageChanged: onPageChanged,
      ),
      bottomNavigationBar: BottomNavigationBar(
        items: [
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: "My Area",
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.exposure_plus_1),
            label: "New Area",
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.pie_chart),
            label: "Stats",
          ),
        ],
        onTap: navigateToPage,
        currentIndex: _page,
      ),
    );
  }

  void navigateToPage(int page) {
    _pageController?.animateToPage(page,
        duration: Duration(milliseconds: 300), curve: Curves.ease);
  }

  void onPageChanged(int page) {
    String _temptitle = "";
    Color _tempColor = Colors.blue;
    switch (page) {
      case 0:
        _temptitle = "My Area";
        _tempColor = Colors.pink;
        break;
      case 1:
        _temptitle = "New Area";
        _tempColor = Colors.green;
        break;
      case 2:
        _temptitle = "Stats";
        _tempColor = Colors.deepPurple;
        break;
    }
    setState(() {
      this._page = page;
      this._title = _temptitle;
      this._appBarColor = _tempColor;
    });
  }

  @override
  void initState() {
    super.initState();
    _pageController = new PageController();
    _title = "My Area";
  }

  @override
  void dispose() {
    super.dispose();
    _pageController?.dispose();
  }
}
