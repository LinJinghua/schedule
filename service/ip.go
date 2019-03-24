package service

import (
    "fmt"
    "sync"
    "sync/atomic"
)

type ipPool struct {
    ips []string
    cur uint64
    lock sync.RWMutex
}

const (
    INIT_SIZE int = 10;
)

var (
    read_pool ipPool
    write_pool ipPool
)

func init()  {
    read_pool.init()
    write_pool.init()
}

func getRPool() *ipPool {
    return &read_pool
}

func getWPool() *ipPool {
    return &write_pool
}

func (pool *ipPool) init()  {
    pool.ips = make([]string, 0, INIT_SIZE)
    pool.cur = 0
}

// GetIP : GetIP
func (pool *ipPool) GetIP() string {
    if (len(pool.ips) == 0) {
        return ""
    }
    atomic.AddUint64(&pool.cur, 1)
    // fmt.Println("GetIP", pool.ips)
    pool.lock.RLock()
    defer pool.lock.RUnlock()
    return pool.ips[pool.cur % uint64(len(pool.ips))]
}

// GetIPFree : GetIPFree
func (pool *ipPool) GetIPFree() string {
    if (len(pool.ips) == 0) {
        return ""
    }
    atomic.AddUint64(&pool.cur, 1)
    // fmt.Println("GetIPFree", pool.ips)
    return pool.ips[pool.cur % uint64(len(pool.ips))]
}


// AddIP : AddIP
func (pool *ipPool) AddIP(ip string) {
    pool.lock.Lock()
    defer pool.lock.Unlock()
    for _, v := range pool.ips {
        if v == ip {
            return
        }
    }
    pool.ips = append(pool.ips, ip)
    fmt.Println("AddIP", pool.ips)
}

// DelIP : DelIP
func (pool *ipPool) DelIP(ip string) {
    pool.lock.Lock()
    defer pool.lock.Unlock()
    for i := 0; i < len(pool.ips); i++ {
        if pool.ips[i] == ip {
            pool.ips[i] = pool.ips[len(pool.ips) - 1]
            pool.ips = pool.ips[:len(pool.ips) - 1]
            break
        }
    }
    fmt.Println("DelIP", pool.ips)
}

