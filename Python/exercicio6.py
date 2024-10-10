class Motor:
    def __init__(self, potencia):
        self.potencia = potencia

class Carro:
    def __init__(self, marca, modelo, motor):
        self.marca = marca
        self.modelo = modelo
        self.motor = motor

motor1 = Motor(150)
carro = Carro("Toyota", "Corolla", motor1)
print(f"{carro.marca} {carro.modelo} com motor de {carro.motor.potencia} HP")