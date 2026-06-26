package cisco_ios

import (
	"github.com/dgethings/chunter/internal/keyword"
	"github.com/dgethings/chunter/internal/protocol"
)

var Keywords = []keyword.Keyword{
    {
        Keyword: "activation-character",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define the character you ent er at a vacant terminal to begin a terminal session, use the activation-character command in line configuration mode. To make any character activate a terminal, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "activation-character ${1:ascii-number}", "no activation-character", },
    },
    {
        Keyword: "alias",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To create a command alias, use the alias command in global configuration mode. To delete all aliases in a command mode or to delete a specific alias, and to revert to the original command syntax, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "alias ${1:mode} ${2:command-alias} ${3:original-command}", "no alias ${1:mode} ${2: [command-alias] }", },
    },
    {
        Keyword: "archive",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enter archive configuration mode, use the archive command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "archive", },
    },
    {
        Keyword: "async-bootp",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure extended BOOTP requests for asynchronous interfaces as defined in RFC 1084, use the async-bootp command in global configuration mode. To restore the default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "async-bootp ${1:tag} :${2:hostname} ${3:data}", "no async-bootp", },
    },
    {
        Keyword: "autobaud",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the line for automatic baud rate detection (autobaud), use the autobaud command in line configuration mode. To disable automatic baud detection, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "autobaud", "no autobaud", },
    },
    {
        Keyword: "auto-sync",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable automatic synchronization of the configuration files in NVRAM, use the auto-sync command in main-cpu redundancy configuration mode. To disable automatic synchronization, use the no form of this command.",
        },
        Section: "config-r",
        Snippets: []string{ "auto-sync startup-configconfig-registerbootvarrunning-configstandard", "no auto-sync startup-configconfig-registerbootvarstandard", },
    },
    {
        Keyword: "autoupgrade disk-cleanup",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the Cisco IOS Auto-Upgrade Manager disk cleanup utility, use the autoupgrade disk-cleanup command in global configuration mode. To disable this configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "autoupgrade disk-cleanup crashinfocoreimageirrecoverable", "no autoupgrade disk-cleanup crashinfocoreimageirrecoverable", },
    },
    {
        Keyword: "autoupgrade ida url",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the URL of the Intelligent Download Application (IDA) running on www.cisco.com, use the autoupgrade ida url command in global configuration mode. To disable this URL, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "autoupgrade ida url ${1:url}", "no autoupgrade ida url ${1:url}", },
    },
    {
        Keyword: "autoupgrade status email",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the address to which status email is to be sent and the outgoing email server, use the autoupgrade status email command in global configuration mode. To disable status email, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "autoupgrade status email recipient ${1: [email-address] } smtp-server ${2: [smtp-server] }", "no autoupgrade status email recipient ${1: [email-address] } smtp-server ${2: [smtp-server] }", },
    },
    {
        Keyword: "banner exec",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify and enable a message to be displayed when an EXEC process is created (an EXEC banner), use the banner exec command in global configuration mode. To delete the existing EXEC banner, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "banner exec ${1:d} ${2:message} ${3:d}", "no banner exec", },
    },
    {
        Keyword: "banner incoming",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define and enable a banner to be displayed when there is an incoming connection to a terminal line from a host on the network, use the banner incoming command in global configuration mode. To delete the incoming connection banner, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "banner incoming ${1:d} ${2:message} ${3:d}", "no banner incoming", },
    },
    {
        Keyword: "banner login",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define and enable a customized banner to be displayed before the username and password login prompts, use the banner login command in global configuration mode. To disable the login banner, use no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "banner login ${1:d} ${2:message} ${3:d}", "no banner login", },
    },
    {
        Keyword: "banner motd",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define and enable a message-of-the-day (MOTD) banner, use the banner motd command in global configuration mode. To delete the MOTD banner, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "banner motd ${1:d} ${2:message} ${3:d}", "no banner motd", },
    },
    {
        Keyword: "banner slip-ppp",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To customize the banner that is displayed when a Serial Line Internet Protocol (SLIP) or PPP connection is made, use the banner slip-ppp command in global configuration mode. To restore the default SLIP or PPP banner, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "banner slip-ppp ${1:d} ${2:message} ${3:d}", "no banner slip-ppp", },
    },
    {
        Keyword: "boot bootldr",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the location of the boot image that ROM uses for booting, use the boot bootldr command in global configuration mode. To remove this boot image specification, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "boot bootldr ${1:file-url} boot bootldr command", "no boot bootldr", },
    },
    {
        Keyword: "boot bootstrap",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the filename that is used to boot a secondary bootstrap image, use the boot bootstrap command in global configuration mode. To disable booting from a secondary bootstrap image, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "boot bootstrap ${1:file-url}", "no boot bootstrap ${1:file-url}", "boot bootstrap flash ${1: [filename] }", "no boot bootstrap flash ${1: [filename] }", "boot bootstrap tftp ${1:filename} ${2: [ip-address] }", "no boot bootstrap tftp ${1:filename} ${2: [ip-address] }", "boot bootstrap mop ${1:filename} ${2:interface-type} ${3:interface-number}", "no boot bootstrap mop ${1:filename} ${2:interface-type} ${3:interface-number}", },
    },
    {
        Keyword: "boot config",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To s pecify the device and filename of the configuration file from which the system configures itself during initialization (startup), use the boot config command in global configuration mode. To return to the default location for the configuration file, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "boot config ${1:file-system-prefix}:${2: [directory/] }${3:filename} nvbypass", "no boot config", "boot config ${1:device}:${2:filename} nvbypass", "no boot config", },
    },
    {
        Keyword: "boot host",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the host-specific configuration file to be used at the next system startup, use the boot host command in global configuration mode. To restore the host configuration filename to the default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "boot host commandboot host ${1:remote-url}", "no boot host ${1:remote-url}", },
    },
    {
        Keyword: "boot network",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To change the default name of the network configuration file from which to load configuration commands, use the boot network command in global configuration mode. To restore the network configuration filename to the default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "boot network ${1:remote-url}", "no boot network ${1:remote-url}", },
    },
    {
        Keyword: "boot system",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the system image that the router loads at startup, use one of the following boot system command in global configuration mode. To remove the startup system image specification, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "boot system ${1:file-url}${2: | filename}", "no boot system ${1:file-url}${2: | filename}", "boot system flash ${1:flash-fs}: ${2:partition-number}: ${3: [filename] }", "no boot system flash ${1:flash-fs}: ${2:partition-number}: ${3: [filename] }", "boot system mop ${1:filename} ${2: [mac-address] } ${3: [interface] }", "no boot system mop ${1:filename} ${2: [mac-address] } ${3: [interface] }", "boot system rom", "no boot system rom", "boot system rcptftpftp ${1:filename} ${2: [ip-address] }", "no boot system rcptftpftp ${1:filename} ${2: [ip-address] }", },
    },
    {
        Keyword: "clock",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the port clocking mode for the 1000BASE-T transceivers, use the clock command in interface configuration mode. To return to the default settings,use the no form of this command.",
        },
        Section: "config-if",
        Snippets: []string{ "clock autoactive preferpassive prefer", "no clock", },
    },
    {
        Keyword: "clock initialize nvram",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To restart the system clock from the last known system clock value, use the clock initialize nvram command in global configuration mode. To disable the restart of the system clock from the last known system clock value, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "clock initialize nvram", "no clock initialize nvram", },
    },
    {
        Keyword: "config-register",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To change the configuration register settings, use the config-register command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "config-register ${1:value}", },
    },
    {
        Keyword: "configuration mode exclusive",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "",
        },
        Section: "config",
        Snippets: []string{ "configuration mode exclusive automanual expire ${1:seconds} lock-show interleave terminate config-wait ${2:seconds} retry-wait ${3:seconds}", "no configuration mode exclusive", },
    },
    {
        Keyword: "databits",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the number of data bits per character that are interpreted and generated by the router hardware, use the databits command in line configuration mode. To restore the default value, use the no form of the command.",
        },
        Section: "config-line",
        Snippets: []string{ "databits 5678", "no databits", },
    },
    {
        Keyword: "data-character-bits",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the number of data bits per character that are interpreted and generated by the Cisco IOS software, use the data-character-bits command in line configuration mode. To restore the default value, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "data-character-bits 78", "no data-character-bits", },
    },
    {
        Keyword: "default-value data-character-bits",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the number of data bits per character that are generated and interpreted by Cisco software to either 7 bits or 8 bits, use the default-value data-character-bits command in global configuration mode. To disable the configured size, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "default-value data-character-bits 78", "no default-value data-character-bits", },
    },
    {
        Keyword: "default-value exec-character-bits",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define the EXEC character width for either 7 bits or 8 bits, use the default-value exec-character-bits command in global configuration mode. To restore the default value, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "default-value exec-character-bits 78", "no default-value exec-character-bits", },
    },
    {
        Keyword: "default-value modem-interval",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the default frequency time to scan modem signals, use the default-value modem-interval command in global configuration mode. To disable the configured frequency, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "default-value modem-interval ${1:milliseconds}", "no default-value modem-interval", },
    },
    {
        Keyword: "default-value special-character-bits",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the flow control default value from a 7-bit width to an 8-bit width, use the default-value special-character-bits command in global configuration mode. To restore the default value, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "default-value special-character-bits commanddefault-value special-character-bits 78", "no default-value special-character-bits", },
    },
    {
        Keyword: "define interface-range",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To create an interface-range macro, use the define interface-range command in global configuration mode. To remove an interface-range macro, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "define interface-range ${1:macro-name} ${2:interface-range}", },
    },
    {
        Keyword: "diagnostic bootup level",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the diagnostic bootup level, use the diagnostic bootup level command in global configuration mode. To skip all diagnostic tests, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "diagnostic bootup level minimalcomplete", "no diagnostic bootup level", },
    },
    {
        Keyword: "diagnostic cns",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the Cisco Networking Services (CNS) diagnostics, use the diagnostic cns command in global configuration mode. To disable sending diagnostic results to the CNS event bus., use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "diagnostic cns publishsubscribe ${1: [subject] }", "no diagnostic cns publishsubscribe ${1: [subject] }", },
    },
    {
        Keyword: "diagnostic event-log size",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To modify the diagnostic event log size dynamically, use the diagnostic event-log size command in global configuration mode. To return to the default settings, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "diagnostic event-log size ${1:size}", "no diagnostic event-log size", },
    },
    {
        Keyword: "diagnostic monitor",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure health-monitoring diagnostic testing, use the diagnostic monitor command in global configuration mode. To disable testing, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "diagnostic monitor intervalmodule ${1:number} test ${2:test-id}${3: | test-id-range}all ${4:hh:mm:ss} ${5:milliseconds} ${6:days}", "diagnostic monitor syslog", "diagnostic monitor module ${1:num} test ${2:test-id}${3: | test-id-range}all", "no diagnostic monitor intervalsyslog", "diagnostic monitor bay ${1:slot}/ ${2:bay}slot ${3:slot} ${4:number}subslot ${5:slot}/ ${6:subslot} test ${7:test-id}${8: | test-id-range}all", "diagnostic monitor intervalbay ${1:slot}/ ${2:bay}slot ${3:slot-no}subslot ${4:slot}/ ${5:subslot} test ${6:test-id}${7: | test-id-range}all ${8:hh:mm:ss} ${9:milliseconds} ${10:days}", "diagnostic monitor syslog", "diagnostic monitor threshold bay ${1:slot}/ ${2:bay}slot ${3:slot-no}subslot ${4:slot}/ ${5:subslot} test ${6:test-id}${7: | test-id-range}all failure count ${8:failures} runsdayshoursminutessecondsmilliseconds ${9:window_size}", },
    },
    {
        Keyword: "diagnostic schedule module",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the scheduling of test-based diagnostic testing for a specific module or schedule a supervisor engine switchover, use the diagnostic schedule module command in global configuration mode. To remove the scheduling, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "diagnostic schedule module ${1:module-number}${2: | slot/subslot} test ${3:test-id}allcompleteminimalnon-disruptiveper-port port${4:interface-port-number}${5: | port-number-list}all on ${6:month} ${7:dd} ${8:yyyy} ${9:hh:mm}daily ${10:hh:mm}weekly ${11:day-of-week} ${12:hh:mm}", "no diagnostic schedule module ${1:module-number}${2: | slot/subslot} test ${3:test-id}allcompleteminimalnon-disruptiveper-port port${4:interface-port-number}${5: | port-number-list}all on ${6:month} ${7:dd} ${8:yyyy} ${9:hh:mm}daily ${10:hh:mm}weekly ${11:day-of-week} ${12:hh:mm}", },
    },
    {
        Keyword: "disconnect-character",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define a character to disconnect a session, use the disconnect-character command in line configuration mode. To remove the disconnect character, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "disconnect-character ${1:ascii-number}", "no disconnect-character", },
    },
    {
        Keyword: "dispatch-character",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define a character that causes a packet to be sent, use the dispatch-character command in line configuration mode. To remove the definition of the specified dispatch character, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "dispatch-character ${1:ascii-number1} ${2:ascii-number2}. . . ", "no dispatch-character ${1:ascii-number1} ${2:ascii-number2}. . . ", },
    },
    {
        Keyword: "dispatch-machine",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify an identifier for a TCP packet dispatch state machine on a particular line, use the dispatch-machine command in line configuration mode. To disable a state machine on a particular line, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "dispatch-machine ${1:name}", "no dispatch-machine", },
    },
    {
        Keyword: "dispatch-timeout",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the character dispatch timer, use the dispatch-timeout command in line configuration mode. To remove the timeout definition, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "dispatch-timeout ${1:milliseconds}", "no dispatch-timeout", },
    },
    {
        Keyword: "do",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To execute user EXEC or privileged EXEC commands from global configuration mode or other configuration modes or submodes, use the do command in any configuration mode.",
        },
        Section: "",
        Snippets: []string{ "do ${1:command}", },
    },
    {
        Keyword: "downward-compatible-config",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To generate a configuration that is compatible with an earlier Cisco IOS release, use the downward-compatible-config command in global configuration mode. To disable this function, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "downward-compatible-config ${1:version}", "no downward-compatible-config", },
    },
    {
        Keyword: "editing",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To reen able Cisco IOS enhanced editing features for a particular line after they have been disabled, use the editing command in line configuration mode. To disable these features, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "editing", "no editing", },
    },
    {
        Keyword: "enable last-resort",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable password parameters as the last resort without specifying the local enable password if no TACACS servers respond, use the enable last-resort command in global configuration mode. To disable the password parameters, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "no enable last-resortpasswordsucceed", "no enable last-resort", },
    },
    {
        Keyword: "end",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To en d the current configuration session and return to privileged EXEC mode, use the end command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "end", },
    },
    {
        Keyword: "environment-monitor shutdown temperature",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable monitoring of the environment sensors, use the environment-monitor shutdown temperature command in global configuration mode. To disable monitoring of the environment sensors, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "environment-monitor shutdown temperature rommonpowerdown", "no environment-monitor shutdown temperature rommonpowerdown", },
    },
    {
        Keyword: "environment temperature-controlled",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the ambient temperature control, use the environment temperature-controlled command in global configuration mode. To disable the ambient temperature control, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "environment temperature-controlled", "no environment temperature-controlled", },
    },
    {
        Keyword: "errdisable detect cause",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable error-disable detection, use the errdisable detect cause command in global configuration mode. To disable error-disable detection, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "errdisable detect cause allbpduguarddtp-flapl2ptguardlink-flappacket-buffer-errorpagp-flaprootguardudld", "no errdisable detect cause allbpduguarddtp-flapl2ptguardlink-flappagp-flaprootguardudld", },
    },
    {
        Keyword: "errdisable recovery",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure recovery mechanism variables, use the errdisable recovery command in global configuration mode. To return to the default state, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "errdisable recovery cause allarp-inspectionbpduguardchannel-misconfigdhcp-rate-limitdtp-flapgbic-invalidl2ptguardlink-flappagp-flappsecure-violationsecurity-violationrootguardudldunicast-floodinterval ${1:seconds}", "no errdisable recovery cause allarp-inspectionbpduguardchannel-misconfigdhcp-rate-limitdtp-flapgbic-invalidl2ptguardlink-flappagp-flappsecure-violationsecurity-violationrootguardudldunicast-floodinterval ${1:seconds}", },
    },
    {
        Keyword: "escape-character",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define a system escape character, use the escape-character command in line configuration mode. To set the escape character to Break, use the no or default form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "escape-character break${1: | char}defaultnonesoft", "no escape-character soft", "default escape-character soft", },
    },
    {
        Keyword: "exec",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To allow an EXEC process on a line, use the exec command in line configuration mode. To turn off the EXEC process for the specified line, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "exec", "no exec", },
    },
    {
        Keyword: "exec-banner",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To reenable the display of EXEC and message-of-the-day (MOTD) banners on the specified line or lines, use the exec-banner command in line configuration mode. To suppress the banners on the specified line or lines, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "exec-banner", "no exec-banner", },
    },
    {
        Keyword: "exec-character-bits",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the character widths of EXEC and configuration command characters, use the exec-character-bits command in line configuration mode. To restore the default value, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "exec-character-bits 78", "no exec-character-bits", },
    },
    {
        Keyword: "exec-timeout",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the interval that the EXEC command interpreter waits until user input is detected, use the exec-timeout command in line configuration mode. To remove the timeout definition, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "exec-timeout ${1:minutes} ${2: [seconds] }", "no exec-timeout", },
    },
    {
        Keyword: "exit (global)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To exit any configuration mode to the next highest mode in the CLI mode hierarchy, use the exit command in any configuration mode.",
        },
        Section: "",
        Snippets: []string{ "exit", },
    },
    {
        Keyword: "file privilege",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "",
        },
        Section: "config",
        Snippets: []string{ "file privilege level ${1:level}", "no file privilege level ${1:level}", },
    },
    {
        Keyword: "file prompt",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the level of prompting, use the file prompt command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "file prompt prompt alertnoisyquiet", },
    },
    {
        Keyword: "file verify auto",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable automatic image verification, use the file verify auto command in global configuration mode. To disable automatic image verification, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "file verify auto", "no file verify auto", },
    },
    {
        Keyword: "full-help",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To get help f or the full set of user-level commands, use the full-help command in line configuration mode.",
        },
        Section: "config-line",
        Snippets: []string{ "full-help", },
    },
    {
        Keyword: "hidekeys",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To suppress the display of password information in configuration log files, use the hidekeys command in configuration change logger configuration mode. To allow the display of password information in configuration log files, use the no form of this command.",
        },
        Section: "config-archive-log-config",
        Snippets: []string{ "hidekeys", "no hidekeys", },
    },
    {
        Keyword: "history",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the command history function, use the history command in line configuration mode. To disable the command history function, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "history", "no history", },
    },
    {
        Keyword: "history size",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To change the command history buffer size for a particular line, use the history size command in line configuration mode. To reset the command history buffer size to ten lines, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "history size ${1:number-of-lines}", "no history size", },
    },
    {
        Keyword: "hold-character",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define the local hold character used to pause output to the terminal screen, use the hold-character command in line configuration mode. To restore the default, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "hold-character ${1:ascii-number}", "no hold-character", },
    },
    {
        Keyword: "hostname",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify or modify the hostname for the network server, use the hostname command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "hostname ${1:name}", },
    },
    {
        Keyword: "insecure",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure a line as insecure, use the insecure command in line configuration mode. To disable this function, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "insecure", "no insecure", },
    },
    {
        Keyword: "international",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "If you are using Telnet to access a Cisco IOS platform and you want to display 8-bit and multibyte international characters (for example, Kanji) and print the Escape character as a single character instead of as the caret and bracket symbols (^[), use the international command in line configuration mode. To display characters in 7-bit format, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "international", "no international", },
    },
    {
        Keyword: "ip bootp server",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the Bootstrap Protocol (BOOTP) service on your routing device, use the ip bootp server command in global configuration mode. To disable BOOTP services, use the no form of the command.",
        },
        Section: "config",
        Snippets: []string{ "ip bootp server", "no ip bootp server", },
    },
    {
        Keyword: "ip finger",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure a system to accept Finger protocol requests (defined in RFC 742), use the ip finger command in global configuration mode. To disable this service, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip finger rfc-compliant", "no ip finger", },
    },
    {
        Keyword: "ip ftp passive",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the router to use only passive FTP connections, use the ip ftp passive command in global configuration mode . To allow all types of FTP connections, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip ftp passive", "no ip ftp passive", },
    },
    {
        Keyword: "ip ftp password",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the password to be used for File Transfer Protocol (FTP) connections, use the ip ftp password command in global configuration mode. To return the password to its default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip ftp password ${1: [type] } ${2:password}", "no ip ftp password", },
    },
    {
        Keyword: "ip ftp source-interface",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the source IP address for File Transfer Protocol (FTP) connections, use the ip ftp source-interface command in global configuration mode. To use the address of the interface where the connection is made, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip ftp source-interface ${1:interface-type} ${2:interface-number}", "no ip ftp source-interface", },
    },
    {
        Keyword: "ip ftp username",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the username for File Transfer Protocol (FTP) connections, use the ip ftp username command in global configuration mode . To configure the router to attempt anonymous FTP, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip ftp username ${1:username}", "no ip ftp username", },
    },
    {
        Keyword: "ip rarp-server",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the router to act as a Reverse Address Resolution Protocol (RARP) server, use the ip rarp-server command in interface configuration mode. To restore the interface to the default of no RARP server support, use the no form of this command.",
        },
        Section: "config-if",
        Snippets: []string{ "ip rarp-server ${1:ip-address}", "no ip rarp-server ${1:ip-address}", },
    },
    {
        Keyword: "ip rcmd domain-lookup",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To reena ble the basic Domain Name Service (DNS) security check for rcp and rsh, use the ip rcmd domain-lookup command in global configuration mode. T o disable the basic DNS security check for remote copy protocol (rcp) and remote shell protoco (rsh), use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip rcmd domain-lookup", "no ip rcmd domain-lookup", },
    },
    {
        Keyword: "ip rcmd rcp-enable",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the Cisco IOS software to allow remote users to copy files to and from the router using remote copy protocol (rcp), use the ip rcmd rcp-enable command in global configuration mode. To disable rcp on the device, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip rcmd rcp-enable", "no ip rcmd rcp-enable", },
    },
    {
        Keyword: "ip rcmd remote-host",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To create an entry for the remote user in a local authentication database so that remote users can execute commands on the router using remote shell protocol (rsh) or remote copy protocol (rcp), use the ip rcmd remote-host command in global configuration mode. To remove an entry for a remote user from the local authentication database, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip rcmd remote-host ${1:local-username} ${2:ip-address}${3: | host-name} ${4:remote-username} enable ${5: [level] }", "no ip rcmd remote-host ${1:local-username} ${2:ip-address}${3: | host-name} ${4:remote-username} enable ${5: [level] }", },
    },
    {
        Keyword: "ip rcmd remote-username",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the remote username to be used when requesting a remote copy using remote copy protocol (rcp), use the ip rcmd remote-username command in global configuration mode . To remove from the configuration the remote username, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip rcmd remote-username ${1:username}", "no ip rcmd remote-username ${1:username}", },
    },
    {
        Keyword: "ip rcmd rsh-enable",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the router to allow remote users to execute commands on it using remote shell protocol (rsh), use the ip rcmd rsh-enable command in global configuration mode. To disable a router that is enabled for rsh, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip rcmd rsh-enable", "no ip rcmd rsh-enable", },
    },
    {
        Keyword: "ip rcmd source-interface",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To force remote copy protocol (rcp) or remote shell protocol (rsh) to use the IP address of a specified interface for all outgoing rcp/rsh communication packets, use the ip rcmd source-interface command in global configuration mode. To disable a previously configured ip rcmd source-interface command, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip rcmd source-interface ${1:interface-id}", "no ip rcmd source-interface ${1:interface-id}", },
    },
    {
        Keyword: "ip telnet source-interface",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the IP address of an interface as the source address for Telnet connections, use the ip telnet source-interface command in global configuration mode. To reset the source address to the default for each connection, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip telnet source-interface ${1:interface}", "no ip telnet source-interface", },
    },
    {
        Keyword: "ip tftp blocksize",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To negotiate a transfer TFTP blocksize, use the ip tftp blocksize command in global configuration mode. To disable this configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip tftp blocksize ${1:bytes}", "no ip tftp blocksize", },
    },
    {
        Keyword: "ip tftp boot-interface",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To use an interface for TFTP booting, use the ip tftp boot-interface command in global configuration mode. To disable this configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip tftp boot-interface ${1:type} ${2:number}", "no ip tftp boot-interface", },
    },
    {
        Keyword: "ip tftp min-timeout",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the minimum timeout period for retransmission of data using TFTP, use the ip tftp min-timeout command in global configuration mode. To disable, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip tftp min-timeout ${1:seconds}", "no ip tftp min-timeout", },
    },
    {
        Keyword: "ip tftp source-interface",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the IP address of an interface as the source address for TFTP connections, use the ip tftp source-interface command in global configuration mode. To return to the default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip tftp source-interface ${1:interface-type} ${2:interface-number}", "no ip tftp source-interface", },
    },
    {
        Keyword: "ip wccp web-cache accelerated",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the hardware acceleration for WCCP version 1, use the ip wccp web-cache accelerated command in global configuration mode. To disable hardware acceleration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "ip wccp web-cache accelerated group-address ${1:groupaddress}redirect-list ${2:access-list}group-list ${3:access-list}password ${4:password}", "no ip wccp web-cache accelerated", },
    },
    {
        Keyword: "length",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the terminal screen length, use the length command in line configuration mode. To restore the default value, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "length ${1:screen-length}", "no length", },
    },
    {
        Keyword: "load-interval",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To change the length of time for which data is used to compute load statistics, use the load-interval command in interface configuration, Frame Relay DLCI configuration, or template configuration modes. To revert to the default setting, use the no form of this command.",
        },
        Section: "config-template",
        Snippets: []string{ "load-interval ${1:seconds}", "no load-interval ${1:seconds}", },
    },
    {
        Keyword: "location",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To provide a description of the location of a serial device, use the location command in line configuration mode. To remove the description, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "location ${1:text}", "no location", },
    },
    {
        Keyword: "lockable",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable use of the lock EXEC command, use the lockable command in line configuration mode. To reinstate the default (the terminal session cannot be locked), use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "lockable", "no lockable", },
    },
    {
        Keyword: "log config",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enter configuration change logger configuration mode, use the log config command in archive configuration mode.",
        },
        Section: "config-archive",
        Snippets: []string{ "log config", },
    },
    {
        Keyword: "logging buffered",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable system message logging to a local buffer, use the logging buffered command in global configuration mode. To cancel the use of the buffer, use the no form of this command. To return the buffer size to its default value, use the default form of this command.",
        },
        Section: "config",
        Snippets: []string{ "logging buffered discriminator ${1:discriminator-name} ${2: [buffer-size] } ${3: [severity-level] }", "no logging buffered", "default logging buffered", },
    },
    {
        Keyword: "logging buginf",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To allow debug messages to be generated for the standard system logging buffer, use the logging buginf command in global configuration mode. To disable the logging for debugging functionality, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "logging buginf", "no logging buginf", },
    },
    {
        Keyword: "logging enable",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the logging of configuration changes, use the logging enable command in configuration change logger configuration mode. To disable the logging of configuration changes, use the no form of this command.",
        },
        Section: "config-archive-log-config",
        Snippets: []string{ "logging enable", "no logging enable", },
    },
    {
        Keyword: "logging esm config",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To permit configuration changes from Embedded Syslog Manager (ESM) filters, use the logging esm config command in global configuration mode. To disable the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "logging esm config", "no logging esm config", },
    },
    {
        Keyword: "logging event bundle-status",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable message bundling, use the logging event bundle-status command in interface configuration mode. To disable message bundling, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "logging event bundle-status", "no logging event bundle-status", },
    },
    {
        Keyword: "logging event link-status (global configuration)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To change the default or set the link-status event messaging during system initialization, use the logging event link-status command in global configuration mode. To disable the link-status event messaging, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "logging event link-status defaultboot", "no logging event link-status defaultboot", },
    },
    {
        Keyword: "logging event link-status (interface configuration)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable link-status event messaging on an interface, use the logging event link-status command in interface configuration mode. To disable link-status event messaging, use the no form of this command.",
        },
        Section: "config-if",
        Snippets: []string{ "logging event link-status bchandchannfas", "no logging event link-status bchandchannfas", },
    },
    {
        Keyword: "logging event subif-link-status",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the link-status event messaging on a subinterface, use the logging event subif-link-status command in interface configuration mode. To disable the link-status event messaging on a subinterface, use the no form of this command.",
        },
        Section: "config-if",
        Snippets: []string{ "logging event subif-link-status", "no logging event subif-link-status", },
    },
    {
        Keyword: "logging event trunk-status",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable trunk status messaging, use the logging event trunk-status command in interface configuration mode. To disable trunk status messaging, use the no form of this command.",
        },
        Section: "config-if",
        Snippets: []string{ "logging event trunk-status", "no logging event trunk-status", },
    },
    {
        Keyword: "logging reload",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the reload logging level, use the logging reload command in global configuration mode. To disable the reload logging, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "logging reload message-limit ${1:number} ${2:severity-level}alertscriticaldebuggingemergencieserrorsinformationalnotificationswarnings", "no logging reload", },
    },
    {
        Keyword: "logging ip access-list cache (global configuration)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the Optimized ACL Logging (OAL) parameters, use the logging ip access-list cache command in global configuration mode. To return to the default settings, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ " logging ip access-list cache entries ${1:entries}interval ${2:seconds}rate-limit ${3:pps}threshold ${4:packets}", "no logging ip access-list cache entriesintervalrate-limitthreshold", },
    },
    {
        Keyword: "logging ip access-list cache (interface configuration)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable an Optimized ACL Logging (OAL)-logging cache on an interface that is based on direction, use the logging ip access-list cache command in interface configuration mode. To disable OAL, use the no form of this command.",
        },
        Section: "config-if",
        Snippets: []string{ "logging ip access-list cache inout", "no logging ip access-list cache", },
    },
    {
        Keyword: "logging persistent (config-archive-log-cfg)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the configuration logging persistent feature and to select how the configuration commands are to be saved to the Cisco IOS secure file system, use the logging persistent command in the log config submode of archive configuration mode. To disable this capability, use the no form of this command.",
        },
        Section: "config-archive-log-config",
        Snippets: []string{ "logging persistent automanual", "no logging persistent automanual", },
    },
    {
        Keyword: "logging persistent reload (config-archive-log-cfg)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To sequentially apply the configuration commands saved in the configuration logger database (since the last write memory command) to the running-config file after a reload, use the logging persistent reload command in configuration change logger configuration mode in archive configuration mode. To disable this capability, use the no form of this command.",
        },
        Section: "config-archive-log-config",
        Snippets: []string{ "logging persistent reload", "no logging persistent reload", },
    },
    {
        Keyword: "logging purge-log buffer days",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To automatically delete the entries from the logging buffer after a configurable time, use the logging purge-log bufferdays command in global configuration mode. To disable this capability, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "logging purge-log buffer days${1:number-of-days}[ time  ${2:deletion-start-time}] ", "nologging purge-log buffer", },
    },
    {
        Keyword: "logging size",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the maximum number of entries retained in the configuration log, use the logging size command in configuration change logger configuration mode. To reset the default value, use the no form of this command.",
        },
        Section: "config-archive-log-config",
        Snippets: []string{ "logging size ${1:entries}", "no logging size", },
    },
    {
        Keyword: "logging synchronous",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To synchronize unsolicited messages and debug output with solicited Cisco IOS software output and prompts for a specific console port line, auxiliary port line, or vty, use the logging synchronous command in line configuration mode. To disable synchronization of unsolicited messages and debug output, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "logging synchronous level ${1:severity-level}all limit ${2:number-of-lines}", "no logging synchronous level ${1:severity-level}all limit ${2:number-of-lines}", },
    },
    {
        Keyword: "logging system",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable System Event Archive (SEA) logging, use the logging system command in global configuration mode. To disable SEA logging, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "logging system disk ${1:name}", "no logging system", },
    },
    {
        Keyword: "logout-warning",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To warn users of an impending forced timeout, use the logout-warning command in line configuration mode. To restore the default, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "logout-warning ${1: [seconds] }", "logout-warning", },
    },
    {
        Keyword: "macro (global configuration)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To create a global command macro, use the macro command in global configuration mode. To remove the macro, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "macro global apply ${1:macro-name}description ${2:text}trace ${3:macro-name} ${4: [keyword-to-value] } ${5:value-first-keyword} ${6: [keyword-to-value] } ${7:value-second-keyword} ${8: [keyword-to-value] } ${9:value-third-keyword} ${10: [keyword-to-value] }name ${11:macro-name}", "no macro global apply ${1:macro-name}description ${2:text}trace ${3:macro-name} ${4: [keyword-to-value] } ${5:value-first-keyword} ${6: [keyword-to-value] } ${7:value-second-keyword} ${8: [keyword-to-value] } ${9:value-third-keyword} ${10: [keyword-to-value] }name ${11:macro-name}", },
    },
    {
        Keyword: "macro (interface configuration)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To create an interface-specific command macro, use the macro command in interface configuration mode. To remove the macro, use the no form of this command.",
        },
        Section: "config-if",
        Snippets: []string{ "macro apply ${1:macro-name}description ${2:text}trace ${3:macro-name} ${4: [keyword-to-value] } ${5:value-first-keyword} ${6: [keyword-to-value] } ${7:value-second-keyword} ${8: [keyword-to-value] } ${9:value-third-keyword} ${10: [keyword-to-value] }", "no macro apply ${1:macro-name}description ${2:text}trace ${3:macro-name} ${4: [keyword-to-value] } ${5:value-first-keyword} ${6: [keyword-to-value] } ${7:value-second-keyword} ${8: [keyword-to-value] } ${9:value-third-keyword} ${10: [keyword-to-value] }", },
    },
    {
        Keyword: "maximum",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the maximum number of archive files of the running configuration to be saved in the Cisco configuration archive, use the maximum command in archive configuration mode. To reset this command to its default, use the no form of this command.",
        },
        Section: "config-archive",
        Snippets: []string{ "maximum ${1:number}", "no maximum ${1:number}", },
    },
    {
        Keyword: "memory cache error-recovery",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To trace error recovery in memory using caches, use the memory cache error-recovery command in global configuration mode. To disable the memory cache error recovery mechanisms, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "memory cache error-recovery L1L2L3 datainst", "no memory cache error-recovery L1L2L3 datainst", },
    },
    {
        Keyword: "memory cache error-recovery options",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To trace error recovery in memory using caches through set options, use the memory cache error-recovery options command in global configuration mode. To disable the set memory cache error recovery mechanisms, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "memory cache error-recovery options abort-if-same-contentblocking-modemax-recoveries ${1:value}nvram-reportparity-checkwindow ${2:seconds}", "no memory cache error-recovery options abort-if-same-contentblocking-modemax-recoveries ${1:value}nvram-reportparity-checkwindow ${2:seconds}", },
    },
    {
        Keyword: "memory free low-watermark",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure a router to issue system logging message notifications when available memory falls below a specified threshold, use the memory free low-watermark command in global configuration mode. To disable memory threshold notifications, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "memory free low-watermark processor ${1:threshold}io ${2:threshold}", "no memory free low-watermark", },
    },
    {
        Keyword: "memory lite",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the memory allocation lite (malloc_lite) feature, use the memory lite command in global configuration mode. To disable this feature, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "memory lite", "no memory lite", },
    },
    {
        Keyword: "memory reserve",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To reserve a specified amount of memory in kilobytes for console access and critical notifications, use the memory reserve command in global configuration mode. To disable the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "memory reserve console ${1:size}critical ${2: [total-size] }", "no memory reserve consolecritical", "memory reserve critical ${1: [total-size] }", "no memory reserve critical", },
    },
    {
        Keyword: "memory reserve critical",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "",
        },
        Section: "config",
        Snippets: []string{ "memory reserve critical ${1:kilobytes}", "no memory reserve critical", },
    },
    {
        Keyword: "memory sanity",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To perform a “sanity check” for corruption in buffers and queues, use the memory sanity command in global configuration mode. To disable this feature, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "memory sanity bufferqueueall", "no memory sanity", },
    },
    {
        Keyword: "memory scan",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the Memory Scan feature, use the memory scan command in global configuration mode. To restore the router configuration to the default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "memory scan", "no memory scan", },
    },
    {
        Keyword: "memory-size iomem",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To reallocate the percentage of DRAM to use for I/O memory and processor memory, use the memory-size iomem command in global configuration mode. To revert to the default memory allocation, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "memory-size iomem ${1:i/o-memory-percentage}", "no memory-size iomem", },
    },
    {
        Keyword: "menu menu-name single-space",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To display menu items single-spaced rather than double-spaced, use the menu <menu-name> single-space command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "menu ${1:menu-name} single-space", },
    },
    {
        Keyword: "menu clear-screen",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To clear the terminal screen before displaying a menu, use the menu clear-screen command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "menu clear-screen ${1:menu-name} clear-screen", },
    },
    {
        Keyword: "menu command",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify underlying commands for user menus, use the menu command command in global configuration mode. To return to default settings, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "menu command menu ${1:menu-name} command ${2:menu-item} ${3:command}menu-exit", },
    },
    {
        Keyword: "menu default",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the menu item to use as the default, use the menu default command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "menu ${1:menu-name} default ${2:menu-item}", },
    },
    {
        Keyword: "menu line-mode",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To require the user to press Enter after specifying an item, use the menu line-mode command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "menu ${1:menu-name} line-mode", },
    },
    {
        Keyword: "menu options",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set options for items in user menus, use the menu options command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "menu ${1:menu-name} options ${2:menu-item} login pause", "menu ${1:menu-name} options ${2:menu-item} loginpause", },
    },
    {
        Keyword: "menu prompt",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the prompt for a user menu, use the menu prompt command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "menu ${1:menu-name} prompt ${2:d} ${3:prompt} ${4:d}", },
    },
    {
        Keyword: "menu status-line",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To display a line of status information about the current user at the top of a menu, use the menu status-line command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "menu ${1:menu-name} status-line", },
    },
    {
        Keyword: "menu text",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify th e text of a menu item in a user menu, use the menu text command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "menu ${1:menu-name} text ${2:menu-item} ${3:menu-text}", },
    },
    {
        Keyword: "menu title",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To create a t itle (banner) for a user menu, use the menu title command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "menu ${1:menu-name} title d menu-title d", },
    },
    {
        Keyword: "microcode (12000)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To load a Cisco IOS software image on a line card from Flash memory or the GRP card on a Cisco 12000 series Gigabit Switch Router (GSR), use the microcode command in global configuration mode. To load the microcode bundled with the GRP system image, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "microcode oc12-atmoc12-posoc3-pos4 flash ${1:file-id} ${2: [slot] }system ${3: [slot] }", "no microcode oc12-atmoc12-posoc3-pos4 flash ${1:file-id} ${2: [slot] }system ${3: [slot] }", },
    },
    {
        Keyword: "microcode (7000/7500)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the location of the microcode that you want to download from Flash memory into the writable control store (WCS) on Cisco 7000 series (including RSP based routers) or Cisco 7500 series routers, use the microcode command in global configuration mode. To load the microcode bundled with the system image, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "microcode ${1:interface-type} ${2:flash-filesystem}:${3:filename} ${4: [slot] }romsystem ${5: [slot] }", "no microcode ${1:interface-type} ${2:flash-filesystem}:${3:filename} ${4: [slot] }romsystem ${5: [slot] }", },
    },
    {
        Keyword: "microcode (7200)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure a default override for the microcode that is downloaded to the hardware on a Cisco 7200 series router, use the microcode command in global configuration mode. To revert to the default microcode for the current running version of the Cisco IOS software, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "microcode ecpapcpa ${1:location}", "no microcode ecpapcpa", },
    },
    {
        Keyword: "microcode reload (12000)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To reload the Cisco IOS image from a line card on Cisco 12000 series routers, use the microcode reload command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "microcode reload ${1: [slot-number] }", },
    },
    {
        Keyword: "microcode reload (7000 7500)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To reload the processor card on the Cisco 7000 series with RSP7000 or Cisco 7500 series routers, use the microcode reload command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "microcode reload ${1: [slot-number] }", },
    },
    {
        Keyword: "mode",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the redundancy mode, use the mode command in redundancy configuration mode.",
        },
        Section: "config-red",
        Snippets: []string{ "mode rprrpr-plussso", "mode rprsso", "mode sso", },
    },
    {
        Keyword: "monitor event-trace (global)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure event tracing for a specified Cisco IOS software subsystem component, use the monitor event-trace command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "monitor event-trace ${1:component} disabledump-file ${2:filename}enablesize ${3:number}stacktrace ${4:number} timestamps datetime localtime msec show-timezoneuptime", "monitor event-trace ${1:component} disabledump-file ${2:filename}enableclearcontinuousone-shot", },
    },
    {
        Keyword: "monitor pcm-tracer capture-destination",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure a location to save the Pulse Code Modulation (PCM) trace information, use the monitor pcm-tracer capture-destination command in global configuration mode. To disable the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "monitor pcm-tracer capture-destination ${1:destination}", "no monitor pcm-tracer capture-destination", },
    },
    {
        Keyword: "monitor pcm-tracer delayed-start",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the delay time to start the Pulse Code Modulation (PCM) trace capture, use the monitor pcm-tracer delayed-start command in global configuration mode. To disable the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "monitor pcm-tracer delayed-start ${1:seconds}", "no monitor pcm-tracer delayed-start", },
    },
    {
        Keyword: "monitor pcm-tracer profile",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To create Pulse Code Modulation (PCM) capture profiles, use the monitor pcm-tracer profile command in global configuration mode. To disable the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "monitor pcm-tracer profile ${1:profile-number} no capture-tdmT1E1analog-voice-portbri-voice-port${2:port}ds0channel-num${3:number}", "no monitor pcm-tracer profile ${1:profile-number}", },
    },
    {
        Keyword: "monitor permit-list",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure a destination port permit list or add to an existing destination port permit list, use the monitor permit-list command in global configuration mode. To delete from or clear an existing destination port permit list, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "monitor permit-list", "no monitor permit-list", "monitor permit-list destination interface ${1:interface-type} ${2:slot}/${3:port}", "no monitor permit-list destination interface ${1:interface-type} ${2:slot}/${3:port}", "monitor permit-list destination interface ${1:interface-type} ${2:slot}/${3:port-last-port}", "no monitor permit-list destination interface ${1:interface-type} ${2:slot}/${3:port-last-port}", "monitor permit-list destination interface ${1:interface-type} ${2:slot}/${3:port-last-port} , ${4: [port-last-port] }", "no monitor permit-list destination interface ${1:interface-type} ${2:slot}/${3:port-last-port} , ${4: [port-last-port] }", },
    },
    {
        Keyword: "monitor session egress replication-mode",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To switch the egress-span mode from the default mode (either centralized or distributed depending on your Cisco IOS software release), use the monitor session egress replication-mode command in global configuration mode. To return to the default mode, use the no form of the command.",
        },
        Section: "config",
        Snippets: []string{ "monitor session egress replication-mode centralized", "no monitor session egress replication-mode centralized", "monitor session egress replication-mode distributed", "no monitor session egress replication-mode distributed", },
    },
    {
        Keyword: "monitor session type",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure a local Switched Port Analyzer (SPAN), RSPAN, or ERSPAN, use the monitor session type command in global configuration mode. To remove one or more source or destination interfaces from the SPAN session, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "monitor session ${1:span-session-number} type erspan-destinationerspan-sourcelocallocal-txrspan-destinationrspan-source", "no monitor session ${1:span-session-number} type erspan-destinationerspan-sourcelocallocal-txrspan-destinationrspan-source", },
    },
    {
        Keyword: "mop device-code",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To identify the type of device sending Maintenance Operation Protocol (MOP) System Identification (sysid) messages and request program messages, use the mop device-code command in global configuration mode. To set the identity to the default value, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "mop device-code commandmop device-code ciscods200", "no mop device-code ciscods200", },
    },
    {
        Keyword: "mop retransmit-timer",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the length of time that the Cisco IOS software waits before resending boot requests to a Maintenance Operation Protocol (MOP) server, use the mop retransmit-timer command in global configuration mode. To reinstate the default value, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "mop retransmit-timer ${1:seconds}", "no mop retransmit-timer", },
    },
    {
        Keyword: "mop retries",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the number of times the Cisco IOS software will resend boot requests to a Maintenance Operation Protocol (MOP) server, use the mop retries command in global configuration mode. To reinstate the default value, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "mop retries ${1:count}", "no mop retries", },
    },
    {
        Keyword: "motd-banner",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the display of message-of-the-day (MOTD) banners on the specified line or lines, use the motd-banner command in line configuration mode. To suppress the MOTD banners on the specified line or lines, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "motd-banner", "no motd-banner", },
    },
    {
        Keyword: "nmsp enable",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable Network Mobility Service Protocol (NMSP) features on the device, use the nmsp enable command in global configuration mode. To disable, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "nmsp enable", "no nmsp enable", },
    },
    {
        Keyword: "nmsp strong-cipher",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the new ciphers, use the nmsp strong-cipher command in global configuration mode. To disable, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "nmsp strong-cipher", "no nmsp strong-cipher", },
    },
    {
        Keyword: "no menu",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To delete a user menu from the configuration file, use the no menu command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "no menu ${1:menu-name}", },
    },
    {
        Keyword: "notify",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable terminal notification about pending output from other Telnet connections, use the notify command in line configuration mode. To disable notifications, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "notify", "no notify", },
    },
    {
        Keyword: "notify syslog",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the sending of notifications of configuration changes to a remote system message logging (syslog), use the notify syslog command in configuration change logger configuration mode. To disable the sending of notifications of configuration changes to the syslog, use the form of this command.",
        },
        Section: "config-archive-log-config",
        Snippets: []string{ "notify syslog contenttype plaintextxml", "no notify syslog contenttype plaintextxml", },
    },
    {
        Keyword: "padding",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the padding on a specific output character, use the padding command in line configuration mode. To remove padding for the specified output character, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "padding ${1:ascii-number} ${2:count}", "no padding ${1:ascii-number}", },
    },
    {
        Keyword: "parity",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define generation of a parity bit, use the parity command in line configuration mode. To specify no parity, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "parity noneevenoddspacemark", "no parity", },
    },
    {
        Keyword: "parser cache",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To reenable the Cisco software parser cache after disabling it, use the parser cache command in global configuration mode. To disable the parser cache, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "parser cache", "no parser cache", },
    },
    {
        Keyword: "parser command serializer",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable configuration access only to the users holding a configuration lock and to prevent other clients from accessing the running configuration, use the parser command serializer command in global configuration mode. To disable this configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "parser command serializer", "no parser command serializer", },
    },
    {
        Keyword: "parser config cache interface",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To reduce the time required for the command-line interpreter to execute commands that manage the running system configuration files, use the parser config cache interface command in global configuration mode. To disable the reduced command execution time functionality, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "parser config cache interface", "no parser config cache interface", },
    },
    {
        Keyword: "parser config partition",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable configuration partitioning, use the parser config partition command. To disable the partitioning of the running configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "parser config partition", "no parser config partition", },
    },
    {
        Keyword: "parser maximum",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify performance maximums for CLI operations use the parser maximum command in global configuration mode. To clear any previously established maximums, us the No form of the command.",
        },
        Section: "config",
        Snippets: []string{ "parser maximum latency${1:limit}utilization${2:limit}", "no parser maximum latencyutilization", },
    },
    {
        Keyword: "partition",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To separate Flash memory into partitions on Class B file system platforms, use the partition command in global configuration mode. To undo partitioning and to restore Flash memory to one partition, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "partition ${1:flash-filesystem}: ${2: [number-of-partitions] } ${3: [partition-size] }", "no partition ${1:flash-filesystem}:", "partition flash ${1:partitions} ${2:size1} ${3:size2}", "no partition flash", },
    },
    {
        Keyword: "path (archive configuration)",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the location and filename prefix for the files in the Cisco configuration archive, use the path command in archive configuration mode. To disable this function, use the no form of this command.",
        },
        Section: "config-archive",
        Snippets: []string{ "path ${1:url}", "no path ${1:url}", },
    },
    {
        Keyword: "periodic",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify a recurring (weekly) time range for functions that support the time-range feature, use the periodic command in time-range configuration mode. To remove the time limitation, use the no form of this command.",
        },
        Section: "config-time-range",
        Snippets: []string{ "periodic ${1:days-of-the-week} ${2:hh:mm} to ${3: [days-of-the-week] } ${4:hh:mm}", "no periodic ${1:days-of-the-week} ${2:hh:mm} to ${3: [days-of-the-week] } ${4:hh:mm}", },
    },
    {
        Keyword: "platform qfp drops threshold",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the warning thresholds for per drop cause and/or total QFP drop in packets per second, use the platform qfp drops threshold command.",
        },
        Section: "config",
        Snippets: []string{ "platform qfp drops threshold per-cause ${1: | drop_id} ${2: | threshold_value}total ${3:threshold_value}", },
    },
    {
        Keyword: "platform shell",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To grant shell access and enter shell access grant configuration mode, use the platform shell command in global configuration mode. To disable this function, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "platform shell", "no platform shell", },
    },
    {
        Keyword: "power enable",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To turn on power for the modules, use the power enable command in global configuration mode. To power down a module, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "power enable module ${1:slot}", "no power enable module ${1:slot}", },
    },
    {
        Keyword: "power redundancy-mode",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the power-supply redundancy mode, use the power redundancy-mode command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "power redundancy-mode combinedredundant", },
    },
    {
        Keyword: "printer",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure a printer and assign a server tty line (or lines) to it, use the printer command in global configuration mode. To disable printing on a tty line, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "printer ${1:printer-name} line ${2:number}rotary ${3:number} formfeed jobtimeout ${4:seconds} newline-convert jobtypes ${5:type}", "no printer ${1:printer-name}", },
    },
    {
        Keyword: "private",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To save user EXEC command changes between terminal sessions, use the private command in line configuration mode. To restore the default condition, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "private", "no private", },
    },
    {
        Keyword: "process cpu statistics limit entry-percentage",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the process entry limit and the size of the history table for CPU utilization statistics, use the process cpu statistics limit entry-percentage command in global configuration mode. To disable CPU utilization statistics, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "process cpu statistics limit entry-percentage ${1:number} size ${2:seconds}", "no process cpu statistics limit entry-percentage", },
    },
    {
        Keyword: "process cpu threshold type",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set CPU thresholding notification types and values, use the process cpu threshold type command in global configuration mode. To disable CPU thresholding notifications, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "process cpu threshold type totalprocessinterrupt rising ${1:percentage} interval ${2:seconds} falling ${3:fall-percentage} interval ${4:seconds}", "no process cpu threshold type totalprocessinterrupt", },
    },
    {
        Keyword: "process-max-time",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the amount of time after which a process should voluntarily yield to another process, use the process-max-time command in global configuration mode. To reset this value to the system default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "process-max-time ${1:milliseconds}", "no process-max-time ${1:milliseconds}", },
    },
    {
        Keyword: "prompt",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To customiz e the CLI prompt, use the prompt command in global configuration mode. To revert to the default prompt, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "prompt ${1:string}", "no prompt ${1: [string] }", },
    },
    {
        Keyword: "prompt config",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the system’s prompt for configuration mode, use the prompt config command in global configuration mode. To disable the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "prompt config hostname-length ${1:number}", "no prompt config", },
    },
    {
        Keyword: "refuse-message",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To define and enable a line-in-use message, use the refuse-message command in line configuration mode. To disable the message, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "refuse-message ${1:d} ${2:message} ${3:d}", "no refuse-message", },
    },
    {
        Keyword: "regexp optimize",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To optimize the compilation of a regular expression access list, use the regexp optimize command in global configuration mode. To disable the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "regexp optimize", "no regexp optimize", },
    },
    {
        Keyword: "remote-span",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure a virtual local area network (VLAN) as a remote switched port analyzer (RSPAN) VLAN, use the remote-span command in config-VLAN mode. To remove the RSPAN designation, use the no form of this command.",
        },
        Section: "config-vlan",
        Snippets: []string{ "remote-span", "no remote-span", },
    },
    {
        Keyword: "revision",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the revision number for the Multiple Spanning Tree (802.1s) (MST) configuration, use the revision command in MST configuration submode. To return to the default settings, use the no form of this command.",
        },
        Section: "config-mst",
        Snippets: []string{ "revision ${1:version}", "no revision", },
    },
    {
        Keyword: "scheduler allocate",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To guarantee CPU time for processes, use the scheduler allocate command in global configuration mode. To restore the default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler allocate ${1:interrupt-time} ${2:process-time}", "no scheduler allocate", },
    },
    {
        Keyword: "scheduler heapcheck enable",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable heapcheck processing, use the scheduler heapcheck enable command in global configuration mode. To disable scheduler heapcheck processing, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler heapcheck enable", "no scheduler heapcheck enable", },
    },
    {
        Keyword: "scheduler heapcheck poll",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To validate the memory and edisms poll routine, use the scheduler heapcheck poll command in global configuration mode. To disable the memory check and edisms poll routine, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler heapcheck poll", "no scheduler heapcheck poll", },
    },
    {
        Keyword: "scheduler heapcheck process",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To perform a “sanity check” for corruption in memory blocks when a process switch occurs, use the scheduler heapcheck process command in global configuration mode. To disable this feature, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler heapcheck process memory fast io multibus pci processor checktype alldatamagicmlite-datapointerrefcountlite-chunks", "no scheduler heapcheck process", },
    },
    {
        Keyword: "scheduler interrupt mask profile",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To start interrupt mask profiling for all processes running on the system, use the scheduler interrupt mask profile command in global configuration mode. To stop interrupt mask profiling, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler interrupt mask profile", "no scheduler interrupt mask profile", },
    },
    {
        Keyword: "scheduler interrupt mask size",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the maximum number of entries that can exist in the interrupt mask buffer, use the scheduler interrupt mask size command in global configuration mode. To reset the maximum number of entries that can exist in the interrupt mask buffer to the default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler interrupt mask size ${1:buffersize}", "no scheduler interrupt mask size", },
    },
    {
        Keyword: "scheduler interrupt mask time",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the maximum time that a process can run with interrupts masked before another entry is created in the interrupt mask buffer, use the scheduler interrupt mask time command in global configuration mode. To reset the threshold time to the default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler interrupt mask time ${1:threshold-time}", "no scheduler interrupt mask time", },
    },
    {
        Keyword: "scheduler interval",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To control the maximum amount of time that can elapse without running system processes, use the scheduler interval command in global configuration mode. To restore the default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler interval ${1:milliseconds}", "no scheduler interval", },
    },
    {
        Keyword: "scheduler isr-watchdog",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To detect if an Interrupt Service Routine (ISR) is suspended or stalled and to schedule and manage a watchdog timeout on an ISR, use the scheduler isr-watchdog command in global configuration mode. To disable the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler isr-watchdog", "no scheduler isr-watchdog", },
    },
    {
        Keyword: "scheduler max-sched-time",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure or change the maximum time, in milliseconds that a scheduler can run without flagging an error or overload of the CPU, use the scheduler max-sched-time command in global configuration mode. To disable this configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler max-sched-time ${1:milliseconds}", "no scheduler max-sched-time", },
    },
    {
        Keyword: "scheduler process-watchdog",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the default action of a watchdog timeout for a process using a scheduler, use the scheduler process-watchdog command in global configuration mode. To disable the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler process-watchdog hangnormalreloadterminate", "no scheduler process-watchdog", },
    },
    {
        Keyword: "scheduler timercheck process",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure process-level timer validation on a scheduler, and check the timer tree of the process after every context switch of the process Packet Identification number (PID) is configured, use the scheduler timercheck process command in global configuration mode. To disable this configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler timercheck process ${1:pid}", "no scheduler timercheck process ${1:pid}", },
    },
    {
        Keyword: "scheduler timercheck system context",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure system-level validation on context switches on a scheduler, and check system level-timers, use the scheduler timercheck system context command in global configuration mode. To disable the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "scheduler timercheck system context", "no scheduler timercheck system context", },
    },
    {
        Keyword: "service compress-config",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To compress startup configuration files, use the service compress-config command in global configuration mode. To disable compression, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service compress-config", "no service compress-config ", },
    },
    {
        Keyword: "service config",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable autoloading of configuration files from a network server, use the service config command in global configuration mode. To restore the default, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service config", "no service config", },
    },
    {
        Keyword: "service counters max age",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the time interval for retrieving statistics, use the service counters max age command in global configuration mode. To return to the default settings, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service counters max age ${1:seconds}", "no service counters max age", },
    },
    {
        Keyword: "service decimal-tty",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify that line numbers be displayed and interpreted as octal numbers rather than decimal numbers, use the no service decimal-tty command in global configuration mode. To restore the default, use the service decimal-tty command.",
        },
        Section: "config",
        Snippets: []string{ "service decimal-tty", "no service decimal-tty", },
    },
    {
        Keyword: "service exec-wait",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To delay the startup of the EXEC on noisy lines, use the service exec-wait command in global configuration mode. To disable the delay function, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service exec-wait", "no service exec-wait", },
    },
    {
        Keyword: "service hide-telnet-address",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To hide addresses while trying to establish a Telnet session, use the service hide-telnet-address command in global configuration mode. To disable this service, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service hide-telnet-address", "no service hide-telnet-address", },
    },
    {
        Keyword: "service linenumber",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the Cisco IOS software to display line number information after the EXEC or incoming banner, use the service linenumber command in global configuration mode. To disable this function, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service linenumber", "no service linenumber", },
    },
    {
        Keyword: "service nagle",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable the Nagle congestion control algorithm, use the service nagle command in global configuration mode. To disable the algorithm, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service nagle ", "no service nagle", },
    },
    {
        Keyword: "service prompt config",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To display the configuration prompt (config), use the service prompt config command in global configuration mode. To remove the configuration prompt, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service prompt config", "no service prompt config", },
    },
    {
        Keyword: "service sequence-numbers",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable visible sequence numbering of system logging messages, use the service sequence-numbers command in global configuration mode. To disable visible sequence numbering of logging messages, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service sequence-numbers", "no service sequence-numbers", },
    },
    {
        Keyword: "service slave-log",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To allow secondary Versatile Interface Processor (VIP) cards to log important error messages to the console, use the service slave-log command in global configuration mode . To disable secondary logging, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service slave-log", "no service slave-log", },
    },
    {
        Keyword: "service tcp-keepalives-in",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To generate keepalive packets on idle incoming network connections (initiated by the remote host), use the service tcp-keepalives-in command in global configuration mode . To disable the keepalives, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service tcp-keepalives-in", "no service tcp-keepalives-in", },
    },
    {
        Keyword: "service tcp-keepalives-out",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To generate keepalive packets on idle outgoing network connections (initiated by a user), use the service tcp-keepalives-out command in global configuration mode . To disable the keepalives, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service tcp-keepalives-out ", "no service tcp-keepalives-out", },
    },
    {
        Keyword: "service tcp-small-servers",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable small TCP servers such as the Echo, use the service tcp-small-servers command in global configuration mode. To disable the TCP server, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service tcp-small-servers max-servers ${1:number}no-limit", "no service tcp-small-servers max-servers ${1:number}no-limit", },
    },
    {
        Keyword: "service telnet-zeroidle",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the TCP window to zero (0) when the Telnet connection is idle, use the service telnet-zeroidle command in global configuration mode. To disable this service, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service telnet-zero-idle  ", "no service telnet-zeroidle", },
    },
    {
        Keyword: "service timestamps",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the system to apply a time stamp to debugging messages or system logging messages, use the service timestamps command in global configuration mode . To disable this service, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service timestamps debuglog uptimedatetime msec localtime show-timezone year", "no service timestamps debuglog", },
    },
    {
        Keyword: "service udp-small-servers",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable small User Datagram Protocol (UDP) servers such as the Echo, use the service udp-small-servers command in global configuration mode. To disable the UDP server, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "service udp-small-servers max-servers ${1:number}no-limit", "no service udp-small-servers max-servers ${1:number}no-limit", },
    },
    {
        Keyword: "service-module apa traffic-management",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure traffic management on the router, use the service-module apa traffic-management command in interface configuration mode.",
        },
        Section: "config-if",
        Snippets: []string{ "service-module apa traffic-management monitorinline", },
    },
    {
        Keyword: "show",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To verify the Multiple Spanning Tree (MST) configuration, use the show command in MST configuration mode.",
        },
        Section: "config-mst",
        Snippets: []string{ "show currentpending", },
    },
    {
        Keyword: "show declassify",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To display the state of the declassify function (enabled, in progress, and so forth) and the sequence of declassification steps that will be performed, use the show declassify command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "show declassify", },
    },
    {
        Keyword: "slave auto-sync config",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To turn on automatic synchronization of configuration files for a Cisco 7507 or Cisco 7513 router that is configured for High System Availability (HSA) using Dual RSP Cards, use the slave auto-sync config global configuration command. To turn off automatic synchronization, use the no form of the command.",
        },
        Section: "config",
        Snippets: []string{ "slave auto-sync config", "no slave auto-sync config", },
    },
    {
        Keyword: "slave default-slot",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the default secondary Route Switch Processor (RSP) card on a Cisco 7507 or Cisco 7513 router, use the slave default-slot global configuration command.",
        },
        Section: "config",
        Snippets: []string{ "slave default-slot ${1:processor-slot-number}", },
    },
    {
        Keyword: "slave image",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the image that the secondary Route Switch Processor (RSP) runs on a Cisco 7507 or Cisco 7513 router, use the slave image command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ "slave image system${1: | file-url}", },
    },
    {
        Keyword: "slave reload",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To force a reload of the image that the secondary Route Switch Processor (RSP) card is running on a Cisco 7507 or Cisco 7513 router, use the slave reload global configuration command.",
        },
        Section: "config",
        Snippets: []string{ "slave reload", },
    },
    {
        Keyword: "slave terminal",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable access to the secondary Route Switch Processor (RSP) console, use the slave terminal global configuration command. To disable access to the secondary RSP console, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "slave terminal", "no slave terminal", },
    },
    {
        Keyword: "software source list",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To create a list of input bundles or directories, use the software source list command in global configuration mode.",
        },
        Section: "config",
        Snippets: []string{ " software source list \n                                 \t\t\t\t ${1: list-name-string \n                                 \t\t\t\t}", "no software source list \n                                 \t\t\t\t ${1: list-name-string \n                                 \t\t\t\t}", },
    },
    {
        Keyword: "special-character-bits",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the number of data bits per character for special characters such as software flow control characters and escape characters, use the special-character-bits command in line configuration mode. To restore the default value, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "special-character-bits 78", "no special-character-bits", },
    },
    {
        Keyword: "stack-mib portname",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify a name string for a port, use the stack-mib portname command in interface configuration mode.",
        },
        Section: "config-if",
        Snippets: []string{ "stack-mib portname ${1:portname}", },
    },
    {
        Keyword: "state-machine",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the transition criteria for the state of a particular state machine, use the state-machine command in global configuration mode . To remove a particular state machine from the configuration, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "state-machine ${1:name} ${2:state} ${3:first-character} ${4:last-character} ${5:next-state} delaytransmit", "no state-machine ${1:name}", },
    },
    {
        Keyword: "stopbits",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the number of the stop bits transmitted per byte, use the stopbits command in line configuration mode. To restore the default value, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "stopbits 11. 52", "no stopbits", },
    },
    {
        Keyword: "storm-control level",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the suppression level, use the storm-control level command in interface configuration mode. To turn off the suppression mode, use the no form of this command.",
        },
        Section: "config-if",
        Snippets: []string{ "storm-control broadcastmulticastunicast level ${1:level} . ${2:level}", "no storm-control broadcastmulticastunicast level", },
    },
    {
        Keyword: "sync-restart-delay",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the synchronization-restart delay timer to ensure accurate status reporting, use the sync-restart-delay command in interface configuration mode. To disable the synchronization-restart delay timer, use the no form of this command.",
        },
        Section: "config-if",
        Snippets: []string{ "sync-restart-delay ${1:timer}", "no sync-restart-delay ${1:timer}", },
    },
    {
        Keyword: "system flowcontrol bus",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the FIFO overflow error count, use the system flowcontrol bus command in global configuration mode. To return to the original FIFO threshold settings, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "default system flowcontrol bus autoon", "no system flowcontrol bus", },
    },
    {
        Keyword: "system jumbomtu",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the maximum size of the Layer 2 and Layer 3 packets, use the system jumbo mtu command in global configuration mode. To revert to the default MTU setting, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "system jumbomtu mtu-size", "no system jumbomtu", },
    },
    {
        Keyword: "tdm clock priority",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the clock source and priority of the clock source used by the time-division multiplexing (TDM) bus on the Cisco AS5350, AS5400, and AS5850 access servers, use the tdm clock priority command in global configuration mode. To return the clock source and priority to the default values, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "tdm clock priority ${1:priority-number} ${2:slot}/${3:ds1-port}${4:slot}/${5:ds3-port}:${6:ds1-port}externalfreerun", "no tdm clock priority ${1:priority-number} ${2:slot}/${3:ds1-port}${4:slot}/${5:ds3-port}:${6:ds1-port}externalfreerun", },
    },
    {
        Keyword: "terminal-queue entry-retry-interval",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To change the retry interval for a terminal port queue, use the terminal-queue entry-rety-interval command in global configuration mode. To restore the default terminal port queue interval, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "terminal-queue entry-retry-interval ${1:seconds}", "no terminal-queue entry-retry-interval", },
    },
    {
        Keyword: "terminal-type",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To specify the type of terminal connected to a line, use the terminal-type command in line configuration mode. To remove any information about the type of terminal and reset the line to the default terminal emulation, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "terminal-type ${1:terminal-name}${2: | terminal-type}", "no terminal-type", },
    },
    {
        Keyword: "tftp-server",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure a router or a Flash memory device on the router as a TFTP server, use one of the following tftp-server commands in global configuration mode. This command replaces the tftp-server system command. To remove a previously defined filename, use the no form of this command with the appropriate filename.",
        },
        Section: "config",
        Snippets: []string{ "tftp-server flash ${1:partition-number} : ${2:filename1} alias ${3:filename2} ${4: [access-list-number] }", "tftp-server rom alias ${1:filename1} ${2: [access-list-number] }", "no tftp-server flash ${1:partition-number} : ${2:filename1}rom alias ${3:filename2}", "tftp-server flash ${1:device} : ${2:partition-number} : ${3:filename}", "no tftp-server flash ${1:device} : ${2:partition-number} : ${3:filename}", "tftp-server flash ${1:device} : ${2:filename}", "no tftp-server flash ${1:device} : ${2:filename}", },
    },
    {
        Keyword: "time-period",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the time increment for automatically saving an archive file of the current running configuration in the Cisco configuration archive, use the time-period command in archive configuration mode. To disable this function, use the no form of this command.",
        },
        Section: "config-archive",
        Snippets: []string{ "time-period ${1:minutes}", "no time-period ${1:minutes}", },
    },
    {
        Keyword: "vacant-message",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To display an idle terminal message, use the vacant-message command in line configuration mode. To remove the default vacant message or any other vacant message that may have been set, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "vacant-message ${1:d} ${2:message} ${3:d}", "no vacant-message", },
    },
    {
        Keyword: "vtp",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To configure the global VLAN Trunking Protocol (VTP) state, use the vtp command in global configuration mode. To return to the default value, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "vtp domain ${1:domain-name}file ${2:filename}interface ${3:interface-name} onlymode clientoffservertransparent vlanmstunknownpassword ${4:password-value} hiddensecretpruningversion 123", "no vtp", },
    },
    {
        Keyword: "warm-reboot",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To enable a router to do a warm-reboot, use the warm-reboot command in global configuration mode. To disable warm rebooting, use the no form of this command.",
        },
        Section: "config",
        Snippets: []string{ "warm-reboot count ${1:number} uptime ${2:minutes}", "no warm-reboot count ${1:number} uptime ${2:minutes}", },
    },
    {
        Keyword: "width",
        Description: keyword.Description{
            Format: protocol.PlainText,
            Value:  "To set the terminal screen width, use the width command in line configuration mode. To return to the default screen width, use the no form of this command.",
        },
        Section: "config-line",
        Snippets: []string{ "width ${1:characters}", "no width", },
    },
}
