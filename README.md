# 🚀 API de Gestão de Usuários e Serviços

API desenvolvida em **Go (Golang)** utilizando o framework **Fiber** para operações de CRUD em banco de dados **MySQL**.

## 📌 Rotas da API

### 🔐 Autenticação

| Rota | Método | Descrição | Autenticado |
| :--- | :--- | :--- | :---: |
| `/login` | `POST` | Valida credenciais e gera token JWT | Não |

### 👥 Usuários

| Rota | Método | Descrição | Autenticado |
| :--- | :--- | :--- | :---: |
| `/usuarios` | `POST` | Cria um novo usuário | Sim |
| `/usuarios` | `GET` | Retorna todos os usuários (sem senha, paginado) | Sim |
| `/usuarios/{id}` | `GET` | Retorna dados de um usuário específico | Sim |
| `/usuarios/{id}` | `PUT` | Atualiza dados de um usuário específico | Sim |
| `/usuarios/{id}` | `DELETE` | Exclui um usuário específico | Sim |

### 🛠️ Serviços

| Rota | Método | Descrição | Autenticado |
| :--- | :--- | :--- | :---: |
| `/servicos` | `POST` | Cria um novo serviço | Sim |
| `/servicos` | `GET` | Retorna todos os serviços (paginado) | Sim |
| `/servicos/{id}` | `GET` | Retorna dados de um serviço específico | Sim |
| `/servicos/{id}` | `PUT` | Atualiza dados de um serviço específico | Sim |
| `/servicos/{id}` | `DELETE` | Exclui um serviço específico | Sim |

### 📜 Logs (Apenas ADMIN)

| Rota | Método | Descrição |
| :--- | :--- | :--- |
| `/logs/{id}` | `GET` | Retorna dados de um log específico |
| `/logs?data={d}&tipo={t}&id-recurso={ir}` | `GET` | Filtro de logs (Data obrigatória, paginado) |

---

## 📅 Cronograma de Desenvolvimento


| Fase | Entrega | Tarefas |
| :--- | :---: | :--- |
| **1. Banco de Dados** | 30/01 | Criar tabelas e inserir usuário ADMIN padrão. |
| **2. Cadastro de Usuários** | 06/02 | CRUD completo de usuários. |
| **3. Cadastro de Serviços** | 13/02 | CRUD completo de serviços. |
| **4. Logs** | 20/02 | Implementação de registros de ações e busca. |
| **5. Autenticação** | 06/03 | Login com JWT e controle de permissões. |

## 🚀 Extras e Deploy
- [ ] Documentação com **Swagger**
- [ ] Deploy na **AWS**
- [ ] Desenvolvimento do **Frontend**

## 🛠️ Tecnologias
- **Linguagem:** Go 1.x
- **Framework:** Fiber
- **Banco de Dados:** MySQL
- **Autenticação:** JWT (JSON Web Token)
