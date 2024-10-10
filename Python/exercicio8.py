class Empregado:
    def __init__(self, nome, cargo, salario):
        self.nome = nome
        self.cargo = cargo
        self.salario = salario

class Empresa:
    def __init__(self):
        self.empregados = []

    def adicionar_empregado(self, empregado):
        self.empregados.append(empregado)

empresa = Empresa()
empresa.adicionar_empregado(Empregado("João", "Gerente", 5000))
empresa.adicionar_empregado(Empregado("Ana", "Desenvolvedora", 4000))

for empregado in empresa.empregados:
    print(f"{empregado.nome}, {empregado.cargo}, R${empregado.salario}")