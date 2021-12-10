# Exercício #1: Quiz Game

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/quiz)

## Detalhes do exercício

Este exercício é dividido em duas partes para ajudar a simplificar o processo de explicá-lo e também para torná-lo mais fácil de resolver. A segunda parte é mais difícil do que a primeira, portanto, se você ficar preso, fique à vontade para passar para outro problema e volte para a parte 2 mais tarde.

### Parte 1

Crie um programa que irá ler um questionário fornecido por meio de um arquivo CSV (mais detalhes abaixo) e, em seguida, entregará o questionário a um usuário para controlar quantas perguntas acertou e quantas errou. Independentemente de saber se a resposta está correta ou errada, a próxima pergunta deve ser feita imediatamente depois.

O arquivo CSV deve ser padronizado como `problems.csv` (exemplo mostrado abaixo), mas o usuário deve ser capaz de personalizar o nome do arquivo por meio de um sinalizador.

O arquivo CSV estará em um formato como o abaixo, onde a primeira coluna é uma pergunta e a segunda coluna na mesma linha é a resposta para essa pergunta.

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

Você pode presumir que os questionários serão relativamente curtos (<100 perguntas) e terão respostas com uma única palavra / número.

No final do questionário, o programa deve produzir o número total de perguntas corretas e quantas perguntas havia no total. Perguntas com respostas inválidas são consideradas incorretas.

**Nota:** *Os arquivos CSV podem ter perguntas com vírgulas. Por exemplo: `" what 2 + 2, sir? ", 4` é uma linha válida em um CSV. Eu sugiro que você dê uma olhada no pacote CSV em Go*

### Parte 2

Adapte seu programa da parte 1 adicionando um cronômetro. O limite de tempo padrão deve ser 30 segundos, mas também deve ser personalizável por meio de um sinalizador.


Seu questionário deve parar assim que o tempo limite for excedido. Ou seja, você não deve esperar que o usuário responda uma pergunta final, mas o ideal é interromper totalmente o questionário, mesmo se estiver aguardando uma resposta do usuário final.

Os usuários devem ser solicitados a pressionar enter (ou alguma outra tecla) antes de o cronômetro iniciar e, em seguida, as perguntas devem ser impressas na tela, uma de cada vez, até que o usuário forneça uma resposta. Independentemente de saber se a resposta está correta ou errada, a próxima pergunta deve ser feita.

No final do questionário, o programa ainda deve produzir o número total de perguntas corretas e quantas perguntas havia no total. Perguntas com respostas inválidas ou sem resposta são consideradas incorretas.

