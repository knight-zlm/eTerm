package cmd

import (
	"github.com/knight-zlm/eTerm/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var addSSHCmd = &cobra.Command{
	Use:   "add",
	Short: "add a ssh connection configuration",
	Long:  `add a ssh connection,usage: eterm add -p my_password -k ~/.ssh/id_rsa -n mySSH -a 192.168.0.01:22 -u root --auth=key`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := model.MachineAdd(name, addr, ip, user, password, key, auth, port); err != nil {
			logrus.Fatalf("addSSHCmd MachineAdd err:%v", err)
		}
	},
}

var name, addr, ip, user, password, key, auth string
var port uint

func init() {
	rootCmd.AddCommand(addSSHCmd)
	// StringVarP 支持短命名
	addSSHCmd.Flags().StringVarP(&name, "name", "n", "", "显示名称")
	addSSHCmd.Flags().StringVarP(&addr, "addr", "a", "", "远端的ip:port 或者 domain")
	addSSHCmd.Flags().StringVarP(&ip, "ip", "", "", "远端的ip")
	addSSHCmd.Flags().StringVarP(&user, "user", "u", "root", "登陆的用户名")
	addSSHCmd.Flags().StringVarP(&password, "password", "p", "", "远端的登陆密码")
	addSSHCmd.Flags().StringVarP(&key, "key", "k", "~/.ssh/id_rsa", "远端登陆用的密钥文件路径")
	addSSHCmd.Flags().StringVarP(&auth, "auth", "", "password", "远端登陆的类型,密码(password)或者密钥(key)")
	addSSHCmd.Flags().UintVar(&port, "port", 22, "远端登陆端口")
	// 确定必填字段
	if err := addSSHCmd.MarkFlagRequired("addr"); err != nil {
		logrus.Fatalf("addSSHCmd init err:%v", err)
	}
	if err := addSSHCmd.MarkFlagRequired("name"); err != nil {
		logrus.Fatalf("addSSHCmd init err:%v", err)
	}
}
