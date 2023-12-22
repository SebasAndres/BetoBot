import redis

redis_client = redis.Redis(
    host='redis-16247.c308.sa-east-1-1.ec2.cloud.redislabs.com',
    port=16247,
    username="default", # use your Redis user. More info https://redis.io/docs/management/security/acl/
    password="flFkN850wrtNG2VyM9aRiwtd3XcrsXDc", # use your Redis password
)

redis_client.set("variable_prueba", "hola")

print(redis_client.get("variable_prueba"))