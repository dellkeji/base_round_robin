package apis

import (
	"encoding/json"
	"fmt"
	"time"

	"round_robin_with_weight/storage"
	"round_robin_with_weight/utils"

	lock "github.com/bsm/redis-lock"
	"github.com/gin-gonic/gin"
)

// Peer : domain info
type Peer struct {
	Domain          string `json:"domain"`           //地址
	CurrentWeight   int    `json:"current_weight"`   //当前权重
	EffectiveWeight int    `json:"effective_weight"` //有效权重
	Weight          int    `json:"weight"`           //配置权重
}

// Peers : domain list info
type Peers struct {
	Number int     `json:"number"` //服务器数量
	Peer   []*Peer `json:"peer"`   //服务器节点数组
}

//PeerData : domain list
type PeerData struct {
	Peers *Peers `json:"peer"` //服务器池数据
}

// GetAPIDomain : get gwateway api's domain
func GetAPIDomain(c *gin.Context) {
	withoutSmooth := c.DefaultQuery("without_smooth", "0")
	withLock := c.DefaultQuery("with_lock", "0")
	apigwID := c.Query("apigw_id")
	if apigwID == "" {
		utils.CommonResponse(c, 400, "Gateway's id is not null", gin.H{})
		return
	}
	// 限制支持平滑的方式
	if withoutSmooth != "0" {
		ok := domainWithWeight(c)
		if !ok {
			return
		}
	}
	// 获取上次调整后的权重值
	domain, ok := domainWithSmoothWeight(c, apigwID, withLock)
	if !ok {
		return
	}
	utils.CommonResponse(c, 0, "success", gin.H{
		"domain": domain,
	})
}

func domainWithSmoothWeight(c *gin.Context, apigwID string, withLock string) (domain string, ok bool) {
	// get redis info
	redisClient := storage.GetDefaultRedisSession()
	if withLock != "0" {
		lockKey := fmt.Sprintf("lock:foo:%s", apigwID)
		// retry one time
		locker, err := lock.Obtain(redisClient.Client, lockKey, &lock.Options{
			RetryCount:  1,
			RetryDelay:  25 * time.Millisecond,
			LockTimeout: 2 * time.Second,
		})
		if err != nil {
			message := fmt.Sprintf("Obtain lock failed, detail : %v", err)
			utils.CommonResponse(c, 400, message, gin.H{})
			return "", false
		} else if locker == nil {
			message := fmt.Sprintf("Obtain lock failed, detail : lock is null")
			utils.CommonResponse(c, 400, message, gin.H{})
			return "", false
		}
		defer locker.Unlock()
	}

	redisField := fmt.Sprintf("apigw_domain_info:%s", apigwID)
	result, ok := redisClient.HGetRedisKey(c, "apigw_domain", redisField)
	if !ok {
		return "", ok
	}
	// 解析数据
	var peers PeerData
	json.Unmarshal([]byte(result), &peers)
	// 获取下一个权重的domain
	domainInfo := getPeer(&peers)
	// 存储变更后的值
	peerStr, err := json.Marshal(peers)
	if err != nil {
		message := fmt.Sprintf("Json serialize error, detail: %v", err)
		utils.CommonResponse(c, 400, message, gin.H{})
		return "", false
	}
	ok = redisClient.HSetRedisKey(c, "apigw_domain", redisField, string(peerStr))
	if !ok {
		return "", false
	}

	return domainInfo.Domain, true
}

func domainWithWeight(c *gin.Context) bool {
	utils.CommonResponse(c, 400, "Only smooth, other not support", gin.H{})
	return false
}

func getPeer(rrp *PeerData) *Peer {
	var best *Peer
	total := 0

	for i := 0; i < rrp.Peers.Number; i++ {
		peer := rrp.Peers.Peer[i]
		//将当前权重与有效权重相加
		peer.CurrentWeight += peer.EffectiveWeight
		//累加总权重
		total += peer.EffectiveWeight

		if best == nil || peer.CurrentWeight > best.CurrentWeight {
			best = peer
		}
	}

	if best == nil {
		return nil
	}
	//将当前权重改为当前权重-总权重
	best.CurrentWeight -= total
	return best
}
