package middleware

import (
    "fmt"
    "github.com/dulei0726/gateway/pkg"
    "github.com/dulei0726/gateway/pkg/errcode"
    "github.com/gin-gonic/gin"
    "github.com/hashicorp/consul/api"
    "math/rand"
    "net"
    "strconv"
    "sync"
    "time"
)

var (
    serviceCache = NewServiceCache()
    client       = NewServiceDiscoveryClient()
)

func ServiceDiscovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        serviceName := c.Query("service")
        value, _ := serviceCache.onceMap.LoadOrStore(serviceName, new(sync.Once))
        once := value.(*sync.Once)
        once.Do(func() {
            go func() {
                loop(serviceName, serviceCache.svcMap)
            }()
        })
        value, ok := serviceCache.svcMap.Load(serviceName)
        if !ok {
            pkg.NewResponse(c).ToErrorResponse(errcode.ServiceNotFound)
            c.Abort()
            return
        }
        addrs := value.([]string)
        if len(addrs) == 0 {
            pkg.NewResponse(c).ToErrorResponse(errcode.ServiceEmpty)
            c.Abort()
            return
        }
        lb := NewRandom(time.Now().UnixNano())
        addr := lb.Get(addrs)
        c.Set("service_addr", addr)
        c.Next()
    }
}

type ServiceCache struct {
    svcMap  *sync.Map
    onceMap *sync.Map
}

func NewServiceCache() *ServiceCache {
    return &ServiceCache{
        svcMap:  &sync.Map{},
        onceMap: &sync.Map{},
    }
}

func NewServiceDiscoveryClient() *api.Client {
    config := api.DefaultConfig()
    config.Address = pkg.ServiceDiscoveryAddress //consul server
    c, err := api.NewClient(config)
    if err != nil {
        panic(fmt.Sprint("new ServiceDiscovery client error:", err))
    }
    return c
}

func loop(service string, svcMap *sync.Map) {
    var lastIndex uint64
    for {
        services, meta, err := client.Health().Service(service, "", true, &api.QueryOptions{
            WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
        })
        if err != nil {
            fmt.Printf("error retrieving instances from Consul: %v", err)
            return
        }
        lastIndex = meta.LastIndex

        var addrs []string
        for _, service := range services {
            addr := net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))
            addrs = append(addrs, addr)
        }
        svcMap.Store(service, addrs)
    }
}

type LoadBalancer interface {
    Get(addrs []string) string
}

// NewRandom returns a load balancer that selects services randomly.
func NewRandom(seed int64) LoadBalancer {
    return &random{
        r: rand.New(rand.NewSource(seed)),
    }
}

type random struct {
    r *rand.Rand
}

func (r *random) Get(addrs []string) string {
    return addrs[r.r.Intn(len(addrs))]
}
