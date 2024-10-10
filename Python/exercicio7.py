class Escola:
    def __init__(self):
        self.professores = []

    def adicionar_professor(self, professor):
        self.professores.append(professor)

class Professor:
    def __init__(self, nome):
        self.nome = nome

escola = Escola()
prof1 = Professor("Maria")
prof2 = Professor("José")

escola.adicionar_professor(prof1)
escola.adicionar_professor(prof2)

for professor in escola.professores:
    print(professor.nome)