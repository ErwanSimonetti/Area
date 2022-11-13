import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'globals.dart' as globals;
import 'dart:async';
import 'dart:convert';
import 'dart:io';
import 'package:url_launcher/url_launcher.dart';

class ListPlateform extends StatefulWidget {
  ListPlateform();

  @override
  _ListPlateformState createState() => new _ListPlateformState();
}

class _ListPlateformState extends State<ListPlateform> {
  TextEditingController mailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();

  Widget build(BuildContext context) {
    mailController.text = "";
    passwordController.text = "";
    return Scaffold(
      body: ListView(
        scrollDirection: Axis.vertical,
        children: <Widget>[
          MaterialButton(
            minWidth: double.infinity,
            height: 60,
            onPressed: () {
              authenticate();
            },
            color: Colors.redAccent,
            shape:
                RoundedRectangleBorder(borderRadius: BorderRadius.circular(40)),
            child: const Text(
              "Discord",
              style: TextStyle(
                fontWeight: FontWeight.w600,
                fontSize: 16,
              ),
            ),
          ),
          makeInput(mailController, label: "Email"),
          makeInput(passwordController, label: "Password", obsureText: true),
          MaterialButton(
            minWidth: double.infinity,
            height: 60,
            onPressed: () {
              loginEmail();
            },
            color: Colors.redAccent,
            shape:
                RoundedRectangleBorder(borderRadius: BorderRadius.circular(40)),
            child: const Text(
              "Submit",
              style: TextStyle(
                fontWeight: FontWeight.w600,
                fontSize: 16,
              ),
            ),
          ),
        ],
      ),
    );
  }

  Future<void> loginEmail() async {
    print(mailController.text);
    print(passwordController.text);
    var response = await http.post(
        Uri.parse("http://10.0.2.2:8080/email/login"),
        headers: {'Content-Type': 'text/plain', 'cookie': globals.tokenUser},
        body: json.encode({
          'email': mailController.text,
          'password': passwordController.text,
        }));
    print(response.statusCode);
    print(response.headers);
    if (response.statusCode == 200) {
      print("sucess");
    } else {
      print("Ratio");
    }
  }

  launchURL(String url) async {
    await launchURL(url);
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

  void authenticate() async {
    var response = await http.get(
      Uri.parse("http://10.0.2.2:8080/discord/auth/url"),
      headers: {
        'Content-Type': 'text/html; charset=utf-8',
        'cookie': globals.tokenUser
      },
    );
    print(response.statusCode);
    if (response.statusCode == 200) {
      launchURL(jsonDecode(utf8.decode(response.bodyBytes)));
      print("sau6");

      // final currentUrl = Uri.base;
      // if (!currentUrl.fragment.contains('access_token=')) {
      //   // You are not connected so redirect to the Twitch authentication page.
      //   WidgetsBinding.instance.addPostFrameCallback((_) {
      //     html.window.location.assign(
      //       'https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=$clientId&redirect_uri=${currentUrl.origin}&scope=viewing_activity_read',
      //     );
      //   });
      // } else {
      //   // You are connected, you can grab the code from the url.
      //   final fragments = currentUrl.fragment.split('&');
      //   _token = fragments
      //       .firstWhere((e) => e.startsWith('access_token='))
      //       .substring('access_token='.length);
    }
  }

  // final result = await FlutterWebAuth.authenticate(
  //   url: jsonDecode(utf8.decode(response.bodyBytes)),
  //   callbackUrlScheme: 'http://localhost:8080/discord/auth',
  // );

  // Extract token from resulting url
}
