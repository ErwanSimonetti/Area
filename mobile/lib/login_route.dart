import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:flutter/gestures.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:http/http.dart' as http;
import './colors.dart';
import './widgets.dart';
import './home_page.dart';
import './create_account.dart';
import 'dart:async';
import 'dart:convert';
import 'dart:io';

class MyStatefulWidget extends StatefulWidget {
  const MyStatefulWidget({Key? key}) : super(key: key);

  @override
  State<MyStatefulWidget> createState() => _MyStatefulWidgetState();
}

class _MyStatefulWidgetState extends State<MyStatefulWidget> {
  TextEditingController nameController = TextEditingController();
  TextEditingController passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    nameController.text = "";
    passwordController.text = "";
    return new Scaffold(
        body: Padding(
            padding: const EdgeInsets.all(10),
            child: ListView(
              children: <Widget>[
                Container(
                    alignment: Alignment.center,
                    padding: const EdgeInsets.all(10),
                    child: const Text(
                      'Area',
                      style: TextStyle(
                          color: Colors.blue,
                          fontWeight: FontWeight.bold,
                          fontSize: 30),
                    )),
                Container(
                    alignment: Alignment.center,
                    padding: const EdgeInsets.all(10),
                    child: const Text(
                      'Sign in',
                      style: TextStyle(fontSize: 25),
                    )),
                makeInput(nameController, label: "Email"),
                makeInput(passwordController,
                    label: "Password", obsureText: true),
                Column(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: [
                    CustomWidgets.socialButtonRect('Login with Google',
                        googleColor, FontAwesomeIcons.googlePlusG,
                        onTap: () {}),
                    CustomWidgets.socialButtonRect('Login with Facebook',
                        facebookColor, FontAwesomeIcons.facebookF, onTap: () {
                      Fluttertoast.showToast(msg: 'I am Facebook');
                    }),
                    CustomWidgets.socialButtonRect('Login with GitHub',
                        githubColor, FontAwesomeIcons.github, onTap: () {
                      Fluttertoast.showToast(msg: 'I am guy teub');
                    }),
                    MaterialButton(
                      minWidth: double.infinity,
                      height: 60,
                      onPressed: () {
                        login(context);
                      },
                      color: Colors.redAccent,
                      shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(40)),
                      child: const Text(
                        "Login",
                        style: TextStyle(
                          fontWeight: FontWeight.w600,
                          fontSize: 16,
                        ),
                      ),
                    ),
                  ],
                ),
                const SizedBox(
                  height: 20,
                ),
                Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: <Widget>[
                      makeClickableText(
                          'Does not have an account ? ', 'Sign up', context)
                    ])
              ],
            )));
  }

  Widget makeClickableText(
      String textNonClickable, String textClickable, BuildContext context) {
    return RichText(
      text: TextSpan(children: [
        TextSpan(
          text: textNonClickable,
          style: TextStyle(
            color: Colors.black,
          ),
        ),
        TextSpan(
            text: textClickable,
            style: TextStyle(
              color: Colors.blue,
            ),
            recognizer: TapGestureRecognizer()
              ..onTap = () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (context) => SignupPage()),
                );
              }),
      ]),
    );
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

  Future<void> login(BuildContext context) async {
    var response = await http.post(Uri.parse("http://10.0.2.2:8080/login/"),
        headers: {'Content-Type': 'text/plain'},
        body: json.encode({
          'email': nameController.text,
          'password': passwordController.text,
        }));
    print(response.statusCode);
    print(response.headers);
    if (response.statusCode == 200) {
      final cookies = response.headers['set-cookie'];

      if (cookies != null) {
        final authToken = cookies.split(';')[0];
      }
      Navigator.push(
        context,
        MaterialPageRoute(builder: (context) => HomePage()),
      );
    } else {
      Fluttertoast.showToast(msg: 'Error: Invalid login');
    }
  }
}
