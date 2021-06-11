package machine

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/knight-zlm/eTerm/model"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func RunSSHTerminal(id int) {
	mch, err := model.GetMachineByID(id)
	if err != nil {
		log.Fatalf("GetMachineByID(%v) err:%v", id, err)
	}

	runSSHTerminal(mch)
}

func NewSSHClient(h *model.Machine) (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		Timeout:         time.Second * 5,
		User:            h.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if h.Type == "password" {
		sshConfig.Auth = []ssh.AuthMethod{ssh.Password(h.Password)}
	} else {
		sshConfig.Auth = []ssh.AuthMethod{publicKeyAuthFunc(h.Key)}
	}

	var addr string
	if h.Host != "" {
		addr = h.Host
	} else {
		addr = fmt.Sprintf("%s:%d", h.Ip, h.Port)
	}

	cli, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatalf("find key path failed err:%v", err)
	}

	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatalf("read key path failed err:%v", err)
	}

	singer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("ssh key signer failed err:%v", err)
	}

	return ssh.PublicKeys(singer)
}

func runSSHTerminal(h *model.Machine) {
	cli, err := NewSSHClient(h)
	if err != nil {
		log.Fatalf("NewSSHClient err:%v", err)
	}

	session, err := cli.NewSession()
	if err != nil {
		log.Fatalf("ssh NewSession err:%v", err)
	}
	defer session.Close()

	err = interactiveSession(session)
	if err != nil {
		log.Printf("interactiveSession err:%v", err)
	}
}

func interactiveSession(sess *ssh.Session) error {
	// 拿到当前终端文件描述符
	fd := int(os.Stdin.Fd())
	if !terminal.IsTerminal(fd) {
		osName := runtime.GOOS
		return fmt.Errorf("%s fd %d is not a terminal,can't create pty of ssh", osName, fd)
	}
	// make raw 解决回显问题
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer terminal.Restore(fd, state)

	// 设置term颜色
	termW, termH, err := terminal.GetSize(fd)
	if err != nil {
		return err
	}

	termType := os.Getenv("TERM")
	if termType == "" {
		termType = "xterm-256color"
	}

	// request pty
	err = sess.RequestPty(termType, termH, termW, ssh.TerminalModes{})
	if err != nil {
		return err
	}

	// 监听本地窗口大小变动，调整session的窗口大小。否则会出现显示问题
	go updateTerminalSize(sess)

	// 拷贝sess的输出到本地的输出
	stdout, err := sess.StdoutPipe()
	if err != nil {
		return err
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := sess.StderrPipe()
	if err != nil {
		return err
	}
	go io.Copy(os.Stderr, stderr)

	// 拷贝本地输入到远端输入
	stdin, err := sess.StdinPipe()
	go io.Copy(stdin, os.Stdin)
	//go func() {
	//	buf := make([]byte, 128)
	//	for {
	//		n, err := os.Stdin.Read(buf)
	//		if err != nil {
	//			log.Printf("read stdin err:%v", err)
	//			return
	//		}
	//
	//		if n == 0 {
	//			continue
	//		}
	//
	//		_, err = stdin.Write(buf)
	//		if err != nil {
	//			log.Printf("send stdin err:%v", err)
	//			return
	//		}
	//	}
	//}()

	err = sess.Shell()
	if err != nil {
		return err
	}

	err = sess.Wait()
	if err != nil {
		return err
	}

	return nil
}

func updateTerminalSize(sess *ssh.Session) {
	sigWinCh := make(chan os.Signal, 1)
	// 注册chan 监听窗口大小变动消息
	signal.Notify(sigWinCh, syscall.SIGWINCH)

	fd := int(os.Stdin.Fd())
	termW, termH, err := terminal.GetSize(fd)
	if err != nil {
		log.Printf("get terminal win size failed, err:%v", err)
	}

	for {
		select {
		case sig := <-sigWinCh:
			if sig == nil {
				return
			}
			curW, curH, err := terminal.GetSize(fd)
			if err != nil {
				log.Printf("get terminal win size failed, err:%v", err)
				continue
			}

			// 判断一下窗口尺寸是否有改变
			if curW == termW && curH == termH {
				continue
			}

			// 更新远端的窗口大小
			err = sess.WindowChange(curH, curW)
			if err != nil {
				log.Printf("change terminal win size failed, err:%v", err)
				continue
			}

			// 记录当前窗口大小
			termW, termH = curW, curH
		}
	}
}
