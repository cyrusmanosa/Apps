import 'package:dating_your_date/client/grpc_services.dart';
import 'package:dating_your_date/core/app_export.dart';
import 'package:dating_your_date/pb/rpc_ResetPasswordEmail.pb.dart';
import 'package:dating_your_date/widgets/app_bar/Custom_App_bar.dart';
import 'package:dating_your_date/widgets/Custom_IconLogoBox.dart';
import 'package:dating_your_date/widgets/app_bar/custom_Input_bar.dart';
import 'package:dating_your_date/widgets/Custom_Input_Form_Bar.dart';
import 'package:dating_your_date/widgets/Custom_WarningLogoBox.dart';
import 'package:dating_your_date/widgets/button/custom_outlined_button.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';

class NewPasswordEmailSetup extends StatefulWidget {
  NewPasswordEmailSetup({Key? key}) : super(key: key);

  @override
  _NewPasswordSetupEmailState createState() => _NewPasswordSetupEmailState();
}

class _NewPasswordSetupEmailState extends State<NewPasswordEmailSetup> {
  TextEditingController newPasswordSetupController = TextEditingController();
  TextEditingController newPasswordSetupConfirmController = TextEditingController();

  // Grpc
  void resetPasswordGrpcRequest(BuildContext context) async {
    if (newPasswordSetupController.text != newPasswordSetupConfirmController.text) {
      await showErrorDialog(context, "パスワード（確認）とパスワードは一致しません");
    } else if (isPureText(newPasswordSetupController.text) || isPureNumber(newPasswordSetupController.text)) {
      await showErrorDialog(context, "パスワードの組み合わせは英数字は必要です");
    } else {
      try {
        final request = ResetPwEmailConfirmRequest(
          email: null,
          password: newPasswordSetupController.text,
        );
        await GrpcInfoService.client.resetPwEmailConfirm(request);
        await showLogoDialog(context, "新しいパスワード設定しました", false);
        await Future.delayed(Duration(seconds: 1));
        Navigator.pop(context);
        Navigator.popAndPushNamed(context, AppRoutes.login);
      } on GrpcError {
        Navigator.pop(context);
        await showErrorDialog(context, "エラー：検証可能な入力データがありません。");
        throw Exception("新しいパスワードの設定中にエラーが発生しました。");
      }
    }
  }

  bool passwordVisible = false;
  bool isPureText(String value) {
    final pattern = RegExp(r'^[a-zA-Z]+$');
    return pattern.hasMatch(value);
  }

  bool isPureNumber(String value) {
    final pattern = RegExp(r'^\d+$');
    return pattern.hasMatch(value);
  }

  @override
  Widget build(BuildContext context) {
    MediaQueryData mediaQueryData = MediaQuery.of(context);
    double mediaH = mediaQueryData.size.height;
    double mediaW = mediaQueryData.size.width;
    return Scaffold(
      appBar: buildAppBar(context, "", true),
      backgroundColor: appTheme.bgColor,
      resizeToAvoidBottomInset: true,
      body: GestureDetector(
        onTap: () {
          FocusScope.of(context).requestFocus(FocusNode());
        },
        child: Container(
          padding: EdgeInsets.symmetric(horizontal: mediaW / 13),
          child: Column(
            children: [
              // Logo and Slogan
              CustomImageView(imagePath: ImageConstant.imgLogo, width: mediaW / 4),
              CustomImageView(imagePath: ImageConstant.imgSlogan, width: mediaW / 3.5),
              SizedBox(height: mediaH / 30),

              // New Password
              CustomInputBar(titleName: "新しいパスワード", backendPart: _buildNewPasswordInput(context)),

              // msg
              Align(alignment: Alignment.centerLeft, child: Text("＊半角英数字の組合せ（8桁以上15桁以下）", style: CustomTextStyles.pwRulegrey500)),
              SizedBox(height: mediaH / 25),

              // New Password Confirm
              CustomInputBar(titleName: "新しいパスワード（確認）", backendPart: _buildNewPasswordConfirm(context)),
              SizedBox(height: mediaH / 25),

              // Button
              _buildNextButton(context)
            ],
          ),
        ),
      ),
    );
  }

  /// NPW Input
  Widget _buildNewPasswordInput(BuildContext context) {
    return CustomInputFormBar(
      controller: newPasswordSetupController,
      textInputAction: TextInputAction.done,
      textInputType: TextInputType.visiblePassword,
      hintText: "Ankdl2332",
      obscureText: passwordVisible ? false : true,
      focusNode: FocusNode(),
      onTap: () {
        FocusNode().requestFocus();
      },
      suffix: IconButton(
        icon: Icon(passwordVisible ? Icons.visibility : Icons.visibility_off),
        onPressed: () {
          setState(() {
            passwordVisible = !passwordVisible;
          });
        },
      ),
    );
  }

  /// NPW Confirm Input
  Widget _buildNewPasswordConfirm(BuildContext context) {
    return CustomInputFormBar(
      controller: newPasswordSetupConfirmController,
      textInputAction: TextInputAction.done,
      textInputType: TextInputType.visiblePassword,
      hintText: "Ankdl2332",
      obscureText: passwordVisible ? false : true,
      focusNode: FocusNode(),
      onTap: () {
        FocusNode().requestFocus();
      },
      suffix: IconButton(
        icon: Icon(passwordVisible ? Icons.visibility : Icons.visibility_off),
        onPressed: () {
          setState(() {
            passwordVisible = !passwordVisible;
          });
        },
      ),
    );
  }

  /// Next Button
  Widget _buildNextButton(BuildContext context) {
    return CustomOutlinedButton(
      text: "設定",
      onPressed: () {
        resetPasswordGrpcRequest(context);
      },
    );
  }
}