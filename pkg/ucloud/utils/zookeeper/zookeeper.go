/*=============================================================================
#     FileName: zookeeper.go
#         Desc:
#       Author: ato.ye
#        Email: ato.ye@ucloud.cn
#     HomePage: http://www.ucloud.cn
#      Version: 0.0.1
#   LastChange: 2018-01-24 15:20:22
#      History:
=============================================================================*/
package zookeeper

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	uflog "server/ucloud/log"
	"server/ucloud/utils/zookeeper/zk"
)

var (
	ZkConnPoolMu      sync.Mutex
	ZkConnPool        = make(map[string]*ZkConn)
	ErrInvalidServers = errors.New("invalid connect string")
	ErrConnect        = errors.New("can not connect remote servers")
)

type ZkConn struct {
	cstr string
	conn *zk.Conn
}

func initZkConn(connstr string) (zkConn *ZkConn, err error) {
	if connstr == "" {
		err = ErrInvalidServers
		uflog.ERRORF("Connect zk server failed:%s", err.Error())
		return
	}
	conn_strs := strings.Split(connstr, ",")
	conn, ec, err := zk.Connect(conn_strs, 3*time.Second)
	if err != nil {
		err = ErrInvalidServers
		uflog.ERRORF("Connect zk servers %s failed:%s", conn_strs, err.Error())
		return nil, err
	}

	for {
		select {
		case connEvent, ok := <-ec:
			if ok {
				switch connEvent.State {
				case zk.StateHasSession:
					zkConn = &ZkConn{
						cstr: connstr,
						conn: conn,
					}
					return zkConn, nil
				default:
					continue
				}
			} else {
				err = ErrConnect
				uflog.ERRORF("Connect zk Servers %s failed:%s", conn_strs, err.Error())
				return nil, err
			}
		default:
			continue
		}
	}
}

// 获取zk集群连接实例，连接字符串格式:ip:port,ip:port
func GetZkInstance(connstr string) (zkConn *ZkConn, err error) {
	ZkConnPoolMu.Lock()
	defer ZkConnPoolMu.Unlock()
	zkConn, ok := ZkConnPool[connstr]
	if ok {
		if zkConn.conn.State() != zk.StateHasSession {
			zkConn, err = initZkConn(connstr)
			if err != nil {
				return nil, err
			}
			ZkConnPool[connstr] = zkConn
		}
		return zkConn, nil
	}

	zkConn, err = initZkConn(connstr)
	if err != nil {
		return nil, err
	}
	ZkConnPool[connstr] = zkConn
	return zkConn, nil
}

// 创建节点
func (c *ZkConn) CreateNode(path string, data []byte) (resPath string, err error) {
	if path == "" {
		return "", errors.New("Invalid Path")
	}

	// 节点参数
	flag := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)
	childPath := path

	// 创建父节点
	paths := strings.Split(path, "/")
	var parentPath string
	for _, v := range paths[1 : len(paths)-1] {
		parentPath += "/" + v
		exist, _, err := c.NodeExists(parentPath)
		if err != nil {
			return "", err
		}
		if !exist {
			_, err = c.conn.Create(parentPath, nil, 0, acl) // 父节点必须是持久节点
			if err != nil {
				return "", err
			}
		}
	}

	// 创建子节点
	exist, _, err := c.NodeExists(childPath)
	if err != nil {
		return "", err
	}
	if !exist {
		resPath, err = c.conn.Create(childPath, data, flag, acl)
		if err != nil {
			return "", err
		}
	} else {
		err = fmt.Errorf("[%s]  exists", childPath)
	}
	return
}

func (c *ZkConn) SetNode(path string, data []byte) (err error) {
	if path == "" {
		return errors.New("Invalid Path")
	}
	// 判断节点是否存在
	exist, stat, err := c.NodeExists(path)
	if err != nil {
		return fmt.Errorf("check node[%s] exist fail:%s", path, err)
	}
	if !exist {
		return fmt.Errorf("node [%s] dosen't exist,can't be setted", path)
	}
	_, err = c.conn.Set(path, data, stat.Version)
	if err != nil {
		return
	}
	return
}

func (c *ZkConn) GetNode(path string) (data []byte, err error) {
	// 判断路径是否为空
	if path == "" {
		return nil, errors.New("Invalid Path")
	}

	data, _, err = c.conn.Get(path)
	return
}

func (c *ZkConn) DeleteNode(path string) (err error) {
	// 判断路径是否为空
	if path == "" {
		return errors.New("Invalid Path")
	}

	// 判断节点是否存在
	exist, stat, err := c.NodeExists(path)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("path [\"%s\"] doesn't exist", path)
	}
	// 删除节点
	return c.conn.Delete(path, stat.Version)
}

func (c *ZkConn) ListChildren(path string) (children []string, err error) {
	// 判断路径是否为空
	if path == "" {
		return nil, errors.New("Invalid Path")
	}

	children, _, err = c.conn.Children(path)
	return
}

//获取当前节点（完整路径）的watcher
func (c *ZkConn) GetZNodeWatcher(path string) ([]byte, *zk.Stat, <-chan zk.Event, error) {
	return c.conn.GetW(path)
}

//获取当前节点所有子节点变化的wather
func (c *ZkConn) GetChildrenWatcher(path string) ([]string, *zk.Stat, <-chan zk.Event, error) {
	return c.conn.ChildrenW(path)
}

// 判断节点是否存在
func (c *ZkConn) NodeExists(path string) (exist bool, stat *zk.Stat, err error) {
	// 判断路径是否为空
	if path == "" {
		return false, nil, errors.New("Invalid Path")
	}
	// 判断节点是否存在
	return c.conn.Exists(path)
}
