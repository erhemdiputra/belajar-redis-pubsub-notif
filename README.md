### Belajar redis pub sub notification

### How to Run
1. Install redis
2. Run command `redis-cli CONFIG SET notify-keyspace-events KEA`
3. Run command `go build && ./belajar-redis-pubsub-notif`
4. Output
```
2022/04/04 06:59:25 c.client.Receive() got redis.Subscription: channel=__keyspace*:test_key_*, kind=psubscribe, count=1 - InitSubscriber
2022/04/04 06:59:25 c.client.Receive() got redis.Subscription: channel=__keyspace*:dummy_key_*, kind=psubscribe, count=2 - InitSubscriber
2022/04/04 06:59:25 testCallback: pattern=__keyspace*:test_key_*, channel=__keyspace@0__:test_key_123, message=set
2022/04/04 06:59:25 testCallback: pattern=__keyspace*:test_key_*, channel=__keyspace@0__:test_key_123, message=expire
2022/04/04 06:59:25 testCallback: pattern=__keyspace*:dummy_key_*, channel=__keyspace@0__:dummy_key_123, message=set
2022/04/04 06:59:25 testCallback: pattern=__keyspace*:dummy_key_*, channel=__keyspace@0__:dummy_key_123, message=expire
2022/04/04 06:59:35 testCallback: pattern=__keyspace*:dummy_key_*, channel=__keyspace@0__:dummy_key_123, message=expired
2022/04/04 06:59:35 testCallback: pattern=__keyspace*:test_key_*, channel=__keyspace@0__:test_key_123, message=expired
```

### Source
- https://blog.karatos.in/a?ID=01050-e1be2cad-5425-43ad-96b8-836c3d8b7697
- https://irshadhasmat.medium.com/golang-simple-pub-sub-implementation-using-redis-b57070476d45
- https://gist.github.com/irshadhasmat/abdd99a0126c1971a6bbe3b68a9f37ef
- https://pkg.go.dev/github.com/gomodule/redigo/redis#hdr-Publish_and_Subscribe
- https://medium.com/easyread/build-a-real-time-app-with-redis-pub-sub-a99e46d000d5