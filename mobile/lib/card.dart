import 'package:flutter/material.dart';

Card buildCard(String actionService, String actionFunc, String reactionService, String reactionFunc) {
   return Card(
      elevation: 4.0,
      child: Column(
        children: [
          Text(actionService),
          Text(actionFunc),
          Text(reactionService),
          Text(reactionFunc),
        ],
      )
    );
 }