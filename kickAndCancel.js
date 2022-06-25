const thrift = require('thrift-http')
const TalkService = require('./LineThrift/TalkService')
const {
  CancelChatInvitationRequest,
  DeleteOtherFromChatRequest
} = require('./LineThrift/TalkService_types')

var _client = '';
var gid = '';
var kick = [];
var cancel = []
var method = '';
var token = '';
var app = '';

var APP_VERSION = {
  "DESKTOPMAC": "6.5.1",
  "IOSIPAD": "10.21.5",
  "DESKTOPWIN": "6.5.1",
  "CHROMEOS": "2.4.0",
  "ANDROIDLITE": "2.16.0",
  "IOS": "10.21.5",
  "CHANNELCP": "2.17.0"
}
var SYS_NAME = {
  "DESKTOPMAC": "DEADLINE",
  "IOSIPAD": "iPadOS 14",
  "IOS": "iOS 10",
  "ANDROIDLITE": "Realme 5 Pro",
  "DESKTOPWIN": "DEADLINE",
  "CHROMEOS": "ChromeOS",
  "CHANNELCP": "Android OS"
}

var SYS_VERSION = {
  "DESKTOPMAC": "10.14.0",
  "IOSIPAD": "14.0.1",
  "IOS": "10.3.3",
  "ANDROIDLITE": "10;SECONDARY",
  "DESKTOPWIN": "10.0",
  "CHROMEOS": "88.0",
  "CHANNELCP": "10.0.5"
}

process.argv.forEach(function(val) {
  if (val.includes("gid=")) {
    gid = val.split("gid=").pop();
  } else if (val.includes("uid=")) {
    kick.push(val.split("uid=").pop());
  } else if (val.includes("uids=")) {
    cancel.push(val.split("uids=").pop());
  } else if (val.includes("method=")) {
    method = val.split("method=").pop();
  } else if (val.includes("token=")) {
    token = val.split("token=").pop();
  } else if (val.includes("app=")) {
    app = val.split("app=").pop();
  }
});

var APP_VER = APP_VERSION[app];
var SYSTEM_NAME = SYS_NAME[app];
var SYSTEM_VER = SYS_VERSION[app];
var APP_NAME = `${app} ${APP_VER} ${SYSTEM_NAME} ${SYSTEM_VER}`
var USER_AGENT = `Line/${APP_VER}`

console.info(`\n\
Auth Token : ${token}\n\
Group ID : ${gid}\n\
Member : ${kick}\n\
Invite : ${cancel}\n\
Method : ${method}\n\
App Name : ${APP_NAME}\n\
`)

function setTHttpClient(options) {
  var connection = thrift.createHttpConnection("ga2.line.naver.jp", 443, options);
  connection.on("error", (err) => {
    console.log("err", err);
    return err;
  });
  _client = thrift.createHttpClient(TalkService, connection);
}

setTHttpClient(options = {
  protocol: thrift.TCompactProtocol,
  transport: thrift.TBufferedTransport,
  headers: {
    "User-Agent": USER_AGENT,
    "X-Line-Application": APP_NAME,
    "X-Line-Access": token
  },
  path: "/S4",
  https: true
});

if (method == 'cancel') {
  async function cancelAll() {
    let cancelPromise = new Promise((resolve, reject) => {
      try {
        for (var i = 0; i < cancel.length; i++) {
          var request = new CancelChatInvitationRequest()
          request.reqSeq = 0;
          request.chatMid = gid;
          request.targetUserMids = [cancel[i]]
          _client.cancelChatInvitation(request);
        }
        resolve("Cancel Done ✓")
      } catch (e) {
        reject(e);
      }
    });
    return cancelPromise;
  }
  var cancelPromise = cancelAll();
  Promise.all([cancelPromise])
  .then(results => console.log(results));
} else if (method == 'kick') {
  async function kickAll() {
    let kickPromise = new Promise((resolve, reject) => {
      try {
        for (var i = 0; i < kick.length; i++) {
          var request = new DeleteOtherFromChatRequest()
          request.reqSeq = 0;
          request.chatMid = gid;
          request.targetUserMids = [kick[i]]
          _client.deleteOtherFromChat(request);
        }
        resolve("Kick done ✓")
      } catch (e) {
        reject(e);
      }
    });
    return kickPromise;
  }
  var kickPromise = kickAll();
  Promise.all([kickPromise])
  .then(results => console.log(results));
} else if (method == 'kickandcancel') {
  async function cancelAll() {
    let cancelPromise = new Promise((resolve, reject) => {
      try {
        for (var i = 0; i < cancel.length; i++) {
          var request = new CancelChatInvitationRequest()
          request.reqSeq = 0;
          request.chatMid = gid;
          request.targetUserMids = [cancel[i]]
          _client.cancelChatInvitation(request);
        }
        resolve("Cancel Done ✓")
      } catch (e) {
        reject(e);
      }
    });
    return cancelPromise;
  }
  async function kickAll() {
    let kickPromise = new Promise((resolve, reject) => {
      try {
        for (var i = 0; i < kick.length; i++) {
          var request = new DeleteOtherFromChatRequest()
          request.reqSeq = 0;
          request.chatMid = gid;
          request.targetUserMids = [kick[i]]
          _client.deleteOtherFromChat(request);
        }
        resolve("Kick done ✓")
      } catch (e) {
        reject(e);
      }
    });
    return kickPromise;
  }
  var cancelPromise = cancelAll();
  var kickPromise = kickAll();
  Promise.all([cancelPromise, kickPromise])
  .then(results => console.log(results));
}
