import 'package:dating_your_date/core/app_export.dart';
import 'package:flutter/material.dart';

// ignore: must_be_immutable
class ShownicknamebarItemWidget extends StatelessWidget {
  const ShownicknamebarItemWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 24.v,
      width: 330.h,
      child: Stack(
        alignment: Alignment.centerRight,
        children: [
          // Under bar
          Align(
            alignment: Alignment.bottomCenter,
            child: SizedBox(width: 330.h, child: Divider()),
          ),
          // Data
          Align(
            alignment: Alignment.centerRight,
            child: Text("山崎太一", style: CustomTextStyles.bodyLargeBlack900),
          ),
          // Theme
          Align(
            alignment: Alignment.centerLeft,
            child: Text("ニックネーム", style: CustomTextStyles.bodyLargePinkA100),
          ),
        ],
      ),
    );
  }
}