## Modelo de "Ambos Convierten Gol"

import tensorflow as tf
from keras.utils import plot_model

class f_ambos_gol:
    def __init__(self, load_last): # O(1)
        if load_last:
            self.model = {} # load model saved
        else: 
            self.model = tf.keras.Sequential([
                tf.keras.Dense(3, activation='relu'),
                tf.keras.Dense(5, activation='relu'),
                tf.keras.Dense(3, activation='relu'),
                tf.keras.Dense(1, activation='sigmoid'),
            ])
            self.model.compile(loss="binary_crossentropy", 
                            optimizer="sgd",
                            metrics=["accuracy"])

    def build_dataset(self, matches, redis, league) -> tuple: # O(N)  
        # Arma un dataset para el entrenamiento en base a 
        # los partidos de la liga y el formato de input del modelo 

        X = [] # input
        Y = [] # output
        pz_l = redis.get_pz_league(league)  

        for match in matches:
            # Validar si pasa mucho que las probabilidades conjuntas sean 0
            # en tal caso, agregar como parametros XE1, XE2, QE1, QE2
            # Testear que es mejor!!!
            f_input = (
                redis.getXE(match.team_1),
                redis.getXE(match.team_2),
                redis.getXE1andXE2(match.team_1, match.team_2),
                redis.getQE1andQE2(match.team_1, match.team_2),
                redis.getQE(match.team_1),
                redis.getQE(match.team_2),
                pz_l
            )
            X.append(f_input)
            Y.append(match.ambos_gol())

        return X, Y
    
    def fit(self, X, Y):
        # Entrenamos el modelo cargado con los datos pasados
        self.model.fit(X,Y)

    def save(self):
        # Guardamos el modelo entrenado
        return

    def print_model(self):
        return plot_model(self.model, show_shapes=True)
