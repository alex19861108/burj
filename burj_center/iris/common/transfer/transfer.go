package transfer

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func DownloadFromHttp(ul string, dir string) (string, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	res, err := http.Get(ul)
	if err != nil {
		panic(err)
	}

	u, err := url.Parse(ul)
	if err != nil {
		panic(err)
	}
	localFilePath := path.Join(dir, path.Base(u.Path))
	dst, err := os.Create(localFilePath)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(dst, res.Body)
	if err != nil {
		panic(err)
	}
	return localFilePath, err
}

func DownloadFromSftp(ul string, dir string) (string, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	u, err := url.Parse(ul)
	if err != nil {
		panic(err)
	}

	localFilePath := path.Join(dir, path.Base(u.Path))
	sftp, err := connect(u)
	defer sftp.Close()
	if err != nil {
		panic(err)
	}

	f, err := sftp.Open(u.Path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	dst, err := os.Create(localFilePath)
	defer dst.Close()
	if err != nil {
		panic(err)
	}

	if _, err = f.WriteTo(dst); err != nil {
		panic(err)
	}

	return localFilePath, err
}

func connect(u *url.URL) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	password, _ := u.User.Password()
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    u.User.Username(),
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr = fmt.Sprintf("%s:%s", u.Host, "22")
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}
