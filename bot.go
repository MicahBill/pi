package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"context"
	"math/rand"
	//"net/http"
	"os"
	"os/exec"
	"sync"
	"unicode/utf8"
	"runtime"
	"strconv"
	"strings"
	"time"

	lord "./LINE/LineThrift"
	"./LINE/auth"
	"./LINE/helper"
	"./LINE/service"
	"./LINE/talk"
	con "./LINE/config"
	//"./LINE/thrift"
	crypto_rand "crypto/rand"
	"math/big"
	"bufio"
)

var (
	Config        config
	Permit        permit
	Commands      commands
	Rname  		  string
	Myself        string
	TimeLeft      time.Time
	startBots     = time.Now()
	Lord          = os.Args
	DB            = Lord[1]
	DataBase      = "DataBase/" + DB + ".json"
	CmdData       = "Commands/" + DB + ".json"
	CmdPermit     = "Permit/" + DB + ".json"
	AppName       = ""
	Changepic     = false
	Changecov     = false
	Addban        = false
	Master        = []string{"u0be40d5b6854cce0f78b76a7ace30727","u54c5ecfded8b1983ef97cd4dc8522c88"}
	Invites 	  = []string{}
	Backup        = []string{}
	XBackup       = []string{}
	Proxy    	  = []string{}
	StatusAsist   = []string{}
	Total         int
	Mue           int
	JoinBreak     int
	Killmode      = 0
	Limit    	  = 0 
	Countkick     = 0
	Countinvite   = 0
	Countcancel   = 0
	Akick         = 0
	Ainvite       = 0
	Acancel       = 0
	xsender  string
	Sname    string 
	Key      = "!"
	Respon   = "Yes sir"
	Lkick    string 
	Lcancel  string 
	Linvite  string
	Lupdate  string 
	Lcontact string 
	Lmention string 
	Ljoin 	 string 
	Lleave   string 
	Logmode  string
	Msgsider string
	WarMode  = false
	Logs     = false
	Lockdown = false 
	Purge 	 = false 
	Mute 	 = false
	Blocked  = false 
	JoinNuke = false
	Swith    = false
	Addbl    = false
	Addsl    = false
	Addow    = false
	Addad    = false
	Addst    = false
	Seler    = []string{}
	Owners   = []string{}
	Admins   = []string{}
	Staff    = []string{}
	Bots     = []string{}
	Center   = []string{}
	Linkpro  = []string{}
	Denyinv  = []string{}
	Protect  = []string{} 
	Namelock = []string{}
	Projoin  = []string{} 
	Limiter  = []string{} 
	Banned   = []string{}
	Mytoken  = []string{}
	Hiden    = []string{}
	Sider    = map[string][]string{}
	SiderV2  = map[string]bool{}
	Gname    = make(map[string]string)
	Replyms  = ""//make(map[string]int)
	WordBan  = map[string][]string{}
	Stay     = map[string][]string{}
	Gmaster  = map[string][]string{}
	Gowner   = map[string][]string{}
	Gadmin   = map[string][]string{}
	Gban     = map[string][]string{}
	InRoom   = map[string][]string{}
	XInRoom  = map[string][]string{}
	justgood = "https://yos.imjustgood.com/"
	oupLogo  = "https://gamedaim.com/wp-content/uploads/2021/07/14.jpg"
	oupTit   = "𝐃𝐄𝐕𝐄𝐋𝐎𝐏𝐄𝐑 𝐁𝐎𝐓𝐒"
)

var PInvBots, PNewseller, PUnseller, PNewowner, PUnowner, PJoinall, PUpkey, PUprespon, PUpbio, PUpname, PUpsname, PUnfriends, PNewadmin, PUnadmin, PSetlimit, PAddme, PJoino, PLeaveto, PInvto, PUrljoined, PNewstaff, PUnstaff, PNewcenter, PUncenter, PNewbots, PUnbots, PNewgmaster, PUngmaster, PContact, PMid, PFriends, PGinvited, PGroups, PSpeed, POurl, PCurl, PUnsend, PUpgname, PNewban, PUnban, PNewgowner, PUngowner, PRuntime, PNewgadmin, PUngadmin, PKick, PHere, PTagall, PRes, PAccess, PLinkpro, PNamelock, PDenyin, PProjoin, PProtect, PAutopurge, PLockdown, PJoinNuke, PLogmode, PPurge, PKillmode, PHelp, PList, PClear, PCancel, PInvite, PNewfuck, PUnfuck, PSider, PMsgsider, PHiden, PUnhiden, PUpimage, PBye, PTimeleft, PExtend, PCleanse, PBreak, PCenterstay, PCheckcenter, PSet, PReduce int

var (
	CNewseller 	   = "newseller"
	CUnseller 	   = "unseller"
	CNewowner 	   = "newowner"
	CUnowner 	   = "unowner"
	CJoinall 	   = "joinall"
	CUpkey 		   = "upkey"
	CUprespon 	   = "uprespon"
	CUpbio 		   = "upbio"
	CUpname 	   = "upname"
	CUpsname 	   = "upsname"
	CUnfriends 	   = "unfriends"
	CNewadmin 	   = "newadmin"
	CUnadmin 	   = "unadmin"
	CSetlimit 	   = "setlimit"
	CAddme 		   = "addme"
	CJoino 		   = "joino"
	CLeaveto 	   = "leaveto"
	CInvto 		   = "invto"
	CUrljoined 	   = "urljoined"
	CNewstaff 	   = "newstaff"
	CUnstaff 	   = "unstaff"
	CNewcenter 	   = "newcenter"
	CUncenter 	   = "uncenter"
	CNewbots 	   = "newbots"
	CUnbots 	   = "unbots"
	CNewgmaster    = "newgmaster"
	CUngmaster 	   = "ungmaster"
	CContact 	   = "contact"
	CMid 		   = "mid"
	CFriends 	   = "friends"
	CGinvited 	   = "ginvited"
	CGroups 	   = "groups"
	CSpeed 		   = "speed"
	COurl 		   = "ourl"
	CCurl 		   = "curl"
	CUnsend 	   = "unsend"
	CUpgname 	   = "upgname"
	CNewban 	   = "newban"
	CUnban 		   = "unban"
	CNewgowner 	   = "newgowner"
	CUngowner 	   = "ungowner"
	CRuntime 	   = "runtime"
	CNewgadmin 	   = "newgadmin"
	CUngadmin 	   = "ungadmin"
	CKick 		   = "kick"
	CHere 		   = "here"
	CTagall 	   = "tagall"
	CRes 		   = "res"
	CAccess 	   = "access"
	CLinkpro 	   = "linkpro"
	CNamelock 	   = "namelock"
	CDenyin 	   = "denyin"
	CProjoin 	   = "projoin"
	CProtect 	   = "protect"
	CAutopurge 	   = "autopurge"
	CLockdown 	   = "lockdown"
	CJoinNuke 	   = "nukejoin"
	CLogmode 	   = "logmode"
	CPurge 		   = "purge"
	CKillmode 	   = "killmode"
	CHelp 		   = "help"
	CList 		   = "list"
	CClear 		   = "clear"
	CCancel        = "cancel"
	CInvite        = "invite"
	CNewfuck       = "cancel"
	CUnfuck        = "invite"
	CSider   	   = "read"
	CMsgsider	   = "readmsg" 
	CHiden   	   = "hiden"
	CUnhiden 	   = "unhiden"
	CUpimage       = "upimage"
	CBye 		   = "bye"
	CTimeleft 	   = "timeleft"
	CExtend 	   = "extend date"
	CCleanse 	   = "cleanse"
	CBreak 		   = "break"
	CCenterstay    = "center stay"
	CCheckcenter   = "check center"	
	CSet 		   = "set"
	CReduce 	   = "reduce date"
)

type permit struct {
	InvBots     int `json:"invitebots`
	Newseller 	int `json:"newseller"`
	Unseller 	int `json:"unseller"`
	Newowner 	int `json:"newowner"`
	Unowner 	int `json:"unowner"`
	Joinall 	int `json:"joinall"`
	Upkey 		int `json:"upkey"`
	Uprespon 	int `json:"uprespon"`
	Upbio 		int `json:"upbio"`
	Upname 		int `json:"upname"`
	Upsname 	int `json:"upsname"`
	Unfriends 	int `json:"unfriends"`
	Newadmin 	int `json:"newadmin"`
	Unadmin 	int `json:"unadmin"`
	Setlimit 	int `json:"setlimit"`
	Addme 		int `json:"addme"`
	Joino 		int `json:"joino"`
	Leaveto 	int `json:"leaveto"`
	Invto 		int `json:"invto"`
	Urljoined 	int `json:"urljoined"`
	Newstaff 	int `json:"newstaff"`
	Unstaff 	int `json:"unstaff"`
	Newcenter 	int `json:"newcenter"`
	Uncenter 	int `json:"uncenter"`
	Newbots 	int `json:"newbots"`
	Unbots 		int `json:"unbots"`
	Newgmaster 	int `json:"newgmaster"`
	Ungmaster 	int `json:"ungmaster"`
	Contact 	int `json:"contact"`
	Mid 		int `json:"mid"`
	Friends 	int `json:"friends"`
	Ginvited 	int `json:"ginvited"`
	Groups 		int `json:"groups"`
	Speed 		int `json:"speed"`
	Ourl 		int `json:"ourl"`
	Curl 		int `json:"curl"`
	Unsend 		int `json:"unsend"`
	Upgname 	int `json:"upgname"`
	Newban 		int `json:"newban"`
	Unban 		int `json:"unban"`
	Newgowner 	int `json:"newgowner"`
	Ungowner 	int `json:"ungowner"`
	Runtime 	int `json:"runtime"`
	Newgadmin 	int `json:"newgadmin"`
	Ungadmin 	int `json:"ungadmin"`
	Kick 		int `json:"kick"`
	Here 		int `json:"here"`
	Tagall 		int `json:"tagall"`
	Res 		int `json:"res"`
	Access 		int `json:"access"`
	Linkpro 	int `json:"linkpro"`
	Namelock 	int `json:"namelock"`
	Denyin 		int `json:"denyin"`
	Projoin 	int `json:"projoin"`
	Protect 	int `json:"protect"`
	Autopurge 	int `json:"autopurge"`
	Lockdown 	int `json:"lockdown"`
	JoinNuke 	int `json:"joinNuke"`
	Logmode 	int `json:"logmode"`
	Purge 		int `json:"purge"`
	Killmode 	int `json:"killmode"`
	Help 		int `json:"help"`
	List 		int `json:"list"`
	Clear 		int `json:"clear"`
	Cancel      int `json:"cancel"`
	Invite      int `json:"invite"`
	Newfuck     int `json:"newfuck"`
	Unfuck      int `json:"unfuck"`
	Sider   	int `json:"sider"`
	Msgsider	int `json:"msgsider"`
	Hiden   	int `json:"hiden"`
	Unhiden 	int `json:"unhiden"`
	Upimage 	int `json:"upimage"`
	Bye 		int `json:"bye"`
	Timeleft 	int `json:"timeleft"`
	Extend 		int `json:"extend"`
	Cleanse 	int `json:"cleanse"`
	Break 		int `json:"break"`
	Centerstay  int `json:"centerstay"`
	Checkcenter int `json:"checkcenter"`
	Set 		int `json:"set"`
	Reduce 		int `json:"reduce"`
}

func PermitLoad() {
	jsonFile, err := os.Open(CmdPermit)
	if err != nil {
		Error := fmt.Sprintf("** ERROR DATABASE **\n* Reason: \n%s", err)
		fmt.Println(Error)
		os.Exit(1)
	}
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(srcJSON), &Permit)
	fmt.Println(err)
	PInvBots = Permit.InvBots
	PNewseller = Permit.Newseller
	PUnseller = Permit.Unseller
	PNewowner = Permit.Newowner
	PUnowner = Permit.Unowner
	PJoinall = Permit.Joinall
	PUpkey = Permit.Upkey
	PUprespon = Permit.Uprespon
	PUpbio = Permit.Upbio
	PUpname = Permit.Upname
	PUpsname = Permit.Upsname
	PUnfriends = Permit.Unfriends
	PNewadmin = Permit.Newadmin
	PUnadmin = Permit.Unadmin
	PSetlimit = Permit.Setlimit
	PAddme = Permit.Addme
	PJoino = Permit.Joino
	PLeaveto = Permit.Leaveto
	PInvto = Permit.Invto
	PUrljoined = Permit.Urljoined
	PNewstaff = Permit.Newstaff
	PUnstaff = Permit.Unstaff
	PNewcenter = Permit.Newcenter
	PUncenter = Permit.Uncenter
	PNewbots = Permit.Newbots
	PUnbots = Permit.Unbots
	PNewgmaster = Permit.Newgmaster
	PUngmaster = Permit.Ungmaster
	PContact = Permit.Contact
	PMid = Permit.Mid
	PFriends = Permit.Friends
	PGinvited = Permit.Ginvited
	PGroups = Permit.Groups
	PSpeed = Permit.Speed
	POurl = Permit.Ourl
	PCurl = Permit.Curl
	PUnsend = Permit.Unsend
	PUpgname = Permit.Upgname
	PNewban = Permit.Newban
	PUnban = Permit.Unban
	PNewgowner = Permit.Newgowner
	PUngowner = Permit.Ungowner
	PRuntime = Permit.Runtime
	PNewgadmin = Permit.Newgadmin
	PUngadmin = Permit.Ungadmin
	PKick = Permit.Kick
	PHere = Permit.Here
	PTagall = Permit.Tagall
	PRes = Permit.Res
	PAccess = Permit.Access
	PLinkpro = Permit.Linkpro
	PNamelock = Permit.Namelock
	PDenyin = Permit.Denyin
	PProjoin = Permit.Projoin
	PProtect = Permit.Protect
	PAutopurge = Permit.Autopurge
	PLockdown = Permit.Lockdown
	PJoinNuke = Permit.JoinNuke
	PLogmode = Permit.Logmode
	PPurge = Permit.Purge
	PKillmode = Permit.Killmode
	PHelp = Permit.Help
	PList = Permit.List
	PClear = Permit.Clear
	PCancel = Permit.Cancel
	PInvite = Permit.Invite
	PNewfuck = Permit.Newfuck
	PUnfuck = Permit.Unfuck
	PSider    = Permit.Sider   
	PMsgsider = Permit.Msgsider
	PHiden    = Permit.Hiden   
	PUnhiden  = Permit.Unhiden
	PUpimage  = Permit.Upimage
	PBye = Permit.Bye
	PTimeleft = Permit.Timeleft
	PExtend = Permit.Extend
	PCleanse = Permit.Cleanse
	PBreak = Permit.Break
	PCenterstay = Permit.Centerstay
	PCheckcenter = Permit.Checkcenter
	PSet = Permit.Set
	PReduce = Permit.Reduce
}

func PermitSave(){
	Permit.InvBots = PInvBots
	Permit.Newseller = PNewseller
	Permit.Unseller = PUnseller
	Permit.Newowner = PNewowner
	Permit.Unowner = PUnowner
	Permit.Joinall = PJoinall
	Permit.Upkey = PUpkey
	Permit.Uprespon = PUprespon
	Permit.Upbio = PUpbio
	Permit.Upname = PUpname
	Permit.Upsname = PUpsname
	Permit.Unfriends = PUnfriends
	Permit.Newadmin = PNewadmin
	Permit.Unadmin = PUnadmin
	Permit.Setlimit = PSetlimit
	Permit.Addme = PAddme
	Permit.Joino = PJoino
	Permit.Leaveto = PLeaveto
	Permit.Invto = PInvto
	Permit.Urljoined = PUrljoined
	Permit.Newstaff = PNewstaff
	Permit.Unstaff = PUnstaff
	Permit.Newcenter = PNewcenter
	Permit.Uncenter = PUncenter
	Permit.Newbots = PNewbots
	Permit.Unbots = PUnbots
	Permit.Newgmaster = PNewgmaster
	Permit.Ungmaster = PUngmaster
	Permit.Contact = PContact
	Permit.Mid = PMid
	Permit.Friends = PFriends
	Permit.Ginvited = PGinvited
	Permit.Groups = PGroups
	Permit.Speed = PSpeed
	Permit.Ourl = POurl
	Permit.Curl = PCurl
	Permit.Unsend = PUnsend
	Permit.Upgname = PUpgname
	Permit.Newban = PNewban
	Permit.Unban = PUnban
	Permit.Newgowner = PNewgowner
	Permit.Ungowner = PUngowner
	Permit.Runtime = PRuntime
	Permit.Newgadmin = PNewgadmin
	Permit.Ungadmin = PUngadmin
	Permit.Kick = PKick
	Permit.Here = PHere
	Permit.Tagall = PTagall
	Permit.Res = PRes
	Permit.Access = PAccess
	Permit.Linkpro = PLinkpro
	Permit.Namelock = PNamelock
	Permit.Denyin = PDenyin
	Permit.Projoin = PProjoin
	Permit.Protect = PProtect
	Permit.Autopurge = PAutopurge
	Permit.Lockdown = PLockdown
	Permit.JoinNuke = PJoinNuke
	Permit.Logmode = PLogmode
	Permit.Purge = PPurge
	Permit.Killmode = PKillmode
	Permit.Help = PHelp
	Permit.List = PList
	Permit.Clear = PClear
	Permit.Cancel = PCancel
	Permit.Invite = PInvite
	Permit.Newfuck = PNewfuck
	Permit.Unfuck = PUnfuck
	Permit.Sider    = PSider   
	Permit.Msgsider = PMsgsider
	Permit.Hiden    = PHiden   
	Permit.Unhiden  = PUnhiden 
	Permit.Upimage = PUpimage
	Permit.Bye = PBye
	Permit.Timeleft = PTimeleft
	Permit.Extend = PExtend
	Permit.Cleanse = PCleanse
	Permit.Break = PBreak
	Permit.Centerstay = PCenterstay
	Permit.Checkcenter = PCheckcenter
	Permit.Set = PSet
	Permit.Reduce = PReduce
	encjson, _ := json.MarshalIndent(Permit, "", "  ")
	ioutil.WriteFile(CmdPermit, encjson, 0644)
}

type commands struct {
	Newseller 	string `json:"newseller"`
	Unseller 	string `json:"unseller"`
	Newowner 	string `json:"newowner"`
	Unowner 	string `json:"unowner"`
	Joinall 	string `json:"joinall"`
	Upkey 		string `json:"upkey"`
	Uprespon 	string `json:"uprespon"`
	Upbio 		string `json:"upbio"`
	Upname 		string `json:"upname"`
	Upsname 	string `json:"upsname"`
	Unfriends 	string `json:"unfriends"`
	Newadmin 	string `json:"newadmin"`
	Unadmin 	string `json:"unadmin"`
	Setlimit 	string `json:"setlimit"`
	Addme 		string `json:"addme"`
	Joino 		string `json:"joino"`
	Leaveto 	string `json:"leaveto"`
	Invto 		string `json:"invto"`
	Urljoined 	string `json:"urljoined"`
	Newstaff 	string `json:"newstaff"`
	Unstaff 	string `json:"unstaff"`
	Newcenter 	string `json:"newcenter"`
	Uncenter 	string `json:"uncenter"`
	Newbots 	string `json:"newbots"`
	Unbots 		string `json:"unbots"`
	Newgmaster 	string `json:"newgmaster"`
	Ungmaster 	string `json:"ungmaster"`
	Contact 	string `json:"contact"`
	Mid 		string `json:"mid"`
	Friends 	string `json:"friends"`
	Ginvited 	string `json:"ginvited"`
	Groups 		string `json:"groups"`
	Speed 		string `json:"speed"`
	Ourl 		string `json:"ourl"`
	Curl 		string `json:"curl"`
	Unsend 		string `json:"unsend"`
	Upgname 	string `json:"upgname"`
	Newban 		string `json:"newban"`
	Unban 		string `json:"unban"`
	Newgowner 	string `json:"newgowner"`
	Ungowner 	string `json:"ungowner"`
	Runtime 	string `json:"runtime"`
	Newgadmin 	string `json:"newgadmin"`
	Ungadmin 	string `json:"ungadmin"`
	Kick 		string `json:"kick"`
	Here 		string `json:"here"`
	Tagall 		string `json:"tagall"`
	Res 		string `json:"res"`
	Access 		string `json:"access"`
	Linkpro 	string `json:"linkpro"`
	Namelock 	string `json:"namelock"`
	Denyin 		string `json:"denyin"`
	Projoin 	string `json:"projoin"`
	Protect 	string `json:"protect"`
	Autopurge 	string `json:"autopurge"`
	Lockdown 	string `json:"lockdown"`
	JoinNuke 	string `json:"joinNuke"`
	Logmode 	string `json:"logmode"`
	Purge 		string `json:"purge"`
	Killmode 	string `json:"killmode"`
	Help 		string `json:"help"`
	List 		string `json:"list"`
	Clear 		string `json:"clear"`
	Cancel      string `json:"cancel"`
	Invite      string `json:"invite"`
	Newfuck     string `json:"newfuck"`
	Unfuck      string `json:"unfuck"`
	Sider   	string `json:"sider"`
	Msgsider	string `json:"msgsider"`
	Hiden   	string `json:"hiden"`
	Unhiden 	string `json:"unhiden"`
	Upimage     string `json:"upimage"`
	Bye 		string `json:"bye"`
	Timeleft 	string `json:"timeleft"`
	Extend 		string `json:"extend"`
	Cleanse 	string `json:"cleanse"`
	Break 		string `json:"break"`
	Centerstay  string `json:"centerstay"`
	Checkcenter string `json:"checkcenter"`
	Set 		string `json:"set"`
	Reduce      string `json:"reduce"`
}

func CmdsLoad(){
	jsonFile, err := os.Open(CmdData)
	if err != nil {
		Error := fmt.Sprintf("** ERROR DATABASE **\n* Reason: \n%s", err)
		fmt.Println(Error)
		os.Exit(1)
	}
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(srcJSON), &Commands)
	fmt.Println(err)
	CNewseller = Commands.Newseller
	CUnseller = Commands.Unseller
	CNewowner = Commands.Newowner
	CUnowner = Commands.Unowner
	CJoinall = Commands.Joinall
	CUpkey = Commands.Upkey
	CUprespon = Commands.Uprespon
	CUpbio = Commands.Upbio
	CUpname = Commands.Upname
	CUpsname = Commands.Upsname
	CUnfriends = Commands.Unfriends
	CNewadmin = Commands.Newadmin
	CUnadmin = Commands.Unadmin
	CSetlimit = Commands.Setlimit
	CAddme = Commands.Addme
	CJoino = Commands.Joino
	CLeaveto = Commands.Leaveto
	CInvto = Commands.Invto
	CUrljoined = Commands.Urljoined
	CNewstaff = Commands.Newstaff
	CUnstaff = Commands.Unstaff
	CNewcenter = Commands.Newcenter
	CUncenter = Commands.Uncenter
	CNewbots = Commands.Newbots
	CUnbots = Commands.Unbots
	CNewgmaster = Commands.Newgmaster
	CUngmaster = Commands.Ungmaster
	CContact = Commands.Contact
	CMid = Commands.Mid
	CFriends = Commands.Friends
	CGinvited = Commands.Ginvited
	CGroups = Commands.Groups
	CSpeed = Commands.Speed
	COurl = Commands.Ourl
	CCurl = Commands.Curl
	CUnsend = Commands.Unsend
	CUpgname = Commands.Upgname
	CNewban = Commands.Newban
	CUnban = Commands.Unban
	CNewgowner = Commands.Newgowner
	CUngowner = Commands.Ungowner
	CRuntime = Commands.Runtime
	CNewgadmin = Commands.Newgadmin
	CUngadmin = Commands.Ungadmin
	CKick = Commands.Kick
	CHere = Commands.Here
	CTagall = Commands.Tagall
	CRes = Commands.Res
	CAccess = Commands.Access
	CLinkpro = Commands.Linkpro
	CNamelock = Commands.Namelock
	CDenyin = Commands.Denyin
	CProjoin = Commands.Projoin
	CProtect = Commands.Protect
	CAutopurge = Commands.Autopurge
	CLockdown = Commands.Lockdown
	CJoinNuke = Commands.JoinNuke
	CLogmode = Commands.Logmode
	CPurge = Commands.Purge
	CKillmode = Commands.Killmode
	CHelp = Commands.Help
	CList = Commands.List
	CClear = Commands.Clear
	CCancel = Commands.Cancel
	CInvite = Commands.Invite
	CNewfuck = Commands.Newfuck
	CUnfuck  = Commands.Unfuck 
	CSider    = Commands.Sider   
	CMsgsider = Commands.Msgsider
	CHiden    = Commands.Hiden   
	CUnhiden  = Commands.Unhiden
	CUpimage  = Commands.Upimage
	CBye = Commands.Bye
	CTimeleft = Commands.Timeleft
	CExtend = Commands.Extend
	CCleanse = Commands.Cleanse
	CBreak = Commands.Break
	CCenterstay = Commands.Centerstay
	CCheckcenter = Commands.Checkcenter
	CSet = Commands.Set
	CReduce = Commands.Reduce
}
func CmdsSave(){
	Commands.Newseller = CNewseller
	Commands.Unseller = CUnseller
	Commands.Newowner = CNewowner
	Commands.Unowner = CUnowner
	Commands.Joinall = CJoinall
	Commands.Upkey = CUpkey
	Commands.Uprespon = CUprespon
	Commands.Upbio = CUpbio
	Commands.Upname = CUpname
	Commands.Upsname = CUpsname
	Commands.Unfriends = CUnfriends
	Commands.Newadmin = CNewadmin
	Commands.Unadmin = CUnadmin
	Commands.Setlimit = CSetlimit
	Commands.Addme = CAddme
	Commands.Joino = CJoino
	Commands.Leaveto = CLeaveto
	Commands.Invto = CInvto
	Commands.Urljoined = CUrljoined
	Commands.Newstaff = CNewstaff
	Commands.Unstaff = CUnstaff
	Commands.Newcenter = CNewcenter
	Commands.Uncenter = CUncenter
	Commands.Newbots = CNewbots
	Commands.Unbots = CUnbots
	Commands.Newgmaster = CNewgmaster
	Commands.Ungmaster = CUngmaster
	Commands.Contact = CContact
	Commands.Mid = CMid
	Commands.Friends = CFriends
	Commands.Ginvited = CGinvited
	Commands.Groups = CGroups
	Commands.Speed = CSpeed
	Commands.Ourl = COurl
	Commands.Curl = CCurl
	Commands.Unsend = CUnsend
	Commands.Upgname = CUpgname
	Commands.Newban = CNewban
	Commands.Unban = CUnban
	Commands.Newgowner = CNewgowner
	Commands.Ungowner = CUngowner
	Commands.Runtime = CRuntime
	Commands.Newgadmin = CNewgadmin
	Commands.Ungadmin = CUngadmin
	Commands.Kick = CKick
	Commands.Here = CHere
	Commands.Tagall = CTagall
	Commands.Res = CRes
	Commands.Access = CAccess
	Commands.Linkpro = CLinkpro
	Commands.Namelock = CNamelock
	Commands.Denyin = CDenyin
	Commands.Projoin = CProjoin
	Commands.Protect = CProtect
	Commands.Autopurge = CAutopurge
	Commands.Lockdown = CLockdown
	Commands.JoinNuke = CJoinNuke
	Commands.Logmode = CLogmode
	Commands.Purge = CPurge
	Commands.Killmode = CKillmode
	Commands.Help = CHelp
	Commands.List = CList
	Commands.Clear = CClear
	Commands.Cancel = CCancel
	Commands.Invite = CInvite
	Commands.Newfuck = CNewfuck
	Commands.Unfuck  = CUnfuck
	Commands.Sider    = CSider   
	Commands.Msgsider = CMsgsider
	Commands.Hiden    = CHiden   
	Commands.Unhiden  = CUnhiden 
	Commands.Upimage  = CUpimage
	Commands.Bye = CBye
	Commands.Timeleft = CTimeleft
	Commands.Extend = CExtend
	Commands.Cleanse = CCleanse
	Commands.Break = CBreak
	Commands.Centerstay = CCenterstay
	Commands.Checkcenter = CCheckcenter
	Commands.Set = CSet
	Commands.Reduce = CReduce
	encjson, _ := json.MarshalIndent(Commands, "", "  ")
	ioutil.WriteFile(CmdData, encjson, 0644)
}


