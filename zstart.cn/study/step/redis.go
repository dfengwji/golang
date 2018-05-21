// redis
package step

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 导入redigo扩展包

func StudyRedis() {
	testRedis2()
}

func testRedis1() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer conn.Close()
	v, err := conn.Do("SET", "name", "red")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	v, err = redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}

func testRedis2() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer conn.Close()
	conn.Do("lpush", "redList", "a")
	conn.Do("lpush", "redList", "b")
	conn.Do("lpush", "redList", "c")
	conn.Do("lpush", "redList", "d")
	values, _ := redis.Values(conn.Do("lrange", "redlist", 0, 100))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}

	var v1 string
	redis.Scan(values, &v1)
	fmt.Println(v1)
}

//func test1() {
//	for i, member := range []string{"red", "blue", "green"} {
//		conn.Send("ZADD", "zset", i, member)
//	}
//	if _, err := conn.Do(""); err != nil {
//		fmt.Println(err)
//		return
//	}
//	v, err := zpop("zset")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(v)

//		v, err = redis.String(zpopScript.Do(conn, "zset"))
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(v)
//}

//func test2() {
//	go subscribe()
//	go subscribe()
//	go subscribe()

//	for {
//		var s string
//		fmt.Scanln(&s)
//		_, err := conn.Do("PUBLISH", "redChatRoom", s)
//		if err != nil {
//			fmt.Println("publish err:", err)
//			return
//		}
//	}
//}

func subscribe() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer c.Close()
	psc := redis.PubSubConn{c}
	psc.Subscribe("redChatRoom")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s:message:%s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s:%s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

//var zpopScript = redis.NewScript(1,'
//	local r=redis.call('ZRANGE',KEYS[1],0,0)
//	if r ~=nil then
//		r = r[1]
//		redis.call('ZREM',KEYs[1],r)
//	end
//	return r
//')

func zpop(conn redis.Conn, key string) (result string, err error) {
	defer func() {
		if err != nil {
			conn.Do("DISCARD")
		}
	}()

	for {
		if _, err := conn.Do("WATCH", key); err != nil {
			return "", err
		}
		members, err := redis.Strings(conn.Do("ZRANGE", key, 0, 0))
		if err != nil {
			return "", err
		}
		if len(members) < 1 {
			return "", redis.ErrNil
		}

		conn.Send("MULTI")
		conn.Send("ZREM", key, members[0])
		queued, err := conn.Do("EXEC")
		if err != nil {
			return "", err
		}
		if queued != nil {
			result = members[0]
			break
		}

	}
	return result, nil
}
