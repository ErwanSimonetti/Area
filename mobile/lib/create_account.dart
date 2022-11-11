import 'package:flutter/material.dart';
import 'package:flutter/gestures.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:http/http.dart' as http;
import './home_page.dart';
import './login_route.dart';
import 'dart:async';
import 'dart:convert';

class SignupPage extends StatelessWidget {
  final TextEditingController email = TextEditingController();
  final TextEditingController passWord = TextEditingController();
  final TextEditingController confirmPassword = TextEditingController();
  final TextEditingController firstName = TextEditingController();
  final TextEditingController lastName = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: false,
      backgroundColor: Colors.white,
      body: SafeArea(
        child: SingleChildScrollView(
          child: Container(
            color: Colors.transparent,
            height: MediaQuery.of(context).size.height,
            width: double.infinity,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Column(
                  children: [
                    Column(
                      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                      children: [
                        const SizedBox(
                          height: 20,
                        ),
                        const Text(
                          "Sign up",
                          style: TextStyle(
                            fontSize: 25,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                        const SizedBox(
                          height: 20,
                        )
                      ],
                    ),
                    Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 40),
                      child: Column(
                        children: [
                          makeInput(email, label: "Email"),
                          makeInput(firstName, label: "First Name"),
                          makeInput(lastName, label: "Last Name"),
                          makeInput(passWord,
                              label: "Password", obsureText: true),
                          makeInput(confirmPassword,
                              label: "Confirm Pasword", obsureText: true),
                          MaterialButton(
                            minWidth: double.infinity,
                            height: 60,
                            onPressed: () {
                              final bool emailValid = RegExp(
                                      r"^[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+")
                                  .hasMatch(email.text);
                              if (confirmPassword.text == passWord.text &&
                                  emailValid == true) {
                                register(context);
                              } else if (confirmPassword.text != passWord.text) {
                                Fluttertoast.showToast(
                                    msg:
                                        'Password and password confirmation are different.');
                              } else if (emailValid == false) {
                                Fluttertoast.showToast(msg: 'Invalid mail.');
                              }
                            },
                            color: Colors.redAccent,
                            shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(40)),
                            child: const Text(
                              "Sign Up",
                              style: TextStyle(
                                fontWeight: FontWeight.w600,
                                fontSize: 16,
                              ),
                            ),
                          ),
                        ],
                      ),
                    ),
                    const SizedBox(
                      height: 20,
                    ),
                    Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: <Widget>[
                          makeClickableText(
                              'Already have account? ', 'Login', context)
                        ])
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget makeClickableText(
      String textNonClickable, String textClickable, BuildContext context) {
    return RichText(
      text: TextSpan(children: [
        TextSpan(
          text: textNonClickable,
          style: const TextStyle(
            color: Colors.black,
          ),
        ),
        TextSpan(
            text: textClickable,
            style: const TextStyle(
              color: Colors.blue,
            ),
            recognizer: TapGestureRecognizer()
              ..onTap = () {
                Navigator.push(
                  context,
                  MaterialPageRoute(builder: (context) => const MyStatefulWidget()),
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
          style: const TextStyle(
              fontSize: 10, fontWeight: FontWeight.w400, color: Colors.black87),
        ),
        const SizedBox(
          height: 5,
        ),
        TextField(
          obscureText: obsureText,
          controller: text,
          decoration: InputDecoration(
            contentPadding: const EdgeInsets.symmetric(vertical: 0, horizontal: 4),
            enabledBorder: OutlineInputBorder(
              borderSide: BorderSide(
                color: Colors.grey.shade400,
              ),
            ),
            border: OutlineInputBorder(
                borderSide: BorderSide(color: Colors.grey.shade400)),
          ),
        ),
        const SizedBox(
          height: 30,
        )
      ],
    );
  }

  Future<void> register(BuildContext context) async {
    var response = await http.post(Uri.parse("http://10.0.2.2:8080/register/"),
        headers: {'Content-Type': 'text/plain'},
        body: json.encode({
          'firstname': firstName.text,
          'lastname': lastName.text,
          'email': email.text,
          'password': passWord.text,
        })
    );
    if (response.statusCode == 200) {
      Fluttertoast.showToast(msg: 'Account successfuly created');
      Navigator.push(
        context,
        MaterialPageRoute(builder: (context) => HomePage()),
      );
    } else {
      Fluttertoast.showToast(msg: 'Cannot create account an error occured');
    }
  }
}
