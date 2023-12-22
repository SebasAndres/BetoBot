## Scripts para el entrenamiento de los modelos estadisticos

from argparse import ArgumentParser

from classes.redis_client import RedisClient
from classes.f_ambos_gol import f_ambos_gol

if __name__ == '__main__':
    args = ArgumentParser()
    leagues = args.leagues.split()

    redis = RedisClient() # cargar de un entorno

    for league in leagues:

        # [0] Optional / Invoke worker-data
        # downloadFixture(league)        

        # [1] Modelo "Ambos Equipos hacen Gol"
        ambos_equipos_gol = f_ambos_gol(load_last=True)
        last_k_matches = redis.get_last_k_matches(league, K=30)
        X, Y = ambos_equipos_gol.build_dataset(last_k_matches)     
        metrics = ambos_equipos_gol.fit_n_evaluate(X,Y)
        ambos_equipos_gol.save(filename=f"model_blabla_date_{metrics}")

        # [2] Modelo 
        # "Tarjetas amarillas por partido"
        pass    

