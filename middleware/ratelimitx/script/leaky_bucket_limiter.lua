-- tokens_key 令牌key
local tokens_key = KEYS[1]

-- timestamp_key 时间戳key
local timestamp_key = KEYS[2]

-- rate 每秒通过请求数量
local rate = tonumber(ARGV[1])

-- capacity 桶最大容量
local capacity = tonumber(ARGV[2])

-- now 当前时间戳
local now = tonumber(ARGV[3])

-- requested 请求的令牌数量
local requested = tonumber(ARGV[4])

-- key_lifetime 设置 Redis 中的过期时间为令牌桶满时加上一个单位时间（通常为 1 秒）。
local key_lifetime = math.ceil((capacity / rate) + 1)

-- 获取上次请求的令牌数量，默认值为 0。
local key_bucket_count = tonumber(redis.call("GET", tokens_key)) or 0

-- 获取上次请求的时间戳，默认值为 now。
local last_time = tonumber(redis.call("GET", timestamp_key)) or now

-- millis_since_last_leak：当前时间与上次请求时间的差值。
local millis_since_last_leak = now - last_time

-- leaks：计算泄漏量。
local leaks = millis_since_last_leak * rate

-- 如果泄漏量大于等于当前令牌数量，则令牌数量清零；否则减去泄漏量，并更新时间戳。
if leaks > 0 then
    if leaks >= key_bucket_count then
        key_bucket_count = 0
    else
        key_bucket_count = key_bucket_count - leaks
    end
    last_time = now
end

local is_allow = 0

-- new_bucket_count：计算新的令牌数量。
local new_bucket_count = key_bucket_count + requested

-- 如果新的令牌数量不超过最大容量，则请求允许，否则返回不允许并退出脚本。
if new_bucket_count <= capacity then
    is_allow = 1
else
    return {is_allow, 0}
end

-- 更新令牌数量和时间戳。
redis.call("SETEX", tokens_key, key_lifetime, new_bucket_count)
redis.call("SETEX", timestamp_key, key_lifetime, now)

return {is_allow, capacity-new_bucket_count}