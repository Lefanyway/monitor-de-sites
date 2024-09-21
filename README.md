# Monitorador de sites
Este é um programa simples em Go que monitora o status de sites através de requisições HTTP. 

---
# Como instalar o Golang?
Windows:
Acesse o site oficial do Golang: https://golang.org/dl/.
Baixe o instalador para Windows.
Execute o instalador e siga as instruções.
Após a instalação, abra o "Prompt de Comando" e digite go version para verificar se o Go foi instalado corretamente.
Linux (Debian/Ubuntu):
Abra o terminal.

Execute os seguintes comandos para baixar e instalar o Go:

bash
Copiar código
sudo apt update
sudo apt install -y golang
Após a instalação, verifique a versão instalada com o comando:

bash
Copiar código
go version
Pronto! Agora que o Go está instalado, você pode testar o programa.

# Como usar?

No terminal, rode os seguintes comandos:
```

 git clone https://github.com/Lefanyway/monitorador-de-sites

 go run hello.go

```
No arquivo `sites.txt` deve conter a lista de sites que serão monitorados. Cada site deve estar em uma linha separada no arquivo.

---

Ao iniciar o programa, você verá um menu com as opções disponíveis:

 **1 - Iniciar Monitoramento: Inicia o monitoramento dos sites listados no arquivo sites.txt.**

 **2 - Exibir Logs: Exibe os logs salvos em um arquivo chamado log.txt, mostrando os resultados das requisições de cada site.**

 **0 - Sair do Programa: Encerra a execução do programa.**

 ---

Durante o monitoramento, o programa enviará requisições HTTP para cada site da lista a cada intervalo. Se um site retornar um status de resposta 200, significa que o site foi carregado com sucesso. Caso contrário, será considerado que o site está com problemas.

Os resultados do monitoramento são registrados em um arquivo chamado `log.txt`. Cada registro de log contém a data e hora, o site testado, o status (online ou offline) e o status code da resposta HTTP.

---
# Observações
Certifique-se de ter instalado o Go em sua máquina antes de executar o programa.

Certifique-se de ter criado o arquivo sites.txt com a lista de sites que deseja monitorar.

Os logs são salvos em um arquivo chamado log.txt. Certifique-se de que o programa tenha permissão para criar e escrever nesse arquivo.

Para interromper a execução do programa, você pode selecionar a opção "0 - Sair do Programa" no menu.

> Este programa foi criado através do curso da Alura.
