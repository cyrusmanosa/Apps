import 'package:dating_your_date/presentation/appNavigationScreen.dart';

import 'package:dating_your_date/presentation/Chat/Chat.dart';
import 'package:dating_your_date/presentation/ChatBox.dart';
import 'package:dating_your_date/presentation/CompanionshipConditions.dart';
import 'package:dating_your_date/presentation/CompanionshipConditionsRepair.dart';
import 'package:dating_your_date/presentation/ConfirmationCore.dart';
import 'package:dating_your_date/presentation/ConfirmationCore_Error.dart';
import 'package:dating_your_date/presentation/ContainerScreen.dart';

import 'package:dating_your_date/presentation/DeleteTarget.dart';

import 'package:dating_your_date/presentation/EmailConfirmation.dart';
import 'package:dating_your_date/presentation/EmailConfirmation_Error.dart';

import 'package:dating_your_date/presentation/Home/Home.dart';
import 'package:dating_your_date/presentation/HobbyCondition.dart';
import 'package:dating_your_date/presentation/HobbyConditionRepair.dart';

import 'package:dating_your_date/presentation/Login.dart';
import 'package:dating_your_date/presentation/LowPercentageHeightWarnning.dart';
import 'package:dating_your_date/presentation/LoverConditionsRepair.dart';
import 'package:dating_your_date/presentation/LoverConditions.dart';
import 'package:dating_your_date/presentation/LowPercentageLowWarnning.dart';

import 'package:dating_your_date/presentation/NewPassword_Done.dart';
import 'package:dating_your_date/presentation/NewPassword_Error.dart';
import 'package:dating_your_date/presentation/NewPassword_Setup.dart';
import 'package:dating_your_date/presentation/NewTargetReady.dart';

import 'package:dating_your_date/presentation/PayDone.dart';
import 'package:dating_your_date/presentation/PasswordReset_Email.dart';
import 'package:dating_your_date/presentation/PasswordReset_Error.dart';
import 'package:dating_your_date/presentation/Profile/Profile.dart';
import 'package:dating_your_date/presentation/ProfileEdit.dart';

import 'package:dating_your_date/presentation/SearchTitle.dart';

import 'package:dating_your_date/presentation/SideBar.dart';
import 'package:dating_your_date/presentation/Signup_1.dart';
import 'package:dating_your_date/presentation/Signup_3.dart';
import 'package:dating_your_date/presentation/Signup_2.dart';
import 'package:dating_your_date/presentation/WarningDeleteUser.dart';
import 'package:dating_your_date/presentation/WarningReturnResetPage.dart';
import 'package:dating_your_date/presentation/TargetReady.dart';
import 'package:dating_your_date/presentation/Target/Target.dart';
import 'package:flutter/material.dart';

class AppRoutes {
  static const String appNavigationScreen = '/app_navigation_screen';
  static const String chat = '/chat';
  static const String chatBox = '/ChatBox';
  static const String containerScreen = '/containerScreen';
  static const String confirmationCore = '/confirmationCore';
  static const String confirmationCoreError = '/confirmationCoreError';
  static const String companionshipConditions = '/companionshipConditions';
  static const String companionshipConditionsRepair = '/companionshipConditionsRepair';
  static const String deleteTarget = '/deleteTarget';
  static const String emailConfirmation = '/emailConfirmation';
  static const String emailConfirmationError = '/emailConfirmationError';
  static const String login = '/login';
  static const String lowPercentageHeightWarnning = '/lowPercentageHeightWarnning';
  static const String lowPercentageLowWarnning = '/lowPercentageLowWarnning';
  static const String loverConditions = '/loverConditions';
  static const String loverConditionsRepair = '/loverConditionsRepair';
  static const String home = '/home';
  static const String hobbyCondition = '/hobbyCondition';
  static const String hobbyConditionRepair = '/hobbyConditionRepair';
  static const String newPasswordError = '/newPasswordError';
  static const String newPasswordSetup = '/newPasswordSetup';
  static const String newPasswordDone = '/newPasswordDone';
  static const String newTargetReady = '/newTargetReady';
  static const String passwordResetError = '/passwordResetError';
  static const String passwordResetEmail = '/passwordResetEmail';
  static const String payDone = '/payDone';
  static const String profile = '/profile';
  static const String profileEdit = '/profileEdit';
  static const String searchTitle = '/searchTitle';
  static const String sideBar = '/sideBar';
  static const String signUp_1 = '/signUp_1';
  static const String signUp_2 = '/signUp_2';
  static const String signUp_3 = '/signUp_3';
  static const String target = '/target';
  static const String targetReady = '/targetReady';
  static const String warningDeleteUser = '/warningDeleteUser';
  static const String warningReturnResetPage = '/warningReturnResetPage';

  static Map<String, WidgetBuilder> routes = {
    appNavigationScreen: (context) => AppNavigationScreen(),
    chat: (context) => Chat(),
    chatBox: (context) => ChatBox(),
    containerScreen: (context) => ContainerScreen(),
    companionshipConditions: (context) => CompanionshipConditions(),
    companionshipConditionsRepair: (context) => CompanionshipConditionsRepair(),
    confirmationCore: (context) => ConfirmationCore(),
    confirmationCoreError: (context) => ConfirmationCoreError(),
    deleteTarget: (context) => DeleteTarget(),
    emailConfirmation: (context) => EmailConfirmation(),
    emailConfirmationError: (context) => EmailConfirmationError(),
    hobbyCondition: (context) => HobbyCondition(),
    hobbyConditionRepair: (context) => HobbyConditionRepair(),
    home: (context) => Home(),
    login: (context) => Login(),
    lowPercentageHeightWarnning: (context) => LowPercentageHeightWarnning(),
    lowPercentageLowWarnning: (context) => LowPercentageLowWarnning(),
    loverConditions: (context) => LoverConditions(),
    loverConditionsRepair: (context) => LoverConditionsRepair(),
    newPasswordDone: (context) => NewPasswordDone(),
    newPasswordError: (context) => NewPasswordError(),
    newPasswordSetup: (context) => NewPasswordSetup(),
    newTargetReady: (context) => NewTargetReady(),
    passwordResetError: (context) => PasswordResetError(),
    passwordResetEmail: (context) => PasswordResetEmail(),
    payDone: (context) => PayDone(),
    profile: (context) => Profile(),
    profileEdit: (context) => ProfileEdit(),
    searchTitle: (context) => SearchTitle(),
    sideBar: (context) => SideBar(),
    signUp_1: (context) => SignUp_1(),
    signUp_2: (context) => SignUp_2(),
    signUp_3: (context) => SignUp_3(),
    target: (context) => Target(),
    targetReady: (context) => TargetReady(),
    warningDeleteUser: (context) => WarningDeleteUser(),
    warningReturnResetPage: (context) => WarningReturnResetPage()
  };
}
