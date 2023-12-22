import redis

rdb = redis.Redis(
    host='redis-16247.c308.sa-east-1-1.ec2.cloud.redislabs.com',
    port=16247,
    password='flFkN850wrtNG2VyM9aRiwtd3XcrsXDc',
    db=0
)

teams = rdb.smembers('FB_TEAMS')
for team in teams:
    i = rdb.delete("FB_TEAM_"+team.decode('utf-8').upper()) 
    #print(i)