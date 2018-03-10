package zk_server

import (
	"encoding/json"
	"log"
	"path"
	"sync"
	"time"

	"github.com/alex19861108/burj/burj_center/iris/common/zk_client"
	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/alex19861108/burj/burj_center/iris/services"
	"github.com/samuel/go-zookeeper/zk"
)

var wg *sync.WaitGroup //定义同步等待组

const (
	timeout = 10
)

func watch(e zk.Event) {
	log.Println("*******************")
	log.Println("path:", e.Path)
	log.Println("type:", e.Type.String())
	log.Println("state:", e.State.String())
	log.Println("-------------------")
}

func watchNodeCreate(conn *zk.Conn, path string) {
	for {
		_, _, ch, _ := conn.ExistsW(path)
		e := <-ch
		if e.Type == zk.EventNodeCreated {
			return
		}
	}
}

func watchNodeDataChange(conn *zk.Conn, path string) {
	for {
		_, _, ch, _ := conn.GetW(path)
		e := <-ch
		if e.Type == zk.EventNodeDataChanged {
			return
		}
	}
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string

	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func watchNodeChildrenChanged(conn *zk.Conn, p string, service services.NodeService) {
	for {
		children, _, ch, _ := conn.ChildrenW(p)
		e := <-ch
		if e.Type == zk.EventNodeChildrenChanged {
			// 增加节点 or 删除节点
			currChildren, _, _ := conn.Children(p)
			if len(currChildren) > len(children) {
				// 增加节点
				for _, n := range difference(currChildren, children) {
					childNodePath := path.Join(p, n)
					data, _, err := conn.Get(childNodePath)
					if err != nil {
						continue
					}

					node := proto.Node{}
					if err = json.Unmarshal(data, &node); err != nil {
						//panic(err)
						continue
					}
					// insert into db
					service.InsertOrUpdate(node)

					go watchChildNode(conn, childNodePath, service)
				}
			}
		}
	}
}

func watchChildNode(conn *zk.Conn, path string, service services.NodeService) {
	for {
		_, _, ch, err := conn.GetW(path)
		if err != nil {
			panic(err)
			return
		}
		_, label, addr, ok := zk_client.ParseNodePath(path)
		if !ok {
			return
		}

		e := <-ch
		switch e.Type {
		case zk.EventNodeDataChanged:
			data, _, err := conn.Get(path)
			if err != nil {
				return
			}

			node := proto.Node{}
			if err = json.Unmarshal(data, &node); err != nil {
				return
			}
			node, err = service.InsertOrUpdate(node)
		case zk.EventNodeDeleted:
			service.DeleteByLabelAndAddr(label, addr)
			return
		}
	}
}

func watchNodeDeleted(conn *zk.Conn, path string) {
	for {
		_, _, ch, _ := conn.ExistsW(path)
		e := <-ch
		if e.Type == zk.EventNodeDeleted {
			return
		}
	}
}

func watchNode(conn *zk.Conn, path string, service services.NodeService) {
	wg.Add(1)
	go watchNodeCreate(conn, path)
	go watchNodeDataChange(conn, path)
	go watchNodeChildrenChanged(conn, path, service)
	watchNodeDeleted(conn, path)
	wg.Done()
}

func InitZkServer(cfg config.BurjCenterConfig, service services.NodeService) (err error) {
	options := zk.WithEventCallback(watch)
	conn, _, err := zk.Connect([]string{cfg.ZkServer}, time.Second*time.Duration(timeout), options)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	wg = &sync.WaitGroup{}
	watchNode(conn, cfg.ZkRoot, service)
	wg.Wait()

	return
}
