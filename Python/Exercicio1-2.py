#1. Classes e Objetos Básicos Crie uma classe Carro com atributos como marca, modelo, e
#   ano. Instancie três objetos dessa classe e exiba os detalhes de cada um.

#2. Métodos Adicione um método acelerar e frear à classe Carro que altere um atributo
#   velocidade. Crie um método para exibir a velocidade atual

class Carro:
    def __init__(self,  marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade = 0
    
    def mostrar_informacoes(self):
        print(f"Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}")
    
    def acelerar(self, incremento):
        self.velocidade += incremento
        print(f"\nO carro {self.modelo} acelerou para {self.velocidade} km/h")
    
    def frear(self, decremento):
        self.velocidade -= decremento
        if(self.velocidade > 0 ):
            print(f"\nO carro {self.modelo} freou para {self.velocidade} km/h")
        elif(self.velocidade < decremento):
            print(f"\nO carro {self.modelo} parou")
        else:
            print(f"\nO carro já está parado")

    def velocidade_atual(self):
        print(f"Velocidade atual do carro {self.modelo}: {self.velocidade} km/h ")

def main():
    carro1 = Carro("Renault", "Kwid", 2022)
    carro2 = Carro("Fiat", "Mobi", 2020)
    carro3 = Carro("Citroen", "C3", 2024)

    carro1.mostrar_informacoes()
    carro2.mostrar_informacoes()
    carro3.mostrar_informacoes()
   
    carro1.acelerar(45)
    carro1.velocidade_atual()
    carro1.frear(45)
    carro1.velocidade_atual()
    carro1.frear(10)

if __name__ == '__main__':
    main()