type config struct {
	Limit    int `json:"Limit"`
	Killmode int `json:"killmode"`
	Sname    string `json:"sname"`
	Rname    string `json:"rname"`
	Key      string `json:"key"`
	Respon   string `json:"response"`
	Lkick    string `json:"lkick"`
	Lcancel  string `json:"lcancel"`
	Linvite  string `json:"linvite"`
	Lupdate  string `json:"lupdate"`
	Lcontact string `json:"lcontact"`
	Lmention string `json:"lmention"`
	Ljoin 	 string `json:"ljoin"`
	Lleave   string `json:"lleave"`
	Logmode  string `json:"logmode"`
	Msgsider string `json:"msgsider`
	TimeLeft time.Time `json:"timeleft"`
	Lockdown bool `json:"lockdown"`
	Purge 	 bool `json:"purge"`
	Mute 	 bool `json:"mute"`
	Blocked  bool `json:"blocked"`
	Logs     bool `json:"logs"`
	JoinNuke bool `json:"joinnuke"`
	Seler   []string `json:"seller"`
	Owners  []string `json:"owners"`
	Admins  []string `json:"admins"`
	Staff   []string `json:"staff"`
	Bots    []string `json:"bots"`
	Center  []string `json:"centers"`
	Linkpro []string `json:"linkprotect"`
	Denyinv []string `json:"denyinvite"`
	Protect []string `json:"protect"`
	Namelock[]string `json:"namelock"`
	Projoin []string `json:"projoin"`
	Limiter []string `json:"limiter"`
	Banned  []string `json:"banned"`
	Hiden   []string `json:"hiden"`
	Gname   map[string]string `json:"groupname"`
	WordBan map[string][]string`json:"wordban"`
	Stay    map[string][]string`json:"stay"`
	Gmaster map[string][]string `json:"gmaster"`
	Gowner  map[string][]string `json:"gowner"`
	Gadmin  map[string][]string `json:"gadmin"`
	Gban    map[string][]string `json:"gban"`
	InRoom  map[string][]string `json:"inroom"`
	XInRoom map[string][]string `json:"inroom"`
}

func LoadJson() {
	jsonFile, err := os.Open(DataBase)
	if err != nil {
		Error := fmt.Sprintf("** ERROR DATABASE **\n* Reason: \n%s", err)
		fmt.Println(Error)
		os.Exit(1)
	}
	defer jsonFile.Close()
	srcJSON, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal([]byte(srcJSON), &Config)
	fmt.Println(err)
	Limit 		= Config.Limit 
	Rname 		= Config.Rname
	Sname 		= Config.Sname    
	Key 		= Config.Key     
	Respon 		= Config.Respon  
	Lkick 		= Config.Lkick    
	Lcancel 	= Config.Lcancel  
	Linvite 	= Config.Linvite 
	Lupdate 	= Config.Lupdate  
	Lcontact 	= Config.Lcontact 
	Lmention 	= Config.Lmention 
	Ljoin 		= Config.Ljoin 	 
	Lleave 		= Config.Lleave   
	Lockdown 	= Config.Lockdown 
	Logmode     = Config.Logmode
	Purge 		= Config.Purge 	 
	Mute 		= Config.Mute 	
	Blocked 	= Config.Blocked
	Logs        = Config.Logs
	Seler 		= Config.Seler    
	Owners 		= Config.Owners  
	Admins 		= Config.Admins  
	Staff 		= Config.Staff   
	Bots 		= Config.Bots    
	Center 		= Config.Center  
	Linkpro 	= Config.Linkpro 
	Denyinv 	= Config.Denyinv 
	Protect 	= Config.Protect  
	Namelock 	= Config.Namelock 
	Projoin 	= Config.Projoin  
	Limiter 	= Config.Limiter  
	Banned 		= Config.Banned  
	Gname 		= Config.Gname
	WordBan 	= Config.WordBan
	Stay 		= Config.Stay   
	Gmaster 	= Config.Gmaster
	Gowner 		= Config.Gowner 
	Gadmin 		= Config.Gadmin 
	Gban 		= Config.Gban   
	InRoom 		= Config.InRoom 
	XInRoom 	= Config.XInRoom
	TimeLeft    = Config.TimeLeft
	Msgsider    = Config.Msgsider
	Hiden       = Config.Hiden
}

func SaveJson() {
	Config.Limit 	= Limit
	Config.Rname 	= Rname
	Config.Sname  	= Sname
	Config.Key      = Key
	Config.Respon   = Respon
	Config.Lkick    = Lkick
	Config.Lcancel  = Lcancel
	Config.Linvite  = Linvite
	Config.Lupdate  = Lupdate
	Config.Lcontact = Lcontact
	Config.Lmention = Lmention
	Config.Ljoin 	= Ljoin
	Config.Lleave   = Lleave
	Config.Lockdown = Lockdown
	Config.Logmode  = Logmode
	Config.Purge 	= Purge
	Config.Mute 	= Mute
	Config.Blocked  = Blocked
	Config.Logs     = Logs
	Config.Seler    = Seler
	Config.Owners   = Owners
	Config.Admins   = Admins
	Config.Staff    = Staff
	Config.Bots     = Bots
	Config.Center   = Center
	Config.Linkpro  = Linkpro
	Config.Denyinv  = Denyinv
	Config.Protect  = Protect
	Config.Namelock = Namelock
	Config.Projoin  = Projoin
	Config.Limiter  = Limiter
	Config.Banned   = Banned
	Config.Gname    = Gname
	Config.WordBan  = WordBan
	Config.Stay     = Stay
	Config.Gmaster  = Gmaster
	Config.Gowner   = Gowner
	Config.Gadmin   = Gadmin
	Config.Gban     = Gban
	Config.InRoom   = InRoom
	Config.XInRoom  = XInRoom
	Config.TimeLeft = TimeLeft
	Config.Msgsider = Msgsider
	Config.Hiden    = Hiden   
	encjson, _ := json.MarshalIndent(Config, "", "  ")
	ioutil.WriteFile(DataBase, encjson, 0644)
}

type tagdata struct {
	S string `json:"S"`
	E string `json:"E"`
	M string `json:"M"`
}

type mentions struct {
	MENTIONEES []struct {
		Start string `json:"S"`
		End string `json:"E"`
		Mid string `json:"M"`
	}`json:"MENTIONEES"`
}

func StripOut(text string) string {
	text = strings.TrimPrefix(text, " ")
	text = strings.TrimSuffix(text, " ")
	return text
}

func GetInt(max int) (int, error) {
	if max <= 0 {
		return 0, fmt.Errorf("can't define input as <=0")
	}
	nbig, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return max, err
	}
	n := int(nbig.Int64())

	return n, err
}

func Choice(j []string) (string, error) {
	i, err := GetInt(len(j))
	if err != nil {
		return "", err
	}

	return j[i], nil
}

func InArrayInt(arr []int, str int) bool {
   	for a := 0; a < len(arr); a++ {
	  	if arr[a] == str {
		 	return true
		 	break
	  	}
   	}
   	return false
}

func contains(arr []string, str string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return true
			break
		}
	}
	return false
}

func uncontains(arr []string, str string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			return false
			break
		}
	}
	return true
}

func Remove(s []string, r string) []string {
	new := make([]string, len(s))
	copy(new, s)
	for i, v := range new {
		if v == r {
			return append(new[:i], new[i+1:]...)
		}
	}
	return s
}

func globalBl(to string, korban []string) {
	//_, found := Gban[to]
	//if found == false { Gban[to] = []string{} }
	for _, cox := range korban {
		if AllAccess(to, cox) > 9 {
			if uncontains(Gban[to], cox) {
				Gban[to] = append(Gban[to], cox)
			}
			SaveJson()
		}
	}
}

func Expired() bool {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	TimeNow := time.Now()
	now := TimeNow.In(loc)
	expired := now.After(TimeLeft.AddDate(0, 0, -12))
	if expired == true {
		return true
	}
	return false
}

func Getout() {
	fmt.Println("BOTS EXPIRED")
	for _, z := range Master {
		exp := fmt.Sprintf("%s Expired Please Tell Creator To Extend\nBots LeaveGroup and Clear Access", DB)
		talk.SendText(z, exp, 2)
	}
	for _, z := range Seler {
		exp := fmt.Sprintf("%s Expired Please Tell Creator To Extend\nBots LeaveGroup and Clear Access", DB)
		talk.SendText(z, exp, 2)
	}
	gr, _ := talk.GetGroupIdsJoined()
	for _, v := range gr {
		talk.SendText(v, "Bots Expired Please Tell Creator\nGood Bye..", 2)
		talk.LeaveGroup(v)
	}
	Owners   = []string{}
	Admins   = []string{}
	Staff    = []string{}
	Gmaster  = map[string][]string{}
	Gowner   = map[string][]string{}
	Gadmin   = map[string][]string{}
	SaveJson()
}

func AddedGban(to string, mid string) {
	//_, found := Gban[to]
	//if found == false { Gban[to] = []string{} }
	if uncontains(Gban[to], mid) && AllAccess(to, mid) > 9 {
		Gban[to] = append(Gban[to], mid)
		SaveJson()
	}
}

