class SaldoInsuficienteException(Exception):
    pass

class ContaBancaria:
    def __init__(self):
        self._saldo = 0

    def depositar(self, valor):
        self._saldo += valor

    def sacar(self, valor):
        if valor > self._saldo:
            raise SaldoInsuficienteException("Saldo insuficiente")
        self._saldo -= valor

conta = ContaBancaria()
conta.depositar(100)
try:
    conta.sacar(150)
except SaldoInsuficienteException as e:
    print(e)
