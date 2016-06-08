//****************************************************/
// https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package models

import (
	"encoding/base64"
	"flag"
	// "io"

	"os"
	"os/exec"
	"strings"

	"github.com/kr/pty"
	"github.com/revel/revel"
	"golang.org/x/net/websocket"
)

var addrFlag, cmdFlag, staticFlag string

func init() {
	// go sshroom()
}

type wsPty struct {
	Cmd *exec.Cmd // pty builds on os.exec
	Pty *os.File  // a pty is simply an os.File
}

func (wp *wsPty) Start() {
	var err error
	args := flag.Args()
	wp.Cmd = exec.Command(cmdFlag, args...)
	wp.Pty, err = pty.Start(wp.Cmd)
	if err != nil {
		revel.ERROR.Println("Failed to start command: %s\n", err)
	}
}

func (wp *wsPty) Stop() {
	wp.Pty.Close()
	wp.Cmd.Wait()
}

type SSHConfig struct {
	Id    string
	Agent *Agent
	WsPty *wsPty
}

func (sshConfig *SSHConfig) Login() string {
	var err error
	sshConfig.WsPty.Cmd = exec.Command("ssh", "-o", "StrictHostKeyChecking=no", sshConfig.Agent.LoginName+"@"+sshConfig.Agent.Host)
	sshConfig.WsPty.Pty, err = pty.Start(sshConfig.WsPty.Cmd)
	if err != nil {
		revel.ERROR.Println("Failed to start command: %s\n", err)
	}

	i := 0
	for {
		if i >= 10 {
			revel.ERROR.Println("login error")
			sshConfig.WsPty.Pty.Close()
			return "login error"
		}
		buf := make([]byte, 1024)
		size, err := sshConfig.WsPty.Pty.Read(buf)
		if err != nil {
			revel.ERROR.Println("login Read error")
		}

		if !strings.Contains(string([]byte(buf[:size])), "password") {
			i++
			continue
		}
		var loginPass = desUtil.Decrypt(sshConfig.Agent.LoginPass)
		_, err = sshConfig.WsPty.Pty.Write([]byte(loginPass + "\n"))
		if err != nil {
			revel.ERROR.Println("login Write error")
		}

		if sshConfig.Agent.InitShellCmd != "" {
			_, err = sshConfig.WsPty.Pty.Write([]byte(sshConfig.Agent.InitShellCmd + "\n"))
			if err != nil {
				revel.ERROR.Println("InitShellCmd Error")
			}
		}
		return ""
	}
}

func NewSSHConfig(agentId string) SSHConfig {
	agent := SelectAgentOneByAgentId(agentId)
	return SSHConfig{
		Agent: &agent,
		Id:    agentId,
	}
}

func Shell(conn *websocket.Conn, sshConfig SSHConfig) {
	defer conn.Close()
	client := conn.Request().RemoteAddr
	revel.INFO.Println("Client connected:", client)

	sshConfig.WsPty = &wsPty{}
	if sshConfig.Login() != "" {
		return
	}

	go func() {
		buf := make([]byte, 128)
		for {
			n, err := sshConfig.WsPty.Pty.Read(buf)
			if err != nil {
				revel.ERROR.Printf("Failed to read from pty master: %s", err)
				return
			}

			out := make([]byte, base64.StdEncoding.EncodedLen(n))
			base64.StdEncoding.Encode(out, buf[0:n])

			conn.Write(out)
			if err != nil {
				revel.ERROR.Printf("Failed to send %d bytes on websocket: %s", n, err)
				return
			}
		}
	}()

	var newMessages string
	for {
		err := websocket.Message.Receive(conn, &newMessages)
		if err != nil {
			return
		}
		buf := make([]byte, base64.StdEncoding.DecodedLen(len(newMessages)))
		_, err = base64.StdEncoding.Decode(buf, []byte(newMessages))
		if err != nil {
			revel.ERROR.Printf("base64 decoding of newMessages failed: %s\n", err)
		}
		sshConfig.WsPty.Pty.Write(buf)
	}
	sshConfig.WsPty.Stop()
}
