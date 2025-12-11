# go-api-sample-project
a simple api project using go


## Maracutaias divertidas de se ver
Dado um ambiente onde não é possível executar o docker, uma alternativa curiosa, e que funcionou nesse projeto, é:

1. Abrir o repo no GitHub Codespaces
2. No canto superior direito, no menu, escolher a opção **"Abrir no VS Code Desktop"**
3. Baixar no seu VS Code local, a extensão [GitHub Codespaces](https://marketplace.visualstudio.com/items?itemName=GitHub.codespaces)
4. Realizar a autenticação com o seu GitHub
    - Pode demorar um pouco o ter alguns problemas no meio do caminho, mas apenas tentar novamente a opção do 2
5. Ao realizar a autenticação e conexão com o codespaces, rode o docker como faria normalmente pelo terminal
6. No terminal, na aba **Portas**, verifique se está tudo rodando, localmente e a porta (o comando ``docker container ls`` também ajuda nessa validação)

- ps: para todos os efeitos, seus containers estarão rodando localmente na sua máquina
    - Nesta etapa ocorreram dificuldades de conexão com o banco de dados. Então tive a ideia de que poderia ser por conta do bd local já estar rodando na tradicional porta 5432. Sendo assim, no meu docker-compose mudei a porta de saída para 5500
