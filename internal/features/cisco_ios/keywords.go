package cisco_ios

import (
	"github.com/dgethings/chunter/internal/keyword"
	"github.com/dgethings/chunter/internal/protocol"
)

var Keywords = keyword.Keywords{
	{
		Keyword: "activation-character",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line console\nRouter(config-line)# activation-character 127",
		},
		Section: "config-line",
	},
	{
		Keyword: "alias",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)#alias exec fixmyrt clear ip route 192.168.116.16\n\nRouter#configure terminal\nEnter configuration commands, one per line.  End with CNTL/Z.\nRouter(config)#interface e0\nRouter(config-if)#?\nInterface configuration commands:\n  access-expression        Build a bridge boolean access expression\n  .\n  .\n  .\nRouter(config-if)#exit\nRouter(config)#alias ?\n  accept-dialin         VPDN group accept dialin configuration mode\n  accept-dialout        VPDN group accept dialout configuration mode\n  address-family        Address Family configuration mode\n  call-discriminator    Call Discriminator Configuration\n  cascustom             Cas custom configuration mode\n  clid-group            CLID group configuration mode\n  configure             Global configuration mode\n  congestion            Frame Relay congestion configuration mode\n  controller            Controller configuration mode\n  cptone-set            custom call progress tone configuration mode\n  customer-profile      customer profile configuration mode\n  dhcp                  DHCP pool configuration mode\n  dnis-group            DNIS group configuration mode\n  exec                  Exec mode\n  flow-cache            Flow aggregation cache config mode\n  fr-fr                 FR/FR connection configuration mode\n  interface             Interface configuration mode\n .\n .\n .\nRouter(config)#alias interface express access-expression\nRouter(config)#int e0\nRouter(config-if)#exp?\n*express=access-expression  \nRouter(config-if)#express ?\n  input   Filter input packets\n  output   Filter output packets\n!Note that the true form of the command/keyword alias appears on the screen after issuing\n!the express ? command.\nRouter(config-if)#access-expression ?\n  input   Filter input packets\n  output  Filter output packets\nRouter(config-if)#ex?\n*express=access-expression  exit  \n!Note that in the following line, a space is used before the ex? command\n!so the alias is not displayed.\nRouter(config-if)# ex?\nexit\n!Note that in the following line, the alias cannot be recognized because\n!a space is used before the command.\nRouter#(config-if)# express ?\n% Unrecognized command\nRouter(config-if)# end \nRouter# show alias interface\nInterface configuration mode aliases:\n  express               access-expression",
		},
		Section: "config",
	},
	{
		Keyword: "archive",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device# configure terminal\n!\nDevice(config)# archive\nDevice(config-archive)#",
		},
		Section: "config",
	},
	{
		Keyword: "async-bootp",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "async-bootp bootfile :172.30.1.1 “pcboot”\nasync-bootp bootfile :mac “macboot”\n\nasync-bootp subnet-mask 255.255.0.0\n\nasync-bootp time-offset -3600\n\nasync-bootp time-server 172.16.1.1",
		},
		Section: "config",
	},
	{
		Keyword: "autobaud",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line aux\n \nRouter(config-line)# autobaud",
		},
		Section: "config-line",
	},
	{
		Keyword: "auto-sync",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "",
		},
		Section: "config-r",
	},
	{
		Keyword: "autoupgrade disk-cleanup",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device(config)# autoupgrade disk-cleanup irrecoverable\n\nDevice(config)# no autoupgrade disk-cleanup image",
		},
		Section: "config",
	},
	{
		Keyword: "autoupgrade ida url",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device(config)# autoupgrade ida url https://www.cisco.com/cgi-bin/ida/locator/locator.pl",
		},
		Section: "config",
	},
	{
		Keyword: "autoupgrade status email",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device(config)# autoupgrade status email recipient tree@abc.com\nDevice(config)# autoupgrade status email smtp-server smtpserver.abc.com",
		},
		Section: "config",
	},
	{
		Keyword: "banner exec",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# banner exec %\nEnter TEXT message.  End with the character '%'.\nSession activated on line $(line), $(line-desc). Enter commands at the prompt. \n% \n\nUser Access Verification\nUsername: joeuser\nPassword: <password>\n Session activated on line 50, vty default line. Enter commands at the prompt.\nRouter>",
		},
		Section: "config",
	},
	{
		Keyword: "banner incoming",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# banner incoming #\nThis is the Reuses router.\n#\n\ndarkstar(config)# banner incoming %\nEnter TEXT message.  End with the character '%'.\nYou have entered $(hostname).$(domain) on line $(line) ($(line-desc)) %\n \nYou have entered darkstar.ourdomain.com on line 5 (Dialin Modem)",
		},
		Section: "config",
	},
	{
		Keyword: "banner login",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# banner login \" Access for authorized users only. Please enter your username and password. \"\n\ndarkstar(config)# banner login %\nEnter TEXT message. End with the character '%'.\nYou have entered $(hostname).$(domain) on line $(line) ($(line-desc)) %\n \nYou have entered darkstar.ourdomain.com on line 5 (Dialin Modem)",
		},
		Section: "config",
	},
	{
		Keyword: "banner motd",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# banner motd # Building power will be off from 7:00 AM until 9:00 AM this coming Tuesday. \n\ndarkstar(config)# banner motd %\nEnter TEXT message.  End with the character '%'.\nNotice: all routers in $(domain) will be upgraded beginning April 20\n%\n \nNotice: all routers in ourdomain.com will be upgraded beginning April 20",
		},
		Section: "config",
	},
	{
		Keyword: "banner slip-ppp",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# banner slip-ppp %\nEnter TEXT message.  End with the character '%'.\nStarting $(encap) connection from $(gate-ip) to $(peer-ip) using a maximum packet size of $(mtu) bytes... %\n\nRouter# slip\nStarting SLIP connection from 172.16.69.96 to 192.168.1.200 using a maximum packet size of 1500 bytes...",
		},
		Section: "config",
	},
	{
		Keyword: "boot bootldr",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "boot bootldr bootflash:boot-image\n\nboot bootldr slot0:boot-image",
		},
		Section: "config",
	},
	{
		Keyword: "boot bootstrap",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# boot bootstrap bootflash:sysimage-2",
		},
		Section: "config",
	},
	{
		Keyword: "boot config",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# boot config flash:router-config\nRouter(config)# end\nRouter# copy system:running-config nvram:startup-config\n\nRouter (config)# boot config slot1:router-config\nRouter (config)# end\nRouter# copy system:running-config nvram:startup-config",
		},
		Section: "config",
	},
	{
		Keyword: "boot host",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# boot host tftp://192.168.7.19/usr/local/tftpdir/wilma-confg\nRouter(config)# service config",
		},
		Section: "config",
	},
	{
		Keyword: "boot network",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# boot network tftp:bridge_9.1\nRouter(config)# service config\n\nRouter(config)# service config\nRouter(config)# boot network rcp://172.16.1.111/bridge_9.1",
		},
		Section: "config",
	},
	{
		Keyword: "boot system",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# boot system tftp://192.168.7.24/cs3-rx.90-1\n \nRouter(config)# boot system tftp://192.168.7.19/cs3-rx.83-2\n \nRouter(config)# boot system rom\n\nRouter(config)# boot system flash:2:igs-bpx-l\n\nRouter(config)# boot system slot0:new-config\n\nRouter(config)# boot system slot0:4:dirt/images/new-ios-image\n\nRouter(config)# boot system flash:2:c1600-y-l",
		},
		Section: "config",
	},
	{
		Keyword: "clock",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-if)# clock active\nRouter(config-if)#",
		},
		Section: "config-if",
	},
	{
		Keyword: "clock initialize nvram",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# clock initialize nvram",
		},
		Section: "config",
	},
	{
		Keyword: "config-register",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "config-register 0x2102",
		},
		Section: "config",
	},
	{
		Keyword: "configuration mode exclusive",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device# configure terminal \nEnter configuration commands, one per line. End with CNTL/Z.\nDevice(config)# configuration mode exclusive auto \nDevice(config)# end\nDevice# show running-configuration\n | include config\n\n\nBuilding configuration...\nCurrent configuration : 2296 bytes\nconfiguration mode exclusive auto <========== auto policy\nDevice# configure terminal ?\n <======== lock option not displayed when in auto policy\nDevice# configure terminal\n <======= acquires the lock\n\nEnter configuration commands, one per line. End with CNTL/Z.\nDevice(config)# \nDevice(config)# show configuration lock\n\n\nParser Configure Lock\n---------------------\nOwner PID : 3\nUser : unknown\nTTY : 0\nType : EXCLUSIVE\nState : LOCKED\nClass : EXPOSED\nCount : 1\nPending Requests : 0\nUser debug info : configure terminal \nSession idle state : TRUE\nNo of exec cmds getting executed : 0\nNo of exec cmds blocked : 0\nConfig wait for show completion : FALSE\nRemote ip address : Unknown\nLock active time (in Sec) : 6\nLock Expiration timer (in Sec) : 593\nDevice(config)#\nDevice(config)# end\n <========= releases the lock\nDevice#\nDevice# show configuration lock\n\n\nParser Configure Lock\n---------------------\nOwner PID : -1\nUser : unknown\nTTY : -1\nType : NO LOCK\nState : FREE\nClass : unknown\nCount : 0\nPending Requests : 0\nUser debug info : \nSession idle state : TRUE\nNo of exec cmds getting executed : 0\nNo of exec cmds blocked : 0\nConfig wait for show completion : FALSE\nRemote ip address : Unknown\nLock active time (in Sec) : 0\nLock Expiration timer (in Sec) : 0\n\nDevice# configure terminal\n \nConfiguration mode locked exclusively. The lock will be cleared once you exit out of configuration mode using end/exit\nEnter configuration commands, one per line. End with CNTL/Z.\nDevice(config)# configuration mode exclusive manual\n \nDevice(config)# end\n\nDevice#\nDevice# show running-configuration\n | include configuration\n\nBuilding configuration...\nCurrent configuration : 2298 bytes\nconfiguration mode exclusive manual <==== 'manual' policy\nDevice# show configuration lock\n\nParser Configure Lock\n---------------------\nOwner PID : -1\nUser : unknown\nTTY : -1\nType : NO LOCK\nState : FREE\nClass : unknown\nCount : 0\nPending Requests : 0\nUser debug info : \nSession idle state : TRUE\nNo of exec cmds getting executed : 0\nNo of exec cmds blocked : 0\nConfig wait for show completion : FALSE\nRemote ip address : Unknown\nLock active time (in Sec) : 0\nLock Expiration timer (in Sec) : 0\nDevice#\nDevice# configure terminal ?\n \nlock Lock configuration mode <========= 'lock' option displayed in 'manual' policy\nDevice# configure terminal <============ ‘configure terminal’ won't acquire lock automatically\nEnter configuration commands, one per line. End with CNTL/Z.\nDevice(config)# show configuration lock\n \nParser Configure Lock\n---------------------\nOwner PID : -1\nUser : unknown\nTTY : -1\nType : NO LOCK\nState : FREE\nClass : unknown\nCount : 0\nPending Requests : 0\nUser debug info : \nSession idle state : TRUE\nNo of exec cmds getting executed : 0\nNo of exec cmds blocked : 0\nConfig wait for show completion : FALSE\nRemote ip address : Unknown\nLock active time (in Sec) : 0\nLock Expiration timer (in Sec) : 0\nDevice(config)# end\n \nDevice# show configuration lock \n\n\nParser Configure Lock\n---------------------\nOwner PID : -1\nUser : unknown\nTTY : -1\nType : NO LOCK\nState : FREE\nClass : unknown\nCount : 0\nPending Requests : 0\nUser debug info : \nSession idle state : TRUE\nNo of exec cmds getting executed : 0\nNo of exec cmds blocked : 0\nConfig wait for show completion : FALSE\nRemote ip address : Unknown\nLock active time (in Sec) : 0\nLock Expiration timer (in Sec) : 0\nDevice#\nDevice# configure\n \nDevice# configure terminal \n\nDevice# configure terminal ?\n \nlock Lock configuration mode <======= 'lock' option displayed when in 'manual' policy\nDevice# configure terminal lock\n \nDevice# configure terminal lock\n <============ acquires exclusive configuration lock\n\nDevice(config)# show configuration lock\n \nParser Configure Lock\n---------------------\nOwner PID : 3\nUser : unknown\nTTY : 0\nType : EXCLUSIVE\nState : LOCKED\nClass : EXPOSED\nCount : 1\nPending Requests : 0\nUser debug info : configure terminal lock\nSession idle state : TRUE\nNo of exec cmds getting executed : 0\nNo of exec cmds blocked : 0\nConfig wait for show completion : FALSE\nRemote ip address : Unknown\nLock active time (in Sec) : 5\nLock Expiration timer (in Sec) : 594\nDevice(config)# end\n <================ 'end' releases exclusive configuration lock\nDevice# show configuration lock \n\n\nParser Configure Lock\n---------------------\nOwner PID : -1\nUser : unknown\nTTY : -1\nType : NO LOCK\nState : FREE\nClass : unknown\nCount : 0\nPending Requests : 0\nUser debug info : \nSession idle state : TRUE\nNo of exec cmds getting executed : 0\nNo of exec cmds blocked : 0\nConfig wait for show completion : FALSE\nRemote ip address : Unknown\nLock active time (in Sec) : 0\nLock Expiration timer (in Sec) : 0\nDevice#",
		},
		Section: "config",
	},
	{
		Keyword: "databits",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 4\nRouter(config-line)# databits 7",
		},
		Section: "config-line",
	},
	{
		Keyword: "data-character-bits",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line vty 1\nRouter(config-line)# data-character-bits 7",
		},
		Section: "config-line",
	},
	{
		Keyword: "default-value data-character-bits",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# default-value data-character-bits 8",
		},
		Section: "config",
	},
	{
		Keyword: "default-value exec-character-bits",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# default-value exec-character-bits 8",
		},
		Section: "config",
	},
	{
		Keyword: "default-value modem-interval",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# default-value modem-signal 345",
		},
		Section: "config",
	},
	{
		Keyword: "default-value special-character-bits",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# default-value special-character-bits 8",
		},
		Section: "config",
	},
	{
		Keyword: "define interface-range",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device(config)# define interface-range macro1 ethernet 1/2 - 5, fastethernet 5/5 - 10",
		},
		Section: "config",
	},
	{
		Keyword: "diagnostic bootup level",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \ndiagnostic bootup level complete",
		},
		Section: "config",
	},
	{
		Keyword: "diagnostic cns",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \ndiagnostic cns publish my.cns.publish\nRouter(config)#\n\nRouter(config)# \ndiagnostic cns subscribe my.cns.subscribe\nRouter(config)#\n\nRouter(config)# \ndefault\ndiagnostic cns publish\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "diagnostic event-log size",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \ndiagnostic event-log size 600",
		},
		Section: "config",
	},
	{
		Keyword: "diagnostic monitor",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# diagnostic monitor interval module 5 test 7 09:07:05 45 3\n\nRouter(config)# \ndiagnostic monitor syslog",
		},
		Section: "config",
	},
	{
		Keyword: "diagnostic schedule module",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# diagnostic schedule module 1 test 5 on may 27 2010 10:30\n\nRouter(config)# diagnostic schedule module 1 test 5 daily 12:25\nRouter(config)# diagnostic schedule module 1 test 5 weekly friday 09:23",
		},
		Section: "config",
	},
	{
		Keyword: "disconnect-character",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line vty 4\nRouter(config-line)# disconnect-character 27",
		},
		Section: "config-line",
	},
	{
		Keyword: "dispatch-character",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line vty 4\nRouter(config-line)# dispatch-character 13",
		},
		Section: "config-line",
	},
	{
		Keyword: "dispatch-machine",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# state-machine linefeed 0 0 9 0\nRouter(config)# state-machine linefeed 0 11 255 0\nRouter(config)# state-machine linefeed 0 10 10 transmit\nRouter(config)# line 1\nRouter(config-line)# dispatch-machine linefeed",
		},
		Section: "config-line",
	},
	{
		Keyword: "dispatch-timeout",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line vty 0 4\nRouter(config-line)# dispatch-timeout 80",
		},
		Section: "config-line",
	},
	{
		Keyword: "do",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# do show interfaces serial 3/0\nSerial3/0 is up, line protocol is up\n  Hardware is M8T-RS232\n  MTU 1500 bytes, BW 1544 Kbit, DLY 20000 usec, rely 255/255, load 1/255\n  Encapsulation HDLC, loopback not set, keepalive set (10 sec)\n  Last input never, output 1d17h, output hang never\n  Last clearing of “show interface” counters never\n.\n.\n.\n\nRouter(config-vpdn)# do clear vpdn tunnel",
		},
		Section: "",
	},
	{
		Keyword: "downward-compatible-config",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# downward-compatible-config 10.2",
		},
		Section: "config",
	},
	{
		Keyword: "editing",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 3\nRouter(config-line)# no editing",
		},
		Section: "config-line",
	},
	{
		Keyword: "enable last-resort",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# enable last-resort succeed",
		},
		Section: "config",
	},
	{
		Keyword: "end",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# interface serial 1:1\n \nRouter(config-if)# alps ascu 4B\n \nRouter(config-alps-ascu)# end\nRouter# show interface serial 1:1",
		},
		Section: "config",
	},
	{
		Keyword: "environment-monitor shutdown temperature",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \nenvironment-monitor shutdown temperature rommon\nRouter(config)# \n\nRouter(config)# \nenvironment-monitor shutdown temperature powerdown\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "environment temperature-controlled",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \nenvironment temperature-controlled\nRouter(config)# \n\nRouter(config)# \nno environment temperature-controlled\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "errdisable detect cause",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \nerrdisable detect cause l2ptguard",
		},
		Section: "config",
	},
	{
		Keyword: "errdisable recovery",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)#\n errdisable recovery cause bpduguard\n\nRouter(config)#\n errdisable recovery interval 300",
		},
		Section: "config",
	},
	{
		Keyword: "escape-character",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line console\nRouter(config-line)# escape-character 16\n\nRouter(config)# line 1\nRouter(config-line)# escape-character !\nRouter(config-line)# end\nRouter# show running-config\nBuilding configuration...\n.\n.\n.\nline 1\n autoselect during-login\n autoselect ppp\n modem InOut\n transport preferred none\n transport output telnet\n escape-character 33",
		},
		Section: "config-line",
	},
	{
		Keyword: "exec",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 7\n \nRouter(config-line)# no exec",
		},
		Section: "config-line",
	},
	{
		Keyword: "exec-banner",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line vty 0 4\n \nRouter(config-line)# no exec-banner",
		},
		Section: "config-line",
	},
	{
		Keyword: "exec-character-bits",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# default-value exec-character-bits 8\nRouter(config)# line 0\nRouter(config-line)# exec-character-bits 7",
		},
		Section: "config-line",
	},
	{
		Keyword: "exec-timeout",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line console 0\nRouter(config-line)# exec-timeout 2 30\n\nRouter(config)# line console 0\n \nRouter(config-line)# exec-timeout 0 10",
		},
		Section: "config-line",
	},
	{
		Keyword: "exit (global)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-subif)# exit\nRouter(config-if)#\nThe following example displays an exit from the interface configuration mode to return to the global configuration mode:\nRouter(config-if)# exit\nRouter(config)#",
		},
		Section: "",
	},
	{
		Keyword: "file privilege",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device(config)# file privilege ? \n <0-15>  Privilege level\n\nDevice(config)# file privilege 3\nDevice(config)# end\n\nDevice# show running-config | i file priv\nfile privilege 3",
		},
		Section: "config",
	},
	{
		Keyword: "file prompt",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# file prompt noisy",
		},
		Section: "config",
	},
	{
		Keyword: "file verify auto",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# file verify auto",
		},
		Section: "config",
	},
	{
		Keyword: "full-help",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> show ?\n  bootflash  Boot Flash information\n  calendar   Display the hardware calendar\n  clock      Display the system clock\n  context    Show context information\n  dialer     Dialer parameters and statistics\n  history    Display the session command history\n  hosts      IP domain-name, lookup style, nameservers, and host table\n  isdn       ISDN information\n  kerberos   Show Kerberos Values\n  modemcap   Show Modem Capabilities database\n  ppp        PPP parameters and statistics\n  rmon       rmon statistics\n  sessions   Information about Telnet connections\n  snmp       snmp statistics\n  terminal   Display terminal configuration parameters\n  users      Display information about terminal lines\n  version    System hardware and software status\nRouter> enable\nPassword:<letmein>\n\nRouter# configure terminal\nEnter configuration commands, one per line.  End with CNTL/Z.\nRouter(config)# line console 0\nRouter(config-line)# full-help\nRouter(config-line)# exit\n \nRouter#\n%SYS-5-CONFIG_I: Configured from console by console\nRouter# disable\nRouter> show ?\n  access-expression  List access expression\n  access-lists       List access lists\n  aliases            Display alias commands\n  apollo             Apollo network information\n  appletalk          AppleTalk information\n  arp                ARP table\n  async              Information on terminal lines used as router interfaces\n  bootflash          Boot Flash information\n  bridge             Bridge Forwarding/Filtering Database [verbose]\n  bsc                BSC interface information\n  bstun              BSTUN interface information\n  buffers            Buffer pool statistics\n  calendar           Display the hardware calendar\n .\n .\n .\n  translate          Protocol translation information\n  ttycap             Terminal capability tables\n  users              Display information about terminal lines\n  version            System hardware and software status\n  vines              VINES information\n  vlans              Virtual LANs Information\n  whoami             Info on current tty line\n  x25                X.25 information\n  xns                XNS information\n  xremote            XRemote statistics",
		},
		Section: "config-line",
	},
	{
		Keyword: "hidekeys",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device# configure terminal\n!\nDevice(config)# archive\nDevice(config-archive)# log config\nDevice(config-archive-log-config)# hidekeys\nDevice(config-archive-log-config)# end",
		},
		Section: "config-archive-log-config",
	},
	{
		Keyword: "history",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 4\n \nRouter(config-line)# no history",
		},
		Section: "config-line",
	},
	{
		Keyword: "history size",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 4\nRouter(config-line)# history size 35",
		},
		Section: "config-line",
	},
	{
		Keyword: "hold-character",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 8\nRouter(config-line)# hold-character 19",
		},
		Section: "config-line",
	},
	{
		Keyword: "hostname",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "To specify or modify the hostname for the network server, use the hostname command in global configuration mode.",
		},
		Section: "config",
	},
	{
		Keyword: "insecure",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 10\nRouter(config-line)# insecure",
		},
		Section: "config-line",
	},
	{
		Keyword: "international",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "line vty 4\n international",
		},
		Section: "config-line",
	},
	{
		Keyword: "ip bootp server",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# no ip bootp server\n  \nRouter(config)# no service dhcp",
		},
		Section: "config",
	},
	{
		Keyword: "ip finger",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# no ip finger",
		},
		Section: "config",
	},
	{
		Keyword: "ip ftp passive",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# ip ftp passive",
		},
		Section: "config",
	},
	{
		Keyword: "ip ftp password",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# ip ftp username red\n \nRouter(config)# ip ftp password blue",
		},
		Section: "config",
	},
	{
		Keyword: "ip ftp source-interface",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# ip ftp source-interface ethernet 0\n\nRouter# configure terminal\nRouter(config)# ip ftp source-interface ethernet 0\nRouter(config)# ip vrf vpn1\nRouter(config-vrf)# rd 200:1\nRouter(config-vrf)# route-target both 200:1\nRouter(config-vrf)# interface ethernet 0\nRouter(config-if)# ip vrf forwarding vpn1\nRouter(config-if)# end",
		},
		Section: "config",
	},
	{
		Keyword: "ip ftp username",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# ip ftp username red\n \nRouter(config)# ip ftp password blue",
		},
		Section: "config",
	},
	{
		Keyword: "ip rarp-server",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "arp 172.30.2.5 0800.2002.ff5b arpa\ninterface ethernet 0\nip address 172.30.3.100 255.255.255.0\nip rarp-server 172.30.3.100\n\n! Allow the router to forward broadcast portmapper requests\nip forward-protocol udp 111\n! Provide the router with the IP address of the diskless sun\narp 172.30.2.5 0800.2002.ff5b arpa\ninterface ethernet 0\n! Configure the router to act as a RARP server, using the Sun Server's IP\n! address in the RARP response packet.\nip rarp-server 172.30.3.100\n! Portmapper broadcasts from this interface are sent to the Sun Server.\nip helper-address 172.30.3.100",
		},
		Section: "config-if",
	},
	{
		Keyword: "ip rcmd domain-lookup",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# no ip rcmd domain-lookup",
		},
		Section: "config",
	},
	{
		Keyword: "ip rcmd rcp-enable",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# ip rcmd rcp-enable\n \nRouter(config)# ip rcmd source-interface Loopback0\n \nRouter(config)# ip rcmd remote-host router1 172.16.101.101 netadmin3",
		},
		Section: "config",
	},
	{
		Keyword: "ip rcmd remote-host",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# ip rcmd remote-host router1 172.16.101.101 netadmin3 enable",
		},
		Section: "config",
	},
	{
		Keyword: "ip rcmd remote-username",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# ip rcmd remote-username netadmin1",
		},
		Section: "config",
	},
	{
		Keyword: "ip rcmd rsh-enable",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# ip rcmd rsh-enable",
		},
		Section: "config",
	},
	{
		Keyword: "ip rcmd source-interface",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  ".\n.\n.",
		},
		Section: "config",
	},
	{
		Keyword: "ip telnet source-interface",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# ip telnet source-interface Ethernet1",
		},
		Section: "config",
	},
	{
		Keyword: "ip tftp blocksize",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# ip tftp bblocksize 1024",
		},
		Section: "config",
	},
	{
		Keyword: "ip tftp boot-interface",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# ip tftp boot-interface",
		},
		Section: "config",
	},
	{
		Keyword: "ip tftp min-timeout",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# ip tftp min-timeout 5",
		},
		Section: "config",
	},
	{
		Keyword: "ip tftp source-interface",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# ip tftp source-interface loopback0\n\nRouter# configure terminal\nRouter(config)# ip tftp source-interface ethernet 1/0\nRouter(config)# ip vrf vpn1\nRouter(config-vrf)# rd 100:1\nRouter(config-vrf)# route-target both 100:1\nRouter(config-vrf)# interface ethernet 1/0\nRouter(config-if)# ip vrf forwarding vpn1\nRouter(config-if)# end",
		},
		Section: "config",
	},
	{
		Keyword: "ip wccp web-cache accelerated",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# ip wccp web-cache accelerated",
		},
		Section: "config",
	},
	{
		Keyword: "length",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 6\nRouter(config-line)# terminal-type VT220\nRouter(config-line)# length 0",
		},
		Section: "config-line",
	},
	{
		Keyword: "load-interval",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "",
		},
		Section: "config-template",
	},
	{
		Keyword: "location",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line console\nRouter(config-line)# location Building 3, Basement",
		},
		Section: "config-line",
	},
	{
		Keyword: "lockable",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# line console 0\nRouter(config-line)# lockable\nRouter(config)# ^Z\nRouter# lock\nPassword: <password>\nAgain: <password>\n                      Locked\n \nPassword: <password>\nRouter#",
		},
		Section: "config-line",
	},
	{
		Keyword: "log config",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device# configure terminal\n!\nDevice(config)# archive\nDevice(config-archive)# log config\nDevice(config-archive-log-config)#",
		},
		Section: "config-archive",
	},
	{
		Keyword: "logging buffered",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# logging buffered\n\nRouter(config)# logging buffered discriminator buffer1 critical",
		},
		Section: "config",
	},
	{
		Keyword: "logging buginf",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# logging buginf",
		},
		Section: "config",
	},
	{
		Keyword: "logging enable",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device# configure terminal\n!\nDevice(config)# archive\nDevice(config-archive)# log config\nDevice(config-archive-log-config)# logging enable\nDevice(config-archive-log-config)# end\n\nDevice# configure terminal\n!\nDevice(config)# archive\nDevice(config-archive)# log config\nDevice(config-archive-log-config)# no logging enable\nDevice(config-archive-log-config)# logging enable\nDevice(config-archive-log-config)# end",
		},
		Section: "config-archive-log-config",
	},
	{
		Keyword: "logging esm config",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# logging esm config",
		},
		Section: "config",
	},
	{
		Keyword: "logging event bundle-status",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# logging event bundle-status\nRouter(config)# end\nRouter # show logging event bundle-status\n*Aug 4 17:36:48.240 UTC: %EC-SP-5-UNBUNDLE: Interface FastEthernet9/23 left the port-channel Port-channel2\n*Aug 4 17:36:48.256 UTC: %LINK-SP-5-CHANGED: Interface FastEthernet9/23, changed state to administratively down\n*Aug 4 17:36:47.865 UTC: %EC-SPSTBY-5-UNBUNDLE: Interface FastEthernet9/23 left the port-channel Port-channel2\nRouter # show logging event bundle-status\n*Aug 4 17:37:35.845 UTC: %EC-SP-5-BUNDLE: Interface FastEthernet9/23 joined port-channel Port-channel2\n*Aug 4 17:37:35.533 UTC: %EC-SPSTBY-5-BUNDLE: Interface FastEthernet9/23 joined port-channel Port-channel2",
		},
		Section: "config",
	},
	{
		Keyword: "logging event link-status (global configuration)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# logging event link-status default\nRouter(config)# \n\nRouter(config)# logging event link-status boot\nRouter(config)# \n\nRouter(config)# no logging event link-status default\nRouter(config)# \n\nRouter(config)# no logging event link-status boot\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "logging event link-status (interface configuration)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-if)# logging event link-status\n\nRouter(config-if)# no logging event link-status",
		},
		Section: "config-if",
	},
	{
		Keyword: "logging event subif-link-status",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-if)# logging event subif-link-status\nRouter(config-if)# \n\nRouter(config-if)# no logging event subif-link-status\nRouter(config-if)#",
		},
		Section: "config-if",
	},
	{
		Keyword: "logging event trunk-status",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# logging event trunk-status\nRouter(config)# end\nRouter# show logging event trunk-status\n*Aug 4 17:27:01.404 UTC: %DTP-SPSTBY-5-NONTRUNKPORTON: Port Gi3/3 has become non-trunk\n*Aug 4 17:27:00.773 UTC: %DTP-SP-5-NONTRUNKPORTON: Port Gi3/3 has become non-trunk\nRouter#",
		},
		Section: "config-if",
	},
	{
		Keyword: "logging reload",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# logging  reload message-limit 100",
		},
		Section: "config",
	},
	{
		Keyword: "logging ip access-list cache (global configuration)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \nlogging ip access-list cache entries 200\n\nRouter(config)# \nlogging ip access-list cache interval 350\n\nRouter(config)# \nlogging ip access-list cache rate-limit 100\n\nRouter(config)# \nlogging ip access-list cache threshold 125",
		},
		Section: "config",
	},
	{
		Keyword: "logging ip access-list cache (interface configuration)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-if)# \nlogging ip access-list cache in\n\nRouter(config-if)# \nlogging ip access-list cache out",
		},
		Section: "config-if",
	},
	{
		Keyword: "logging persistent (config-archive-log-cfg)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# archive \nRouter(config-archive)# log config \nRouter(config-archive-log-cfg)# logging enable \nRouter(config-archive-log-cfg)# logging persistent auto",
		},
		Section: "config-archive-log-config",
	},
	{
		Keyword: "logging persistent reload (config-archive-log-cfg)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-archive-log-cfg)# logging persistent reload",
		},
		Section: "config-archive-log-config",
	},
	{
		Keyword: "logging purge-log buffer days",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "",
		},
		Section: "config",
	},
	{
		Keyword: "logging size",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device(config-archive-log-config)# logging size 200\n\nDevice(config)# archive\nDevice(config-archive)# log config\nDevice(config-archive-log-config)# logging size 1\nDevice(config-archive-log-config)# logging size 200",
		},
		Section: "config-archive-log-config",
	},
	{
		Keyword: "logging synchronous",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-line)# end\nRouter# show ru\n2w1d: %SYS-5-CONFIG_I: Configured from console by consolenning-config\n .\n .\n .\nRouter# show line\n \n   Tty Typ     Tx/Rx    A Modem  Roty AccO AccI   Uses   Noise  Overruns   Int\n*    0 CTY              -    -      -    -    -      0       3     0/0       -\n .\n .\n .\nRouter# configure terminal \nEnter configuration commands, one per line.  End with CNTL/Z.\nRouter(config)# line 0 \nRouter(config-line)# logging syn\n<tab> \nRouter(config-line)# logging synchronous\n \nRouter(config-line)# end\n \nRouter# show ru\n \n2w1d: %SYS-5-CONFIG_I: Configured from console by console\nRouter# show running-config\n\nRouter(config)# line 4 \nRouter(config-line)# logging synchronous level 6 \nRouter(config-line)# exit \nRouter(config)# line 2 \nRouter(config-line)# logging synchronous level 7 limit 1000 \nRouter(config-line)# end \nRouter#",
		},
		Section: "config-line",
	},
	{
		Keyword: "logging system",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# logging system disk disk1:\n \nRouter(config)# end",
		},
		Section: "config",
	},
	{
		Keyword: "logout-warning",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 5 \nRouter(config-line)# logout-warning 30",
		},
		Section: "config-line",
	},
	{
		Keyword: "macro (global configuration)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# macro global apply snmp ADDRESS test-server VALUE 7 \n\nRouter(config)# macro global trace snmp VALUE 7 VALUE 8 VALUE 9 \nApplying command...`snmp-server enable traps port-security'\nApplying command...`snmp-server enable traps linkup'\nApplying command...`snmp-server enable traps linkdown'\nApplying command...`snmp-server host'\n%Error Unknown error.\nApplying command...`snmp-server ip precedence 7'\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "macro (interface configuration)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# interface fastethernet1/2\n \nRouter(config-if)# macro apply desktop-config\n \nRouter(config-if)# macro apply desktop-config vlan 25",
		},
		Section: "config-if",
	},
	{
		Keyword: "maximum",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "configure terminal\n!\narchive\n path disk0:myconfig\n maximum 5\n end",
		},
		Section: "config-archive",
	},
	{
		Keyword: "memory cache error-recovery",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# memory cache error-recovery",
		},
		Section: "config",
	},
	{
		Keyword: "memory cache error-recovery options",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# memory cache error-recovery options abort-if-same-content",
		},
		Section: "config",
	},
	{
		Keyword: "memory free low-watermark",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# memory free low-watermark processor 200000\n\n000029: *Aug 12 22:31:19.559: %SYS-4-FREEMEMLOW: Free Memory has dropped below 20000k\nPool: Processor  Free: 66814056  freemem_lwm: 204800000\n\n000032: *Aug 12 22:33:29.411: %SYS-5-FREEMEMRECOVER: Free Memory has recovered 20000k\nPool: Processor  Free: 66813960  freemem_lwm: 0",
		},
		Section: "config",
	},
	{
		Keyword: "memory lite",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "no memory lite",
		},
		Section: "config",
	},
	{
		Keyword: "memory reserve",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# memory reserve console 2",
		},
		Section: "config",
	},
	{
		Keyword: "memory reserve critical",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# memory reserve critical 1000",
		},
		Section: "config",
	},
	{
		Keyword: "memory sanity",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "memory sanity all",
		},
		Section: "config",
	},
	{
		Keyword: "memory scan",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# memory scan",
		},
		Section: "config",
	},
	{
		Keyword: "memory-size iomem",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# \nconfigure terminal\nRouter(config)# \nmemory-size iomem 40\nSmart-init will be disabled and new I/O memory size will take effect upon reload.",
		},
		Section: "config",
	},
	{
		Keyword: "menu menu-name single-space",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "menu Access1 single-space",
		},
		Section: "config",
	},
	{
		Keyword: "menu clear-screen",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# menu Access1 clear-screen",
		},
		Section: "config",
	},
	{
		Keyword: "menu command",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device (config) #menu Access1 command 1 tn3270 vms.cisco.com\nDevice (config) #menu Access1 command 2 rlogin unix.cisco.com\nDevice (config) #menu Access1 command 3 menu-exit\n\nmenu Access1 text Exit Exit\nmenu Access1 command Exit menu-exit",
		},
		Section: "config",
	},
	{
		Keyword: "menu default",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "menu Access1 9 text Exit the menu\nmenu Access1 9 command menu-exit\nmenu Access1 default \n9",
		},
		Section: "config",
	},
	{
		Keyword: "menu line-mode",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "menu Access1 line-mode",
		},
		Section: "config",
	},
	{
		Keyword: "menu options",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \nmenu Access1 options 3 login",
		},
		Section: "config",
	},
	{
		Keyword: "menu prompt",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# menu Access1 prompt /\nEnter TEXT message.  End with the character '/'.\nSelect an item. /\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "menu status-line",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "menu Access1 status-line",
		},
		Section: "config",
	},
	{
		Keyword: "menu text",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "menu Access1 text 1 IBM Information Systems\nmenu Access1 text 2 UNIX Internet Access\nmenu Access1 text 3 Exit menu system",
		},
		Section: "config",
	},
	{
		Keyword: "menu title",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# menu Access1 title /^[[H^[[J\nEnter TEXT message.  End with the character '/'.\n               Welcome to Access1 Internet Services\n               \n                  Type a number to select an option;\n                           Type 9 to exit the menu.\n/\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "microcode (12000)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# microcode oc3-POS-4 flash slot0:fip.v141-7 10\n \nRouter(config)# microcode reload 10",
		},
		Section: "config",
	},
	{
		Keyword: "microcode (7000/7500)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \nmicrocode fip slot0:fip.v141-7\nRouter(config)# end\nRouter# copy system:running-config nvram:startup-config",
		},
		Section: "config",
	},
	{
		Keyword: "microcode (7200)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "microcode ecpa slot0:xcpa26-1",
		},
		Section: "config",
	},
	{
		Keyword: "microcode reload (12000)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# microcode reload 10",
		},
		Section: "config",
	},
	{
		Keyword: "microcode reload (7000 7500)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# microcode reload",
		},
		Section: "config",
	},
	{
		Keyword: "mode",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# redundancy\nRouter(config-red)# mode rpr-plus\n\nRouter(config)# redundancy\nRouter(config-red)# mode sso",
		},
		Section: "config-red",
	},
	{
		Keyword: "monitor event-trace (global)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "configure terminal\n!\nmonitor event-trace ipc enable\nmonitor event-trace ipc dump-file slot0:ipc-dump \nmonitor event-trace ipc size 4096\n\nconfigure terminal\n!\nmonitor event-trace cef ipv4 enable\nconfigure terminal\n!\nmonitor event-trace cef ipv6 enable\nexit\nThe following example shows what happens when you try to enable event tracing for a component (in this case, adjacency events) when it is already enabled: \nconfigure terminal\n!\nmonitor event-trace adjacency enable\n%EVENT_TRACE-6-ENABLE: Trace already enabled.",
		},
		Section: "config",
	},
	{
		Keyword: "monitor pcm-tracer capture-destination",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# monitor pcm-tracer capture-destination flash:",
		},
		Section: "config",
	},
	{
		Keyword: "monitor pcm-tracer delayed-start",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# monitor pcm-tracer delayed-start 1000",
		},
		Section: "config",
	},
	{
		Keyword: "monitor pcm-tracer profile",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# monitor pcm-tracer profile 1",
		},
		Section: "config",
	},
	{
		Keyword: "monitor permit-list",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\n \nRouter(config)# monitor permit-list destination interface gigabitethernet 5/1-4\nRouter(config)# monitor permit-list\n\nRouter# configure terminal\n \nRouter(config)# monitor permit-list destination interface fastEthernet 1/1-48 , fastEthernet 2/1-48 , gigabitEthernet 3/1-4",
		},
		Section: "config",
	},
	{
		Keyword: "monitor session egress replication-mode",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "",
		},
		Section: "config",
	},
	{
		Keyword: "monitor session type",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# monitor session 55 type erspan-source\n \nRouter(config-mon-erspan-src)#\n\nRouter(config)# monitor session 55 type erspan-destination\n \nRouter(config-mon-erspan-dst)#\n\nRouter(config-mon-erspan-dst) destination interface fastethernet 1/2 , 2/3\n\nRouter(config-mon-erspan-dst)# source\nRouter(config-mon-erspan-dst-src)#\n\nRouter(config-mon-erspan-dst)# source\nRouter(config-mon-erspan-dst-src)#\n\nRouter(config-mon-erspan-src)# source interface fastethernet 5/15 , 7/3 rx\nRouter(config-mon-erspan-src)# source interface gigabitethernet 1/2 tx \nRouter(config-mon-erspan-src)# source interface port-channel 102 \nRouter(config-mon-erspan-src)# source filter vlan 2 - 3\nRouter(config-mon-erspan-src)#\n\nRouter(config-mon-erspan-src)# destination\nRouter(config-mon-erspan-src-dst)#\n\nRouter(config-mon-erspan-src-dst)# erspan-id 1005\nRouter(config-mon-erspan-src-dst)#\n\nRouter(config)# monitor session 1 type local \nRouter(config-mon-local)# source interface gigabitethernet 1/1 rx \nRouter(config-mon-local)# destination interface gigabitethernet 1/2 \n\nRouter(config)# monitor session 1 type local-tx \nRouter(config-mon-local)# source interface gigabitethernet 5/1 rx \nRouter(config-mon-local)# destination interface gigabitethernet 5/2 \n\nRouter(config)# no monitor session 1 type local-tx",
		},
		Section: "config",
	},
	{
		Keyword: "mop device-code",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "mop device-code ds200",
		},
		Section: "config",
	},
	{
		Keyword: "mop retransmit-timer",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "mop retransmit-timer 10",
		},
		Section: "config",
	},
	{
		Keyword: "mop retries",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# mop retries 11",
		},
		Section: "config",
	},
	{
		Keyword: "motd-banner",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "line vty 0 4\n no motd-banner",
		},
		Section: "config-line",
	},
	{
		Keyword: "nmsp enable",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device> enable\nDevice> configure terminal\nDevice(config)# nmsp enable",
		},
		Section: "config",
	},
	{
		Keyword: "nmsp strong-cipher",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device> enable\nDevice> configure terminal\nDevice(config)# nmsp strong-cipher",
		},
		Section: "config",
	},
	{
		Keyword: "no menu",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "no menu Access1",
		},
		Section: "config",
	},
	{
		Keyword: "notify",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line vty 0 4\nRouter(config-line)# notify",
		},
		Section: "config-line",
	},
	{
		Keyword: "notify syslog",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device# configure terminal\n!\nDevice(config)# archive\nDevice(config-archive)# log config\nDevice(config-archive-log-config)# notify syslog contenttype xml\nDevice(config-archive-log-config)# end",
		},
		Section: "config-archive-log-config",
	},
	{
		Keyword: "padding",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line console\n \nRouter(config-line)# padding 13 25",
		},
		Section: "config-line",
	},
	{
		Keyword: "parity",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 34\nRouter(config-line)# parity even",
		},
		Section: "config-line",
	},
	{
		Keyword: "parser cache",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device(config)# no parser cache",
		},
		Section: "config",
	},
	{
		Keyword: "parser command serializer",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# parser command serializer",
		},
		Section: "config",
	},
	{
		Keyword: "parser config cache interface",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device(config)# parser config cache interface",
		},
		Section: "config",
	},
	{
		Keyword: "parser config partition",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device> enable\nDevice# config t\n \nEnter configuration commands, one per line.  End with CNTL/Z.\nDevice(config)# no parser config partition\nSystem configured",
		},
		Section: "config",
	},
	{
		Keyword: "parser maximum",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)#paser maximum latency 100\n      \n        Router(config)#no paser maximum latency",
		},
		Section: "config",
	},
	{
		Keyword: "partition",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# partition flash 2 4 4\n\nRouter(config)# \npartition slot0: 2 8 8\n\nRouter(config)# partition flash: 4",
		},
		Section: "config",
	},
	{
		Keyword: "path (archive configuration)",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "configure terminal\n!\narchive \n path disk0:$h$t\n time-period 20\n end\n\nDevice# show archive\nThere are currently 3 archive configurations saved.\nThe next archive file will be named routerJan-16-01:12:23.019-4\nArchive #  Name\n   0        \n   1       disk0:routerJan-16-00:12:23.019-1\n   2       disk0:routerJan-16-00:32:23.019-2\n   3       disk0:routerJan-16-00:52:23.019-3 <- Most Recent\n   4        \n   5        \n   6        \n   7        \n   8        \n   9        \n   10        \n   11        \n   12        \n   13        \n   14",
		},
		Section: "config-archive",
	},
	{
		Keyword: "periodic",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# show startup-config\n \n.\n.\n.\ntime-range no-http\n periodic weekdays 8:00 to 18:00\n!\nip access-list extended strict\n deny tcp any any eq http time-range no-http\n!\ninterface ethernet 0\n ip access-group strict in\n.\n.\n.\nRouter# show startup-config\n \n.\n.\n.\ntime-range testing\n periodic Monday Tuesday Friday 9:00 to 17:00\n!\nip access-list extended legal\n permit tcp any any eq telnet time-range testing\n!\ninterface ethernet 0\n ip access-group legal in\n.\n.\n.",
		},
		Section: "config-time-range",
	},
	{
		Keyword: "platform qfp drops threshold",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "",
		},
		Section: "config",
	},
	{
		Keyword: "platform shell",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# platform shell\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "power enable",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \npower enable module 5\nRouter(config)#\n\nRouter(config)# \nno power enable module 5\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "power redundancy-mode",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \npower redundancy-mode combined\nRouter(config)#\n\nRouter(config)# \npower redundancy-mode redundant\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "printer",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# printer printer1 line 4",
		},
		Section: "config",
	},
	{
		Keyword: "private",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 15\nRouter(config-line)# private",
		},
		Section: "config-line",
	},
	{
		Keyword: "process cpu statistics limit entry-percentage",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "configure terminal\n!\nprocess cpu statistics limit entry-percentage 40 size 300\nend",
		},
		Section: "config",
	},
	{
		Keyword: "process cpu threshold type",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "configure terminal\n!\nprocess cpu threshold type total rising 80 interval 5 falling 20 interval 5\nend",
		},
		Section: "config",
	},
	{
		Keyword: "process-max-time",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# process-max-time 100",
		},
		Section: "config",
	},
	{
		Keyword: "prompt",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# prompt TTY%n@%h%s%p\n\nTTY17@Router1 > enable\nTTY17@Router1 #",
		},
		Section: "config",
	},
	{
		Keyword: "prompt config",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)#\n prompt config hostname-length 4",
		},
		Section: "config",
	},
	{
		Keyword: "refuse-message",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "line 5\nrefuse-message  /The dial-out modem is currently in use.\nPlease try again later./",
		},
		Section: "config-line",
	},
	{
		Keyword: "regexp optimize",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# regexp optimize",
		},
		Section: "config",
	},
	{
		Keyword: "remote-span",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-vlan)# remote-span\nRouter(config-vlan)\n\nRouter(config-vlan)# no remote-span\nRouter(config-vlan)",
		},
		Section: "config-vlan",
	},
	{
		Keyword: "revision",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-mst)# revision 5\nRouter(config-mst)#",
		},
		Section: "config-mst",
	},
	{
		Keyword: "scheduler allocate",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# scheduler allocate 2000 500",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler heapcheck enable",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# scheduler heapcheck enable",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler heapcheck poll",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nRouter(config)# scheduler heapcheck poll",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler heapcheck process",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "scheduler heapcheck process memory io checktype magic\n\nscheduler heapcheck process memory processor checktype pointer",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler interrupt mask profile",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# scheduler interrupt mask profile",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler interrupt mask size",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# scheduler interrupt mask size 100",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler interrupt mask time",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# scheduler interrupt mask time 100",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler interval",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# scheduler interval 750",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler isr-watchdog",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# scheduler isr-watchdog",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler max-sched-time",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# scheduler max-sched-time 1000",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler process-watchdog",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# scheduler process-watchdog normal",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler timercheck process",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# scheduler timercheck process 5\nRouter# show processes timer\nSystem timer check not configured.\nProcess timer check configuration follows.\nPID   Configuration            Name\n1     On every context switch. Chunk Manager",
		},
		Section: "config",
	},
	{
		Keyword: "scheduler timercheck system context",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# scheduler timercheck system context",
		},
		Section: "config",
	},
	{
		Keyword: "service compress-config",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nEnter configuration commands, one per line.  End with CNTL/Z.\nRouter(config)# service compress-config\nRouter(config)# end\nRouter#\n%SYS-5-CONFIG_I: Configured from console by console\nRouter# copy system:running-config nvram:startup-config\nBuilding configuration...\nCompressing configuration from 1179 bytes to 674 bytes\n[OK]",
		},
		Section: "config",
	},
	{
		Keyword: "service config",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# service config\n\nRouter(config)# service config\nRouter(config)# boot network rcp://172.16.1.111/bridge_9.1",
		},
		Section: "config",
	},
	{
		Keyword: "service counters max age",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# service counters max age 10\nRouter(config)# \n\nRouter(config)# no service counters max age\nRouter(config)#",
		},
		Section: "config",
	},
	{
		Keyword: "service decimal-tty",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# service decimal-tty",
		},
		Section: "config",
	},
	{
		Keyword: "service exec-wait",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# service exec-wait",
		},
		Section: "config",
	},
	{
		Keyword: "service hide-telnet-address",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# service hide-telnet-address",
		},
		Section: "config",
	},
	{
		Keyword: "service linenumber",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router1> telnet Router2\nTrying Router2 (172.30.162.131)... Open\nWelcome to Router2.\nUser Access Verification\nPassword:\nRouter2> enable\nPassword:\nRouter2# configure terminal\nEnter configuration commands, one per line.  End with CNTL/Z.\nRouter2(config)# service linenumber\nRouter2(config)# end\nRouter2# logout\n[Connection to Router2 closed by foreign host]\nRouter1> telnet Router2\nTrying Router2 (172.30.162.131)... Open\nWelcome to Router2.\nRouter2 line 10\nUser Access Verification\nPassword:\nRouter2>",
		},
		Section: "config",
	},
	{
		Keyword: "service nagle",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# service nagle",
		},
		Section: "config",
	},
	{
		Keyword: "service prompt config",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# configure terminal\nEnter configuration commands, one per line.  End with CNTL/Z.\nRouter(config)# no service prompt config\nhostname newname\nend\nnewname# configure terminal\nEnter configuration commands, one per line.  End with CNTL/Z.\nservice prompt config\nnewname(config)# hostname Router\nRouter(config)# end\nRouter#",
		},
		Section: "config",
	},
	{
		Keyword: "service sequence-numbers",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  ".Mar 22 15:28:02 PST: %SYS-5-CONFIG_I: Configured from console by console\nRouter# config terminal\n \nEnter configuration commands, one per line.  End with CNTL/Z.\nRouter(config)# service sequence-numbers\n \nRouter(config)# end\n \nRouter#\n000066: .Mar 22 15:35:57 PST: %SYS-5-CONFIG_I: Configured from console by console",
		},
		Section: "config",
	},
	{
		Keyword: "service slave-log",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# service slave-log\n\n%IPC-5-SLAVELOG: VIP-SLOT2:\n IPC-2-NOMEM: No memory available for IPC system initialization",
		},
		Section: "config",
	},
	{
		Keyword: "service tcp-keepalives-in",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# service tcp-keepalives-in",
		},
		Section: "config",
	},
	{
		Keyword: "service tcp-keepalives-out",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# service tcp-keepalives-out",
		},
		Section: "config",
	},
	{
		Keyword: "service tcp-small-servers",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)#\n service tcp-small-servers max-servers 14",
		},
		Section: "config",
	},
	{
		Keyword: "service telnet-zeroidle",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# service telnet-zeroidle",
		},
		Section: "config",
	},
	{
		Keyword: "service timestamps",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# show running-config | include time\n \nno service timestamps debug uptime\nno service timestamps log uptime\nRouter# config terminal\n \nRouter(config)# service timestamps\n \n! issue the show running-config command in config mode using  do  Router(config)# do show running-config | inc time\n \n! shows that debug timestamping is enabled, log timestamping is disabled  \nservice timestamps debug uptime\nno service timestamps log uptime\n! enable timestamps for logging messages  \nRouter(config)# service timestamps log \nRouter(config)# do show run | inc time\n \nservice timestamps debug uptime\nservice timestamps log uptime\nRouter(config)# service sequence-numbers\n \nRouter(config)# end\n \n000075: 5w0d: %SYS-5-CONFIG_I: Configured from console by console\n! The following is a level 5 system logging message  \n! The leading number comes from the  service sequence-numbers command.  \n! 4w6d indicates the timestamp of 4 weeks, 6 days 000075: 4w6d: %SYS-5-CONFIG_I: Configured from console by console\n\nRouter(config)#\n! The following line shows the timestamp with uptime (1 week 0 days)  \n1w0d: %SYS-5-CONFIG_I: Configured from console by console\nRouter(config)# service timestamps log datetime show-timezone\n year\n \nRouter(config)# end\n! The following line shows the timestamp with datetime (11:13 PM March 22nd)  \n.Mar 22 2004 23:13:25 UTC: %SYS-5-CONFIG_I: Configured from console by console\n\nRouter# configure terminal\n \n! Logging output can be quite long; first changing line width to show ful l \n! logging message  \nRouter(config)# line 0\n \nRouter(config-line)# width 180\n \nRouter(config-line)# logging synchronous\n \nRouter(config-line)# end\n \n! Timestamping already enabled for logging messages; time shown in UTC.  \nOct 13 23:20:05 UTC: %SYS-5-CONFIG_I: Configured from console by console\nRouter# show clock\n \n23:20:53.919 UTC Wed Oct 13 2004\nRouter# configure terminal\n \nEnter configuration commands, one per line.  End with the end command. \n! Timezone set as Pacific Standard Time, with an 8 hour offset from UTC  \nRouter(config)# clock timezone PST -8\n \nRouter(config)# \nOct 13 23:21:27 UTC: %SYS-6-CLOCKUPDATE: \nSystem clock has been updated from 23:21:27 UTC Wed Oct 13 2004 \nto 15:21:27 PST Wed Oct 13 2004, configured from console by console.\nRouter(config)# \n! Pacific Daylight Time (PDT) configured to start in April and end in October.  \n! Default offset is +1 hour.  \nRouter(config)# clock summer-time PDT recurring first Sunday April 2:00 last Sunday October 2:00 \nRouter(config)#\n! Time changed from 3:22 P.M. Pacific Standard Time (15:22 PST)\n \n! to 4:22 P.M. Pacific Daylight (16:22 PDT)\n \nOct 13 23:22:09 UTC: %SYS-6-CLOCKUPDATE: \nSystem clock has been updated from 15:22:09 PST Wed Oct 13 2004 \nto 16:22:09 PDT Wed Oct 13 2004, configured from console by console.\n! Change the timestamp to show the local time and timezone.  \nRouter(config)# service timestamps log datetime localtime show-timezone\n \nRouter(config)# end\n \nOct 13 16:23:19 PDT: %SYS-5-CONFIG_I: Configured from console by console \nRouter# show clock\n \n16:23:58.747 PDT Wed Oct 13 2004 \nRouter# config t\n \nEnter configuration commands, one per line.  End with the end command. \nRouter(config)# service sequence-numbers\n \nRouter(config)# end\n \nRouter#\n\nRouter(config)# service timestamps log datetime localtime show-timezone \nRouter(config)# end \n! The year is not displayed.  \nOct 13 15:44:46 PDT: %SYS-5-CONFIG_I: Configured from console by console \nRouter# config t\n \nEnter configuration commands, one per line.  End with the end command.\nRouter(config)# service timestamps log datetime show-timezone year\n \nRouter(config)# end\n \n! note: because the \nlocaltime option was not specified again, that option is\n \n! removed from the output, and time is displayed in UTC (the default)\n \nOct 13 2004 22:45:31 UTC: %SYS-5-CONFIG_I: Configured from console by console",
		},
		Section: "config",
	},
	{
		Keyword: "service udp-small-servers",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)#\n service udp-small-servers max-servers 10",
		},
		Section: "config",
	},
	{
		Keyword: "service-module apa traffic-management",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router> enable\nRouter# configure terminal\nRouter(config)# interface gigabitethernet 0/1\nRouter(config-if)# ip address\n 10.10.10.43 255.255.255.0\nRouter(config-if)# service-module apa traffic-management inline\nRouter(config-if)# exit",
		},
		Section: "config-if",
	},
	{
		Keyword: "show",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-mst)# show pending\nPending MST configuration\nName      [zorglub]\nVersion   31415\nInstance  Vlans Mapped\n-------- ---------------------------------------------------------------------\n0        4001-4096\n2        1010, 1020, 1030, 1040, 1050, 1060, 1070, 1080, 1090, 1100, 1110\n         1120\n3        1-1009, 1011-1019, 1021-1029, 1031-1039, 1041-1049, 1051-1059\n         1061-1069, 1071-1079, 1081-1089, 1091-1099, 1101-1109, 1111-1119\n         1121-4000\n------------------------------------------------------------------------------\nRouter(config-mst)# \n\nRouter(config-mst)# show current \nCurrent MST configuration \nName [] \nRevision 0 \nInstance Vlans mapped \n-------- --------------------------------------------------------------------- \n0 1-4094 \n-------------------------------------------------------------------------------",
		},
		Section: "config-mst",
	},
	{
		Keyword: "show declassify",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# show declassify\nDeclassify facility: Enabled=Yes  In Progress=No\n    Erase flash=Yes  Erase nvram=Yes\n    Obtain memory size\n    Shutdown Interfaces\n    Declassify Console and Aux Ports\n    Erase flash\n    Declassify NVRAM\n    Declassify Communications Processor Module\n    Declassify RAM, D-Cache, and I-Cache",
		},
		Section: "config",
	},
	{
		Keyword: "slave auto-sync config",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# slave auto-sync config\nRouter(config)# end\nRouter# copy system:running-config nvram:startup-config",
		},
		Section: "config",
	},
	{
		Keyword: "slave default-slot",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "c7507(config)# slave default-slot 2",
		},
		Section: "config",
	},
	{
		Keyword: "slave image",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# slave image slot0:rsp-dw-mz.ucode.111-3.2",
		},
		Section: "config",
	},
	{
		Keyword: "slave reload",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "c7507(config)# slave reload",
		},
		Section: "config",
	},
	{
		Keyword: "slave terminal",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "c7507(config)# no slave terminal",
		},
		Section: "config",
	},
	{
		Keyword: "software source list",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "(config)software source list my-list-123\n  (config-source-list)tftp://my-big-bundle.bin\n  (config-source-list)bootflash:/packages1\n  (config-source-list)end",
		},
		Section: "config",
	},
	{
		Keyword: "special-character-bits",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 5\nRouter(config-line)# special-character-bits 8",
		},
		Section: "config-line",
	},
	{
		Keyword: "stack-mib portname",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-if)# \nstack-mib portname portall\nRouter(config-if)#",
		},
		Section: "config-if",
	},
	{
		Keyword: "state-machine",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 1 20\nRouter(config-line)# dispatch-machine function\nRouter(config-line)# exit\nRouter(config)# state-machine function 0 0 255 6 transmit",
		},
		Section: "config",
	},
	{
		Keyword: "stopbits",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 4\nRouter(config-line)# stopbits 1",
		},
		Section: "config-line",
	},
	{
		Keyword: "storm-control level",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-if)# \nstorm-control broadcast level 30\n\nRouter(config-if)# \nno storm-control multicast level",
		},
		Section: "config-if",
	},
	{
		Keyword: "sync-restart-delay",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config-if)# sync-restart-delay 2000",
		},
		Section: "config-if",
	},
	{
		Keyword: "system flowcontrol bus",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# system flowcontrol bus auto\n\nRouter(config)# system flowcontrol bus on",
		},
		Section: "config",
	},
	{
		Keyword: "system jumbomtu",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# system jumbomtu 1550\n\nRouter(config)# no system jumbomtu",
		},
		Section: "config",
	},
	{
		Keyword: "tdm clock priority",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "AS5400(config)# tdm clock priority priority 1 external\n\nAS5400(config)# tdm clock priority priority 2 4/6\n\nAS5400(config)# tdm clock priority priority 2 1/0:19\n\nAS5400(config)# tdm clock priority priority 3 freerun",
		},
		Section: "config",
	},
	{
		Keyword: "terminal-queue entry-retry-interval",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router# terminal-queue entry-retry-interval 10",
		},
		Section: "config",
	},
	{
		Keyword: "terminal-type",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 7\n \nRouter(config-line)# terminal-type VT220",
		},
		Section: "config-line",
	},
	{
		Keyword: "tftp-server",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "tftp-server flash version-10.3 22\n\ntftp-server rom alias gs3-k.101\n\ntftp-server flash slot0:version-11.0\n\nRouter# \nconfigure terminal\nEnter configuration commands, one per line.  End with CNTL/Z.\nrouter(config)# tftp-server flash flash:2:dirt/gate/c3640-i-mz\n\nRouter# \nconfigure terminal\nEnter configuration commands, one per line.  End with CNTL/Z.\nRouter(config)# \ntftp-server flash slot0:2:dirt/gate/c3640-j-mz\n\nrouter# \nconfigure terminal\nEnter configuration commands, one per line.  End with CNTL/Z.\nrouter(config)# tftp-server flash flash:2:dirt/gate/c1600-i-mz",
		},
		Section: "config",
	},
	{
		Keyword: "time-period",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Device# configure terminal\n!\nDevice(config)# archive\nDevice(config-archive)# path disk0:myconfig\nDevice(config-archive)# time-period 20\nDevice(config-archive)# end",
		},
		Section: "config-archive",
	},
	{
		Keyword: "vacant-message",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 0\nRouter(config-line)# vacant-message %\n                Welcome to Cisco Systems, Inc.\n                 Press Return to get started.\n%",
		},
		Section: "config-line",
	},
	{
		Keyword: "vtp",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# \nvtp domain DomainName1\n\nRouter(config)# \nvtp file vtpconfig\nSetting device to store VLAN database at filename vtpconfig. \n\nRouter(config)# \nvtp mode client\nSetting device to VTP CLIENT mode.\n\nRouter(config)# vtp mode off\nSetting device to VTP OFF mode.\n\nRouter(config)# no vtp mode off\nSetting device to VTP OFF mode.",
		},
		Section: "config",
	},
	{
		Keyword: "warm-reboot",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router#(config) warm-reboot count 10 uptime 10",
		},
		Section: "config",
	},
	{
		Keyword: "width",
		Description: keyword.Description{
			Format: protocol.PlainText,
			Value:  "Router(config)# line 7\nRouter(config-line)# location console terminal\nRouter(config-line)# width 132",
		},
		Section: "config-line",
	},
}
