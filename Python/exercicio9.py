from abc import ABC, abstractmethod

class Imprimivel(ABC):
    @abstractmethod
    def imprimir(self):
        pass

class Relatorio(Imprimivel):
    def imprimir(self):
        print("Imprimindo relatório")

class Contrato(Imprimivel):
    def imprimir(self):
        print("Imprimindo contrato")

itens = [Relatorio(), Contrato()]
for item in itens:
    item.imprimir()