//********** CMDKEY *************//
func CmdKey(pesan string, rname string, sname string, key string, Mid string, MentionMsg []string) []string {
	var result = []string{}
	var txt string
	if strings.HasPrefix(pesan, rname+" ") {
		txt = strings.Replace(pesan, rname+" ", "", 1)
		if strings.Contains(txt, ",") {
			hasil := strings.Split(txt, ",")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else if strings.Contains(txt, "&") {
			hasil := strings.Split(txt, "&")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else if strings.Contains(txt, ";") {
			hasil := strings.Split(txt, ";")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else {
			result = []string{txt}
		}
	} else if strings.HasPrefix(pesan, sname+" ") {
		txt = strings.Replace(pesan, sname+" ", "", 1)
		if strings.Contains(txt, ",") {
			hasil := strings.Split(txt, ",")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else if strings.Contains(txt, "&") {
			hasil := strings.Split(txt, "&")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else if strings.Contains(txt, ";") {
			hasil := strings.Split(txt, ";")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else {
			result = []string{txt}
		}
	} else if strings.HasPrefix(pesan, key) {
		txt = strings.Replace(pesan, key, "", 1)
		if strings.Contains(txt, ",") {
			hasil := strings.Split(txt, ",")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else if strings.Contains(txt, "&") {
			hasil := strings.Split(txt, "&")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else if strings.Contains(txt, ";") {
			hasil := strings.Split(txt, ";")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else {
			result = []string{txt}
		}
	} else if strings.HasPrefix(pesan, sname) {
		txt = strings.Replace(pesan, sname, "", 1)
		if strings.Contains(pesan, ",") {
			hasil := strings.Split(txt, ",")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else if strings.Contains(pesan, "&") {
			hasil := strings.Split(txt, "&")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else if strings.Contains(pesan, ";") {
			hasil := strings.Split(txt, ";")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else {
			result = []string{txt}
		}
	} else if strings.HasPrefix(pesan, rname) {
		txt = strings.Replace(pesan, rname, "", 1)
		if strings.Contains(pesan, ",") {
			hasil := strings.Split(txt, ",")
			for _, id := range hasil {
				id = StripOut(id)
			}
		} else if strings.Contains(pesan, "&") {
			hasil := strings.Split(txt, "&")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else if strings.Contains(pesan, ";") {
			hasil := strings.Split(txt, ";")
			for _, id := range hasil {
				id = StripOut(id)
				result = append(result, id)
			}
		} else {
			result = []string{txt}
		}
	} else {
		if contains(MentionMsg, Mid) {
			pr, _ := talk.GetProfile()
			name := pr.DisplayName
			Vs := fmt.Sprintf("@%v", name)
			Vs = strings.ToLower(Vs)
			Vs = strings.TrimSuffix(Vs, " ")
			txt = strings.Replace(pesan, Vs, "", 1)
			txt = strings.TrimPrefix(txt, " ")
			for _, men := range MentionMsg {
				prs, _ := talk.GetContact(men)
				names := prs.DisplayName
				jj := fmt.Sprintf("@%v", names)
				jj = strings.ToLower(jj)
				jj = strings.TrimSuffix(jj, " ")
				txt = strings.Replace(txt, jj, "", 1)
				txt = strings.TrimPrefix(txt, " ")
			}

			if strings.Contains(txt, ",") {
				hasil := strings.Split(txt, ",")
				for _, id := range hasil {
					id = StripOut(id)
					result = append(result, id)
				}
			} else if strings.Contains(txt, "&") {
				hasil := strings.Split(txt, "&")
				for _, id := range hasil {
					id = StripOut(id)
					result = append(result, id)
				}
			} else if strings.Contains(txt, ";") {
				hasil := strings.Split(txt, ";")
				for _, id := range hasil {
					id = StripOut(id)
					result = append(result, id)
				}
			} else {
				result = []string{txt}
			}
		}
	}
	return result
}

//********** CMDLIST *************//
func CmdList(s string, list []string) []string {
	ln := len(list)
	ls := []int{}
	ind := []int{}
	hasil := []string{}
	if strings.Contains(s,",") {
		logics := strings.Split(s,",")
		fmt.Println(logics)
		for _,logic := range logics {
			fmt.Println(logic)
			if strings.Contains(logic,">") {
				su := strings.TrimPrefix(logic, ">")
				si,_ := strconv.Atoi(su)
				si -= 1
				for i := (si+1); i > si && i <= ln ; i++ {
					ls = append(ls,i)
				}
			} else if strings.Contains(logic,"<") {
				su := strings.TrimPrefix(logic, "<")
				si,_ := strconv.Atoi(su)
				si -= 1
				for i := 0; i <= si; i++ {
					ls = append(ls,i)
				}
			} else if strings.Contains(logic,"-") {
				las := strings.Split(logic,"-")
				si := las[0]
				siu,_ := strconv.Atoi(si)
				siu -= 1
				sa := las[1]
				sau,_ := strconv.Atoi(sa)
				sau -= 1
				for i := (siu); i >= siu && i <= sau ; i++ {
					ls = append(ls,i)
				}
			} else {
				sau,_ := strconv.Atoi(logic)
				sau -= 1
				ls = append(ls,sau)
			}
		}
	} else {
			logic := s
			if strings.Contains(logic,">") {
				su := strings.TrimPrefix(logic, ">")
				si,_ := strconv.Atoi(su)
				si -= 1
				for i := (si+1); i > si && i <= ln ; i++ {
					ls = append(ls,i)
				}
			} else if strings.Contains(logic,"<") {
				su := strings.TrimPrefix(logic, "<")
				si,_ := strconv.Atoi(su)
				si -= 1
				for i := 0; i <= si; i++ {
					ls = append(ls,i)
				}
			} else if strings.Contains(logic,"-") {
				las := strings.Split(logic,"-")
				si := las[0]
				siu,_ := strconv.Atoi(si)
				siu -= 1
				sa := las[1]
				sau,_ := strconv.Atoi(sa)
				sau -= 1
				for i := (siu); i >= siu && i <= sau ; i++ {
					ls = append(ls,i)
				}
			} else {
				sau,_ := strconv.Atoi(logic)
				sau -= 1
				ls = append(ls,sau)
			}
	}
	for _,do := range ls {
		if !InArrayInt(ind,do) && do < ln {
			jo := list[do]
			ind = append(ind,do)
			hasil = append(hasil,jo)
		}
	}
	return hasil
}
///==============================CHANNEL
func LoginChannel(channelId string) string {
	var ChannelResult *lord.ChannelToken
	var err error
	if service.IsLogin != true {
		panic("[Error]Not yet logged in.")
	}
	CS := service.ConnectChannel()
	ChannelResult, err = CS.ApproveChannelAndIssueChannelToken(context.TODO(), channelId)
	if err != nil {
		panic(err)
	}
	return ChannelResult.ChannelAccessToken
}

func SendReplyContact(relatedMessageId string, to string, mid string) {
	TS := service.TalkService()
	M := &lord.Message{
		From_: Myself,
		To: to,
		Text: "",
		ContentType: 13,
		ContentMetadata: map[string]string{"mid":mid},
		RelatedMessageServiceCode: 1,
		MessageRelationType: 3,
		RelatedMessageId: relatedMessageId,
	}
	_, e := TS.SendMessage(context.TODO(), int32(0), M)
	fmt.Println(e)
}

func SendReplyMessage(relatedMessageId string, to string, text string) {
	client := service.TalkService()
	M := &lord.Message {
		From_: Myself,
		To: to,
		Text: text,
		ContentType: 0,
		ContentMetadata: nil,
		RelatedMessageServiceCode: 1,
		MessageRelationType: 3,
		RelatedMessageId: relatedMessageId,
	}
	_, e := client.SendMessage(context.TODO(), int32(0), M)
	fmt.Println(e)
}

func SendTextMention(toID string,msgText string,mids []string) {
    client := service.TalkService()
    arr := []*tagdata{}
    mentionee := "@lord"
    texts := strings.Split(msgText, "@!")
    textx := ""
    for i := 0; i < len(mids); i++ {
        textx += texts[i]
        arr = append(arr, &tagdata{S: strconv.Itoa(len(textx)), E: strconv.Itoa(len(textx) + 5), M:mids[i]})
        textx += mentionee
    }
    textx += texts[len(texts)-1]
    allData,_ := json.MarshalIndent(arr, "", " ")
    msgObj := lord.NewMessage()
    msgObj.ContentType = 0
    msgObj.RelatedMessageId = "0"
    msgObj.To = toID
    msgObj.Text = textx
    msgObj.ContentMetadata = map[string]string{"MENTION": "{\"MENTIONEES\":"+string(allData)+"}"}
    _, e := client.SendMessage(context.TODO(), int32(0), msgObj)
    fmt.Println(e)
}

func SendReplyMention(relatedMessageId string,toID string,msgText string,mids []string) {
    client := service.TalkService()
    arr := []*tagdata{}
    mentionee := "@lord"
    texts := strings.Split(msgText, "@!")
    textx := ""
    for i := 0; i < len(mids); i++ {
        textx += texts[i]
        arr = append(arr, &tagdata{S: strconv.Itoa(utf8.RuneCountInString(textx)), E: strconv.Itoa(utf8.RuneCountInString(textx) + 5), M:mids[i]})
        textx += mentionee
    }
    textx += texts[len(texts)-1]
    allData,_ := json.MarshalIndent(arr, "", " ")
    M := &lord.Message{
        From_: Myself,
        To: toID,
        Text: textx,
        ContentType: 0,
        ContentMetadata: map[string]string{"MENTION": "{\"MENTIONEES\":"+string(allData)+"}"},
        RelatedMessageServiceCode: 1,
        MessageRelationType: 3,
        RelatedMessageId: relatedMessageId,
    }
    _, e := client.SendMessage(context.TODO(), int32(0), M)
    fmt.Println(e)
}

func SendReplyMentionByList(relatedMessageId string,to string,msgText string,targets []string){
    listMid := targets
    listMid2 := []string{}
    listChar := msgText
    listNum := 0
    loopny := len(listMid) / 20 + 1
    limiter := 0
    limiter2 := 20
    for a:=0;a<loopny;a++{
        for c:=limiter;c<len(listMid);c++{
            if c < limiter2{
                listNum = int(listNum) + 1
                listMid2 = append(listMid2,listMid[c])
                limiter = limiter + 1
            }else{
                limiter2 = limiter + 20
                break
            }
        }
        SendReplyMention(relatedMessageId,to,listChar,listMid2)
        listChar = ""
        listMid2 = []string{}
    }
}

func SendReplyMentionByList2(relatedMessageId string,to string,msgText string,targets []string){
    listMid := targets
    listMid2 := []string{}
    listChar := msgText
    listNum := 0
    loopny := len(listMid) / 20 + 1
    limiter := 0
    limiter2 := 20
    for a:=0;a<loopny;a++{
        for c:=limiter;c<len(listMid);c++{
            if c < limiter2{
                listNum = int(listNum) + 1
                listChar += "\n" + strconv.Itoa(listNum) + ". @!"
                listMid2 = append(listMid2,listMid[c])
                limiter = limiter + 1
            }else{
                limiter2 = limiter + 20
                break
            }
        }
        SendReplyMention(relatedMessageId,to,listChar,listMid2)
        listChar = ""
        listMid2 = []string{}
    }
}

func SendTextMentionByList(to string,msgText string,targets []string){
    listMid := targets
    listMid2 := []string{}
    listChar := msgText
    listNum := 0
    loopny := len(listMid) / 20 + 1
    limiter := 0
    limiter2 := 20
    for a:=0;a<loopny;a++{
        for c:=limiter;c<len(listMid);c++{
            if c < limiter2{
                listNum = int(listNum) + 1
                listMid2 = append(listMid2,listMid[c])
                limiter = limiter + 1
            }else{
                limiter2 = limiter + 20
                break
            }
        }
        SendTextMention(to,listChar,listMid2)
        listChar = ""
        listMid2 = []string{}
    }
}

// **send message multi mentions** //
func SendTextMentionByList2(to string,msgText string,targets []string){
    listMid := targets
    listMid2 := []string{}
    listChar := msgText
    listNum := 0
    loopny := len(listMid) / 20 + 1
    limiter := 0
    limiter2 := 20
    for a:=0;a<loopny;a++{
        for c:=limiter;c<len(listMid);c++{
            if c < limiter2{
                listNum = int(listNum) + 1
                listChar += "\n" + strconv.Itoa(listNum) + ". @!"
                listMid2 = append(listMid2,listMid[c])
                limiter = limiter + 1
            }else{
                limiter2 = limiter + 20
                break
            }
        }
        SendTextMention(to,listChar,listMid2)
        listChar = ""
        listMid2 = []string{}
    }
}
//
func AllBanned(to string, from string) bool {
	if contains(Gban[to], from) == true || contains(Banned, from) == true {
		return true
	}
	return false
}
func IsGban(to string, from string) bool {
	if contains(Gban[to], from) == true {
		return true
	}
	return false
}

func IsBanned(from string) bool {
	if contains(Banned, from) == true {
		return true
	}
	return false
}

func IsHiden(from string) bool {
	if contains(Hiden, from) == true {
		return true
	}
	return false
}

func IsMaster(from string) bool {
	if contains(Master, from) == true {
		return true
	}
	return false
}

func IsSeler(from string) bool {
	if contains(Seler, from) == true {
		return true
	}
	return false
}

func IsOnwer(from string) bool {
	if contains(Owners, from) == true {
		return true
	}
	return false
}

func IsAdmin(from string) bool {
	if contains(Admins, from) == true {
		return true
	}
	return false
}

func IsStaff(from string) bool {
	if contains(Staff, from) == true {
		return true
	}
	return false
}

func IsCenter(from string) bool {
	if contains(Center, from) == true {
		return true
	}
	return false
}

func IsBots(from string) bool {
	if contains(Bots, from) == true {
		return true
	}
	return false
}

func IsBackup(from string) bool {
	if contains(Backup, from) == true {
		return true
	}
	return false
}

func IsXBackup(from string) bool {
	if contains(XBackup, from) == true {
		return true
	}
	return false
}

func IsLimit(from string) bool {
	if contains(Limiter, from) == true {
		return true
	}
	return false
}

func IsGmaster(to string, from string) bool {
	if contains(Gmaster[to], from) == true {
		return true
	}
	return false
}

func IsInRoom(to string, from string) bool {
	if contains(InRoom[to], from) == true {
		return true
	}
	return false
}

func IsGonwer(to string, from string) bool {
	if contains(Gowner[to], from) == true {
		return true
	}
	return false
}

func IsGadmin(to string, from string) bool {
	if contains(Gadmin[to], from) == true {
		return true
	}
	return false
}

func IsWordText(to string, text string) bool {
	if contains(WordBan[to], text) == true {
		return true
	}
	return false
}

func IsAccept(from []string) bool {
	if contains(from, Myself) == true{
		return true
	}
	return false
}

func callProfile(cmd1 string, cmd2 string) {
	cmd, _ := exec.Command("python3","LINE/profile/profile.py",service.AuthToken,Myself,cmd1,cmd2).Output()
	fmt.Println("\033[33m"+string(cmd)+"\033[39m")
}

func nodejs(gid string,meta string, cms string){
	cmo :=fmt.Sprintf("node /root/test/kickAndCancel.js token=%s app=IOS gid=%s method=%s %s",service.AuthToken, gid, meta, cms)
	parts := strings.Fields(cmo)
	cmd, _ := exec.Command(parts[0],parts[1:]...).Output()
	fmt.Println(string(cmd))
}

func AddedText(msg string, babi string, sname string, rname string) string {
	var str string
	su := babi
	if strings.Contains(msg, rname+" ") {
		str = strings.Replace(msg, rname+" "+su+" ", "", 1)
	} else if strings.Contains(msg, sname+" ") {
		str = strings.Replace(msg, sname+" "+su+" ", "", 1)
	} else if strings.Contains(msg, rname) {
		str = strings.Replace(msg, rname+su+" ", "", 1)
	} else if strings.Contains(msg, sname) {
		str = strings.Replace(msg, sname+su+" ", "", 1)
	}
	return str
}

func AllAccessV1(to string, from string) int {
	if contains(Master, from) {
		return 0
	} else if contains(Seler, from) {
		return 1
	} else if contains(Owners, from) {
		return 2
	} else if contains(Admins, from) {
		return 3
	} else if contains(Staff, from) {
		return 4
	} else if contains(Gmaster[to], from) {
		return 5
	} else if contains(Gowner[to], from) {
		return 6
	} else if contains(Gadmin[to], from) {
		return 7
	} else if contains(Bots, from) {
		return 8
	} else if contains(Center, from) {
		return 9
	}
	return 100
}

func AllAccess(to string, from string) int {
	if contains(Master, from) {
		return 0
	} else if contains(Seler, from) {
		return 1
	} else if contains(Owners, from) {
		return 2
	} else if contains(Admins, from) {
		return 3
	} else if contains(Staff, from) {
		return 4
	} else if contains(Gmaster[to], from) {
		return 5
	} else if contains(Gowner[to], from) {
		return 6
	} else if contains(Gadmin[to], from) {
		return 7
	} else if contains(Center, from) {
		return 8
	} else if contains(Bots, from) {
		return 9
	} else if contains(Backup, from) {
		return 10
	}
	return 100
}

func InvitedAjs(param1 string) {
	if !IsCenter(Myself) {
		for _, v := range Center {
			if IsPending(v, param1) == true {
				fmt.Println("Already On InRoom")
			} else {
				cok := ExecuteClient(param1)
				if Myself == cok {
					err := talk.InviteIntoGroupV2(param1, Center)
					if err != nil {
						talk.SendText(param1, "i Limit", 2)
					}
				}
			}
		}
	}
}

//// ================= BACKUP
func GenerateTimeLog(id string, to string){
	loc, _ := time.LoadLocation("Asia/Jakarta")
	a:=time.Now().In(loc)
	yyyy := strconv.Itoa(a.Year())
	MM := a.Month().String()
	dd := strconv.Itoa(a.Day())
	hh := a.Hour()
	mm := a.Minute()
	ss := a.Second()
	var hhconv string
	var mmconv string
	var ssconv string
	if hh < 10 {
		hhconv = "0"+strconv.Itoa(hh)
	}else {
		hhconv = strconv.Itoa(hh)
	}
	if mm < 10 {
		mmconv = "0"+strconv.Itoa(mm)
	}else {
		mmconv = strconv.Itoa(mm)
	}
	if ss < 10 {
		ssconv = "0"+strconv.Itoa(ss)
	}else {
		ssconv = strconv.Itoa(ss)
	}
	times := "Date : "+dd+"-"+MM+"-"+yyyy+"\nTime : "+hhconv+":"+mmconv+":"+ssconv
	SendReplyMessage(id, to, times)
}

func DueDate(id string, to string, a time.Time){
	yyyy := strconv.Itoa(a.Year())
	MM := a.Month().String()
	dd := strconv.Itoa(a.Day())
	hh := a.Hour()
	mm := a.Minute()
	ss := a.Second()
	var hhconv string
	var mmconv string
	var ssconv string
	if hh < 10 {
		hhconv = "0"+strconv.Itoa(hh)
	}else {
		hhconv = strconv.Itoa(hh)
	}
	if mm < 10 {
		mmconv = "0"+strconv.Itoa(mm)
	}else {
		mmconv = strconv.Itoa(mm)
	}
	if ss < 10 {
		ssconv = "0"+strconv.Itoa(ss)
	}else {
		ssconv = strconv.Itoa(ss)
	}
	times := "**Time Left:\nDate : "+dd+"-"+MM+"-"+yyyy+"\nTime : "+hhconv+":"+mmconv+":"+ssconv
	SendReplyMessage(id, to, times)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	return fmt.Sprintf("ʀᴜɴᴛɪᴍᴇ: %02dD %02dH %02dM", h/24, h%24, m)
}

func JoinQrV2(to string) {
	fmt.Println("JoinQrV2")
	ticket,_ := talk.ReissueGroupTicket(to)
	grup,_ := talk.GetCompactGroup(to)
	if grup.PreventedJoinByTicket {
		grup.PreventedJoinByTicket = false
		talk.UpdateGroup(grup)
	}
	for _, o := range Bots {
		if !IsCenter(o) {
			talk.SendMessage(o, "jointicket_ "+to+" "+ticket, map[string]string{})
		}
	}
	fmt.Println("jointicket_ "+to+" "+ticket)
	time.Sleep(500 * time.Millisecond)
	grup.PreventedJoinByTicket = true
	talk.UpdateGroup(grup)
}

func JoinAj(to string, from string) {
	ticket,_ := talk.ReissueGroupTicket(to)
	grup,_ := talk.GetCompactGroup(to)
	if grup.PreventedJoinByTicket {
		grup.PreventedJoinByTicket = false
		talk.UpdateGroup(grup)
	}
	talk.SendMessage(from, "jointicket_ "+to+" "+ticket, map[string]string{})
	fmt.Println("jointicket_ "+to+" "+ticket)
	time.Sleep(500 * time.Millisecond)
	grup.PreventedJoinByTicket = true
	talk.UpdateGroup(grup)
}

func addAsFriendContact(target string){
	if nothingInMyContacts(target){
		talk.FindAndAddContactsByMid(target)
	}
}

func nothingInMyContacts(target string) bool {
	friends,_ := talk.GetAllContactIds()
	if uncontains(friends, target){
		return true
	}
	return false
}

////////// BACKUP
func InMembersV2(from string, groups string) bool {
	runtime.GOMAXPROCS(runtime.NumCPU())
	res,_ := talk.GetCompactGroup(groups)
	memlist := res.Members
	for a := 0; a < len(memlist); a++ {
		if memlist[a].Mid == from {
			return true
			break
		}
	}
	return false
}


func IsPending(from string, groups string) bool {
	res, _ := talk.GetCompactGroup(groups)
	memlist := res.Invitee
	for _, a := range memlist {
		if a.Mid == from {
			return true
			break
		}
	}
	return false
}

func IsKillMode(groups string, slayer string, victim string) {
	if IsBots(victim) && AllAccess(groups, slayer) < 10 {
		boom, _ := talk.GetContact(slayer)
		cokss := boom.DisplayName
		var str string
		if strings.Contains(cokss, "1") {
			str = strings.Replace(cokss, "1", "", 1)
		} else if strings.Contains(cokss, "2") {
			str = strings.Replace(cokss, "2", "", 1)
		} else if strings.Contains(cokss, "3") {
			str = strings.Replace(cokss, "3", "", 1)
		} else if strings.Contains(cokss, "4") {
			str = strings.Replace(cokss, "4", "", 1)
		} else if strings.Contains(cokss, "5") {
			str = strings.Replace(cokss, "5", "", 1)
		} else if strings.Contains(cokss, "6") {
			str = strings.Replace(cokss, "6", "", 1)
		} else if strings.Contains(cokss, "7") {
			str = strings.Replace(cokss, "7", "", 1)
		} else if strings.Contains(cokss, "8") {
			str = strings.Replace(cokss, "8", "", 1)
		} else if strings.Contains(cokss, "9") {
			str = strings.Replace(cokss, "9", "", 1)
		} else if strings.Contains(cokss, "0") {
			str = strings.Replace(cokss, "0", "", 1)
		} else {
			str = cokss
		}
		res, _ := talk.GetGroup(groups)
		memlist := res.Members
		cokz := []string{}
		for _, a := range memlist {
			if strings.Contains(a.DisplayName, str) {
				fmt.Println("INI NAMA: " + a.DisplayName)
				fmt.Println("INI STR: " + str)
				cokz = append(cokz, a.Mid)
			} else {
				fmt.Println("Notting")
			}
		}
		for _, v := range cokz {
			if AllAccess(groups, v) < 9 {
				if v != Myself {
					go func(v string) {
						talk.KickoutFromGroup(groups, []string{v})
					}(v)
					go func(v string) {
						AddedGban(groups, v)
					}(v)
				}
			}
		}
	}
}
func InMembers(from string, groups string) bool {
	runtime.GOMAXPROCS(runtime.NumCPU())
	res,_ := talk.GetCompactGroup(groups)
	memlist := res.Members
	for _, a := range memlist {
		if a.Mid == from {
			return true
			break
		}
	}
	return false
}

func BackupV2(to string, mid string, korban string) {
	//for x := 0; x < len(Bots); x++ {
	for x := range Bots {
		if InMembers(Bots[x], to) {
			if Myself == Bots[x] {
				go func() { talk.KickoutFromGroup(to, []string{mid}) }()
				go func() { talk.InviteIntoGroup(to, []string{korban}) }()
			}
			break
		} else {
			continue
		}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func InvitesDel(pelaku string){
	Invites = Remove(Invites, pelaku)
}

func checkEqual(to string, korban []string) bool {
	_, found := Gban[to]
	if found == true {
		//for a := 0; a < len(korban); a++ {
		for a := range korban {
			if contains(Gban[to], korban[a]) {
				return true
				break
			}
		}
	}
	return false
}

func checkEqualV1(korban []string) bool {
	//for a := 0; a < len(korban); a++ {
	for a := range korban {
		if contains(Banned, korban[a]) {
			return true
			break
		}
	}
	return false
}

func PurgeCancelV2(to string) {
	_, found := Gban[to]
	if found == true {
		//for x := 0; x < len(Bots); x++ {
		for x := range Bots {
			if InMembers(Bots[x], to) {
				if Bots[x] == Myself {
					var wg sync.WaitGroup
					wg.Add(len(Gban[to]))
					for i := 0; i < len(Gban[to]); i++ {
						go func(i int) {
							defer wg.Done()
							talk.CancelGroupInvitation(to, []string{Gban[to][i]})
						}(i)
					}
					wg.Wait()
				}
				break
			} else {
				continue
			}
		}
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}
func PurgeCancelV1(to string) {
	//for x := 0; x < len(Bots); x++ {
	for x := range Bots {
		if InMembers(Bots[x], to) {
			if Bots[x] == Myself {
				var wg sync.WaitGroup
				wg.Add(len(Banned))
				for i := 0; i < len(Banned); i++ {
					go func(i int) {
						defer wg.Done()
						talk.CancelGroupInvitation(to, []string{Banned[i]})
					}(i)
				}
				wg.Wait()
			}
			break
		} else {
			continue
		}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func FastCancelV3(to string, korban []string) {
	data, err := Choice(Bots)
	if err != nil {
		fmt.Println(err)
	} else if data == Myself {
		for i := 0; i < len(korban); i++ {
			go func(i int) {
				talk.CancelGroupInvitation(to, []string{korban[i]})
			}(i)
		}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func InKick(to string, from string) {
	//for z := 0; z < len(Bots); z++ {
	for z := range Bots {
		if InMembers(Bots[z], to) {
			if Bots[z] == Myself && Blocked == false {
				go func() { talk.KickoutFromGroup(to, []string{from}) }()
			}
			break
		} else {
			continue
		}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func InCancel(to string, from string) {
	//for z := 0; z < len(Bots); z++ {
	for z := range Bots {
		if InMembers(Bots[z], to) {
			if Bots[z] == Myself && Blocked == false {
				go func() { talk.CancelGroupInvitation(to, []string{from}) }()
			}
			break
		} else {
			continue
		}
	}
}

func InInvite(to string, from string) {
	//for z := 0; z < len(Bots); z++ {
	for z := range Bots {
		if InMembers(Bots[z], to) {
			if Bots[z] == Myself && Blocked == false {
				go func() { talk.InviteIntoGroup(to, []string{from}) }()
			}
			break
		} else {
			continue
		}
	}
}

func InKickV1(to string, from []string) {
	//for z := 0; z < len(Bots); z++ {
	for z := range Bots {
		if InMembers(Bots[z], to) {
			if Bots[z] == Myself && Blocked == false {
				for a := 0; a < len(from); a++ {
					go func() { talk.KickoutFromGroup(to, []string{from[a]}) }()
				}
			}
			break
		} else {
			continue
		}
	}
}

func InCancelV1(to string, korban []string) {
	for x := range Bots {
		if InMembers(Bots[x], to) {
			if Bots[x] == Myself {
				for i := 0; i < len(korban); i++ {
					go func(i int) {
						talk.CancelGroupInvitation(to, []string{korban[i]})
					}(i)
				}
			}
			break
		} else {
			continue
		}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func InInviteV1(to string, from []string) {
	//for z := 0; z < len(Bots); z++ {
	for z := range Bots {
		if InMembers(Bots[z], to) {
			if Bots[z] == Myself && Blocked == false {
				for a := 0; a < len(from); a++ {
					go func() { talk.InviteIntoGroup(to, []string{from[a]}) }()
				}
			}
			break
		} else {
			continue
		}
	}
}

func Autopurge(to string) {
	g, _ := talk.GetCompactGroup(to)
	memb := g.Members
	for z := 0; z < len(memb); z++ {
		fck := memb[z].Mid
		if AllBanned(to, fck) == true {
			go func(fck string) {
				talk.KickoutFromGroup(to, []string{fck})
			}(fck)
			if z >= Limit {
				break
			}
		}
	}
}

func NUKEJOIN(to string) {
	g, _ := talk.GetCompactGroup(to)
	memb := g.Members
	var batas = 150
	cms := ""
	for z := 0; z < len(memb); z++ {
		fck := memb[z].Mid
		if AllAccess(to, fck) > 10 {
			cms += fmt.Sprintf(" uid=%s",fck)
			if z >= batas {
				break
			}
		}
	}
	talk.SendMessage(to, "🌪️🌪️🌪️ Just Some Casual Cleanse 🌪️🌪️🌪️", map[string]string{})
	nodejs(to,"kick",cms)
}

func CLEANSEMEMBERS(to string) {
	g, _ := talk.GetCompactGroup(to)
	memb := g.Members
	var kbatas = 150
	var cbatas = 15
	cms := ""
	for z := 0; z < len(memb); z++ {
		fck := memb[z].Mid
		if AllAccess(to, fck) > 10 {
			cms += fmt.Sprintf(" uid=%s",fck)
			if z >= kbatas {
				break
			}
		}
	}
	invit := g.Invitee
	for z := 0; z < len(invit); z++ {
		fck := invit[z].Mid
		if AllAccess(to, fck) > 10 {
			cms += fmt.Sprintf(" uids=%s",fck)
			if z >= cbatas {
				break
			}
		}
	}
	talk.SendMessage(to, "🌪️🌪️🌪️ Just Some Casual Cleanse 🌪️🌪️🌪️", map[string]string{})
	nodejs(to,"kickandcancel",cms)
}

func ProLink(group string, p2 string){
	runtime.GOMAXPROCS(runtime.NumCPU())
	go func(){InKick(group,p2)}()
	go func(){g,_:= talk.GetGroupWithoutMembers(group);if g.PreventedJoinByTicket == false{g.PreventedJoinByTicket = true;talk.UpdateGroup(g)}}()
	go func(){AddedGban(group, p2)}()
}

func proNameGroup(group string, pl string, kr string){
	if kr == "1"{
		go func(){
			InKick(group,pl)
		}()
		go func(){
			ChangeGname(group,[]string{pl})
		}()
		go func(){
			AddedGban(group, pl)
		}()
	}
}
func ChangeGname(group string, p2 []string) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go func(){
		for i := range Bots {
			if InMembers(Bots[i], group) {
				if Myself == Bots[i] {
					for x := range p2 {
						GnamePoll(group, []string{p2[x]})
					}
				}
				break
			}else{
				continue
			}
		}
	}()
}
func GnamePoll(lc string,pd []string) {
	var wg sync.WaitGroup
	wg.Add(len(pd))
	for i:=0;i<len(pd);i++ {
		go func(i int) {
			defer wg.Done()
			g,_ := talk.GetGroupWithoutMembers(lc)
			if g.Name != Gname[lc]{
				g.Name = Gname[lc]
				talk.UpdateGroup(g)
			}
		}(i)
	}
	wg.Wait()
}

func AjsCoks(groups string, slayer string, victim string) {
	runtime.GOMAXPROCS(4)
	if AllAccessV1(groups, slayer) > 8 {
		if AllAccessV1(groups, victim) < 9 {
			go func() {
				talk.AcceptGroupInvitation(groups)
			}()
			JoinQrV2(groups)
			go func() {
				addAsFriendContact(victim)
			}()
			go func() {
				talk.InviteIntoGroup(groups, []string{victim})
			}()
			go func() {
				talk.KickoutFromGroup(groups, []string{slayer})
				talk.LeaveGroup(groups)
			}()
			go func() {
				AddedGban(groups, slayer)
			}()
		}
	} else {
		go func() {
			talk.AcceptGroupInvitation(groups)
		}()
		JoinQrV2(groups)
		go func() {
			addAsFriendContact(victim)
		}()
		go func() {
			talk.InviteIntoGroup(groups, []string{victim})
		}()
		go func() {
			talk.LeaveGroup(groups)
		}()
	}
}

func InMessage(id string, to string, reply bool, message string) {
	for z := range Bots {
		if InMembers(Bots[z], to) {
			if Bots[z] == Myself {
				if reply == true {
					SendReplyMessage(id, to, message)
				} else {
					talk.SendMessage(to, message, map[string]string{})
				}
			}
			break
		} else {
			continue
		}
	}
}
func ExecuteClient(to string) string {
	var saya string
	for z := range Bots {
		if InMembers(Bots[z], to) {
			if Bots[z] == Myself {
				saya = Bots[z]
			}
			break
		} else {
			continue
		}
	}
	return saya
}

func print(message string) {
	fmt.Println(message)
}

func IsBlockKick(from string) bool {
	if helper.InArray(Protect, from) == true {
		return true
	}
	return false
}

func IsBlockInvite(from string) bool {
	if helper.InArray(Denyinv, from) == true {
		return true
	}
	return false
}

func IsBlockQr(from string) bool {
	if helper.InArray(Linkpro, from) == true {
		return true
	}
	return false
}

func IsBlockName(from string) bool {
	if helper.InArray(Namelock, from) == true {
		return true
	}
	return false
}

func IsBlockJoin(from string) bool {
	if helper.InArray(Projoin, from) == true {
		return true
	}
	return false
}
///====================================
func Executor(op *lord.Operation) {
	var param1, param2, param3, operation = op.Param1, op.Param2, op.Param3, op.Type
	if operation == 19 || operation == 133 { // OpType_NOTIFIED_KICKOUT_FROM_GROUP
		Lkick = param3
		if Myself == param3 {
			if AllAccess(param1, param2) > 10 {
				go func() {
					AddedGban(param1, param2)
				}()
				WarMode = true
			}
		}
		if IsBots(param3) && AllAccess(param1, param2) > 10 {
			go func() { 
				BackupV2(param1, param2, param3)
			}()
			go func() {
				AddedGban(param1, param2)
			}()
		} 
		if AllBanned(param1, param2) {
			go func() {
				InInvite(param1, param3)				
			}()
			go func() {
				InKick(param1, param2)
			}()
		} 
		if AllAccess(param1, param3) < 8 && AllAccess(param1, param2) > 10 {
			go func() {
				BackupV2(param1, param2, param3)
			}()
			go func() {
				AddedGban(param1, param2)
			}()
		} 
		if IsBlockKick(param1) && AllAccess(param1, param2) > 10 {
			go func() {
				InKick(param1, param2)
			}()
			go func() {
				AddedGban(param1, param2)
			}()
			if AllAccess(param1, param2) < 10 {
				go func(){
					addAsFriendContact(param3)
				}()
				go func() {
					InInvite(param1, param3)
				}()
			}
		} else if Logs == true && WarMode == false {
			ListMid := []string{}
			if AllAccess(param1, param3) < 8  || AllAccess(param1, param2) < 8 {
				ListMid = append(ListMid, param2)
				ListMid = append(ListMid, param3)
				xa, _ := talk.GetGroupWithoutMembers(param1)
				text := fmt.Sprintf("** NOTIFIED_KICKOUT_FROM_GROUP:\n Group: %s\n From: @!\n To: @! ", xa.Name)
				SendTextMention(Logmode, text, ListMid)
				ListMid = ListMid[:0]
			}
		}
		if IsCenter(Myself) == true {
			go func() {
				AjsCoks(op.Param1, op.Param2, op.Param3)
			}()
		}
	}
	if operation == 13 || operation == 124 { //OpType_NOTIFIED_INVITE_INTO_GROUP
		Linvite = param3
		korban := strings.Split(param3, "\x1e")
		if IsBots(param2) || IsCenter(param2) && IsAccept(korban) {
			if IsCenter(Myself) == false {
				go func() { 
					talk.AcceptGroupInvitation(param1) 
				}()
				go func() { 
					Autopurge(param1)
				}()
			}
		} else if IsAccept(korban) && AllAccess(param1, param2) == PInvBots {
			expired := Expired()
			if expired == true {
				go talk.AcceptGroupInvitation(param1)
				if PInvBots != 0 {
					talk.LeaveGroup(param1)
				}
			} else {
				if IsCenter(Myself) == false {
					if JoinNuke == true {
						go talk.AcceptGroupInvitation(param1)
						go NUKEJOIN(param1)
					} else {
						go talk.AcceptGroupInvitation(param1)
					}
				}
			}
		}
		if checkEqual(param1, korban) || checkEqualV1(korban) && AllAccess(param1, param2) > 9 {
			go func() {
				InKick(param1, param2)
			}()
			if Lockdown == true {
				go func() { PurgeCancelV2(param1) }()
			} else {
				go func() { InCancelV1(param1, korban) }()
			}
			go func() {
				AddedGban(param1, param2)
			}()
		} 
		if AllBanned(param1, param2) {
			go func() {
				InKick(param1, param2)
			}()
			if Lockdown == true {
				go func() { PurgeCancelV2(param1) }()
			} else {
				go func() { InCancelV1(param1, korban) }()
			}
			go func() {
				globalBl(param1, korban)
			}()
		} 
		if IsBlockInvite(param1) && AllAccess(param1, param2) > 9 {
			go func() {
				InCancelV1(param1, korban)
			}()
			go func() {
				InKick(param1, param2)
			}()
			go func() {
				AddedGban(param1, param2)
			}()
			go func() {
				globalBl(param1, korban)
			}()
		} else if Logs == true && WarMode == false {
			ListMid := []string{}
			if AllAccess(param1, param3) < 8 || AllAccess(param1, param2) < 8 {
				ListMid = append(ListMid, param2)
				ListMid = append(ListMid, param3)
				xa, _ := talk.GetGroupWithoutMembers(param1)
				text := fmt.Sprintf("** NOTIFIED_INVITE_INTO_GROUP:\n Group: %s\n From: @!\n To: @! ", xa.Name)
				SendTextMention(Logmode, text, ListMid)
				ListMid = ListMid[:0]
			}
		}
	}
	if operation == 16 || operation == 129 { //OpType_ACCEPT_GROUP_INVITATION
		if len(Gban[param1]) >= 1 {
			go func() {
				Autopurge(param1)
			}()
			if len(Gban[param1]) >= 3 {
				for x := 0; x < len(Bots); x++ {
					if InMembers(Bots[x], param1) == false {
						talk.InviteIntoGroup(param1, []string{Bots[x]})
						break
					}
				}
			}
		}
	}
	if operation == 11 || operation == 122 { // OpType_NOTIFIED_UPDATE_GROUP
		Lupdate = param2
		if IsGban(param1, param2) || IsBanned(param2) {
			go func(){ProLink(param1, param2)}()
		}
		if contains(Namelock, param1) && AllAccess(param1, param2) > 10 {
			go func(){proNameGroup(param1, param2, param3)}()
		}
		if contains(Linkpro, param1) && AllAccess(param1, param2) > 10 {
			go func(){ProLink(param1, param2)}()
		}
		if Logs == true && WarMode == false {
			ListMid := []string{}
			if AllAccess(param1, param2) < 8 {
				ListMid = append(ListMid, param2)
				xa, _ := talk.GetGroupWithoutMembers(param1)
				text := fmt.Sprintf("** NOTIFIED_UPDATE_GROUP:\n Group: %s\n From: @!", xa.Name)
				SendTextMention(Logmode, text, ListMid)
				ListMid = ListMid[:0]
			}
		}
	}
	if operation == 17 || operation == 130 { // OpType_NOTIFIED_ACCEPT_GROUP_INVITATION
		Ljoin = param2
		if AllBanned(param1, param2) {
			go func() {
				InKick(param1, param2)
			}()
		}
		if contains(Projoin, param1) && AllAccess(param1, param2) > 10 {
			go func() {
				InKick(param1, param2)
			}()
		}
		if contains(Invites, param2) {
			go func() {
				InKick(param1, param2)
			}()
			go func(){InvitesDel(param2)}()
		}
		if Logs == true && WarMode == false {
			ListMid := []string{}
			if AllAccess(param1, param2) < 8 {
				ListMid = append(ListMid, param2)
				xa, _ := talk.GetGroupWithoutMembers(param1)
				text := fmt.Sprintf("** NOTIFIED_ACCEPT_GROUP_INVITATION:\n Group: %s\n From: @!", xa.Name)
				SendTextMention(Logmode, text, ListMid)
				ListMid = ListMid[:0]
			}
		}

	}
	if operation == 15 || operation == 128 { // OpType_NOTIFIED_LEAVE_GROUP
		Lleave = param2
		if IsCenter(param2) {
			go func() {
				InvitedAjs(param1)
			}()
		}
		if Logs == true && WarMode == false {
			ListMid := []string{}
			if AllAccess(param1, param2) < 8 {
				ListMid = append(ListMid, param2)
				xa, _ := talk.GetGroupWithoutMembers(param1)
				text := fmt.Sprintf("** NOTIFIED_LEAVE_GROUP:\n\n Group: %s\n From: @!", xa.Name)
				SendTextMention(Logmode, text, ListMid)
				ListMid = ListMid[:0]
			}
		}
	}
	if operation == 32 || operation == 126 { // OpType_NOTIFIED_CANCEL_INVITATION_GROUP
		Lcancel = param3
		if param3 == Myself {
			if AllAccess(param1, param2) > 10 {
				go func() {
					AddedGban(param1, param2)
				}()
			}
			return
		}
		if AllBanned(param1, param2) {
			go func() {
				BackupV2(param1, param2, param3)
			}()
		} 
		if IsBots(param3) && AllAccess(param1, param2) > 9 {
			go func() {
				BackupV2(param1, param2, param3)
			}()
			go func() {
				AddedGban(param1, param2)
			}()
		} 
		if AllAccess(param1, param3) < 9 {
			go func() {
				BackupV2(param1, param2, param3)
			}()
			go func() {
				AddedGban(param1, param2)
			}()
		} 
		if IsBlockInvite(param1) && AllAccess(param1, param2) > 9 {
			go func() {
				InKick(param1, param2)
			}()
			go func() {
				AddedGban(param1, param2)
			}()
		} else if IsCenter(param3) && AllAccess(param1, param2) > 9 {
			go func() {
				InInvite(param1, param3)
			}()
			go func() {
				JoinQrV2(param1)
			}()
			go func() {
				InKick(param1, param2)
			}()
		}

	}
	if operation == 5 { // OpType_NOTIFIED_ADD_CONTACT
		fmt.Println(op)
		if AllAccess(param1, param1) < 10 {
			addAsFriendContact(param1)
		}
		if Logs == true {
			ListMid := []string{}
			if AllAccess(param1, param1) < 10 {
				ListMid = append(ListMid, param1)
				text := "** NOTIFIED_ADD_CONTACT:\n From: @! "
				SendTextMention(Logmode, text, ListMid)
				ListMid = ListMid[:0]
			}
		}
	}
	if operation == 55 || operation == 28 {
		if SiderV2[param1] == true {
			if helper.InArray(Sider[param1], param2) == false {
				if !IsHiden(param2) {
					asu := "** @!\n"
					if Msgsider != "" {
						jaran := asu + Msgsider
						SendTextMention(param1, jaran, []string{param2})
					} else {
						cok := "Hello @!\n I See You!!"
						SendTextMention(param1, cok, []string{param2})
					}
					Sider[op.Param1] = append(Sider[op.Param1], op.Param2)
				}
			}
		}
	}
	if operation == 26 {
	 expired := Expired()
	 if expired == true && AllAccess(op.Message.To, op.Message.From_) > 0 {
	 	Getout()
	 } else {
		var cmds = []string{}
		//var silent = false
		msg := op.Message
	    text := msg.Text
	    var Tcm = []string{}
	    var MentionMsg = helper.MentionList(op)
	    var txt string
	    var pesan = strings.ToLower(text)
	    if Logs == true {
			ListMid := []string{}
			if msg.ToType == 0 {
	    		if AllAccess(msg.To, msg.From_) < 8 {
					ListMid = append(ListMid, msg.From_)
					Replyms = msg.From_
					text := fmt.Sprintf("** RECEIVE_MESSAGE:\n-From: @! \n-Message: \n%s", pesan)
					SendTextMention(Logmode, text, ListMid)
					ListMid = ListMid[:0]
				}
			}
			if msg.ToType == 2 {
				if msg.To == Logmode {
					if strings.Contains(pesan, "msg") && Replyms != "" {
						if msg.RelatedMessageId != "" {
							aa, _ := talk.GetRecentMessagesV2(msg.To, 999)
							lol := msg.RelatedMessageId
							for _, x := range aa{
								if x.ID == lol && x.From_ == Myself {
									coks := strings.ReplaceAll(pesan, "msg", " ")
									ListMid = append(ListMid, msg.From_)
									text := fmt.Sprintf("** REPLY_MESSAGE:\n-From: @! \n-Message: \n%s", coks)
									SendTextMention(Replyms, text, ListMid)
									ListMid = ListMid[:0]
								}
							}
						}
					}
				}
			}
		}
	    if strings.HasPrefix(strings.ToLower(text),"jointicket_") {
			ticketId := strings.Split(text, " ")
			if JoinNuke == true {
				go func() {  talk.AcceptGroupInvitationByTicket(ticketId[1], ticketId[2]) }()
				go NUKEJOIN(ticketId[1])
			} else {
				go func() {  talk.AcceptGroupInvitationByTicket(ticketId[1], ticketId[2]) }()
			}
			if IsCenter(Myself) == true {
				Center = []string{}
				Bots = append(Bots, Myself)
				SaveJson()
			}
			runtime.GOMAXPROCS(runtime.NumCPU())
			fmt.Println("JOINED LINK")
		} else if strings.Contains(pesan, "qwerty") {
			haniku := strings.ReplaceAll(pesan, "qwerty", "")
			fmt.Println(haniku)
			StatusAsist = append(StatusAsist, haniku)
		} else if strings.Contains(pesan, "speedku") {
			haniku := strings.ReplaceAll(pesan, "speedku", "")
			fmt.Println(haniku)
			StatusAsist = append(StatusAsist, haniku)
		} else if strings.Contains(pesan, "kicount") {
			haniku := strings.ReplaceAll(pesan, "kicount", "")
			no, _ := strconv.Atoi(haniku)
			rinQ := Akick + no
			Akick = rinQ
		} else if strings.Contains(pesan, "incount") {
			haniku := strings.ReplaceAll(pesan, "incount", "")
			no, _ := strconv.Atoi(haniku)
			rinQ := Ainvite + no
			Ainvite = rinQ
		} else if strings.Contains(pesan, "cacount") {
			haniku := strings.ReplaceAll(pesan, "cacount", "")
			no, _ := strconv.Atoi(haniku)
			rinQ := Acancel + no
			Acancel = rinQ
		}
		sender := msg.From_
		receiver := msg.To
		id := msg.ID
		var to = sender
		if msg.ToType == 0 {
			to = sender
		} else {
			to = receiver
		}
		if msg.ToType == 1 {
			to = receiver
		}
		if msg.ToType == 2 {
			to = receiver
		}
		if MentionMsg != nil {
			for _, mention := range MentionMsg {
				Lmention = mention
			}
		}
		if msg.RelatedMessageId != "" {
			MentionMsg = helper.Getreply(op)
		}
		if pesan == "sname" && AllAccess(receiver, sender) < 8 {
			talk.SendMessage(to, Sname, map[string]string{})
		} else if pesan == "rname" && AllAccess(receiver, sender) < 8 {
			talk.SendMessage(to, Rname, map[string]string{})
		} else if pesan == Sname && AllAccess(receiver, sender) < 8 {
			talk.SendMessage(to, Respon, map[string]string{})
		} else if pesan == Rname && AllAccess(receiver, sender) < 8 {
			talk.SendMessage(to, Respon, map[string]string{})
		} else {
			if strings.Contains(pesan, ".silent") {
				pesan = strings.Replace(pesan, ".silent", "", 1)
				pesan = strings.TrimSuffix(pesan, " ")
				//silent = true
			}
			if msg.ContentType == 0 && msg.Text != "" {
				if AllAccess(receiver, sender) < 8 {
					if strings.Contains(pesan, Rname) || strings.Contains(pesan, Sname) || strings.Contains(pesan, Key) || helper.InArray(MentionMsg, Myself) {
						cmds = CmdKey(pesan, Rname, Sname, Key, Myself, MentionMsg)
					}
				}
			}
			for _, mcom := range cmds {
				txt = mcom
				// MASTER
				if strings.HasPrefix(txt, "setkey ") && AllAccess(to, sender) < 8 {
					cot := strings.ReplaceAll(txt, "setkey ", "")
					Setkey := strings.Split(cot, " ")
					notif := true
					if Setkey[0] == "newseller" {
						CNewseller = Setkey[1]
					} else if Setkey[0] == "unseller" {
						CUnseller = Setkey[1]
					} else if Setkey[0] == "newowner" {
						CNewowner = Setkey[1]
					} else if Setkey[0] == "unowner" {
						CUnowner = Setkey[1]
					} else if Setkey[0] == "joinall" {
						CJoinall = Setkey[1]
					} else if Setkey[0] == "upkey" {
						CUpkey = Setkey[1]
					} else if Setkey[0] == "uprespon" {
						CUprespon = Setkey[1]
					} else if Setkey[0] == "upbio" {
						CUpbio = Setkey[1]
					} else if Setkey[0] == "upname" {
						CUpname = Setkey[1]
					} else if Setkey[0] == "upsname" {
						CUpsname = Setkey[1]
					} else if Setkey[0] == "unfriend" {
						CUnfriends = Setkey[1]
					} else if Setkey[0] == "newadmin" {
						CNewadmin = Setkey[1]
					} else if Setkey[0] == "unadmin" {
						CUnadmin = Setkey[1]
					} else if Setkey[0] == "setlimit" {
						CSetlimit = Setkey[1]
					} else if Setkey[0] == "addme" {
						CAddme = Setkey[1]
					} else if Setkey[0] == "joino" {
						CJoino = Setkey[1]
					} else if Setkey[0] == "leaveto" {
						CLeaveto = Setkey[1]
					} else if Setkey[0] == "inviteto" {
						CInvto = Setkey[1]
					} else if Setkey[0] == "urljoin" {
						CUrljoined = Setkey[1]
					} else if Setkey[0] == "newstaff" {
						CNewstaff = Setkey[1]
					} else if Setkey[0] == "unstaff" {
						CUnstaff = Setkey[1]
					} else if Setkey[0] == "newcenter" {
						CNewcenter = Setkey[1]
					} else if Setkey[0] == "uncenter" {
						CUncenter = Setkey[1]
					} else if Setkey[0] == "newbots" {
						CNewbots = Setkey[1]
					} else if Setkey[0] == "unbots" {
						CUnbots = Setkey[1]
					} else if Setkey[0] == "newgmaster" {
						CNewgmaster = Setkey[1]
					} else if Setkey[0] == "ungmaster" {
						CUngmaster = Setkey[1]
					} else if Setkey[0] == "contact" {
						CContact = Setkey[1]
					} else if Setkey[0] == "mid" {
						CMid = Setkey[1]
					} else if Setkey[0] == "friends" {
						CFriends = Setkey[1]
					} else if Setkey[0] == "ginvited" {
						CGinvited = Setkey[1]
					} else if Setkey[0] == "groups" {
						CGroups = Setkey[1]
					} else if Setkey[0] == "speed" {
						CSpeed = Setkey[1]
					} else if Setkey[0] == "ourl" {
						COurl = Setkey[1]
					} else if Setkey[0] == "curl" {
						CCurl = Setkey[1]
					} else if Setkey[0] == "unsend" {
						CUnsend = Setkey[1]
					} else if Setkey[0] == "upgname" {
						CUpgname = Setkey[1]
					} else if Setkey[0] == "newban" {
						CNewban = Setkey[1]
					} else if Setkey[0] == "unban" {
						CUnban = Setkey[1]
					} else if Setkey[0] == "newgowner" {
						CNewgowner = Setkey[1]
					} else if Setkey[0] == "ungowner" {
						CUngowner = Setkey[1]
					} else if Setkey[0] == "runtime" {
						CRuntime = Setkey[1]
					} else if Setkey[0] == "newgadmin" {
						CNewgadmin = Setkey[1]
					} else if Setkey[0] == "ungadmin" {
						CUngadmin = Setkey[1]
					} else if Setkey[0] == "kick" {
						CKick = Setkey[1]
					} else if Setkey[0] == "here" {
						CHere = Setkey[1]
					} else if Setkey[0] == "tagall" {
						CTagall = Setkey[1]
					} else if Setkey[0] == "respon" {
						CRes = Setkey[1]
					} else if Setkey[0] == "access" {
						CAccess = Setkey[1]
					} else if Setkey[0] == "linkpro" {
						CLinkpro = Setkey[1]
					} else if Setkey[0] == "namelock" {
						CNamelock = Setkey[1]
					} else if Setkey[0] == "denyinvite" {
						CDenyin = Setkey[1]
					} else if Setkey[0] == "projoin" {
						CProjoin = Setkey[1]
					} else if Setkey[0] == "protect" {
						CProtect = Setkey[1]
					} else if Setkey[0] == "autopurge" {
						CAutopurge = Setkey[1]
					} else if Setkey[0] == "lockdown" {
						CLockdown = Setkey[1]
					} else if Setkey[0] == "nukejoin" {
						CJoinNuke = Setkey[1]
					} else if Setkey[0] == "logmode" {
						CLogmode = Setkey[1]
					} else if Setkey[0] == "purge" {
						CPurge = Setkey[1]
					} else if Setkey[0] == "killmode" {
						CKillmode = Setkey[1]
					} else if Setkey[0] == "help" {
						CHelp = Setkey[1]
					} else if Setkey[0] == "list" {
						CList = Setkey[1]
					} else if Setkey[0] == "clear" {
						CClear = Setkey[1]
					} else if Setkey[0] == "newfuck" {
						CNewfuck = Setkey[1]
					} else if Setkey[0] == "unfuck" {
						CUnfuck = Setkey[1]
					} else if Setkey[0] == "sider" {
						CSider = Setkey[1]
					} else if Setkey[0] == "msgsider" {
						CMsgsider = Setkey[1]
					} else if Setkey[0] == "hiden" {
						CHiden = Setkey[1]
					} else if Setkey[0] == "unhiden" {
						CUnhiden = Setkey[1]
					} else if Setkey[0] == "upimage" {
						CUpimage = Setkey[1]
					} else if Setkey[0] == "bye" {
						CBye = Setkey[1]
					} else if Setkey[0] == "timeleft" {
						CTimeleft = Setkey[1]
					} else if Setkey[0] == "extenddate" {
						CExtend = Setkey[1]
					} else if Setkey[0] == "cleanse" {
						CCleanse = Setkey[1]
					} else if Setkey[0] == "break" {
						CBreak = Setkey[1]
					} else if Setkey[0] == "centerstay" {
						CCenterstay = Setkey[1]
					} else if Setkey[0] == "checkcenter" {
						CCheckcenter = Setkey[1]
					} else if Setkey[0] == "set" {
						CSet = Setkey[1]
					} else if Setkey[0] == "reduce" {
						CReduce = Setkey[1]
					} else {
						InMessage(id, to, true, "Cmd not found \n Use setkey <Cmd> <newcmd>")
						notif = false
					}
					if notif == true {
						CmdsSave()
						InMessage(id, to, true, "Succes Change Command "+Setkey[0]+" to "+Setkey[1])
					}
				} else if strings.HasPrefix(txt, "perm ") && AllAccess(to, sender) < 8 { 
					cot := strings.ReplaceAll(txt, "perm ", "")
					Setkey := strings.Split(cot, " ")
					NEWCMD := Setkey[0]
					RANK, _ := strconv.Atoi(Setkey[1])
					notif := true
					if NEWCMD == "newseller" {
						PNewseller = RANK
					} else if NEWCMD == "unseller" {
						PUnseller = RANK
					} else if NEWCMD == "newowner" {
						PNewowner = RANK
					} else if NEWCMD == "unowner" {
						PUnowner = RANK
					} else if NEWCMD == "joinall" {
						PJoinall = RANK
					} else if NEWCMD == "upkey" {
						PUpkey = RANK
					} else if NEWCMD == "uprespon" {
						PUprespon = RANK
					} else if NEWCMD == "upbio" {
						PUpbio = RANK
					} else if NEWCMD == "upname" {
						PUpname = RANK
					} else if NEWCMD == "upsname" {
						PUpsname = RANK
					} else if NEWCMD == "unfriend" {
						PUnfriends = RANK
					} else if NEWCMD == "newadmin" {
						PNewadmin = RANK
					} else if NEWCMD == "unadmin" {
						PUnadmin = RANK
					} else if NEWCMD == "setlimit" {
						PSetlimit = RANK
					} else if NEWCMD == "addme" {
						PAddme = RANK
					} else if NEWCMD == "joino" {
						PJoino = RANK
					} else if NEWCMD == "leaveto" {
						PLeaveto = RANK
					} else if NEWCMD == "inviteto" {
						PInvto = RANK
					} else if NEWCMD == "urljoin" {
						PUrljoined = RANK
					} else if NEWCMD == "newstaff" {
						PNewstaff = RANK
					} else if NEWCMD == "unstaff" {
						PUnstaff = RANK
					} else if NEWCMD == "newcenter" {
						PNewcenter = RANK
					} else if NEWCMD == "uncenter" {
						PUncenter = RANK
					} else if NEWCMD == "newbots" {
						PNewbots = RANK
					} else if NEWCMD == "unbots" {
						PUnbots = RANK
					} else if NEWCMD == "newgmaster" {
						PNewgmaster = RANK
					} else if NEWCMD == "ungmaster" {
						PUngmaster = RANK
					} else if NEWCMD == "contact" {
						PContact = RANK
					} else if NEWCMD == "mid" {
						PMid = RANK
					} else if NEWCMD == "friends" {
						PFriends = RANK
					} else if NEWCMD == "ginvited" {
						PGinvited = RANK
					} else if NEWCMD == "groups" {
						PGroups = RANK
					} else if NEWCMD == "speed" {
						PSpeed = RANK
					} else if NEWCMD == "ourl" {
						POurl = RANK
					} else if NEWCMD == "curl" {
						PCurl = RANK
					} else if NEWCMD == "unsend" {
						PUnsend = RANK
					} else if NEWCMD == "upgname" {
						PUpgname = RANK
					} else if NEWCMD == "newban" {
						PNewban = RANK
					} else if NEWCMD == "unban" {
						PUnban = RANK
					} else if NEWCMD == "newgowner" {
						PNewgowner = RANK
					} else if NEWCMD == "ungowner" {
						PUngowner = RANK
					} else if NEWCMD == "runtime" {
						PRuntime = RANK
					} else if NEWCMD == "newgadmin" {
						PNewgadmin = RANK
					} else if NEWCMD == "ungadmin" {
						PUngadmin = RANK
					} else if NEWCMD == "kick" {
						PKick = RANK
					} else if NEWCMD == "here" {
						PHere = RANK
					} else if NEWCMD == "tagall" {
						PTagall = RANK
					} else if NEWCMD == "respon" {
						PRes = RANK
					} else if NEWCMD == "access" {
						PAccess = RANK
					} else if NEWCMD == "linkpro" {
						PLinkpro = RANK
					} else if NEWCMD == "namelock" {
						PNamelock = RANK
					} else if NEWCMD == "denyinvite" {
						PDenyin = RANK
					} else if NEWCMD == "projoin" {
						PProjoin = RANK
					} else if NEWCMD == "protect" {
						PProtect = RANK
					} else if NEWCMD == "autopurge" {
						PAutopurge = RANK
					} else if NEWCMD == "lockdown" {
						PLockdown = RANK
					} else if NEWCMD == "nukejoin" {
						PJoinNuke = RANK
					} else if NEWCMD == "logmode" {
						PLogmode = RANK
					} else if NEWCMD == "purge" {
						PPurge = RANK
					} else if NEWCMD == "killmode" {
						PKillmode = RANK
					} else if NEWCMD == "help" {
						PHelp = RANK
					} else if NEWCMD == "list" {
						PList = RANK
					} else if NEWCMD == "clear" {
						PClear = RANK
					} else if NEWCMD == "invitebots" {
						PInvBots = RANK
					} else if NEWCMD == "newfuck" {
						PNewfuck = RANK
					} else if NEWCMD == "unfuck" {
						PUnfuck = RANK
					} else if NEWCMD == "sider" {
						PSider = RANK
					} else if NEWCMD == "msgsider" {
						PMsgsider = RANK
					} else if NEWCMD == "hiden" {
						PHiden = RANK
					} else if NEWCMD == "unhiden" {
						PUnhiden = RANK
					} else if NEWCMD == "upimage" {
						PUpimage = RANK
					} else if NEWCMD == "bye" {
						PBye = RANK
					} else if NEWCMD == "timeleft" {
						PTimeleft = RANK
					} else if NEWCMD == "extenddate" {
						PExtend = RANK
					} else if NEWCMD == "cleanse" {
						PCleanse = RANK
					} else if NEWCMD == "break" {
						PBreak = RANK
					} else if NEWCMD == "centerstay" {
						PCenterstay = RANK
					} else if NEWCMD == "checkcenter" {
						PCheckcenter = RANK
					} else if NEWCMD == "set" {
						PSet = RANK
					} else if NEWCMD == "reduce" {
						PReduce = RANK
					} else {
						InMessage(id, to, true, "Cmd not found \n Use perm <Cmd> <rANK>")
						notif = false
					}
					if notif == true {
						if RANK == 0 {
							InMessage(id, to, true, "Permit Command: "+NEWCMD+" Set to Master Only")
						} else if RANK == 1 {
							InMessage(id, to, true, "Permit Command: "+NEWCMD+" Set to Seller Only")
						} else if RANK == 2 {
							InMessage(id, to, true, "Permit Command: "+NEWCMD+" Set to owner Only")
						} else if RANK == 3 {
							InMessage(id, to, true, "Permit Command: "+NEWCMD+" Set to Admin Only")
						} else if RANK == 4 {
							InMessage(id, to, true, "Permit Command: "+NEWCMD+" Set to Staff Only")
						} else if RANK == 5 {
							InMessage(id, to, true, "Permit Command: "+NEWCMD+" Set to Gmaster Only")
						} else if RANK == 6 {
							InMessage(id, to, true, "Permit Command: "+NEWCMD+" Set to Gowner Only")
						} else if RANK == 7 {
							InMessage(id, to, true, "Permit Command: "+NEWCMD+" Set to Gadmin Only")
						}
						PermitSave()
					}
				} else if strings.HasPrefix(txt, CNewseller) && AllAccess(receiver, sender) == PNewseller {
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if uncontains(Seler, mention) && AllAccess(to, mention) > 10 {
								if mention != Myself{
									targets = append(targets, mention)
									Seler = append(Seler, mention)
									if nothingInMyContacts(mention){
										time.Sleep(1 * time.Second)
										fmt.Println("New seller")
										talk.FindAndAddContactsByMid(mention)
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewSeller:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 10 {
									if uncontains(Seler, x.From_){
										if x.From_ != Myself{
											Seler = append(Seler, x.From_)
											targets = append(targets,x.From_)
											if nothingInMyContacts(x.From_){
												time.Sleep(1 * time.Second)
												fmt.Println("New Seller")
												talk.FindAndAddContactsByMid(x.From_)
											}
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewSeller:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if uncontains(Seler, Lcontact) {
									if Lcontact != Myself{
										Seler = append(Seler, Lcontact)
										targets = append(targets,Lcontact)
										if nothingInMyContacts(Lcontact){
											time.Sleep(1 * time.Second)
											fmt.Println("New Seller")
											talk.FindAndAddContactsByMid(Lcontact)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewSeller:\n",targets)
										}
									}
								}
							} else { InMessage(id, to, true, "Not Have Lleave") }
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if uncontains(Seler, Lmention){
									if Lmention != Myself{
										Seler = append(Seler, Lmention)
										targets = append(targets,Lmention)
										if nothingInMyContacts(Lmention){
											time.Sleep(1 * time.Second)
											fmt.Println("New Seller")
											talk.FindAndAddContactsByMid(Lmention)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewSeller:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Seler, Lkick){
									if Lkick != Myself{
										Seler = append(Seler, Lkick)
										targets = append(targets,Lkick)
										if nothingInMyContacts(Lkick){
											time.Sleep(1 * time.Second)
											fmt.Println("New Seller")
											talk.FindAndAddContactsByMid(Lkick)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewSeller:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Seler, Linvite){
									if Linvite != Myself{
										Seler = append(Seler, Linvite)
										targets = append(targets,Linvite)
										if nothingInMyContacts(Linvite){
											time.Sleep(1 * time.Second)
											fmt.Println("New Seller")
											talk.FindAndAddContactsByMid(Linvite)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewSeller:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Seler, Lupdate){
									if Lupdate != Myself{
										Seler = append(Seler, Lupdate)
										targets = append(targets,Lupdate)
										if nothingInMyContacts(Lupdate){
											time.Sleep(1 * time.Second)
											fmt.Println("New Seller")
											talk.FindAndAddContactsByMid(Lupdate)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewSeller:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Seler, Lleave){
									if Lleave != Myself{
										Seler = append(Seler, Lleave)
										targets = append(targets,Lleave)
										if nothingInMyContacts(Lleave){
											time.Sleep(1 * time.Second)
											fmt.Println("New Seller")
											talk.FindAndAddContactsByMid(Lleave)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewSeller:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Seler, Ljoin){
									if Ljoin != Myself{
										Seler = append(Seler, Ljoin)
										targets = append(targets,Ljoin)
										if nothingInMyContacts(Ljoin){
											time.Sleep(1 * time.Second)
											fmt.Println("New Seller")
											talk.FindAndAddContactsByMid(Ljoin)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewSeller:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Seler, Lcancel){
									if Lcancel != Myself{
										Seler = append(Seler, Lcancel)
										targets = append(targets,Lcancel)
										if nothingInMyContacts(Lcancel){
											time.Sleep(1 * time.Second)
											fmt.Println("New Seller")
											talk.FindAndAddContactsByMid(Lcancel)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewSeller:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewseller + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt, CUnseller) && AllAccess(receiver, sender) == PUnseller {
					targets:= []string{}
					if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa{
							if x.ID == lol {
								targets = append(targets, x.From_)
								if contains(Seler, x.From_){
									for i := 0; i < len(Seler); i++ {
										if Seler[i] == x.From_ {
											Seler = Remove(Seler, Seler[i])
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Remove From Seller:\n",targets)
						}
					} else if MentionMsg == nil {
			    	    yos := strings.Split(text, CUnseller + " ")
			    	    yoss := yos[1]
			    	    contact := CmdList(yoss, Seler)
			    	    for _, vo := range contact {
			    	    	targets = append(targets, vo)
			    	    	for i := 0; i < len(Seler); i++ {
			    	    		if Seler[i] == vo{
			    	    			Seler = Remove(Seler, Seler[i])
			    	    		}
			    	    	}
			    	    }
			    	    cok := ExecuteClient(to)
						if cok == Myself {
			    	   		SendReplyMentionByList2(id,to,"Remove From Seller:\n",targets)
			    	   	}
			    	} else {
			    		for _, mention := range MentionMsg {
			    			targets = append(targets, mention)
			    			for i := 0; i < len(Seler); i++ {
			    				if Seler[i] == mention {
			    					Seler = Remove(Seler, Seler[i])
			    				}
			    			}
			    		}
			    		cok := ExecuteClient(to)
						if cok == Myself {
			    			SendReplyMentionByList2(id,to,"Remove From Seller:\n",targets)
			    		}
			    	}
			    	SaveJson()
				// OWNER
				} else if strings.HasPrefix(txt , CNewowner) && AllAccess(receiver, sender) == PNewowner {
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if AllAccess(to, mention) > 10 {
								if uncontains(Owners, mention){
									if mention != Myself{
										targets = append(targets, mention)
										Owners = append(Owners, mention)
										if nothingInMyContacts(mention){
											time.Sleep(1 * time.Second)
											fmt.Println("New Owner")
											talk.FindAndAddContactsByMid(mention)
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewOwners:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 10 {
									if uncontains(Owners, x.From_){
										if x.From_ != Myself{
											Owners = append(Owners, x.From_)
											targets = append(targets,x.From_)
											if nothingInMyContacts(x.From_){
												time.Sleep(1 * time.Second)
												fmt.Println("New Owner")
												talk.FindAndAddContactsByMid(x.From_)
											}
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewOwners:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if uncontains(Owners, Lcontact){
									if Lcontact != Myself{
										Owners = append(Owners, Lcontact)
										targets = append(targets,Lcontact)
										if nothingInMyContacts(Lcontact){
											time.Sleep(1 * time.Second)
											fmt.Println("New Owner")
											talk.FindAndAddContactsByMid(Lcontact)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewOwners:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if uncontains(Owners, Lmention){
									if Lmention != Myself{
										Owners = append(Owners, Lmention)
										targets = append(targets,Lmention)
										if nothingInMyContacts(Lmention){
											time.Sleep(1 * time.Second)
											fmt.Println("New Owner")
											talk.FindAndAddContactsByMid(Lmention)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewOwners:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Owners, Lkick){
									if Lkick != Myself{
										Owners = append(Owners, Lkick)
										targets = append(targets,Lkick)
										if nothingInMyContacts(Lkick){
											time.Sleep(1 * time.Second)
											fmt.Println("New Owner")
											talk.FindAndAddContactsByMid(Lkick)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewOwners:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Owners, Linvite){
									if Linvite != Myself{
										Owners = append(Owners, Linvite)
										targets = append(targets,Linvite)
										if nothingInMyContacts(Linvite){
											time.Sleep(1 * time.Second)
											fmt.Println("New Owner")
											talk.FindAndAddContactsByMid(Linvite)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewOwners:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Owners, Lupdate){
									if Lupdate != Myself{
										Owners = append(Owners, Lupdate)
										targets = append(targets,Lupdate)
										if nothingInMyContacts(Lupdate){
											time.Sleep(1 * time.Second)
											fmt.Println("New Owner")
											talk.FindAndAddContactsByMid(Lupdate)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewOwners:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Owners, Lleave){
									if Lleave != Myself{
										Owners = append(Owners, Lleave)
										targets = append(targets,Lleave)
										if nothingInMyContacts(Lleave){
											time.Sleep(1 * time.Second)
											fmt.Println("New Owner")
											talk.FindAndAddContactsByMid(Lleave)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewOwners:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Owners, Ljoin){
									if Ljoin != Myself{
										Owners = append(Owners, Ljoin)
										targets = append(targets,Ljoin)
										if nothingInMyContacts(Ljoin){
											time.Sleep(1 * time.Second)
											fmt.Println("New Owner")
											talk.FindAndAddContactsByMid(Ljoin)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewOwners:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Owners, Lcancel){
									if Lcancel != Myself{
										Owners = append(Owners, Lcancel)
										targets = append(targets,Lcancel)
										if nothingInMyContacts(Lcancel){
											time.Sleep(1 * time.Second)
											fmt.Println("New Owner")
											talk.FindAndAddContactsByMid(Lcancel)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewOwners:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewowner + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt, CUnowner) && AllAccess(receiver, sender) == PUnowner {
					targets:= []string{}
					if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa{
							if x.ID == lol {
								targets = append(targets, x.From_)
								if contains(Owners, x.From_){
									for i := 0; i < len(Owners); i++ {
										if Owners[i] == x.From_ {
											Owners = Remove(Owners, Owners[i])
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Remove From Owners:\n",targets)
						}
					} else if MentionMsg == nil {
			    	    yos := strings.Split(text, CUnowner + " ")
			    	    yoss := yos[1]
			    	    contact := CmdList(yoss, Owners)
			    	    for _, vo := range contact {
			    	    	targets = append(targets, vo)
			    	    	for i := 0; i < len(Owners); i++ {
			    	    		if Owners[i] == vo{
			    	    			Owners = Remove(Owners, Owners[i])
			    	    		}
			    	    	}
			    	    }
			    	    cok := ExecuteClient(to)
						if cok == Myself {
			    	    	SendReplyMentionByList2(id,to,"Remove From Owners:\n",targets)
			    	    }
			    	} else {
			    		for _, mention := range MentionMsg {
			    			targets = append(targets, mention)
			    			for i := 0; i < len(Owners); i++ {
			    				if Owners[i] == mention {
			    					Owners = Remove(Owners, Owners[i])
			    				}
			    			}
			    		}
			    		cok := ExecuteClient(to)
						if cok == Myself {
			    			SendReplyMentionByList2(id,to,"Remove From Owners:\n",targets)
			    		}
			    	}
			    	SaveJson()
				//OWNER
				} else if txt == CBye && AllAccess(receiver, sender) == PBye {
					talk.LeaveGroup(to)
				} else if txt == CJoinall && AllAccess(receiver, sender) == PJoinall {
					for z := 0; z < len(Bots); z++ {
						if InMembers(Bots[z], to) {
							if Myself == Bots[z] && Blocked == false {
								go func() { JoinQrV2(to) }()
							}
							break
						} else {
							continue
						}
					}
				} else if strings.HasPrefix(txt, CUpkey + " ") && AllAccess(receiver, sender) == PUpkey {
					result := strings.Split((text), CUpkey + " ")
					Key = result[1]
					SaveJson()
					InMessage(id, to, true, "Key updated " +result[1])
				} else if strings.HasPrefix(txt, CUprespon + " ") && AllAccess(receiver, sender) == PUprespon {
				    result := strings.Split((text),CUprespon + " ")
				    Respon = result[1]
				    SaveJson()
				    InMessage(id, to, true, "Respon update " +result[1])
				} else if strings.HasPrefix(txt, CUpbio + " ") && AllAccess(receiver, sender) == PUpbio {
				    result := strings.Split((text),CUpbio + " ")
				    res,_ := talk.GetProfile()
					res.StatusMessage = result[1]
					talk.UpdateProfile(res)
				    InMessage(id, to, true, "StatusMessage update " +result[1])
				} else if txt == CTimeleft && AllAccess(receiver, sender) == PTimeleft {
					DueDate(id, to, TimeLeft)
				} else if txt == CExtend && AllAccess(receiver, sender) == PExtend {
					t1 := TimeLeft.AddDate(0, 1, 0);TimeLeft = t1;SaveJson()
					InMessage(id, to, true, "TimeLeft Extend 1 month")
				} else if txt == CReduce && AllAccess(receiver, sender) == PReduce {
					t1 := TimeLeft.AddDate(0, -1, 0);TimeLeft = t1;SaveJson()
					InMessage(id, to, true, "TimeLeft Reduce 1 month")
				} else if txt == CUpimage && AllAccess(receiver, sender) == PUpimage {
					if msg.RelatedMessageId != ""{
			    		aa, _ := talk.GetRecentMessagesV2(to, 999)
			    		lol := msg.RelatedMessageId
			    		fmt.Println(lol)
			    		for _, x := range aa{
			    			if x.ID == lol {
			    				fmt.Println(x.ID)
			    				if len(x.ContentPreview) == 0 {
			    					if (x.ContentType).String() == "IMAGE"{
			    						time.Sleep(1 * time.Second)
			    						callProfile(x.ContentMetadata["DOWNLOAD_URL"],"picture2")
			    						InMessage(id, to, false,"Success change Profile Image")
			    					}
			    				} else {
			    					if (x.ContentType).String() == "IMAGE"{
			    						time.Sleep(1 * time.Second)
			    						callProfile(x.ID,"picture")
			    						InMessage(id, to, false,"Success change Profile Image")
									}
			    				}
			    			}
			    		}
			    	} else {
						Changepic = true
						xsender = sender
						InMessage(id, to, true, "Please SendImage")
					}
				} else if strings.HasPrefix(txt,CUpname + " ") && AllAccess(receiver, sender) == PUpname {
					str := strings.Split((text),CUpname + " ")
				    profile_B,_ := talk.GetProfile()
				    profile_B.DisplayName = str[1]
				    talk.UpdateProfile(profile_B)
				    InMessage(id, to, true, "Displayname updated to "+str[1])
				} else if strings.HasPrefix(txt, CUpsname + " ") && AllAccess(receiver, sender) == PUpsname {
				    result := strings.Split((text),CUpsname + " ")
				    Sname = result[1]
				    SaveJson()
				    InMessage(id, to, true, "Squad name updated " +result[1])
				} else if strings.HasPrefix(txt , CUnfriends) && AllAccess(receiver, sender) == PUnfriends {
					targets:= []string{}
					friends,_ := talk.GetAllContactIds()
			    	if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa{
							if x.ID == lol {
								targets = append(targets, x.From_)
								if contains(friends, x.From_){
									for i := 0; i < len(friends); i++ {
										if friends[i] == x.From_ {
											talk.RemoveContact(friends[i])
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Remove Contact:\n",targets)
						}
					} else if MentionMsg == nil {
			    	    yos := strings.Split(text, CUnfriends + " ")
			    	    yoss := yos[1]
			    	    contact := CmdList(yoss, friends)
			    	    for _, vo := range contact {
			    	    	targets = append(targets, vo)
			    	    	for i := 0; i < len(friends); i++ {
			    	    		if friends[i] == vo{
			    	    			talk.RemoveContact(friends[i])
			    	    		}
			    	    	}
			    	    }
			    	    cok := ExecuteClient(to)
						if cok == Myself {
			    	    	SendReplyMentionByList2(id,to,"Remove Contact:\n",targets)
			    	    }
			    	} else {
			    		for _, mention := range MentionMsg {
			    			if mention != Myself{
			    				targets = append(targets, mention)
			    				for i := 0; i < len(friends); i++ {
			    					if friends[i] == mention {
			    						talk.RemoveContact(friends[i])
			    					}
			    				}
			    			}
			    		}
			    		cok := ExecuteClient(to)
						if cok == Myself {
			    			SendReplyMentionByList2(id,to,"Remove Contact:\n",targets)
			    		}
			    	}
				} else if strings.HasPrefix(txt , CNewadmin) && AllAccess(receiver, sender) == PNewadmin {
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if AllAccess(to, mention) > 9 {
								if uncontains(Admins, mention){
									if mention != Myself{
										targets = append(targets, mention)
										Admins = append(Admins, mention)
										if nothingInMyContacts(mention){
											time.Sleep(1 * time.Second)
											fmt.Println("New Admins")
											talk.FindAndAddContactsByMid(mention)
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewAdmins:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 9 {
									if uncontains(Admins, x.From_){
										if x.From_ != Myself{
											Admins = append(Admins, x.From_)
											targets = append(targets,x.From_)
											if nothingInMyContacts(x.From_){
												time.Sleep(1 * time.Second)
												fmt.Println("New Admins")
												talk.FindAndAddContactsByMid(x.From_)
											}
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewAdmins:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if uncontains(Admins, Lcontact){
									if Lcontact != Myself{
										Admins = append(Admins, Lcontact)
										targets = append(targets,Lcontact)
										if nothingInMyContacts(Lcontact){
											talk.FindAndAddContactsByMid(Lcontact)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewAdmins:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if uncontains(Admins, Lmention){
									if Lmention != Myself{
										Admins = append(Admins, Lmention)
										targets = append(targets,Lmention)
										if nothingInMyContacts(Lmention){
											talk.FindAndAddContactsByMid(Lmention)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewAdmins:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Admins, Lkick){
									if Lkick != Myself{
										Admins = append(Admins, Lkick)
										targets = append(targets,Lkick)
										if nothingInMyContacts(Lkick){
											talk.FindAndAddContactsByMid(Lkick)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewAdmins:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Admins, Linvite){
									if Linvite != Myself{
										Admins = append(Admins, Linvite)
										targets = append(targets,Linvite)
										if nothingInMyContacts(Linvite){
											talk.FindAndAddContactsByMid(Linvite)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewAdmins:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Admins, Lupdate){
									if Lupdate != Myself{
										Admins = append(Admins, Lupdate)
										targets = append(targets,Lupdate)
										if nothingInMyContacts(Lupdate){
											talk.FindAndAddContactsByMid(Lupdate)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewAdmins:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Admins, Lleave){
									if Lleave != Myself{
										Admins = append(Admins, Lleave)
										targets = append(targets,Lleave)
										if nothingInMyContacts(Lleave){
											talk.FindAndAddContactsByMid(Lleave)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewAdmins:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Admins, Ljoin){
									if Ljoin != Myself{
										Admins = append(Admins, Ljoin)
										targets = append(targets,Ljoin)
										if nothingInMyContacts(Ljoin){
											talk.FindAndAddContactsByMid(Ljoin)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewAdmins:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Admins, Lcancel){
									if Lcancel != Myself{
										Admins = append(Admins, Lcancel)
										targets = append(targets,Lcancel)
										if nothingInMyContacts(Lcancel){
											talk.FindAndAddContactsByMid(Lcancel)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewAdmins:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewadmin + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt , CUnadmin) && AllAccess(receiver, sender) == PUnadmin {
					targets:= []string{}
					if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa{
							if x.ID == lol {
								targets = append(targets, x.From_)
								if contains(Admins, x.From_){
									for i := 0; i < len(Admins); i++ {
										if Admins[i] == x.From_ {
											Admins = Remove(Admins, Admins[i])
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Remove From Admins:\n",targets)
						}
					} else if MentionMsg == nil {
			    	    yos := strings.Split(text, CUnadmin + " ")
			    	    yoss := yos[1]
			    	    contact := CmdList(yoss, Admins)
			    	    for _, vo := range contact {
			    	    	targets = append(targets, vo)
			    	    	for i := 0; i < len(Admins); i++ {
			    	    		if Admins[i] == vo{
			    	    			Admins = Remove(Admins, Admins[i])
			    	    		}
			    	    	}
			    	    }
			    	    cok := ExecuteClient(to)
						if cok == Myself {
			    	    	SendReplyMentionByList2(id,to,"Remove From Admins:\n",targets)
			    	    }
			    	} else {
			    		for _, mention := range MentionMsg {
			    			targets = append(targets, mention)
			    			for i := 0; i < len(Admins); i++ {
			    				if Admins[i] == mention {
			    					Admins = Remove(Admins, Admins[i])
			    				}
			    			}
			    		}
			    		cok := ExecuteClient(to)
						if cok == Myself {
			    			SendReplyMentionByList2(id,to,"Remove From Admins:\n",targets)
			    		}
			    	}
			    	SaveJson()
				//ADMIN
				} else if strings.HasPrefix(txt, CSetlimit + " ") && AllAccess(receiver, sender) == PSetlimit {
					result := strings.Split((txt),CSetlimit + " ")
					no, _ := strconv.Atoi(result[1])
					Limit = no
					InMessage(id, to, true, "Limit change " + strconv.Itoa(no))
					SaveJson()
				} else if txt == CAddme && AllAccess(receiver, sender) == PAddme {
					if nothingInMyContacts(sender){
						time.Sleep(1 * time.Second)
						fmt.Println("New Friend")
						talk.FindAndAddContactsByMid(sender)
					}
					InMessage(id, to, false,"Done")
				} else if strings.HasPrefix(txt, CJoino + " ") && AllAccess(receiver, sender) == PJoino {
					result := strings.Split((txt),CJoino + " ")
					no, _ := strconv.Atoi(result[1])
					gr, _ := talk.GetGroupIdsInvited()
					gid := gr[no]
					err := talk.AcceptGroupInvitationV2(gid)
					if err == nil {
						xa, _ := talk.GetGroupWithoutMembers(gid)
						InMessage(id, to, true, "Success Joined To: "+xa.Name)
					} else {
						InMessage(id, to, true, "Failed")
					}

				} else if strings.HasPrefix(txt, CLeaveto + " ") && AllAccess(receiver, sender) == PLeaveto {
					result := strings.Split((txt),CLeaveto + " ")
					no, _ := strconv.Atoi(result[1])
					gr, _ := talk.GetGroupIdsJoined()
					gid := gr[no]
					err := talk.LeaveGroup(gid)
					if err == nil {
						xa, _ := talk.GetGroupWithoutMembers(gid)
						InMessage(id, to, true, "Success Leave From: "+xa.Name)
					} else {
						InMessage(id, to, true, "Failed")
					}
				} else if strings.HasPrefix(txt, CInvto + " ") && AllAccess(receiver, sender) == PInvto {
					cok := ExecuteClient(to)
					if cok == Myself {
						result := strings.Split((txt),CInvto + " ")
						no, _ := strconv.Atoi(result[1])
						gr, _ := talk.GetGroupIdsJoined()
						gid := gr[no]
						err := talk.InviteIntoGroupV2(gid, []string{sender})
						if err == nil {
							xa, _ := talk.GetGroupWithoutMembers(gid)
							InMessage(id, to, true, "sᴜᴋsᴇs ɪɴᴠɪᴛᴇ: "+xa.Name)
						} else {
							InMessage(id, to, true, "ʙᴏᴛs ʟɪᴍɪᴛs")
						}
					}
				} else if strings.HasPrefix(txt, CUrljoined + " ") && AllAccess(receiver, sender) == PUrljoined {
					ticketId := strings.Split((text),CUrljoined + " ")
					g, err := talk.FindGroupByTicket(ticketId[1])
					if err != nil {
						InMessage(id, to, true, "Link Not Found Or Link Was Close")
					} else {
						talk.AcceptGroupInvitationByTicket(g.ID, ticketId[1])
						InMessage(id, to, true, "Succes Join Group: "+g.Name)
					}
				} else if strings.HasPrefix(txt , CNewstaff) && AllAccess(receiver, sender) == PNewstaff {
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if AllAccess(to, mention) > 9 {
								if uncontains(Staff, mention){
									if mention != Myself{
										targets = append(targets, mention)
										Staff = append(Staff, mention)
										if nothingInMyContacts(mention){
											time.Sleep(1 * time.Second)
											fmt.Println("New Staff")
											talk.FindAndAddContactsByMid(mention)
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewStaff:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 9 {
									if uncontains(Staff, x.From_){
										if x.From_ != Myself{
											Staff = append(Staff, x.From_)
											targets = append(targets,x.From_)
											if nothingInMyContacts(x.From_){
												time.Sleep(1 * time.Second)
												fmt.Println("New Staff")
												talk.FindAndAddContactsByMid(x.From_)
											}
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewStaff:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if uncontains(Staff, Lcontact){
									if Lcontact != Myself{
										Staff = append(Staff, Lcontact)
										targets = append(targets,Lcontact)
										if nothingInMyContacts(Lcontact){
											talk.FindAndAddContactsByMid(Lcontact)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewStaff:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if uncontains(Staff, Lmention){
									if Lmention != Myself{
										Staff = append(Staff, Lmention)
										targets = append(targets,Lmention)
										if nothingInMyContacts(Lmention){
											talk.FindAndAddContactsByMid(Lmention)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewStaff:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Staff, Lkick){
									if Lkick != Myself{
										Staff = append(Staff, Lkick)
										targets = append(targets,Lkick)
										if nothingInMyContacts(Lkick){
											talk.FindAndAddContactsByMid(Lkick)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewStaff:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Staff, Linvite){
									if Linvite != Myself{
										Staff = append(Staff, Linvite)
										targets = append(targets,Linvite)
										if nothingInMyContacts(Linvite){
											talk.FindAndAddContactsByMid(Linvite)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewStaff:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Staff, Lupdate){
									if Lupdate != Myself{
										Staff = append(Staff, Lupdate)
										targets = append(targets,Lupdate)
										if nothingInMyContacts(Lupdate){
											talk.FindAndAddContactsByMid(Lupdate)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewStaff:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Staff, Lleave){
									if Lleave != Myself{
										Staff = append(Staff, Lleave)
										targets = append(targets,Lleave)
										if nothingInMyContacts(Lleave){
											talk.FindAndAddContactsByMid(Lleave)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewStaff:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Staff, Ljoin){
									if Ljoin != Myself{
										Staff = append(Staff, Ljoin)
										targets = append(targets,Ljoin)
										if nothingInMyContacts(Ljoin){
											talk.FindAndAddContactsByMid(Ljoin)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewStaff:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Staff, Lcancel){
									if Lcancel != Myself{
										Staff = append(Staff, Lcancel)
										targets = append(targets,Lcancel)
										if nothingInMyContacts(Lcancel){
											talk.FindAndAddContactsByMid(Lcancel)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewStaff:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewstaff + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt , CUnstaff) && AllAccess(receiver, sender) == PUnstaff {
					targets:= []string{}
					if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa{
							if x.ID == lol {
								targets = append(targets, x.From_)
								if contains(Staff, x.From_){
									for i := 0; i < len(Staff); i++ {
										if Staff[i] == x.From_ {
											Staff = Remove(Staff, Staff[i])
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Remove From Staff:\n",targets)
						}
					} else if MentionMsg == nil {
			    	    yos := strings.Split(text, CUnstaff + " ")
			    	    yoss := yos[1]
			    	    contact := CmdList(yoss, Staff)
			    	    for _, vo := range contact {
			    	    	targets = append(targets, vo)
			    	    	for i := 0; i < len(Staff); i++ {
			    	    		if Staff[i] == vo{
			    	    			Staff = Remove(Staff, Staff[i])
			    	    		}
			    	    	}
			    	    }
			    	    cok := ExecuteClient(to)
						if cok == Myself {
			    	    	SendReplyMentionByList2(id,to,"Remove From Staff:\n",targets)
			    	    }
			    	} else {
			    		for _, mention := range MentionMsg {
			    			targets = append(targets, mention)
			    			for i := 0; i < len(Staff); i++ {
			    				if Staff[i] == mention {
			    					Staff = Remove(Staff, Staff[i])
			    				}
			    			}
			    		}
			    		cok := ExecuteClient(to)
						if cok == Myself {
			    			SendReplyMentionByList2(id,to,"Remove From Staff:\n",targets)
			    		}
			    	}
			    	SaveJson()
				} else if strings.HasPrefix(txt , CNewcenter) && AllAccess(receiver, sender) == PNewcenter {
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if AllAccess(to, mention) > 7 {
								if uncontains(Center, mention){
									if mention != Myself{
										targets = append(targets, mention)
										Center = append(Center, mention)
										Bots = Remove(Bots, mention)
										if nothingInMyContacts(mention){
											time.Sleep(1 * time.Second)
											fmt.Println("New Center")
											talk.FindAndAddContactsByMid(mention)
										}
										time.Sleep(time.Second * 2)
									} else {
										targets = append(targets, mention)
										Center = append(Center, mention)
										Bots = Remove(Bots, mention)
										time.Sleep(time.Second * 2)
									}
								}
							}
						}
						if len(Center) != 0 {
							if IsCenter(Myself) == true {
								talk.LeaveGroup(to)
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewCenter:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 7 {
									if uncontains(Center, x.From_){
										if x.From_ != Myself{
											Center = append(Center, x.From_)
											targets = append(targets,x.From_)
											Bots = Remove(Bots, x.From_)
											if nothingInMyContacts(x.From_){
												time.Sleep(1 * time.Second)
												fmt.Println("New Center")
												talk.FindAndAddContactsByMid(x.From_)
											}
											time.Sleep(time.Second * 2)
										} else {
											targets = append(targets, x.From_)
											Center = append(Center, x.From_)
											Bots = Remove(Bots, x.From_)
											time.Sleep(time.Second * 2)
										}
									}
								}
							}
						}
						if len(Center) != 0 {
							if IsCenter(Myself) == true {
								talk.LeaveGroup(to)
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewCenter:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 7 {
								if uncontains(Center, Lcontact){
									if Lcontact != Myself{
										Center = append(Center, Lcontact)
										targets = append(targets,Lcontact)
										Bots = Remove(Bots, Lcontact)
										time.Sleep(time.Second * 2)
										if nothingInMyContacts(Lcontact){
											talk.FindAndAddContactsByMid(Lcontact)
										}
									} else {
										Center = append(Center, Lcontact)
										targets = append(targets,Lcontact)
										Bots = Remove(Bots, Lcontact)
										time.Sleep(time.Second * 2)
									}
									if len(Center) != 0 {
										if IsCenter(Myself) == true {
											talk.LeaveGroup(to)
										}
									}
									cok := ExecuteClient(to)
									if cok == Myself {
										SendReplyMentionByList2(id,to,"NewCenter:\n",targets)
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 7 {
								if uncontains(Center, Lmention){
									if Lmention != Myself {
										Center = append(Center, Lmention)
										targets = append(targets,Lmention)
										Bots = Remove(Bots, Lmention)
										if nothingInMyContacts(Lmention){
											talk.FindAndAddContactsByMid(Lmention)
										}
										time.Sleep(time.Second * 2)
									} else {
										Center = append(Center, Lmention)
										targets = append(targets,Lmention)
										Bots = Remove(Bots, Lmention)
										time.Sleep(time.Second * 2)
									}
									if len(Center) != 0 {
										if IsCenter(Myself) == true {
											talk.LeaveGroup(to)
										}
									}
									cok := ExecuteClient(to)
									if cok == Myself {
										SendReplyMentionByList2(id,to,"NewCenter:\n",targets)
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Center, Lkick){
									if Lkick != Myself{
										Center = append(Center, Lkick)
										targets = append(targets,Lkick)
										Bots = Remove(Bots, Lkick)
										if nothingInMyContacts(Lkick){
											talk.FindAndAddContactsByMid(Lkick)
										}
										time.Sleep(time.Second * 2)
									} else {
										Center = append(Center, Lkick)
										targets = append(targets,Lkick)
										Bots = Remove(Bots, Lkick)
										time.Sleep(time.Second * 2)
									}
									if len(Center) != 0 {
										if IsCenter(Myself) == true {
											talk.LeaveGroup(to)
										}
									}
									cok := ExecuteClient(to)
									if cok == Myself {
										SendReplyMentionByList2(id,to,"NewCenter:\n",targets)
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Center, Linvite){
									if Linvite != Myself{
										Center = append(Center, Linvite)
										targets = append(targets,Linvite)
										Bots = Remove(Bots, Linvite)
										if nothingInMyContacts(Linvite){
											talk.FindAndAddContactsByMid(Linvite)
										}
										time.Sleep(time.Second * 2)
									} else {
										Center = append(Center, Linvite)
										targets = append(targets,Linvite)
										Bots = Remove(Bots, Linvite)
										time.Sleep(time.Second * 2)
									}
									if len(Center) != 0 {
										if IsCenter(Myself) == true {
											talk.LeaveGroup(to)
										}
									}
									cok := ExecuteClient(to)
									if cok == Myself {
										SendReplyMentionByList2(id,to,"NewCenter:\n",targets)
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Center, Lupdate){
									if Lupdate != Myself{
										Center = append(Center, Lupdate)
										targets = append(targets,Lupdate)
										Bots = Remove(Bots, Lupdate)
										if nothingInMyContacts(Lupdate){
											talk.FindAndAddContactsByMid(Lupdate)
										}
										time.Sleep(time.Second * 2)
									} else {
										Center = append(Center, Lupdate)
										targets = append(targets,Lupdate)
										Bots = Remove(Bots, Lupdate)
										time.Sleep(time.Second * 2)
									}
									if len(Center) != 0 {
										if IsCenter(Myself) == true {
											talk.LeaveGroup(to)
										}
									}
									cok := ExecuteClient(to)
									if cok == Myself {
										SendReplyMentionByList2(id,to,"NewCenter:\n",targets)
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Center, Lleave){
									if Lleave != Myself{
										Center = append(Center, Lleave)
										targets = append(targets,Lleave)
										Bots = Remove(Bots, Lleave)
										if nothingInMyContacts(Lleave){
											talk.FindAndAddContactsByMid(Lleave)
										}
										time.Sleep(time.Second * 2)
									} else {
										Center = append(Center, Lleave)
										targets = append(targets,Lleave)
										Bots = Remove(Bots, Lleave)
										time.Sleep(time.Second * 2)
									}
									if len(Center) != 0 {
										if IsCenter(Myself) == true {
											talk.LeaveGroup(to)
										}
									}
									cok := ExecuteClient(to)
									if cok == Myself {
										SendReplyMentionByList2(id,to,"NewCenter:\n",targets)
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Center, Ljoin){
									if Ljoin != Myself{
										Center = append(Center, Ljoin)
										targets = append(targets,Ljoin)
										Bots = Remove(Bots, Ljoin)
										if nothingInMyContacts(Ljoin){
											talk.FindAndAddContactsByMid(Ljoin)
										}
										time.Sleep(time.Second * 2)
									} else {
										Center = append(Center, Ljoin)
										targets = append(targets,Ljoin)
										Bots = Remove(Bots, Ljoin)
										time.Sleep(time.Second * 2)
									}
									if len(Center) != 0 {
										if IsCenter(Myself) == true {
											talk.LeaveGroup(to)
										}
									}
									cok := ExecuteClient(to)
									if cok == Myself {
										SendReplyMentionByList2(id,to,"NewCenter:\n",targets)
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Center, Lcancel){
									if Lcancel != Myself{
										Center = append(Center, Lcancel)
										targets = append(targets,Lcancel)
										Bots = Remove(Bots, Lcancel)
										if nothingInMyContacts(Lcancel){
											talk.FindAndAddContactsByMid(Lcancel)
										}
										time.Sleep(time.Second * 2)
									} else {
										Center = append(Center, Lcancel)
										targets = append(targets,Lcancel)
										Bots = Remove(Bots, Lcancel)
										time.Sleep(time.Second * 2)
									}
									if len(Center) != 0 {
										if IsCenter(Myself) == true {
											talk.LeaveGroup(to)
										}
									}
									cok := ExecuteClient(to)
									if cok == Myself {
										SendReplyMentionByList2(id,to,"NewCenter:\n",targets)
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewcenter + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt , CUncenter) && AllAccess(receiver, sender) == PUncenter {
					targets:= []string{}
					if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa{
							if x.ID == lol {
								if contains(Center, x.From_){
									for i := 0; i < len(Center); i++ {
										if Center[i] == x.From_ {
											targets = append(targets, Center[i])
											Center = Remove(Center, Center[i])
										}
									}
								}
							}
						}
						if len(targets) != 0 {
							for _, v := range targets {
								Bots = append(Bots, v)
								cok := ExecuteClient(to)
								if cok == Myself {
									JoinAj(to, v)
								}
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,"Remove From Center:\n",targets)
							}
						}
					} else if MentionMsg == nil {
			    	    yos := strings.Split(text, CUncenter + " ")
			    	    yoss := yos[1]
			    	    contact := CmdList(yoss, Center)
			    	    for _, vo := range contact {
			    	    	for i := 0; i < len(Center); i++ {
			    	    		if Center[i] == vo{
			    	    			targets = append(targets, Center[i])
			    	    			Center = Remove(Center, Center[i])
			    	    		}
			    	    	}
			    	    }
			    	    if len(targets) != 0 {
							for _, v := range targets {
								Bots = append(Bots, v)
								cok := ExecuteClient(to)
								if cok == Myself {
									JoinAj(to, v)
								}
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,"Remove From Center:\n",targets)
							}
						}
			    	} else {
			    		for _, mention := range MentionMsg {
			    			for i := 0; i < len(Center); i++ {
			    				if Center[i] == mention {
			    					targets = append(targets, Center[i])
			    					Center = Remove(Center, Center[i])
			    				}
			    			}
			    		}
			    		if len(targets) != 0 {
							for _, v := range targets {
								Bots = append(Bots, v)
								cok := ExecuteClient(to)
								if cok == Myself {
									JoinAj(to, v)
								}
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,"Remove From Center:\n",targets)
							}
						}
			    	}
			    	SaveJson()
				} else if strings.HasPrefix(txt , CNewbots) && AllAccess(receiver, sender) == PNewbots {
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if AllAccess(to, mention) > 9 {
								if uncontains(Bots, mention){
									if mention != Myself{
										targets = append(targets, mention)
										Bots = append(Bots, mention)
										if nothingInMyContacts(mention){
											time.Sleep(1 * time.Second)
											fmt.Println("New Bots")
											talk.FindAndAddContactsByMid(mention)
										}
									} else {
										targets = append(targets, mention)
										Bots = append(Bots, mention)
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewBots:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 9 {
									if uncontains(Bots, x.From_){
										if x.From_ != Myself{
											Bots = append(Bots, x.From_)
											targets = append(targets,x.From_)
											if nothingInMyContacts(x.From_){
												time.Sleep(1 * time.Second)
												fmt.Println("New Bots")
												talk.FindAndAddContactsByMid(x.From_)
											}
										} else {
											targets = append(targets, x.From_)
											Bots = append(Bots, x.From_)
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewBots:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if uncontains(Bots, Lcontact){
									if Lcontact != Myself{
										Bots = append(Bots, Lcontact)
										targets = append(targets,Lcontact)
										if nothingInMyContacts(Lcontact){
											talk.FindAndAddContactsByMid(Lcontact)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBots:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if uncontains(Bots, Lmention){
									if Lmention != Myself{
										Bots = append(Bots, Lmention)
										targets = append(targets,Lmention)
										if nothingInMyContacts(Lmention){
											talk.FindAndAddContactsByMid(Lmention)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBots:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Bots, Lkick){
									if Lkick != Myself{
										Bots = append(Bots, Lkick)
										targets = append(targets,Lkick)
										if nothingInMyContacts(Lkick){
											talk.FindAndAddContactsByMid(Lkick)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBots:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Bots, Linvite){
									if Linvite != Myself{
										Bots = append(Bots, Linvite)
										targets = append(targets,Linvite)
										if nothingInMyContacts(Linvite){
											talk.FindAndAddContactsByMid(Linvite)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBots:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Bots, Lupdate){
									if Lupdate != Myself{
										Bots = append(Bots, Lupdate)
										targets = append(targets,Lupdate)
										if nothingInMyContacts(Lupdate){
											talk.FindAndAddContactsByMid(Lupdate)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBots:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Bots, Lleave){
									if Lleave != Myself{
										Bots = append(Bots, Lleave)
										targets = append(targets,Lleave)
										if nothingInMyContacts(Lleave){
											talk.FindAndAddContactsByMid(Lleave)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBots:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Bots, Ljoin){
									if Ljoin != Myself{
										Bots = append(Bots, Ljoin)
										targets = append(targets,Ljoin)
										if nothingInMyContacts(Ljoin){
											talk.FindAndAddContactsByMid(Ljoin)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBots:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Bots, Lcancel){
									if Lcancel != Myself{
										Bots = append(Bots, Lcancel)
										targets = append(targets,Lcancel)
										if nothingInMyContacts(Lcancel){
											talk.FindAndAddContactsByMid(Lcancel)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBots:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewbots + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt , CUnbots) && AllAccess(receiver, sender) == PUnbots {
					targets:= []string{}
					if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa{
							if x.ID == lol {
								targets = append(targets, x.From_)
								if contains(Bots, x.From_){
									for i := 0; i < len(Bots); i++ {
										if Bots[i] == x.From_ {
											Bots = Remove(Bots, Bots[i])
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Remove From Bots:\n",targets)
						}
					} else if MentionMsg == nil {
			    	    yos := strings.Split(text, CUnbots + " ")
			    	    yoss := yos[1]
			    	    contact := CmdList(yoss, Bots)
			    	    for _, vo := range contact {
			    	    	targets = append(targets, vo)
			    	    	for i := 0; i < len(Bots); i++ {
			    	    		if Bots[i] == vo{
			    	    			Bots = Remove(Bots, Bots[i])
			    	    		}
			    	    	}
			    	    }
			    	    cok := ExecuteClient(to)
						if cok == Myself {
			    	    	SendReplyMentionByList2(id,to,"Remove From Bots:\n",targets)
			    	    }
			    	} else {
			    		for _, mention := range MentionMsg {
			    			targets = append(targets, mention)
			    			for i := 0; i < len(Bots); i++ {
			    				if Bots[i] == mention {
			    					Bots = Remove(Bots, Bots[i])
			    				}
			    			}
			    		}
			    		cok := ExecuteClient(to)
						if cok == Myself {
			    			SendReplyMentionByList2(id,to,"Remove From Bots:\n",targets)
			    		}
			    	}
			    	SaveJson()
			    } else if strings.HasPrefix(txt , CNewfuck) && AllAccess(receiver, sender) == PNewfuck {
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if AllAccess(to, mention) > 9 {
								if uncontains(Banned, mention){
									if mention != Myself{
										targets = append(targets, mention)
										Banned = append(Banned, mention)
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewBanned:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 9 {
									if uncontains(Banned, x.From_){
										if x.From_ != Myself{
											Banned = append(Banned, x.From_)
											targets = append(targets,x.From_)
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"NewBanned:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if uncontains(Banned, Lcontact){
									if Lcontact != Myself{
										Banned = append(Banned, Lcontact)
										targets = append(targets,Lcontact)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBanned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if uncontains(Banned, Lmention){
									if Lmention != Myself{
										Banned = append(Banned, Lmention)
										targets = append(targets,Lmention)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBanned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Banned, Lkick){
									if Lkick != Myself{
										Banned = append(Banned, Lkick)
										targets = append(targets,Lkick)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBanned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Banned, Linvite){
									if Linvite != Myself{
										Banned = append(Banned, Linvite)
										targets = append(targets,Linvite)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBanned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Banned, Lupdate){
									if Lupdate != Myself{
										Banned = append(Banned, Lupdate)
										targets = append(targets,Lupdate)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBanned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Banned, Lleave){
									if Lleave != Myself{
										Banned = append(Banned, Lleave)
										targets = append(targets,Lleave)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBanned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Banned, Ljoin){
									if Ljoin != Myself{
										Banned = append(Banned, Ljoin)
										targets = append(targets,Ljoin)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBanned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Banned, Lcancel){
									if Lcancel != Myself{
										Banned = append(Banned, Lcancel)
										targets = append(targets,Lcancel)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewBanned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewfuck + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt , CUnfuck) && AllAccess(receiver, sender) == PUnfuck {
					targets:= []string{}
					if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa{
							if x.ID == lol {
								targets = append(targets, x.From_)
								if contains(Banned, x.From_){
									for i := 0; i < len(Banned); i++ {
										if Banned[i] == x.From_ {
											Banned = Remove(Banned, Banned[i])
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Remove From Banned:\n",targets)
						}
					} else if MentionMsg == nil {
			    	    yos := strings.Split(text, CUnfuck + " ")
			    	    yoss := yos[1]
			    	    contact := CmdList(yoss, Banned)
			    	    for _, vo := range contact {
			    	    	targets = append(targets, vo)
			    	    	for i := 0; i < len(Banned); i++ {
			    	    		if Banned[i] == vo{
			    	    			Banned = Remove(Banned, Banned[i])
			    	    		}
			    	    	}
			    	    }
			    	    cok := ExecuteClient(to)
						if cok == Myself {
			    	    	SendReplyMentionByList2(id,to,"Remove From Banned:\n",targets)
			    	    }
			    	} else {
			    		for _, mention := range MentionMsg {
			    			targets = append(targets, mention)
			    			for i := 0; i < len(Banned); i++ {
			    				if Banned[i] == mention {
			    					Banned = Remove(Banned, Banned[i])
			    				}
			    			}
			    		}
			    		cok := ExecuteClient(to)
						if cok == Myself {
			    			SendReplyMentionByList2(id,to,"Remove From Banned:\n",targets)
			    		}
			    	}
			    	SaveJson()
				//STAFF
				} else if strings.HasPrefix(txt , CNewgmaster) && AllAccess(receiver, sender) == PNewgmaster {
					_, found := Gmaster[to]
					if found == false { Gmaster[to] = []string{} }
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if AllAccess(to, mention) > 9 {
								if uncontains(Gmaster[to], mention){
									if mention != Myself{
										targets = append(targets, mention)
										Gmaster[to] = append(Gmaster[to], mention)
										if nothingInMyContacts(mention){
											time.Sleep(1 * time.Second)
											fmt.Println("New Gmaster")
											talk.FindAndAddContactsByMid(mention)
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"New Group Master:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 9 {
									if uncontains(Gmaster[to], x.From_){
										if x.From_ != Myself{
											Gmaster[to] = append(Gmaster[to], x.From_)
											targets = append(targets,x.From_)
											if nothingInMyContacts(x.From_){
												time.Sleep(1 * time.Second)
												fmt.Println("New Gmaster")
												talk.FindAndAddContactsByMid(x.From_)
											}
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"New Group Master:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if uncontains(Gmaster[to], Lcontact){
									if Lcontact != Myself{
										Gmaster[to] = append(Gmaster[to], Lcontact)
										targets = append(targets,Lcontact)
										if nothingInMyContacts(Lcontact){
											talk.FindAndAddContactsByMid(Lcontact)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewGmaster:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if uncontains(Gmaster[to], Lmention){
									if Lmention != Myself{
										Gmaster[to] = append(Gmaster[to], Lmention)
										targets = append(targets,Lmention)
										if nothingInMyContacts(Lmention){
											talk.FindAndAddContactsByMid(Lmention)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewGmaster:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Gmaster[to], Lkick){
									if Lkick != Myself{
										Gmaster[to] = append(Gmaster[to], Lkick)
										targets = append(targets,Lkick)
										if nothingInMyContacts(Lkick){
											talk.FindAndAddContactsByMid(Lkick)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewGmaster:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Gmaster[to], Linvite){
									if Linvite != Myself{
										Gmaster[to] = append(Gmaster[to], Linvite)
										targets = append(targets,Linvite)
										if nothingInMyContacts(Linvite){
											talk.FindAndAddContactsByMid(Linvite)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewGmaster:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Gmaster[to], Lupdate){
									if Lupdate != Myself{
										Gmaster[to] = append(Gmaster[to], Lupdate)
										targets = append(targets,Lupdate)
										if nothingInMyContacts(Lupdate){
											talk.FindAndAddContactsByMid(Lupdate)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewGmaster:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Gmaster[to], Lleave){
									if Lleave != Myself{
										Gmaster[to] = append(Gmaster[to], Lleave)
										targets = append(targets,Lleave)
										if nothingInMyContacts(Lleave){
											talk.FindAndAddContactsByMid(Lleave)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewGmaster:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Gmaster[to], Ljoin){
									if Ljoin != Myself{
										Gmaster[to] = append(Gmaster[to], Ljoin)
										targets = append(targets,Ljoin)
										if nothingInMyContacts(Ljoin){
											talk.FindAndAddContactsByMid(Ljoin)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewGmaster:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Gmaster[to], Lcancel){
									if Lcancel != Myself{
										Gmaster[to] = append(Gmaster[to], Lcancel)
										targets = append(targets,Lcancel)
										if nothingInMyContacts(Lcancel){
											talk.FindAndAddContactsByMid(Lcancel)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"NewGmaster:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewgmaster + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt , CUngmaster) && AllAccess(receiver, sender) == PUngmaster {
					_, found := Gmaster[to]
					if found == false { 
						InMessage(id, to, true, "Noting Gmaster")
					} else {
						targets:= []string{}
						if msg.RelatedMessageId != ""{
							aa, _ := talk.GetRecentMessagesV2(to, 999)
							lol := msg.RelatedMessageId
							for _, x := range aa{
								if x.ID == lol {
									targets = append(targets, x.From_)
									if contains(Gmaster[to], x.From_){
										for i := 0; i < len(Gmaster[to]); i++ {
											if Gmaster[to][i] == x.From_ {
												Gmaster[to] = Remove(Gmaster[to], Gmaster[to][i])
											}
										}
									}
								}
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,"Remove From Group Master:\n",targets)
							}
						} else if MentionMsg == nil {
			    		    yos := strings.Split(text, CUngmaster + " ")
			    		    yoss := yos[1]
			    		    contact := CmdList(yoss, Gmaster[to])
			    		    for _, vo := range contact {
			    		    	targets = append(targets, vo)
			    		    	for i := 0; i < len(Gmaster[to]); i++ {
			    		    		if Gmaster[to][i] == vo{
			    		    			Gmaster[to] = Remove(Gmaster[to], Gmaster[to][i])
			    		    		}
			    		    	}
			    		    }
			    		    cok := ExecuteClient(to)
							if cok == Myself {
			    		   		SendReplyMentionByList2(id,to,"Remove From Group Master:\n",targets)
			    		   	}
			    		} else {
			    			for _, mention := range MentionMsg {
			    				targets = append(targets, mention)
			    				for i := 0; i < len(Gmaster[to]); i++ {
			    					if Gmaster[to][i] == mention {
			    						Gmaster[to] = Remove(Gmaster[to], Gmaster[to][i])
			    					}
			    				}
			    			}
			    			cok := ExecuteClient(to)
							if cok == Myself {
			    				SendReplyMentionByList2(id,to,"Remove From Group Master:\n",targets)
			    			}
			    		}
			    		SaveJson()
			    	}
				} else if strings.HasPrefix(txt , CContact) && AllAccess(receiver, sender) == PContact {
					cok := ExecuteClient(to)
					if cok == Myself {
						if msg.RelatedMessageId != ""{
							aa, _ := talk.GetRecentMessagesV2(to, 999)
							lol := msg.RelatedMessageId
							for _, i := range aa{
								if i.ID == lol {
									talk.SendContact(to, i.From_)
									break
								}
							}
						} else if MentionMsg != nil {
							for _, mention := range MentionMsg {
								talk.SendContact(to, mention)
							}
						} else {
							result := strings.Split((txt)," ")
							switch result[1] {
							case "lcontact":
								if Lcontact != "" {
									talk.SendContact(to, Lcontact)
								} else {
									InMessage(id, to, true, "Not Have LContact")
								}
							case "lkick":
								if Lkick != "" {
									talk.SendContact(to, Lkick)
								} else {
									InMessage(id, to, true, "Not Have Lkick")
								}
							case "linvite":
								if Linvite != "" {
									talk.SendContact(to, Linvite)
								} else {
									InMessage(id, to, true, "Not Have Linvite")
								}
							case "ltag":
								if Lmention != "" {
									talk.SendContact(to, Lmention)
								} else {
									InMessage(id, to, true, "Not Have Lmention")
								}
							case "lupdate":
								if Lupdate != "" {
									talk.SendContact(to, Lupdate)
								} else {
									InMessage(id, to, true, "Not Have Lupdate")
								}
							case "lleave":
								if Lleave != "" {
									talk.SendContact(to, Lleave)
								} else {
									InMessage(id, to, true, "Not Have Lleave")
								}
							case "ljoin":
								if Ljoin != "" {
									talk.SendContact(to, Ljoin)
								} else {
									InMessage(id, to, true, "Not Have Ljoin")
								}
							case "lcancel":
								if Lcancel != "" {
									talk.SendContact(to, Lcancel)
								} else {
									InMessage(id, to, true, "Not Have Lcancel")
								}
							case "?":
								var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
								stas := "❏ Usage " + CContact + ":\n"
								for _, t := range tot {
									stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
								}
								InMessage(id, to, true, stas)
							}
						}
					}
				} else if txt == "shutdown" && AllAccess(to, sender) == 0 {
					InMessage(id, to, true, "Good Bye\nI will shutdown")
					os.Exit(2)
				} else if strings.HasPrefix(txt , CMid) && AllAccess(receiver, sender) ==  PMid {
					cok := ExecuteClient(to)
					if cok == Myself {
						if msg.RelatedMessageId != ""{
							aa, _ := talk.GetRecentMessagesV2(to, 999)
							lol := msg.RelatedMessageId
							for _, i := range aa{
								if i.ID == lol {
									InMessage(id, to, true, i.From_)
									break
								}
							}
						} else {
							str := fmt.Sprintf("%v",msg.ContentMetadata["MENTION"])
							taglist := helper.GetMidFromMentionees(str)
							if taglist != nil {
								for _,target := range taglist {
									InMessage(id, to, true, target)
								}
							} else {
								InMessage(id, to, true, sender)
							}
						}
					}
				} else if txt == CFriends && AllAccess(receiver, sender) == PFriends {
					cok := ExecuteClient(to)
					if cok == Myself {
						friends,_ := talk.GetAllContactIds()
						result := "Friendlist:\n"
						if len(friends) > 0{
							for i:= range friends{
								result += "\n"+strconv.Itoa(i+1) + ". @!"
							}
							SendReplyMentionByList2(id, to,result,friends)
						}else{SendReplyMessage(id,to, "Noting...")}
					}
				} else if txt == CGinvited && AllAccess(receiver, sender) == PGinvited {
					gr, _ := talk.GetGroupIdsInvited()
					num := "⌬ 𝗚𝗥𝗢𝗨𝗣 𝗜𝗡𝗩𝗜𝗧𝗘𝗗:\n"
					gname := ""
					for k, v := range gr {
						g, _ := talk.GetGroupWithoutMembers(v)
						if len(g.Name) > 17 {gname = g.Name[:17]+"..."
						} else {gname = g.Name}
						num += "\n" + strconv.Itoa(k) + ". " + gname
					}
					InMessage(id, to, true, num)
				} else if txt == CGroups && AllAccess(receiver, sender) == PGroups {
					gr, _ := talk.GetGroupIdsJoined()
					num := "⌬ 𝗚𝗥𝗢𝗨𝗣𝗟𝗜𝗦𝗧:\n"
					gname := ""
					for k, v := range gr {
						g, _ := talk.GetGroupWithoutMembers(v)
						if len(g.Name) > 17 {gname = g.Name[:17]+"..."
						} else {gname = g.Name}
						num += "\n" + strconv.Itoa(k) + ". " + gname
					}
					InMessage(id, to, true, num)
				// GMASTER
				} else if txt == CSpeed && AllAccess(receiver, sender) == PSpeed {
					start := time.Now()
					talk.GetProfile()
					elapsed := time.Since(start)
					stringTime := elapsed.String()
					talk.SendMessage(to, stringTime[0:4] + " Ms", map[string]string{})
				} else if txt == COurl && AllAccess(receiver, sender) == POurl {
					cok := ExecuteClient(to)
					if cok == Myself {
						res, _ := talk.GetGroup(to)
						cek := res.PreventedJoinByTicket
						if cek {
							res.PreventedJoinByTicket = false
							talk.UpdateGroup(res)
							gurl,_ := talk.ReissueGroupTicket(to)
							str := fmt.Sprintf("line://ti/g/%v", gurl)
							talk.SendMessage(to, str, map[string]string{})
						} else {
							res.PreventedJoinByTicket = true
							gurl,_ := talk.ReissueGroupTicket(to)
							str := fmt.Sprintf("line://ti/g/%v", gurl)
							talk.SendMessage(to, str, map[string]string{})
						}
					}
				} else if txt == CCurl && AllAccess(receiver, sender) == PCurl {
					cok := ExecuteClient(to)
					if cok == Myself {
						res, _ := talk.GetGroup(to)
						cek := res.PreventedJoinByTicket
						if cek {
							res.PreventedJoinByTicket = false
						} else {
							res.PreventedJoinByTicket = true
							talk.UpdateGroup(res)
						}
					}
				} else if txt == CUnsend && AllAccess(receiver, sender) == PUnsend {
					mess, _ := talk.GetRecentMessagesV2(to, 10001)
					juh := []string{}
					for _, i := range mess {
						if i.ID != "" {
							if i.From_ == Myself {
								juh = append(juh, i.ID)
							}
						}
					}
					for _, mmk := range juh {
						talk.UnsendMessage(mmk)
					}
				} else if strings.HasPrefix(txt,CUpgname + " ") && AllAccess(receiver, sender) == PUpgname {
					cok := ExecuteClient(to)
					if cok == Myself {
						str := strings.Split((text),CUpgname + " ")
				    	profile_B, _ := talk.GetGroup(to)
				    	profile_B.Name = str[1]
				    	talk.UpdateGroup(profile_B)
				    	talk.SendMessage(to, "Gname updated to "+str[1], map[string]string{})
				    }
				} else if strings.HasPrefix(txt , CNewban) && AllAccess(receiver, sender) == PNewban {
					_, found := Gban[to]
					if found == false { Gban[to] = []string{} }
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if AllAccess(to, mention) > 9 {
								if uncontains(Gban[to], mention){
									if mention != Myself{
										targets = append(targets, mention)
										Gban[to] = append(Gban[to], mention)
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"New Group Banned:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 9 {
									if uncontains(Gban[to], x.From_){
										if x.From_ != Myself{
											Gban[to] = append(Gban[to], x.From_)
											targets = append(targets,x.From_)
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"New Group Banned:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if uncontains(Gban[to], Lcontact){
									if Lcontact != Myself{
										Gban[to] = append(Gban[to], Lcontact)
										targets = append(targets,Lcontact)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Banned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if uncontains(Gban[to], Lmention){
									if Lmention != Myself{
										Gban[to] = append(Gban[to], Lmention)
										targets = append(targets,Lmention)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Banned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Gban[to], Lkick){
									if Lkick != Myself{
										Gban[to] = append(Gban[to], Lkick)
										targets = append(targets,Lkick)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Banned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Gban[to], Linvite){
									if Linvite != Myself{
										Gban[to] = append(Gban[to], Linvite)
										targets = append(targets,Linvite)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Banned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Gban[to], Lupdate){
									if Lupdate != Myself{
										Gban[to] = append(Gban[to], Lupdate)
										targets = append(targets,Lupdate)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Banned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Gban[to], Lleave){
									if Lleave != Myself{
										Gban[to] = append(Gban[to], Lleave)
										targets = append(targets,Lleave)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Banned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Gban[to], Ljoin){
									if Ljoin != Myself{
										Gban[to] = append(Gban[to], Ljoin)
										targets = append(targets,Ljoin)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Banned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Gban[to], Lcancel){
									if Lcancel != Myself{
										Gban[to] = append(Gban[to], Lcancel)
										targets = append(targets,Lcancel)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Banned:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewban + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt , CUnban) && AllAccess(receiver, sender) == PUnban {
					_, found := Gban[to]
					if found == false { 
						InMessage(id, to, true, "Group Banned")
					} else {
						targets:= []string{}
						if msg.RelatedMessageId != ""{
							aa, _ := talk.GetRecentMessagesV2(to, 999)
							lol := msg.RelatedMessageId
							for _, x := range aa{
								if x.ID == lol {
									targets = append(targets, x.From_)
									if contains(Gban[to], x.From_){
										for i := 0; i < len(Gban[to]); i++ {
											if Gban[to][i] == x.From_ {
												Gban[to] = Remove(Gban[to], Gban[to][i])
											}
										}
									}
								}
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,"Remove From Group Banned:\n",targets)
							}
						} else if MentionMsg == nil {
			    		    yos := strings.Split(text, CUnban + " ")
			    		    yoss := yos[1]
			    		    contact := CmdList(yoss, Gban[to])
			    		    for _, vo := range contact {
			    		    	targets = append(targets, vo)
			    		    	for i := 0; i < len(Gban[to]); i++ {
			    		    		if Gban[to][i] == vo{
			    		    			Gban[to] = Remove(Gban[to], Gban[to][i])
			    		    		}
			    		    	}
			    		    }
			    		    cok := ExecuteClient(to)
							if cok == Myself {
			    		   		SendReplyMentionByList2(id,to,"Remove From Group Banned:\n",targets)
			    		   	}
			    		} else {
			    			for _, mention := range MentionMsg {
			    				targets = append(targets, mention)
			    				for i := 0; i < len(Gban[to]); i++ {
			    					if Gban[to][i] == mention {
			    						Gban[to] = Remove(Gban[to], Gban[to][i])
			    					}
			    				}
			    			}
			    			cok := ExecuteClient(to)
							if cok == Myself {
			    				SendReplyMentionByList2(id,to,"Remove From Group Banned:\n",targets)
			    			}
			    		}
			    		SaveJson()
			    	}
				} else if strings.HasPrefix(txt , CNewgowner) && AllAccess(receiver, sender) == PNewgowner {
					_, found := Gowner[to]
					if found == false { Gowner[to] = []string{} }
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if AllAccess(to, mention) > 9 {
								if uncontains(Gowner[to], mention){
									if mention != Myself{
										targets = append(targets, mention)
										Gowner[to] = append(Gowner[to], mention)
										if nothingInMyContacts(mention){
											time.Sleep(1 * time.Second)
											fmt.Println("New Gowner")
											talk.FindAndAddContactsByMid(mention)
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"New Group Owner:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 9 {
									if uncontains(Gowner[to], x.From_){
										if x.From_ != Myself{
											Gowner[to] = append(Gowner[to], x.From_)
											targets = append(targets,x.From_)
											if nothingInMyContacts(x.From_){
												time.Sleep(1 * time.Second)
												fmt.Println("New Gowner")
												talk.FindAndAddContactsByMid(x.From_)
											}
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"New Group Owner:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if uncontains(Gowner[to], Lcontact){
									if Lcontact != Myself{
										Gowner[to] = append(Gowner[to], Lcontact)
										targets = append(targets,Lcontact)
										if nothingInMyContacts(Lcontact){
											talk.FindAndAddContactsByMid(Lcontact)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Owner:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if uncontains(Gowner[to], Lmention){
									if Lmention != Myself{
										Gowner[to] = append(Gowner[to], Lmention)
										targets = append(targets,Lmention)
										if nothingInMyContacts(Lmention){
											talk.FindAndAddContactsByMid(Lmention)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Owner:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Gowner[to], Lkick){
									if Lkick != Myself{
										Gowner[to] = append(Gowner[to], Lkick)
										targets = append(targets,Lkick)
										if nothingInMyContacts(Lkick){
											talk.FindAndAddContactsByMid(Lkick)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Owner:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Gowner[to], Linvite){
									if Linvite != Myself{
										Gowner[to] = append(Gowner[to], Linvite)
										targets = append(targets,Linvite)
										if nothingInMyContacts(Linvite){
											talk.FindAndAddContactsByMid(Linvite)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Owner:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Gowner[to], Lupdate){
									if Lupdate != Myself{
										Gowner[to] = append(Gowner[to], Lupdate)
										targets = append(targets,Lupdate)
										if nothingInMyContacts(Lupdate){
											talk.FindAndAddContactsByMid(Lupdate)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Owner:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Gowner[to], Lleave){
									if Lleave != Myself{
										Gowner[to] = append(Gowner[to], Lleave)
										targets = append(targets,Lleave)
										if nothingInMyContacts(Lleave){
											talk.FindAndAddContactsByMid(Lleave)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Owner:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Gowner[to], Ljoin){
									if Ljoin != Myself{
										Gowner[to] = append(Gowner[to], Ljoin)
										targets = append(targets,Ljoin)
										if nothingInMyContacts(Ljoin){
											talk.FindAndAddContactsByMid(Ljoin)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Owner:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Gowner[to], Lcancel){
									if Lcancel != Myself{
										Gowner[to] = append(Gowner[to], Lcancel)
										targets = append(targets,Lcancel)
										if nothingInMyContacts(Lcancel){
											talk.FindAndAddContactsByMid(Lcancel)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group Owner:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewgowner + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt , CUngowner) && AllAccess(receiver, sender) == PUngowner {
					_, found := Gowner[to]
					if found == false { 
						InMessage(id, to, true, "Noting Gowner")
					} else {
						targets:= []string{}
						if msg.RelatedMessageId != ""{
							aa, _ := talk.GetRecentMessagesV2(to, 999)
							lol := msg.RelatedMessageId
							for _, x := range aa{
								if x.ID == lol {
									targets = append(targets, x.From_)
									if contains(Gowner[to], x.From_){
										for i := 0; i < len(Gowner[to]); i++ {
											if Gowner[to][i] == x.From_ {
												Gowner[to] = Remove(Gowner[to], Gowner[to][i])
											}
										}
									}
								}
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,"Remove From Group Master:\n",targets)
							}
						} else if MentionMsg == nil {
			    		    yos := strings.Split(text, CUngowner + " ")
			    		    yoss := yos[1]
			    		    contact := CmdList(yoss, Gowner[to])
			    		    for _, vo := range contact {
			    		    	targets = append(targets, vo)
			    		    	for i := 0; i < len(Gowner[to]); i++ {
			    		    		if Gowner[to][i] == vo{
			    		    			Gowner[to] = Remove(Gowner[to], Gowner[to][i])
			    		    		}
			    		    	}
			    		    }
			    		    cok := ExecuteClient(to)
							if cok == Myself {
			    		    	SendReplyMentionByList2(id,to,"Remove From Group Master:\n",targets)
			    		    }
			    		} else {
			    			for _, mention := range MentionMsg {
			    				targets = append(targets, mention)
			    				for i := 0; i < len(Gowner[to]); i++ {
			    					if Gowner[to][i] == mention {
			    						Gowner[to] = Remove(Gowner[to], Gowner[to][i])
			    					}
			    				}
			    			}
			    			cok := ExecuteClient(to)
							if cok == Myself {
			    				SendReplyMentionByList2(id,to,"Remove From Group Master:\n",targets)
			    			}
			    		}
			    		SaveJson()
			    	}
				//GOWNER
				} else if strings.HasPrefix(txt, CMsgsider + " ") && AllAccess(receiver, sender) == PMsgsider {
					result := strings.Split((text), CMsgsider + " ")
					Msgsider = result[1]
					SaveJson()
					InMessage(id, to, true, "Sider Message Change to:\n" +result[1])
				} else if txt == CRuntime && AllAccess(receiver, sender) == PRuntime {
					elapsed := time.Since(startBots)
					tme := fmtDuration(elapsed)
					InMessage(id, to, false,tme)
				} else if strings.HasPrefix(txt , CNewgadmin) && AllAccess(receiver, sender) == PNewgadmin {
					_, found := Gadmin[to]
					if found == false { Gadmin[to] = []string{} }
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if AllAccess(to, mention) > 9 {
								if uncontains(Gadmin[to], mention){
									if mention != Myself{
										targets = append(targets, mention)
										Gadmin[to] = append(Gadmin[to], mention)
										if nothingInMyContacts(mention){
											time.Sleep(1 * time.Second)
											fmt.Println("New Gadmin")
											talk.FindAndAddContactsByMid(mention)
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"New Group Admins:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 9 {
									if uncontains(Gadmin[to], x.From_){
										if x.From_ != Myself{
											Gadmin[to] = append(Gadmin[to], x.From_)
											targets = append(targets,x.From_)
											if nothingInMyContacts(x.From_){
												time.Sleep(1 * time.Second)
												fmt.Println("New Gadmin")
												talk.FindAndAddContactsByMid(x.From_)
											}
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"New Group Admins:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if uncontains(Gadmin[to], Lcontact){
									if Lcontact != Myself{
										Gadmin[to] = append(Gadmin[to], Lcontact)
										targets = append(targets,Lcontact)
										if nothingInMyContacts(Lcontact){
											talk.FindAndAddContactsByMid(Lcontact)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group admin:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if uncontains(Gadmin[to], Lmention){
									if Lmention != Myself{
										Gadmin[to] = append(Gadmin[to], Lmention)
										targets = append(targets,Lmention)
										if nothingInMyContacts(Lmention){
											talk.FindAndAddContactsByMid(Lmention)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group admin:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if uncontains(Gadmin[to], Lkick){
									if Lkick != Myself{
										Gadmin[to] = append(Gadmin[to], Lkick)
										targets = append(targets,Lkick)
										if nothingInMyContacts(Lkick){
											talk.FindAndAddContactsByMid(Lkick)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group admin:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if uncontains(Gadmin[to], Linvite){
									if Linvite != Myself{
										Gadmin[to] = append(Gadmin[to], Linvite)
										targets = append(targets,Linvite)
										if nothingInMyContacts(Linvite){
											talk.FindAndAddContactsByMid(Linvite)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group admin:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if uncontains(Gadmin[to], Lupdate){
									if Lupdate != Myself{
										Gadmin[to] = append(Gadmin[to], Lupdate)
										targets = append(targets,Lupdate)
										if nothingInMyContacts(Lupdate){
											talk.FindAndAddContactsByMid(Lupdate)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group admin:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if uncontains(Gadmin[to], Lleave){
									if Lleave != Myself{
										Gadmin[to] = append(Gadmin[to], Lleave)
										targets = append(targets,Lleave)
										if nothingInMyContacts(Lleave){
											talk.FindAndAddContactsByMid(Lleave)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group admin:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if uncontains(Gadmin[to], Ljoin){
									if Ljoin != Myself{
										Gadmin[to] = append(Gadmin[to], Ljoin)
										targets = append(targets,Ljoin)
										if nothingInMyContacts(Ljoin){
											talk.FindAndAddContactsByMid(Ljoin)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group admin:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if uncontains(Gadmin[to], Lcancel){
									if Lcancel != Myself{
										Gadmin[to] = append(Gadmin[to], Lcancel)
										targets = append(targets,Lcancel)
										if nothingInMyContacts(Lcancel){
											talk.FindAndAddContactsByMid(Lcancel)
										}
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"New Group admin:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CNewgadmin + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt , CUngadmin) && AllAccess(receiver, sender) == PUngadmin {
					_, found := Gadmin[to]
					if found == false { 
						InMessage(id, to, true," Noting Gadmin")
					} else {
						targets:= []string{}
						if msg.RelatedMessageId != ""{
							aa, _ := talk.GetRecentMessagesV2(to, 999)
							lol := msg.RelatedMessageId
							for _, x := range aa{
								if x.ID == lol {
									targets = append(targets, x.From_)
									if contains(Gadmin[to], x.From_){
										for i := 0; i < len(Gadmin[to]); i++ {
											if Gadmin[to][i] == x.From_ {
												Gadmin[to] = Remove(Gadmin[to], Gadmin[to][i])
											}
										}
									}
								}
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,"Remove From Group Admins:\n",targets)
							}
						} else if MentionMsg == nil {
			    		    yos := strings.Split(text, CUngadmin + " ")
			    		    yoss := yos[1]
			    		    contact := CmdList(yoss, Gadmin[to])
			    		    for _, vo := range contact {
			    		    	targets = append(targets, vo)
			    		    	for i := 0; i < len(Gadmin[to]); i++ {
			    		    		if Gadmin[to][i] == vo{
			    		    			Gadmin[to] = Remove(Gadmin[to], Gadmin[to][i])
			    		    		}
			    		    	}
			    		    }
			    		    cok := ExecuteClient(to)
							if cok == Myself {
			    		    	SendReplyMentionByList2(id,to,"Remove From Group Admins:\n",targets)
			    		    }
			    		} else {
			    			for _, mention := range MentionMsg {
			    				targets = append(targets, mention)
			    				for i := 0; i < len(Gadmin[to]); i++ {
			    					if Gadmin[to][i] == mention {
			    						Gadmin[to] = Remove(Gadmin[to], Gadmin[to][i])
			    					}
			    				}
			    			}
			    			cok := ExecuteClient(to)
							if cok == Myself {
			    				SendReplyMentionByList2(id,to,"Remove From Group Admins:\n",targets)
			    			}
			    		}
			    		SaveJson()
			    	}
			    } else if txt == CCleanse && AllAccess(receiver, sender) == PCleanse {
			    	CLEANSEMEMBERS(to)
			    } else if txt == CBreak && AllAccess(receiver, sender) == PBreak {
			    	NUKEJOIN(to)
			    } else if strings.HasPrefix(txt, CKick) && AllAccess(receiver, sender) == PKick {
					if MentionMsg != nil {
						for _, mention := range MentionMsg {
							if AllAccess(to, mention) > 10 {
								go InKick(to, mention)
								go AddedGban(to, mention)
							}
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if AllAccess(to, x.From_) > 10 {
									go InKick(to, x.From_)
									go AddedGban(to, x.From_)
								}
							}
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								go InKick(to, Lcontact)
								go AddedGban(to, Lcontact)
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								go InKick(to, Lmention)
								go AddedGban(to, Lmention)
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								go InKick(to, Lkick)
								go AddedGban(to, Lkick)
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								go InKick(to, Linvite)
								go AddedGban(to, Linvite)
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								go InKick(to, Lupdate)
								go AddedGban(to, Lupdate)
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								go InKick(to, Lleave)
								go AddedGban(to, Lleave)
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								go InKick(to, Ljoin)
								go AddedGban(to, Ljoin)
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								go InKick(to, Lcancel)
								go AddedGban(to, Lcancel)
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage Kick"+CKick+":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
				} else if strings.HasPrefix(txt, CCancel) && AllAccess(receiver, sender) == PCancel {
					cok := ExecuteClient(to)
					if cok == Myself {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" {
								talk.CancelGroupInvitation(to, []string{Lcontact})
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "lkick":
							if Lkick != "" {
								talk.CancelGroupInvitation(to, []string{Lkick})
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" {
								talk.CancelGroupInvitation(to, []string{Linvite})
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" {
								talk.CancelGroupInvitation(to, []string{Lupdate})
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" {
								talk.CancelGroupInvitation(to, []string{Lleave})
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" {
								talk.CancelGroupInvitation(to, []string{Ljoin})
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "ltag":
							if Lmention != "" {
								talk.CancelGroupInvitation(to, []string{Lmention})
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "lcancel":
							if Lcancel != "" {
								talk.CancelGroupInvitation(to, []string{Lcancel})
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage "+CCancel+":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
				} else if strings.HasPrefix(txt, CInvite) && AllAccess(receiver, sender) == PInvite {
					cok := ExecuteClient(to)
					if cok == Myself {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" {
								talk.InviteIntoGroup(to, []string{Lcontact})
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "lkick":
							if Lkick != "" {
								talk.InviteIntoGroup(to, []string{Lkick})
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" {
								talk.InviteIntoGroup(to, []string{Linvite})
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" {
								talk.InviteIntoGroup(to, []string{Lupdate})
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" {
								talk.InviteIntoGroup(to, []string{Lleave})
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" {
								talk.InviteIntoGroup(to, []string{Ljoin})
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" {
								talk.InviteIntoGroup(to, []string{Lcancel})
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "ltag":
							if Lmention != "" {
								talk.InviteIntoGroup(to, []string{Lmention})
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CInvite + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
				//GADMIN
				} else if txt == CHere && AllAccess(receiver, sender) == PHere {
					res, _ := talk.GetCompactGroup(to)
					memlist := res.Members
					var cahasw = []string{}
					for _, v := range memlist {
						if IsBots(v.Mid) == true {
							cahasw = append(cahasw, v.Mid)
						}
					}
					anu := len(cahasw) //+ 1
					anumu := len(Bots) //+ 1
					InMessage(id, to, true, "ʙᴏᴛ ɪɴ sǫᴜᴀᴅ: "+strconv.Itoa(anumu)+"\nʙᴏᴛs ɪɴ ɢʀᴏᴜᴘ: "+strconv.Itoa(anu))
				} else if txt == CTagall && AllAccess(receiver, sender) == PTagall {
					members,_ := talk.GetGroup(to)
					target := members.Members
					targets:= []string{}
					for i:= range target{
						targets = append(targets,target[i].Mid)
					}
					cok := ExecuteClient(to)
					if cok == Myself {
						SendReplyMentionByList2(id,to,"Mentions member:\n",targets)
					}
				} else if txt == CRes && AllAccess(receiver, sender) == PRes {
					talk.SendMessage(to, Respon, map[string]string{})
				} else if txt == CAccess && AllAccess(receiver, sender) == PAccess {
					listacc := "Access In Bots:\n"
					targets:= []string{}
					if len(Master) > 0{
						listacc += "\nMaster: "
						for _,i := range Master{
							listacc += fmt.Sprintf("\n● @!")
							targets = append(targets,i)
						}
					}
					if len(Seler) > 0{
						listacc += "\n\nSeler: "
						for _,i := range Seler{
							listacc += fmt.Sprintf("\n● @!")
							targets = append(targets,i)
						}
					}
					if len(Owners) > 0{
						listacc += "\n\nOwners: "
						for _,i := range Owners{
							listacc += fmt.Sprintf("\n● @!")
							targets = append(targets,i)
						}
					}
					if len(Admins) > 0{
						listacc += "\n\nAdmins: "
						for _,i := range Admins{
							listacc += fmt.Sprintf("\n● @!")
							targets = append(targets,i)
						}
					}
					if len(Staff) > 0{
						listacc += "\n\nStaff: "
						for _,i := range Staff{
							listacc += fmt.Sprintf("\n● @!")
							targets = append(targets,i)
						}
					}
					_, found := Gmaster[to]
					if found == false {
						fmt.Println("Gmaster NotFound")
					} else {
						if len(Gmaster[to]) > 0{
							listacc += "\n\nGmaster: "
							for _,i := range Gmaster[to]{
								listacc += fmt.Sprintf("\n● @!")
								targets = append(targets,i)
							}
						}
					}
					_, found1 := Gowner[to]
					if found1 == false {
						fmt.Println("Gowner NotFound")
					} else {
						if len(Gowner[to]) > 0{
							listacc += "\n\nGowner: "
							for _,i := range Gowner[to]{
								listacc += fmt.Sprintf("\n● @!")
								targets = append(targets,i)
							}
						}
					}
					_, found2 := Gadmin[to]
					if found2 == false {
						fmt.Println("Gadmin NotFound")
					} else {
						if len(Gadmin[to]) > 0{
							listacc += "\n\nGadmin: "
							for _,i := range Gadmin[to]{
								listacc += fmt.Sprintf("\n● @!")
								targets = append(targets,i)
							}
						}
					}
					cok := ExecuteClient(to)
					if cok == Myself {
						talk.SendMentionWeb(id, to, listacc , targets, oupTit, oupLogo, justgood)
					}
					
				} else if strings.HasPrefix(txt, CLinkpro + " ") && AllAccess(receiver, sender) == PLinkpro {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "on":
						if !helper.InArray(Linkpro, to) { Linkpro = append(Linkpro, to) }
						InMessage(id, to, true, "LinkProtect Enabled.")
					case "off":
						if helper.InArray(Linkpro, to) { Linkpro = Remove(Linkpro, to) }
						InMessage(id, to, true, "LinkProtect Disabled.")
					}
					SaveJson()
				} else if strings.HasPrefix(txt, CNamelock + " ") && AllAccess(receiver, sender) == PNamelock {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "on":
						g, _ := talk.GetCompactGroup(to)
						if !helper.InArray(Namelock, to) { Gname[to] = g.Name;Namelock = append(Namelock, to) }
						InMessage(id, to, true, "NameLock Enabled.")
					case "off":
						if helper.InArray(Namelock, to) { Namelock = Remove(Namelock, to) ;delete(Gname,to) }
						InMessage(id, to, true, "NameLock Disabled.")
					}
					SaveJson()
				} else if strings.HasPrefix(txt, CDenyin + " ") && AllAccess(receiver, sender) == PDenyin {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "on":
						if !helper.InArray(Denyinv, to) { Denyinv = append(Denyinv, to) }
						InMessage(id, to, true, "Denyinvite Enabled.")
					case "off":
						if helper.InArray(Denyinv, to) { Denyinv = Remove(Denyinv, to) }
						InMessage(id, to, true, "Denyinvite Disabled.")
					}
					SaveJson()
				} else if strings.HasPrefix(txt, CProjoin + " ") && AllAccess(receiver, sender) == PProjoin {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "on":
						if !helper.InArray(Projoin, to) { Projoin = append(Projoin, to) }
						InMessage(id, to, true, "ProJoined Enabled.")
					case "off":
						if helper.InArray(Projoin, to) { Projoin = Remove(Projoin, to) }
						InMessage(id, to, true, "ProJoined Disabled.")
					}
					SaveJson()
				} else if strings.HasPrefix(txt, CProtect + " ") && AllAccess(receiver, sender) == PProtect {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "on":
						if !helper.InArray(Protect, to) { Protect = append(Protect, to) }
						InMessage(id, to, true, "Protect Enabled.")
					case "off":
						if helper.InArray(Protect, to) { Protect = Remove(Protect, to) }
						InMessage(id, to, true, "Protect Disabled.")
					case "max":
						g, _ := talk.GetCompactGroup(to)
						if !helper.InArray(Linkpro, to) { Linkpro = append(Linkpro, to) }
						if !helper.InArray(Protect, to) { Protect = append(Protect, to) }
						if !helper.InArray(Namelock, to) { Gname[to] = g.Name;Namelock = append(Namelock, to) }
						if !helper.InArray(Denyinv, to) { Denyinv = append(Denyinv, to) }
						InMessage(id, to, true, "Protect Max Enabled.")
					case "none":
						if helper.InArray(Linkpro, to) { Linkpro = Remove(Linkpro, to) }
						if helper.InArray(Protect, to) { Protect = Remove(Protect, to) }
						if helper.InArray(Namelock, to) { Namelock = Remove(Namelock, to) ;delete(Gname,to) }
						if helper.InArray(Denyinv, to) { Denyinv = Remove(Denyinv, to) }
						InMessage(id, to, true, "Protect Max Disabled.")
					}
					SaveJson()
				} else if txt == CCenterstay && AllAccess(receiver, sender) == PCenterstay {
					if IsCenter(Myself) == true {
						talk.LeaveGroup(to)
					} else {
						InvitedAjs(to)
					}
				} else if txt == CCheckcenter && AllAccess(receiver, sender) == PCheckcenter {
					res, _ := talk.GetGroup(to)
					memlist := res.Invitee
					aktf := false
					for _, v := range memlist {
						if IsCenter(v.Mid) {
							aktf = true
						}
					}
					if aktf == true {
						InMessage(id, to, true, "Center be in a group invited.")
					}
				} else if strings.HasPrefix(txt, CClear + " ") && AllAccess(receiver, sender) == PClear {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "seller":
						jum := len(Seler)
						Seler = Seler[:0]
						str := fmt.Sprintf("Cleared %v Seller", jum)
						InMessage(id, to, true, str)
					case "owner":
						jum := len(Owners)
						Owners = Owners[:0]
						str := fmt.Sprintf("Cleared %v Owners", jum)
						InMessage(id, to, true, str)
					case "admin":
						jum := len(Admins)
						Admins = Admins[:0]
						str := fmt.Sprintf("Cleared %v Admins", jum)
						InMessage(id, to, true, str)
					case "staff":
						jum := len(Staff)
						Staff = Staff[:0]
						str := fmt.Sprintf("Cleared %v Staff", jum)
						InMessage(id, to, true, str)
					case "gmaster":
						_, found := Gmaster[to]
						if found == false { 
							InMessage(id, to, true, "Noting Gmaster")
						} else {
							jum := len(Gmaster[to])
							Gmaster[to] = Gmaster[to][:0]
							str := fmt.Sprintf("Cleared %v Gmaster", jum)
							InMessage(id, to, true, str)
						}
					case "gowner":
						_, found := Gowner[to]
						if found == false { 
							InMessage(id, to, true, "Noting Gowner")
						} else {
							jum := len(Gowner[to])
							Gowner[to] = Gowner[to][:0]
							str := fmt.Sprintf("Cleared %v Gowner", jum)
							InMessage(id, to, true, str)
						}
					case "gadmin":
						_, found := Gadmin[to]
						if found == false { 
							InMessage(id, to, true, "Noting Gadmin")
						} else {
							jum := len(Gadmin)
							Gadmin[to] = Gadmin[to][:0]
							str := fmt.Sprintf("Cleared %v Gadmin", jum)
							InMessage(id, to, true, str)
						}
					case "bots":
						jum := len(Bots)
						Bots = Bots[:0]
						str := fmt.Sprintf("Cleared %v Bots", jum)
						InMessage(id, to, true, str)
					case "center":
						jum := len(Center)
						if jum != 0 {
							for _, z := range Center {
								Bots = append(Bots, z)
								cok := ExecuteClient(to)
								if cok == Myself {
									JoinAj(to, z)
								}
							}
						}
						Center = Center[:0]
						str := fmt.Sprintf("Cleared %v Center", jum)
						InMessage(id, to, true, str)
					case "fuck":
						jum := len(Banned)
						Banned = Banned[:0]
						str := fmt.Sprintf("Cleared %v Banned", jum)
						InMessage(id, to, true, str)
						WarMode = false
					case "hiden":
						jum := len(Hiden)
						Hiden = Hiden[:0]
						str := fmt.Sprintf("Cleared %v Hiden", jum)
						InMessage(id, to, true, str)
					case "allban":
						Gban = map[string][]string{}
						str := fmt.Sprintf("Cleared All Group Banned")
						InMessage(id, to, true, str)
						WarMode = false
					case "ban":
						_, found := Gban[to]
						if found == false { 
							InMessage(id, to, true, "Noting Gban")
							WarMode = false
						} else {
							jum := len(Gban[to])
							Gban[to] = Gban[to][:0]
							str := fmt.Sprintf("Cleared %v Banned", jum)
							InMessage(id, to, true, str)
							WarMode = false
						}
					case "?":
						var tot = []string{"seller","owner","admin","staff","gmaster","gowner","gadmin","bots","center","fuck","ban","allban","hiden"}
						stas := "❏ Usage " + CClear + ":\n"
						for _, t := range tot {
							stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
						}
						InMessage(id, to, true, stas)
					}
					SaveJson()
				} else if strings.HasPrefix(txt, CList + " ") && AllAccess(receiver, sender) == PList {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "seller":
						listbl := "Seller List:\n"
						targets:= []string{}
						if len(Seler) > 0 {
							for _, i := range Seler{
								targets = append(targets,i)
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,listbl,targets)
							}
						}else{InMessage(id, to, true, "Seller NotFound.")}
					case "owner":
						listbl := "Owner List:\n"
						targets:= []string{}
						if len(Owners) > 0 {
							for _, i := range Owners{
								targets = append(targets,i)
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,listbl,targets)
							}
						}else{InMessage(id, to, true, "Owners NotFound.")}
					case "admin":
						listbl := "Admins List:\n"
						targets:= []string{}
						if len(Admins) > 0 {
							for _, i := range Admins{
								targets = append(targets,i)
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,listbl,targets)
							}
						}else{InMessage(id, to, true, "Admins NotFound.")}
					case "staff":
						listbl := "Staff List:\n"
						targets:= []string{}
						if len(Staff) > 0 {
							for _, i := range Staff{
								targets = append(targets,i)
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,listbl,targets)
							}
						}else{InMessage(id, to, true, "Staff NotFound.")}
					case "gmaster":
						_, found := Gmaster[to]
						if found == false { 
							InMessage(id, to, true, "Noting Gmaster")
						} else {
							listbl := "Gmaster List:\n"
							targets:= []string{}
							if len(Gmaster[to]) > 0 {
								for _, i := range Gmaster[to]{
									targets = append(targets,i)
								}
								cok := ExecuteClient(to)
								if cok == Myself {
									SendReplyMentionByList2(id,to,listbl,targets)
								}
							}else{InMessage(id, to, true, "Gmaster NotFound.")}
						}
					case "gowner":
						_, found := Gowner[to]
						if found == false { 
							InMessage(id, to, true, "Noting Gowner")
						} else {
							listbl := "Gowner List:\n"
							targets:= []string{}
							if len(Gowner[to]) > 0 {
								for _, i := range Gowner[to]{
									targets = append(targets,i)
								}
								cok := ExecuteClient(to)
								if cok == Myself {
									SendReplyMentionByList2(id,to,listbl,targets)
								}
							}else{InMessage(id, to, true, "Gowner NotFound.")}
						}
					case "gadmin":
						_, found := Gadmin[to]
						if found == false { 
							InMessage(id, to, true," Noting Gadmin")
						} else {
							listbl := "Gadmin List:\n"
							targets:= []string{}
							if len(Gadmin[to]) > 0 {
								for _, i := range Gadmin[to]{
									targets = append(targets,i)
								}
								cok := ExecuteClient(to)
								if cok == Myself {
									SendReplyMentionByList2(id,to,listbl,targets)
								}
							}else{InMessage(id, to, true, "Gadmin NotFound.")}
						}
					case "bots":
						listbl := "Bots List:\n"
						targets:= []string{}
						if len(Bots) > 0 {
							for _, i := range Bots{
								targets = append(targets,i)
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,listbl,targets)
							}
						}else{InMessage(id, to, true, "Bots NotFound.")}
					case "center":
						listbl := "Center List:\n"
						targets:= []string{}
						if len(Center) > 0 {
							for _, i := range Center{
								targets = append(targets,i)
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,listbl,targets)
							}
						}else{InMessage(id, to, true, "Center NotFound.")}
					case "fuck":
						listbl := "Fuck List:\n"
						targets:= []string{}
						if len(Banned) > 0 {
							for _, i := range Banned{
								targets = append(targets,i)
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,listbl,targets)
							}
						}else{InMessage(id, to, true, "Banned NotFound.")}
					case "hiden":
						listbl := "Hiden List:\n"
						targets:= []string{}
						if len(Hiden) > 0 {
							for _, i := range Hiden{
								targets = append(targets,i)
							}
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,listbl,targets)
							}
						}else{InMessage(id, to, true, "Hiden NotFound.")}
					case "allban":
						Gban = map[string][]string{}
						InMessage(id, to, true, "Success Clear Group Banned")
					case "ban":
						_, found := Gban[to]
						if found == false { 
							InMessage(id, to, true, "Noting Banned")
						} else {
							listbl := "Banned List:\n"
							targets:= []string{}
							if len(Gban[to]) > 0 {
								for _, i := range Gban[to] {
									targets = append(targets,i)
								}
								cok := ExecuteClient(to)
								if cok == Myself {
									SendReplyMentionByList2(id,to,listbl,targets)
								}
								WarMode = false
							}else{InMessage(id, to, true, "Banned NotFound.")}
						}
					case "?":
						var tot = []string{"seller","owner","admin","staff","gmaster","gowner","gadmin","bots","center","fuck","ban","allban","hiden"}
						stas := "❏ Usage " + CClear + ":\n"
						for _, t := range tot {
							stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
						}
						InMessage(id, to, true, stas)
					}
					
				} else if strings.HasPrefix(txt, CAutopurge + " ") && AllAccess(receiver, sender) == PAutopurge {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "on":
						Purge = true
						InMessage(id, to, true, "AutoPurge Enabled.")
					case "off":
						Purge = false
						InMessage(id, to, true, "AutoPurge Disabled.")
					}
				} else if strings.HasPrefix(txt, CLockdown + " ") && AllAccess(receiver, sender) == PLockdown {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "on":
						Lockdown = true
						InMessage(id, to, true, "Lockdown Enabled.")
					case "off":
						Lockdown = false
						InMessage(id, to, true, "Lockdown Disabled.")
					}
				} else if strings.HasPrefix(txt, CJoinNuke + " ") && AllAccess(receiver, sender) == PJoinNuke {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "on":
						JoinNuke = true
						InMessage(id, to, true, "Join Nuke Enabled.")
					case "off":
						JoinNuke = false
						InMessage(id, to, true, "Join Nuke Disabled.")
					}
				} else if txt == "cek" && AllAccess(receiver, sender) < 8 {
					var tot string
					res := talk.KickoutFromGroupV2(to, []string{"FuckYou"})
					if strings.Contains(res.Error(), "request blocked") {
						Blocked = true
					}
					if Blocked == true {
						tot = "Request Blocked"
					} else {
						tot = "Request Ready"
					}
					talk.SendText(to, tot, 2)
				} else if txt == "time" && AllAccess(receiver, sender) < 8 {
					GenerateTimeLog(id, to)
				} else if txt == CPurge && AllAccess(receiver, sender) == PPurge {
					go Autopurge(to)
				} else if strings.HasPrefix(txt, CLogmode + " ") && AllAccess(receiver, sender) == PLogmode {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "on":
						Logs = true
						Logmode = to
						InMessage(id, to, false, "LogMode Enabled.")
					case "off":
						Logs = false
						Logmode = to
						InMessage(id, to, false, "LogMode Disabled.")
					}
					SaveJson()
				} else if strings.HasPrefix(txt, CSider + " ") && AllAccess(receiver, sender) == PSider {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "on":
						Sider[to] = []string{}
						SiderV2[to] = true
						InMessage(id, to, false, "Reader Enabled.")
					case "off":
						SiderV2[to] = false
						if Sider[to] != nil {
							cok := ExecuteClient(to)
							if cok == Myself {
								SendReplyMentionByList2(id,to,"** Reader Disabled **\nHistory:\n",Sider[to])
							}
						} else { InMessage(id, to, false, "Reader Disabled.")}
						delete(Sider, to)
					}
				} else if strings.HasPrefix(txt, CKillmode + " ") && AllAccess(receiver, sender) == PKillmode {
					result := strings.Split((txt)," ")
					switch result[1] {
					case "random":
						Killmode = 0
						InMessage(id, to, true, "Killmode Set To Random.")
					case "purge":
						Killmode = 1
						InMessage(id, to, true, "Killmode Set To Purge.")
					case "kill":
						Killmode = 2
						InMessage(id, to, true, "Killmode Set To Kill.")
					}
					SaveJson()
				} else if strings.HasPrefix(txt, CHiden) && AllAccess(receiver, sender) == PHiden {
					targets:= []string{}
					if MentionMsg != nil {
						for _, mention := range MentionMsg{
							if !IsHiden(mention) {
								if uncontains(Hiden, mention){
									targets = append(targets, mention)
									Hiden = append(Hiden, mention)
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Hiden:\n",targets)
						}
					} else if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa {
							if x.ID == lol {
								if !IsHiden(x.From_) {
									Hiden = append(Hiden, x.From_)
									targets = append(targets,x.From_)
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Hiden:\n",targets)
						}
					} else {
						result := strings.Split((txt)," ")
						switch result[1] {
						case "lcontact":
							if Lcontact != "" && AllAccess(to, Lcontact) > 10 {
								if !IsHiden(Lcontact) {
									if Lcontact != Myself{
										Hiden = append(Hiden, Lcontact)
										targets = append(targets,Lcontact)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"Hiden:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have LContact")
							}
						case "ltag":
							if Lmention != "" && AllAccess(to, Lmention) > 10 {
								if !IsHiden(Lmention) {
									if Lmention != Myself{
										Hiden = append(Hiden, Lmention)
										targets = append(targets,Lmention)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"Hiden:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lmention")
							}
						case "lkick":
							if Lkick != "" && AllAccess(to, Lkick) > 10 {
								if !IsHiden(Lkick){
									if Lkick != Myself{
										Hiden = append(Hiden, Lkick)
										targets = append(targets,Lkick)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"Hiden:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lkick")
							}
						case "linvite":
							if Linvite != "" && AllAccess(to, Linvite) > 10 {
								if !IsHiden(Linvite){
									if Linvite != Myself{
										Hiden = append(Hiden, Linvite)
										targets = append(targets,Linvite)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"Hiden:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Linvite")
							}
						case "lupdate":
							if Lupdate != "" && AllAccess(to, Lupdate) > 10 {
								if !IsHiden(Lupdate){
									if Lupdate != Myself{
										Hiden = append(Hiden, Lupdate)
										targets = append(targets,Lupdate)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"Hiden:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lupdate")
							}
						case "lleave":
							if Lleave != "" && AllAccess(to, Lleave) > 10 {
								if !IsHiden(Lleave){
									if Lleave != Myself{
										Hiden = append(Hiden, Lleave)
										targets = append(targets,Lleave)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"Hiden:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lleave")
							}
						case "ljoin":
							if Ljoin != "" && AllAccess(to, Ljoin) > 10 {
								if !IsHiden(Ljoin){
									if Ljoin != Myself{
										Hiden = append(Hiden, Ljoin)
										targets = append(targets,Ljoin)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"Hiden:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Ljoin")
							}
						case "lcancel":
							if Lcancel != "" && AllAccess(to, Lcancel) > 10 {
								if !IsHiden(Lcancel){
									if Lcancel != Myself{
										Hiden = append(Hiden, Lcancel)
										targets = append(targets,Lcancel)
										cok := ExecuteClient(to)
										if cok == Myself {
											SendReplyMentionByList2(id,to,"Hiden:\n",targets)
										}
									}
								}
							} else {
								InMessage(id, to, true, "Not Have Lcancel")
							}
						case "?":
							var tot = []string{"lcontact","lkick","linvite","lupdate","lleave","ljoin","lcancel","ltag"}
							stas := "❏ Usage " + CHiden + ":\n"
							for _, t := range tot {
								stas += fmt.Sprintf("\n➥ %s",strings.Title(t))
							}
							InMessage(id, to, true, stas)
						}
					}
					SaveJson()
				} else if strings.HasPrefix(txt , CUnhiden) && AllAccess(receiver, sender) == PUnhiden {
					targets:= []string{}
					if msg.RelatedMessageId != ""{
						aa, _ := talk.GetRecentMessagesV2(to, 999)
						lol := msg.RelatedMessageId
						for _, x := range aa{
							if x.ID == lol {
								targets = append(targets, x.From_)
								if IsHiden(x.From_){
									for i := 0; i < len(Hiden); i++ {
										if Hiden[i] == x.From_ {
											Hiden = Remove(Hiden, Hiden[i])
										}
									}
								}
							}
						}
						cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Remove From Hiden:\n",targets)
						}
					} else if MentionMsg == nil {
			    	    yos := strings.Split(text, CUnhiden + " ")
			    	    yoss := yos[1]
			    	    contact := CmdList(yoss, Hiden)
			    	    for _, vo := range contact {
			    	    	targets = append(targets, vo)
			    	    	for i := 0; i < len(Hiden); i++ {
			    	    		if Hiden[i] == vo{
			    	    			Hiden = Remove(Hiden, Hiden[i])
			    	    		}
			    	    	}
			    	    }
			    	    cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Remove From Hiden:\n",targets)
						}
			    	} else {
			    		for _, mention := range MentionMsg {
			    			targets = append(targets, mention)
			    			for i := 0; i < len(Hiden); i++ {
			    				if Hiden[i] == mention {
			    					Hiden = Remove(Hiden, Hiden[i])
			    				}
			    			}
			    		}
			    		cok := ExecuteClient(to)
						if cok == Myself {
							SendReplyMentionByList2(id,to,"Remove From Hiden:\n",targets)
						}
			    	}
			    	SaveJson()
				} else if txt == CSet && AllAccess(receiver, sender) == PSet {
					cok := ExecuteClient(to)
					if cok == Myself {
						gname := ""
						g, _ := talk.GetCompactGroup(to)
						if len(g.Name) > 17 {gname = g.Name[:17]+"..."
						} else {gname = g.Name}
						stf := fmt.Sprintf("%s\n𝗚𝗥𝗢𝗨𝗣 𝗦𝗘𝗧𝗧𝗜𝗡𝗚𝗦:",gname)
						if helper.InArray(Protect, to) {stf += "\n➥ ⚫️ Protect"
						} else {stf += "\n➥ ⚪️ Protect"}
						if helper.InArray(Linkpro, to) {stf += "\n➥ ⚫️ LinkPro"
						} else {stf += "\n➥ ⚪️ LinkProtect"}
						if helper.InArray(Denyinv, to) {stf += "\n➥ ⚫️ DenyInvite"
						} else {stf += "\n➥ ⚪️ DenyInvite"}
						if helper.InArray(Namelock, to) {stf += "\n➥ ⚫️ NameLock"
						} else {stf += "\n➥ ⚪️ NameLock"}
						if helper.InArray(Projoin, to) {stf += "\n➥ ⚫️ ProJoin"
						} else {stf += "\n➥ ⚪️ ProJoin"}
						stf += "\n\n𝐆𝐋𝐎𝐁𝐀𝐋 𝗦𝗘𝗧𝗧𝗜𝗡𝗚𝗦:"
						if JoinNuke == true {stf += "\n➥ ⚫️ NukeJoin"
						} else {stf += "\n➥ ⚪️ NukeJoin"}
						if Purge == true {stf += "\n➥ ⚫️ AutoPurge"
						} else {stf += "\n➥ ⚪️ AutoPurge"}
						if Lockdown == true {stf += "\n➥ ⚫️ LockDown"
						} else {stf += "\n➥ ⚪️ LockDown"}
						if Logs == true {stf += "\n➥ ⚫️ LogMode"
						} else {stf += "\n➥ ⚪️ LogMode"}
						if Blocked == true {stf += "\n➥ ⚫️ Blocked"
						} else {stf += "\n➥ ⚪️ Blocked"}
						if WarMode == true {stf += "\n➥ ⚫️ WarMode"
						} else {stf += "\n➥ ⚪️ WarMode"}
						if Killmode == 0 {stf += "\n➥ KillMode: Rand"
						} else if Killmode == 1 {stf += "\n➥ KillMode: Purge"
						} else if Killmode == 2 {stf += "\n➥ KillMode: Kill"}
						reg, err := talk.GetCountryWithRequestIp()
						if err != nil { fmt.Println(err) }
						stf += "\n➥ Region: "+reg
						talk.SendFooter(id, to, stf, oupTit, oupLogo, justgood)
					}
				} else if txt == CHelp && AllAccess(receiver, sender) == PHelp{
					cmd := fmt.Sprintf("𝐃𝐄𝐕𝐄𝐋𝐎𝐏𝐄𝐑 𝐁𝐎𝐓𝐒\n©2022 Imjustgood Team\n\nRname: %s\nSname: %s\nKey: `%s`",Rname,Sname,Key)
					var employee = map[string]int{
						CNewseller 	: 	PNewseller,
						CUnseller 	: 	PUnseller,
						CNewowner 	: 	PNewowner,
						CUnowner 	: 	PUnowner,
						CJoinall 	: 	PJoinall,
						CUpkey 		: 	PUpkey,
						CUprespon 	: 	PUprespon,
						CUpbio 		: 	PUpbio,	
						CUpname 	: 	PUpname,
						CUpsname 	: 	PUpsname,
						CUnfriends 	: 	PUnfriends,
						CNewadmin 	: 	PNewadmin,
						CUnadmin 	: 	PUnadmin,
						CSetlimit 	: 	PSetlimit,
						CAddme 		: 	PAddme,
						CJoino 		: 	PJoino,
						CLeaveto 	: 	PLeaveto,
						CInvto 		: 	PInvto,
						CUrljoined 	: 	PUrljoined,
						CNewstaff 	: 	PNewstaff,
						CUnstaff 	: 	PUnstaff,
						CNewcenter 	: 	PNewcenter,
						CUncenter 	: 	PUncenter,
						CNewbots 	: 	PNewbots,
						CUnbots 	: 	PUnbots,
						CNewgmaster : 	PNewgmaster,
						CUngmaster 	: 	PUngmaster,
						CContact 	: 	PContact,
						CMid 		: 	PMid,
						CFriends 	: 	PFriends,
						CGinvited 	: 	PGinvited,
						CGroups 	: 	PGroups,
						CSpeed 		: 	PSpeed,
						COurl 		: 	POurl,
						CCurl 		: 	PCurl,
						CUnsend 	: 	PUnsend,
						CUpgname 	: 	PUpgname,
						CNewban 	: 	PNewban,
						CUnban 		: 	PUnban,
						CNewgowner 	: 	PNewgowner,
						CUngowner 	: 	PUngowner,
						CRuntime 	: 	PRuntime,
						CNewgadmin 	: 	PNewgadmin,
						CUngadmin 	: 	PUngadmin,
						CKick 		: 	PKick,
						CHere 		: 	PHere,
						CTagall 	: 	PTagall,
						CRes 		: 	PRes,
						CAccess 	: 	PAccess,
						CLinkpro 	: 	PLinkpro,
						CNamelock 	: 	PNamelock,
						CDenyin 	: 	PDenyin,
						CProjoin 	: 	PProjoin,
						CProtect 	: 	PProtect,
						CAutopurge 	: 	PAutopurge,
						CLockdown 	: 	PLockdown,
						CJoinNuke 	: 	PJoinNuke,
						CLogmode 	: 	PLogmode,
						CPurge 		: 	PPurge,
						CKillmode 	: 	PKillmode,
						CHelp 		: 	PHelp,
						CList 		: 	PList,
						CClear 		: 	PClear,
						CNewfuck    :   PNewfuck,
						CUnfuck     :   PUnfuck,
						CUpimage    :   PUpimage,
						CBye        : 	PBye,
						CTimeleft   : 	PTimeleft,
						CExtend     : 	PExtend,
						CCleanse    : 	PCleanse,
						CBreak      : 	PBreak,
						CCenterstay : 	PCenterstay,
						CCheckcenter: 	PCheckcenter,
						CSet        : 	PSet,
						CReduce     :   PReduce,
					}
					cmd += "\n\n❏ 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 :\n"
					for cm, perm := range employee {
						cmd += fmt.Sprintf("\n    ➥ %s [%v]",strings.Title(cm), perm)
					}
					cmd += "\n\n𝗜𝗰𝗼𝗻 𝗜𝗻𝗳𝗼𝗿𝗺𝗮𝘁𝗶𝗼𝗻 :\nRank:\n0. Master 1. Seller\n2. Owner 3. Admin\n4. Staff 5. Gmaster\n6. Gowner 7. Gadmin"
        	       	InMessage(id, to, true, cmd)
				}
			}
			if len(Tcm) != 0 {
				strs := strings.Join(Tcm, ",\n")
				if strs != "" {
					cok := ExecuteClient(to)
					if cok == Myself {
						talk.SendMessage(to, strs, map[string]string{})
					} else {
						if cok == "" {
							talk.SendMessage(to, strs, map[string]string{})
						}
					}
				}
				Tcm = []string{}
			}

			if (msg.ContentType).String() == "IMAGE"{
				if Changepic == true {
					if sender == xsender {
						if len(msg.ContentPreview) == 0 {
	    					callProfile(msg.ContentMetadata["DOWNLOAD_URL"],"picture2")
	    					InMessage(id, to, false,"Success change Profile Image")
	    				} else {
	    					callProfile(id,"picture")
	    					InMessage(id, to, false,"Success change Profile Image")
	    				}
					}
					Changepic = false
				}
			} else if (msg.ContentType).String() == "VIDEO"{
			} else if (msg.ContentType).String() == "CONTACT"{
				Lcontact = msg.ContentMetadata["mid"]
			}
		}
	}
 }
}

func InHeader() {
	f, err := os.Open("proxies.txt")
	if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan(){
    	result := strings.Split((scanner.Text()),"] ")[1]
    	tot := strings.Replace(result, ">","",1)
    	Proxy = append(Proxy, tot)
    }
    if err := scanner.Err(); err != nil {
    	fmt.Println(err)
    }
    prox, _ := Choice(Proxy)
    proxx := strings.Split((prox),":")[0]
    service.Proxy = proxx
	con.APP_TYPE = "CHANNELCP"
	con.SYSTEM_NAME = "Android OS"
	con.SYSTEM_VER = "10.0.5"
	con.VER = "2.17.0"
	con.LINE_APPLICATION = con.APP_TYPE + "\t" + con.VER + "\t" + con.SYSTEM_NAME + "\t" + con.SYSTEM_VER
	fmt.Println(con.LINE_APPLICATION)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if _, err := os.Stat(CmdData); os.IsNotExist(err) {
		CmdsSave()
	} else { CmdsLoad() }
	if _, err := os.Stat(CmdPermit); os.IsNotExist(err) {
		PermitSave()
	} else { PermitLoad() }
	filepath := fmt.Sprintf("token/%s.txt", DB)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("Enter Token: ")
		var Token string
		fmt.Scanln(&Token)
		f, err := os.Create(filepath)
		if err != nil { fmt.Println(err) }
		defer f.Close()
		_, err2 := f.WriteString(Token)
		if err2 != nil {fmt.Println(err2) }
		InHeader()
		auth.LoginWithAuthToken(Token)
	} else {
		b, err := ioutil.ReadFile(filepath)
		if err != nil {fmt.Println(err)}
		token := string(b)
		InHeader()
		auth.LoginWithAuthToken(token)
	}
	if _, err := os.Stat(DataBase); os.IsNotExist(err) {
		Sname = "default"
		Rname = DB
		loc, _ := time.LoadLocation("Asia/Jakarta")
		tt := time.Now();now := tt.In(loc)
		t1 := now.AddDate(0, 1, 0);TimeLeft = t1
		ab := "** GOBOT STARTING UP **"
		ab += "\n Rname: " + DB
		pr, _ := talk.GetProfile()
		Myself = pr.Mid
		name := pr.DisplayName
		ab += "\n Name: " + name
		ab += "\n Mid: " + Myself
		abc := ""
		fmt.Println(string(helper.ColorPurple), ab, string(helper.ColorReset))
		fmt.Println(abc)
		SaveJson()
		for {
			fetch, _ := talk.FetchOperations(service.Revision, 50)
			if len(fetch) > 0 {
				rev := fetch[0].Revision
				service.Revision = helper.MaxRevision(service.Revision, rev)
				Executor(fetch[0])
			}
		}
	} else {
		LoadJson()
		expired := Expired()
		ab := "** GOBOT STARTING UP **"
		ab += "\n Rname: " + DB
		pr, _ := talk.GetProfile()
		Myself = pr.Mid
		name := pr.DisplayName
		ab += "\n Name: " + name
		ab += "\n Mid: " + Myself
		abc := ""
		fmt.Println(string(helper.ColorPurple), ab, string(helper.ColorReset))
		fmt.Println(abc)
		fmt.Println(expired)
		if expired == true {
			Getout()
			for {
				fetch, _ := talk.FetchOperations(service.Revision, 50)
				if len(fetch) > 0 {
					rev := fetch[0].Revision
					service.Revision = helper.MaxRevision(service.Revision, rev)
					//fmt.Println(fetch[0].Type)
					Executor(fetch[0])
				}
			}
			//os.Exit(2)
		} else {
			for {
				fetch, _ := talk.FetchOperations(service.Revision, 50)
				if len(fetch) > 0 {
					rev := fetch[0].Revision
					service.Revision = helper.MaxRevision(service.Revision, rev)
					//fmt.Println(fetch[0].Type)
					Executor(fetch[0])
				}
			}
		}
	}
}