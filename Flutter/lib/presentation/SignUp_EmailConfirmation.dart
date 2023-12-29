import 'package:dating_your_date/core/app_export.dart';
import 'package:dating_your_date/client/grpc_services.dart';
import 'package:dating_your_date/pb/rpc_checkEmail.pb.dart';
import 'package:dating_your_date/widgets/app_bar/custom_Input_bar.dart';
import 'package:dating_your_date/widgets/Custom_Outlined_Button.dart';
import 'package:dating_your_date/widgets/Custom_Input_Form_Bar.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

// ignore_for_file: must_be_immutable
class EmailConfirmation extends StatelessWidget {
  EmailConfirmation({Key? key}) : super(key: key);
  RegExp emailRegex = RegExp(r'^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$');
  TextEditingController emailController = TextEditingController();

  // Http
  void resetPasswordEmailHttpRequest(BuildContext context) async {
    var url = "http://127.0.0.1:8080/SignUpEmail";
    var requestBody = {"email": emailController.text};
    var response = await http.post(Uri.parse(url), body: jsonEncode(requestBody), headers: {"Content-Type": "application/json"});

    if (response.statusCode == 200) {
      onTapNextButton(context);
    }
  }

// Grpc
  void resetPasswordEmailGrpcRequest(BuildContext context) async {
    final request = CheckEmailRequest(email: emailController.text);
    try {
      await GrpcInfoService.client.checkEmail(request);
      onTapNextButton(context);
    } on GrpcError catch (e) {
      if (e.code == 6) showErrorDialog(context, "このメールアドレスは登録できません。");
    }
  }

  void showErrorDialog(BuildContext context, String errorMessage) {
    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          shape: RoundedRectangleBorder(borderRadius: BorderRadiusStyle.r15),
          // Error Logo
          title: CustomImageView(
            imagePath: ImageConstant.imgWarning,
            height: mediaQueryData.size.height / 20,
            width: mediaQueryData.size.width / 10,
            alignment: Alignment.center,
          ),

          // Word
          content: Container(
            width: mediaQueryData.size.width / 1.1,
            child: Text(errorMessage, style: CustomTextStyles.msgWordOfMsgBox, textAlign: TextAlign.center),
          ),
          actions: [
            CustomOutlinedButton(
              alignment: Alignment.center,
              text: "OK",
              margin: EdgeInsets.only(bottom: mediaQueryData.size.height / 100),
              onPressed: () {
                onTapReturn(context);
              },
            ),
          ],
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        resizeToAvoidBottomInset: false,
        body: Form(
          child: Container(
            width: double.maxFinite,
            padding: EdgeInsets.symmetric(horizontal: mediaQueryData.size.width / 13, vertical: mediaQueryData.size.height / 20),
            child: Column(
              children: [
                // Logo and Slogan
                SizedBox(height: mediaQueryData.size.height / 15),
                CustomImageView(imagePath: ImageConstant.imgLogo, width: mediaQueryData.size.width / 4.5),
                CustomImageView(imagePath: ImageConstant.imgSlogan, width: mediaQueryData.size.width / 3.5),
                SizedBox(height: mediaQueryData.size.height / 30),

                // input
                CustomInputBar(titleName: "メールアドレス", backendPart: _buildEmailInputSection(context)),
                SizedBox(height: mediaQueryData.size.height / 50),

                // send button
                _buildNextPageButton(context),
                SizedBox(height: mediaQueryData.size.height / 30),

                // 手続き
                Align(
                  alignment: Alignment.centerLeft,
                  child: Container(child: Text("・ご手続きは1回のみです", overflow: TextOverflow.ellipsis, style: CustomTextStyles.titleOfUnderLogo)),
                ),

                // information note
                Align(
                  alignment: Alignment.topLeft,
                  child: Container(
                    child: Text("・メールアドレスの受信確認が必須です。", overflow: TextOverflow.ellipsis, style: CustomTextStyles.titleOfUnderLogo),
                  ),
                ),
                Align(
                  alignment: Alignment.topLeft,
                  child: Container(
                    child: Text("・ご登録済みのお客様は受信確認をお願いします。", overflow: TextOverflow.ellipsis, style: CustomTextStyles.titleOfUnderLogo),
                  ),
                ),
                SizedBox(height: mediaQueryData.size.height / 50),
              ],
            ),
          ),
        ),
      ),
    );
  }

  /// Section Widget
  Widget _buildEmailInputSection(BuildContext context) {
    return CustomInputFormBar(
      controller: emailController,
      hintText: " example@email.com",
      textInputType: TextInputType.emailAddress,
    );
  }

  // Next Button
  Widget _buildNextPageButton(BuildContext context) {
    return CustomOutlinedButton(
      text: "送信",
      onPressed: () {
        if (isEmailValid(emailController.text)) {
          resetPasswordEmailGrpcRequest(context);
        } else {
          showErrorDialog(context, "無効なメールアドレス");
        }
      },
    );
  }

  onTapReturn(BuildContext context) {
    Navigator.pop(context);
  }

  onTapNextButton(BuildContext context) {
    Navigator.pushNamed(context, AppRoutes.confirmationCore);
  }

  bool isEmailValid(String email) {
    final emailRegex = RegExp(r'^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$');
    return emailRegex.hasMatch(email);
  }
}