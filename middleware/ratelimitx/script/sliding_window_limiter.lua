-- tokens_key 令牌key
local tokens_key = KEYS[1]

-- timestamp_key 时间戳key
local timestamp_key = KEYS[2]

-- rate 滑动窗口流量阈值，每秒请求数量
local rate = tonumber(ARGV[1])

-- capacity 限流窗口总请求数量
local capacity = tonumber(ARGV[2])

-- now 当前时间戳
local now = tonumber(ARGV[3])

-- window_size 窗口大小
local window_size = tonumber(capacity / rate)

-- window_time 窗口时间间隔
local window_time = 1

-- last_requested 之前的总请求数量
local last_requested = 0
local exists_key = redis.call('EXISTS', tokens_key)
if (exists_key == 1) then
    -- 清理过期数据
    redis.call('ZREMRANGEBYSCORE', tokens_key, 0, now - window_size / window_time)
    -- 获取之前总请求数量
    last_requested = redis.call('ZCARD', tokens_key)
end

-- remain_request 获取剩余请求数量
local remain_request = capacity - last_requested

-- 如果之前总请求数量小于限流窗口总请求数量，则允许通过，否则不允许通过
-- 允许通过->allowed_num=1, 不允许通过->allowed_num=0
local allowed_num = 0
if (last_requested < capacity) then
    allowed_num = 1
    -- 添加到滑动窗口
    redis.call('ZADD', tokens_key, now, timestamp_key)
    -- 设置过期时间
    redis.call('EXPIRE', tokens_key, window_size)
end

-- 返回是否允许通过和剩余请求数量
return { allowed_num, remain_request }
