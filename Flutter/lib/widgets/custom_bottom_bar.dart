import 'package:dating_your_date/core/app_export.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

// ignore: must_be_immutable
class CustomBottomBar extends StatefulWidget {
  CustomBottomBar({this.onChanged});
  Function(BottomBarEnum)? onChanged;
  @override
  CustomBottomBarState createState() => CustomBottomBarState();
}

class CustomBottomBarState extends State<CustomBottomBar> {
  int selectedIndex = 0;
  List<BottomMenuModel> bottomMenuList = [
    // ホーム
    BottomMenuModel(
      icon: ImageConstant.imgNav,
      activeIcon: ImageConstant.imgNav,
      label: "ホーム",
      type: BottomBarEnum.tf,
    ),
    // ターゲット
    BottomMenuModel(
      icon: ImageConstant.imgNavPrimary,
      activeIcon: ImageConstant.imgNavPrimary,
      label: "ターゲット",
      type: BottomBarEnum.tf1,
    ),
    // チャット
    BottomMenuModel(
      icon: ImageConstant.imgNavPrimary41x41,
      activeIcon: ImageConstant.imgNavPrimary41x41,
      label: "チャット",
      type: BottomBarEnum.tf2,
    ),
    // プロフィール
    BottomMenuModel(
      icon: ImageConstant.imgNav41x41,
      activeIcon: ImageConstant.imgNav41x41,
      label: "プロフィール",
      type: BottomBarEnum.tf3,
    )
  ];

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 80.v,
      decoration: BoxDecoration(color: appTheme.gray500),
      child: BottomNavigationBar(
        backgroundColor: Colors.transparent,
        showSelectedLabels: false,
        showUnselectedLabels: false,
        selectedFontSize: 0,
        elevation: 0,
        currentIndex: selectedIndex,
        type: BottomNavigationBarType.fixed,
        items: List.generate(bottomMenuList.length, (index) {
          // Footage
          return BottomNavigationBarItem(
            // Icon bar
            label: "Test",
            icon: Column(
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                CustomImageView(
                    imagePath: bottomMenuList[index].icon, height: 41.adaptSize, width: 41.adaptSize, color: theme.colorScheme.primary),
                Padding(
                  padding: EdgeInsets.only(top: 2.v),
                  child: Text(
                    bottomMenuList[index].label!,
                    style: CustomTextStyles.roundedMplus1cPrimary.copyWith(color: theme.colorScheme.primary),
                  ),
                ),
              ],
            ),
            // Icon OnTap
            activeIcon: Column(
              mainAxisSize: MainAxisSize.min,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                CustomImageView(
                    imagePath: bottomMenuList[index].activeIcon, height: 41.adaptSize, width: 41.adaptSize, color: appTheme.cyan600),
                Padding(
                  padding: EdgeInsets.only(top: 2.v),
                  child: Text(
                    bottomMenuList[index].label!,
                    style: CustomTextStyles.roundedMplus1cCyan600.copyWith(color: appTheme.cyan600),
                  ),
                ),
              ],
            ),
          );
        }),
        onTap: (index) {
          selectedIndex = index;
          widget.onChanged?.call(bottomMenuList[index].type!);
          setState(() {});
        },
      ),
    );
  }
}

enum BottomBarEnum { tf, tf1, tf2, tf3 }

class BottomMenuModel {
  BottomMenuModel({
    this.label,
    this.activeIcon,
    this.icon,
    this.type,
  });

  String? activeIcon;
  String? icon;
  String? label;
  BottomBarEnum? type;
}

class DefaultWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      padding: EdgeInsets.all(10),
      child: Center(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('Please replace the respective Widget here', style: TextStyle(fontSize: 18)),
          ],
        ),
      ),
    );
  }
}