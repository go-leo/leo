-- tokens_key 令牌key
local tokens_key = KEYS[1]

-- timestamp_key 时间戳key
local timestamp_key = KEYS[2]

-- rate 每秒生成令牌的数量
local rate = tonumber(ARGV[1])

-- capacity 令牌桶容量
local capacity = tonumber(ARGV[2])

-- now 当前时间戳
local now = tonumber(ARGV[3])

-- requested 请求的令牌数量
local requested = tonumber(ARGV[4])

-- 计算令牌桶桶满的时间
local fill_time = capacity/rate

--设置redis中的过期时间为令牌桶满时的两倍，
local ttl = math.floor(fill_time*2)

-- last_tokens 获取上次请求的令牌数量，默认值为 capacity。
local last_tokens = tonumber(redis.call("GET", tokens_key))
if last_tokens == nil then
    last_tokens = capacity
end

-- last_refreshed 获取上次请求的时间戳，默认值为 0。
local last_refreshed = tonumber(redis.call("GET", timestamp_key))
if last_refreshed == nil then
    last_refreshed = 0
end

-- 计算当前时间与上次请求时间的差值
local delta = math.max(0, now-last_refreshed)

-- 计算当前令牌桶中的令牌数量
local filled_tokens = math.min(capacity, last_tokens+(delta*rate))

-- 判断是否有足够的令牌满足请求
local allowed = filled_tokens >= requested

-- 更新令牌桶中的令牌数量
local new_tokens = filled_tokens
local allowed_num = 0
if allowed then
    new_tokens = filled_tokens - requested
    allowed_num = 1
end

-- 更新Redis中的数据
redis.call("SETEX", tokens_key, ttl, new_tokens)
redis.call("SETEX", timestamp_key, ttl, now)

return { allowed_num, new_tokens }
