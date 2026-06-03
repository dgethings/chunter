/**
 * @file Tree sitter grammar for the Cisco IOS network Operating System
 * @author David Gethings <dgethings76@gmail.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

export default grammar({
  name: "cisco_ios",

  conflicts: $ => [
    [$._alias_section, $.service_section]
  ],

  rules: {
    configuration_file: $ => seq(
      "!",
      $._running_version,
      "!",
      optional($.service_section),
      optional($.archive_section),
      optional($.async_bootp_section),
      optional($.line_section),
      optional($.hostname_section),
      optional($._alias_section),
    ),

    // this is a comment showing the loaded OS version the device is running
    _running_version: $ => seq(
      "!",
      "version",
      field("running_version", $.running_version)
    ),

    running_version: $ => /[a-z0-9\.-]+/,

    service_section: $ => seq(
      optional($._configured_version),
      optional($.service_statements),
      "!"
    ),

    _configured_version: $ => seq(
      "version",
      field("configured_version", $.configured_version)
    ),

    configured_version: $ => /[a-z0-9\.-]+/,

    service_statements: $ => seq(
      "no",
      "service",
      "pad"
    ),

    hostname_section: $ => seq(
      "hostname",
      $.value,
      "!"
    ),

    value: $ => /[a-zA-Z0-9]+/,

    _alias_section: $ => seq(
      repeat($.alias_statement),
      "!"
    ),

    alias_statement: $ => seq(
      "alias",
      $.mode,
      $.alias,
      $.command
    ),

    mode: $ => choice("exec", "foo"),

    alias: $ => /[a-z]+/,

    command: $ => /[a-z0-9 ]+/,

    line_section: $ => seq(
      "line",
      $.line_type,
      $.activation_character,
      "!"
    ),

    line_type: $ => choice("console", "tty0"),

    activation_character: $ => seq(
      "activation-character",
      $.ascii_value
    ),

    ascii_value: $ => /[0-9]+/,

    archive_section: $ => seq(
      "archive",
      $._archive_path,
      "end",
      "!"
    ),

    _archive_path: $ => seq(
      "path",
      $.file_path,
    ),

    file_path: $ => /[a-z0-9:\/]+/,

    async_bootp_section: $ => seq(
      "async-bootp",
      $.async_bootp_tag,
      optional($.async_bootp_hostname),
      $.async_bootp_data,
      "!"
    ),

    async_bootp_tag: $ => choice("bootfile"),

    async_bootp_hostname: $ => /\:[a-z0-9\.]+/,

    async_bootp_data: $ => /"[a-z]+"/,
  }
});
