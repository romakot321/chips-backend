package models

const (
  MessageEventAuthenticate = iota
  MessageEventConnected = iota
  MessageEventUserMove = iota
  MessageEventEntityAdd = iota
  MessageEventUsers = iota
  MessageEventChangeScore = iota
  MessageEventWin = iota
  MessageEventRestart = iota
)

type message struct {
  Event int
  Data interface{}
}

type MessageAuthenticateData struct {
  Username string `json:"username"`
  Room string `json:"room"`
}

type MessageAuthenticate struct {
  message
  Event int `json:"event"`
  Data MessageAuthenticateData `json:"data"`
}

func MakeMessageAuthenticate(data map[string]interface{}) MessageAuthenticate {
  return MessageAuthenticate{
    Event: MessageEventAuthenticate,
    Data: MessageAuthenticateData{
      Username: data["data"].(map[string]interface{})["username"].(string),
      Room: data["data"].(map[string]interface{})["room"].(string),
    },
  }
}

type MessageConnected struct {
  message
  Event int `json:"event"`
  Data interface{} `json:"data"`
}

type MessageUserMoveData struct {
  X int `json:"x"`
  Y int `json:"y"`
  Username string `json:"username"`
}

type MessageUserMove struct {
  message
  Event int `json:"event"`
  Data MessageUserMoveData `json:"data"`
}

func MakeMessageUserMove(data map[string]interface{}) MessageUserMove {
  msgData := data["data"].(map[string]interface{});
  username, ok := msgData["username"].(string);
  if (!ok) { username = ""; }

  return MessageUserMove{
    Event: MessageEventUserMove,
    Data: MessageUserMoveData{
      X: int(msgData["x"].(float64)),
      Y: int(msgData["y"].(float64)),
      Username: username,
    },
  }
}

type MessageChangeScoreData struct {
  Amount int `json:"amount"`
  Username string `json:"username"`
}

type MessageChangeScore struct {
  message
  Event int `json:"event"`
  Data MessageChangeScoreData `json:"data"`
}

func MakeMessageChangeScore(data map[string]interface{}) MessageChangeScore {
  msgData := data["data"].(map[string]interface{});
  username, ok := msgData["username"].(string);
  if (!ok) { username = ""; }

  return MessageChangeScore{
    Event: MessageEventChangeScore,
    Data: MessageChangeScoreData{
      Amount: int(msgData["amount"].(float64)),
      Username: username,
    },
  }
}

type MessageWinData struct {
  Name string `json:"name"`
}

type MessageWin struct {
  message
  Event int `json:"event"`
  Data MessageWinData `json:"data"`
}

func MakeMessageWin(data map[string]interface{}) MessageWin {
  msgData := data["data"].(map[string]interface{});

  return MessageWin{
    Event: MessageEventWin,
    Data: MessageWinData{
      Name: msgData["name"].(string),
    },
  }
}

type MessageRestartData struct { }

type MessageRestart struct {
  message
  Event int `json:"event"`
  Data MessageRestartData `json:"data"`
}

func MakeMessageRestart(data map[string]interface{}) MessageRestart {
  return MessageRestart{
    Event: MessageEventRestart,
    Data: MessageRestartData{},
  }
}
