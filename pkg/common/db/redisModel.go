package db

import (
	"github.com/garyburd/redigo/redis"
	"open-im/pkg/common/log"
)

const (
	uidPidToken = "UID_PID_TOKEN_STATUS:"
)

func (d *DataBases) Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := d.redisPool.Get()
	if err := con.Err(); err != nil {
		log.Error("", "", "redis cmd = %v, err = %v", cmd, err)
		return nil, err
	}
	defer con.Close()

	params := make([]interface{}, 0)
	params = append(params, key)

	if len(args) > 0 {
		for _, v := range args {
			params = append(params, v)
		}
	}

	return con.Do(cmd, params...)
}

func (d *DataBases) GetTokenMapByUidPid(userID, platformID string) (map[string]int, error) {
	key := uidPidToken + userID + ":" + platformID
	log.NewDebug("", "get token key is ", key)
	return redis.IntMap(d.Exec("HGETALL", key))
}
