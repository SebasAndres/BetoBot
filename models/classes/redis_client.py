import redis
from classes.match import Match

class RedisClient:
    def __init__(self):
        self.client = redis.Redis(
            host='redis-16247.c308.sa-east-1-1.ec2.cloud.redislabs.com',
            port=16247,
            username="default", 
            password="flFkN850wrtNG2VyM9aRiwtd3XcrsXDc", 
        )

    def getKLastMatches(self, league, K=-1): # O(K log N)
        # Devuelve un iterador con los ultimos K partidos
        # de la liga
        if K==-1:
            K = self.client.get(f"{league}_num_played") # leer_num_partidos_registrados_en_liga

        i = 0
        it = redis.matches(league)
        while i<K and it.has_next():
            yield Match(it.next())            
            i+=1

    def get_pz_league(self, league):
        return self.client.get(f"{league}_pZ")

    """
    def getXE1andXE2(self, e1, e2):
        # return redis->league(l)->TableX[e1][e2]

    def getXE(self, E):
        # return redis.get(E)....

    def getQE(self, Q):
    """
        