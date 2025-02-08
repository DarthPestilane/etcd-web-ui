package controller

import (
	"context"
	"etcd-web-ui/cmd/start/flags"
	"etcd-web-ui/pkg/bizio"
	"github.com/gin-gonic/gin"
	etcdCli "go.etcd.io/etcd/client/v3"
	"time"
)

func NoRoute(c *gin.Context) {
	bizio.HttpFail(c, 404, "api not found")
}

func GetOptions(c *gin.Context) {
	bizio.HttpOK(c, gin.H{
		"endpoint":  flags.Endpoint,
		"username":  flags.Username,
		"password":  flags.Password,
		"keyPrefix": flags.KeyPrefix,
	})
}

func newEtcdCli(endpoint, username, password string) (*etcdCli.Client, error) {
	return etcdCli.New(etcdCli.Config{
		Endpoints: []string{endpoint},
		Username:  username,
		Password:  password,
	})
}

func Refresh(c *gin.Context) {
	var req struct {
		Endpoint  string `json:"endpoint" binding:"required"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		KeyPrefix string `json:"keyPrefix"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		bizio.HttpFail(c, 422, err.Error())
		return
	}

	cli, err := newEtcdCli(req.Endpoint, req.Username, req.Password)
	if err != nil {
		bizio.HttpFail(c, 400, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	resp, err := cli.Get(ctx, req.KeyPrefix, etcdCli.WithPrefix(), etcdCli.WithKeysOnly())
	if err != nil {
		bizio.HttpFail(c, 400, err.Error())
		return
	}

	etcdKeys := make([]string, 0, resp.Count)
	for _, kv := range resp.Kvs {
		etcdKeys = append(etcdKeys, string(kv.Key))
	}
	bizio.HttpOK(c, gin.H{
		"etcdKeys": etcdKeys,
	})
}

func Get(c *gin.Context) {
	var req struct {
		Endpoint string `json:"endpoint" binding:"required"`
		Username string `json:"username"`
		Password string `json:"password"`
		EtcdKey  string `json:"etcdKey" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		bizio.HttpFail(c, 422, err.Error())
		return
	}

	cli, err := newEtcdCli(req.Endpoint, req.Username, req.Password)
	if err != nil {
		bizio.HttpFail(c, 400, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	resp, err := cli.Get(ctx, req.EtcdKey)
	if err != nil {
		bizio.HttpFail(c, 400, err.Error())
		return
	}

	var content string
	var ttl int64
	if resp.Count != 0 {
		kv := resp.Kvs[0]
		content = string(kv.Value)

		// get ttl
		ttlResp, err := cli.TimeToLive(ctx, etcdCli.LeaseID(kv.Lease))
		if err != nil {
			bizio.HttpFail(c, 400, err.Error())
			return
		}
		ttl = ttlResp.TTL
	}

	bizio.HttpOK(c, gin.H{
		"content": content,
		"ttl":     ttl,
	})
}

func Put(c *gin.Context) {
	var req struct {
		Endpoint string `json:"endpoint" binding:"required"`
		EtcdKey  string `json:"etcdKey" binding:"required"`
		Content  string `json:"content"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		bizio.HttpFail(c, 422, err.Error())
		return
	}

	cli, err := newEtcdCli(req.Endpoint, req.Username, req.Password)
	if err != nil {
		bizio.HttpFail(c, 400, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	if _, err := cli.Put(ctx, req.EtcdKey, req.Content); err != nil {
		bizio.HttpFail(c, 400, err.Error())
		return
	}

	bizio.HttpOK(c, "ok")
}

func Delete(c *gin.Context) {
	var req struct {
		Endpoint string `json:"endpoint" binding:"required"`
		EtcdKey  string `json:"etcdKey" binding:"required"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		bizio.HttpFail(c, 422, err.Error())
		return
	}

	cli, err := newEtcdCli(req.Endpoint, req.Username, req.Password)
	if err != nil {
		bizio.HttpFail(c, 400, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	if _, err := cli.Delete(ctx, req.EtcdKey); err != nil {
		bizio.HttpFail(c, 400, err.Error())
		return
	}

	bizio.HttpOK(c, "ok")
}
