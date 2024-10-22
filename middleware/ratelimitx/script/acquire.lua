-- get rate
local rate = ARGV[1];
-- get interval
local interval = ARGV[2];
-- get permits number
local permits = tonumber(ARGV[3])
-- get permits name
local permits_name = KEYS[1];
-- get remaining_permits
local remaining_permits = redis.call('GET', permits_name);

local is_allow = 1;

-- if permits_name exists
if remaining_permits ~= false then
    -- if remaining_permits less than permits
    if tonumber(remaining_permits) < permits then
        is_allow = 0;
        return {is_allow, 0}
    else
        is_allow = 1;
        return {is_allow, redis.call('DECRBY', permits_name, ARGV[3])}
    end;
-- if permits_name not exists
else
    is_allow = 1;
    redis.call('SET', permits_name, rate, 'PX', interval);
    return {is_allow, redis.call('DECRBY', permits_name, ARGV[3])}
end;