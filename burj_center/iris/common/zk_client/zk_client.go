package zk_client

import (
	"encoding/json"
	"log"
	"math/rand"
	"path"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/alex19861108/burj/burj_center/general/common/net"
	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/kataras/iris/core/errors"
	"github.com/samuel/go-zookeeper/zk"
)

func watch(e zk.Event) {
	log.Println("*******************")
	log.Println("path:", e.Path)
	log.Println("type:", e.Type.String())
	log.Println("state:", e.State.String())
	log.Println("-------------------")
}

type HttpServer func(config.Config)

const (
	timeout = 10
)

func InitZkClient(cfg config.Config) (err error) {
	options := zk.WithEventCallback(watch)
	conn, _, err := zk.Connect([]string{cfg.BurjCenterConfig.ZkServer}, time.Duration(timeout)*time.Second, options)
	if err != nil {
		return
	}
	defer conn.Close()

	addr, err := net.GetHostAddr()
	node := proto.Node{
		Addr: &proto.Addr{
			Http: GenServerAddr(addr.String(), cfg.ServerConfig.Port.Http),
			Rpc:  GenServerAddr(addr.String(), cfg.ServerConfig.Port.Rpc),
		},
		Labels: cfg.ServerConfig.Name,
	}
	err = RigisterServer(conn, cfg.BurjCenterConfig.ZkRoot, node)
	if err != nil {
		panic(err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	return
}

func RigisterServer(conn *zk.Conn, root string, node proto.Node) (err error) {
	// create root node
	exist, _, err := conn.Exists(root)
	if err != nil {
		log.Fatalf("get server list failed: %s \n", err)
		panic(err)
		return
	}
	if !exist {
		_, err = conn.Create(root, nil, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			return
		}
	}

	nodeName := GenNodeName(node.Labels, node.Addr.Http)
	nodePath := path.Join(root, nodeName)
	nodeContent, err := json.Marshal(node)
	// create child node
	_, err = conn.Create(nodePath, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Fatalf("ceate node failed: %s \n", err)
		panic(err)
		return
	}
	// set node data
	_, err = conn.Set(nodePath, nodeContent, -1)
	if err != nil {
		log.Fatalf("ceate node failed: %s \n", err)
		panic(err)
		return
	}
	return
}

func SearchOneNode(cfg config.Config, labels string) (proto.Node, error) {
	conn, _, err := zk.Connect([]string{cfg.BurjCenterConfig.ZkServer}, time.Duration(timeout)*time.Second)
	if err != nil {
		return proto.Node{}, err
	}
	defer conn.Close()

	children, _, err := conn.Children(cfg.BurjCenterConfig.ZkRoot)
	if len(children) > 0 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		child := children[r.Intn(len(children))]
		label, _, ok := ParseNodeName(child)
		if ok && label == labels {
			childPath := path.Join(cfg.BurjCenterConfig.ZkRoot, child)
			c, _, _ := conn.Get(childPath)
			node := proto.Node{}
			err := json.Unmarshal(c, &node)
			return node, err
		}
	}
	return proto.Node{}, errors.New("No Children Found!")
}

func NodeGet(cfg config.Config) (node proto.Node, err error) {
	conn, _, err := zk.Connect([]string{cfg.BurjCenterConfig.ZkServer}, time.Duration(timeout)*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()

	addr, err := net.GetHostAddr()
	nodeName := GenNodeName(cfg.ServerConfig.Name, GenServerAddr(addr.String(), cfg.ServerConfig.Port.Http))
	nodePath := path.Join(cfg.BurjCenterConfig.ZkRoot, nodeName)
	c, _, err := conn.Get(nodePath)
	if err != nil {
		return
	}
	err = json.Unmarshal(c, &node)
	return
}

func NodeSet(cfg config.Config, node proto.Node) error {
	conn, _, err := zk.Connect([]string{cfg.BurjCenterConfig.ZkServer}, time.Duration(timeout)*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()

	addr, err := net.GetHostAddr()
	nodeName := GenNodeName(cfg.ServerConfig.Name, GenServerAddr(addr.String(), cfg.ServerConfig.Port.Http))
	nodePath := path.Join(cfg.BurjCenterConfig.ZkRoot, nodeName)
	c, err := json.Marshal(node)
	_, err = conn.Set(nodePath, c, -1)
	return err
}

func GenServerAddr(h string, p string) string {
	return h + ":" + p
}

func GenNodeName(name string, addr string) string {
	return name + "[" + addr + "]"
}

func ParseNodeName(node string) (string, string, bool) {
	reg := regexp.MustCompile(`(?P<name>.*)\[(?P<host>.*)\]`)
	submatchs := reg.FindStringSubmatch(node)
	if len(submatchs) == 3 {
		return submatchs[1], submatchs[2], true
	}
	return "", "", false
}

func ParseNodePath(node string) (string, string, string, bool) {
	reg := regexp.MustCompile(`(?P<root>.*/)(?P<name>.*)\[(?P<host>.*)\]`)
	submatchs := reg.FindStringSubmatch(node)
	if len(submatchs) == 4 {
		return strings.TrimRight(submatchs[1], "/"), submatchs[2], submatchs[3], true
	}
	return "", "", "", false
}
