class ContaBancaria:
    def __init__(self, titular):
        self._saldo = 0
        self.titular = titular

    def depositar(self, valor):
        self._saldo += valor

    def sacar(self, valor):
        if valor > self._saldo:
            raise Exception("Saldo insuficiente")
        self._saldo -= valor

    def obter_saldo(self):
        return self._saldo

conta = ContaBancaria("João")
conta.depositar(100)
print(conta.obter_saldo())
conta.sacar(50)
print(conta.obter_saldo())