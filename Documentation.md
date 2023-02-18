Documentação A Simple-ChatBot

Este é um programa escrito em Go (golang) que lê dois arquivos JSON e permite ao usuário visualizar o menu de um restaurante e fazer um pedido. 
Aqui está uma breve descrição do que cada parte do código faz:

* Item 1

    O pacote principal main é importado.

    As estruturas de dados Menu1 e Escolhas são definidas. 

    Menu1 contém três campos: Entradas, Principais e Sobremesas. 

    Escolhas contém uma lista de opções para cada prato do menu.

    Cada lista é representada por um array de strings.

* Item 2

    A função main é definida. 
    Ela começa abrindo o arquivo chatbot-responses.json e decodificando-o em um valor do tipo Menu1. 

    Em seguida, abre o arquivo chatbot-responses1.json e decodifica-o em um valor do tipo Escolhas. 

    Esses arquivos JSON devem estar presentes no diretório em que o programa é executado.

* Item 3

    Depois de carregar as informações do menu e das opções de escolha, o programa começa a interagir com o usuário. 

    Ele pede ao usuário que digite "next" para avançar para as opções do menu. 

    Se o usuário digitar qualquer outra coisa, o programa exibe uma mensagem de erro e pede novamente para que o usuário digite "next".

* Item 4

    Em seguida, o programa exibe o menu. 

    Ele pede ao usuário que digite "c" para ver as opções do menu. 

    Se o usuário digitar qualquer outra coisa, o programa exibe uma mensagem de erro e pede novamente para que o usuário digite "c".

* Item 5

    Depois que o usuário visualiza o menu, o programa pede que ele escolha um prato digitando o número correspondente. 

    O programa exibe uma mensagem de erro se o usuário digitar um número inválido ou qualquer outra coisa que não seja um número.

* Item 6

    Depois que o usuário faz um pedido, o programa exibe uma mensagem confirmando o pedido. 

    O programa não tem nenhuma funcionalidade real de pedidos ou de comunicação com um restaurante. 

    Ele é apenas um exemplo simples de como ler e decodificar arquivos JSON em Go.
