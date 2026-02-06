# api-go-crud
projeto alex => api em go para efetuar crud em banco de dados mysql


Rotas:



/login:

  Método: POST

  Descrição: Recebe o login e senha do usuário, verifica se estão corretos no banco de dados. Em caso positivo, gera um token jwt.

  Autenticado: Não

 

/usuarios

  Método: POST

  Descrição: Cria um novo usuário

  Autenticado: Sim

 

  Método: GET

  Descrição: Retorna todos os usuários do banco, sem a informação da senha, paginado

  Autenticado: Sim

 

/usuarios/{id}

  Método: GET

  Descrição: Retorna os dados de um usuário específico

  Autenticado: Sim

 

  Método: DELETE

  Descrição: Exclui um usuário específico

  Autenticado: Sim

 

  Método: PUT

  Descrição: Atualiza os dados de um usuário específico

  Autenticado: Sim

 

/servicos

  Método: POST

  Descrição: Cria um novo serviço

  Autenticado: Sim

 

  Método: GET

  Descrição: Retorna todos os serviços do banco, paginado

  Autenticado: Sim

 

/servicos/{id}

  Método: GET

  Descrição: Retorna os dados de um serviço específico

  Autenticado: Sim

 

  Método: DELETE

  Descrição: Exclui um serviço específico

  Autenticado: Sim

 

  Método: PUT

  Descrição: Atualiza os dados de um serviço específico

  Autenticado: Sim

 

/logs/{id}

  Método: GET

  Descrição: Retorna os dados de um log específico

  Autenticado: Sim - Apenas usuários do tipo ADMIN

 

/logs?tipo={tipo}&id-recurso={idrecurso}

  Método: GET

  Descrição: Retorna os dados de um log específico

  Autenticado: Sim - Apenas usuários do tipo ADMIN

  Restrições: Obrigatório informar a data para o filtro. Os query parameters tipo e idrecurso não opcionais. Paginado

 

 

 

1. Estudo e criação do banco de dados. 30/01.

1.1. Criar tabelas

1.2. Inserir usuário ADMIN padrão.

 

2. Cadastro de usuários. 06/02.

2.1. Inserção.

2.2. Deleção.

2.3. Atualização.

2.4. Busca.

 

3. Cadastro de serviços. 13/02.

3.1. Inserção.

3.2. Deleção.

3.3. Atualização.

3.4. Busca.

 

4. Logs. 20/02.

4.1. Inserção de logs na execução de alguma ação.

4.2. Busca de logs.

 

5. Login. 06/03.

5.1. Autenticação.

5.2. Autorização.

 

Opcionais: 

6. Documentação da api com swagger.

7. Subir api na aws.

8. Criar frontend.