import 'package:dating_your_date/core/app_export.dart';
import 'package:dating_your_date/widgets/base_button.dart';
import 'package:flutter/material.dart';

class CustomElevatedButton extends BaseButton {
  CustomElevatedButton({
    Alignment? alignment,
    bool? isDisabled,
    ButtonStyle? buttonStyle,
    double? height,
    double? width,
    EdgeInsets? margin,
    Key? key,
    TextStyle? buttonTextStyle,
    VoidCallback? onPressed,
    this.decoration,
    this.leftIcon,
    this.rightIcon,
    required String text,
  }) : super(
          text: text,
          onPressed: onPressed,
          buttonStyle: buttonStyle,
          isDisabled: isDisabled,
          buttonTextStyle: buttonTextStyle,
          height: height,
          width: width,
          alignment: alignment,
          margin: margin,
        );

  final BoxDecoration? decoration;
  final Widget? leftIcon;
  final Widget? rightIcon;

  @override
  Widget build(BuildContext context) {
    return alignment != null
        ? Align(alignment: alignment ?? Alignment.center, child: buildElevatedButtonWidget)
        : buildElevatedButtonWidget;
  }

  Widget get buildElevatedButtonWidget => Container(
        height: this.height ?? 120.v,
        width: this.width ?? double.maxFinite,
        margin: margin,
        decoration: decoration,
        child: ElevatedButton(
          style: buttonStyle,
          onPressed: isDisabled ?? false ? null : onPressed ?? () {},
          child: Row(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              leftIcon ?? const SizedBox.shrink(),
              Text(text, style: buttonTextStyle ?? theme.textTheme.displayMedium),
              rightIcon ?? const SizedBox.shrink(),
            ],
          ),
        ),
      );
}
