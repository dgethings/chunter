/**
 * @file Tree sitter grammar for the Cisco IOS network Operating System
 * @author David Gethings <dgethings76@gmail.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

export default grammar({
  name: "cisco_ios",

  rules: {
    configuration_file: $ => seq($._section),
    
    _section: $ => choice($.system_section),

    system_section: $ => seq($.hostname_statement),

    hostname_statement: $ => seq(
      "hostname",
      $.value
    ),

    value: $ => /[a-zA-Z0-9]+/,
  }
});
