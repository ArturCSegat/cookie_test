package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Session struct{
    name string
    exp time.Time
}

func (s Session) isExpired() bool {
	return s.exp.Before(time.Now())
}

var sessions = map[string]Session{
    "xxx": Session{name: "artur", exp: time.Now().Add(60 * time.Second)},
    "yyy": Session{name: "jullia", exp: time.Now().Add(60 * time.Second)},
}


func name(c* gin.Context){
    cookie, err := c.Cookie("session_id")
    fmt.Println(cookie)
    name := ""    

    if err != nil{
        name = "muhamad"
    }
    
    session, exists := sessions[cookie]

    if !exists{
        name = "muhamad"
    }
    if session.isExpired(){
            delete(sessions, cookie)
            name = "muhamad"
    }

    if name == ""{
        name = session.name
    }
    c.JSON(http.StatusOK, gin.H{"name":name})
}


func new_session(c * gin.Context){
    id := uuid.NewString()

    name := c.Param("name") 
    exp := time.Now().Add(60 * time.Second)

    sessions[id] = Session{name: name, exp:exp}
    fmt.Println(id)
    c.SetCookie("session_id", id, 60, "/", "localhost", false, false)

    c.JSON(http.StatusOK, gin.H{"message": "session created for 60 seconds"})
}

func clean_expired_sessions(){
    for {
        time.Sleep(30 * time.Second)
        for session_id, session := range sessions{
            if session.isExpired(){
                delete(sessions, session_id)
                println("deleted session:", session_id, "value: ", session.name)
            }
        }
        println("finished a check")
    }
}

func main(){
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:2000")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
        }
        c.Next()
    })

    r.GET("/name", name)
    r.GET("/register/:name", new_session)
    go clean_expired_sessions()
    r.Run(":3000")
}
