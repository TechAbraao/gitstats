# GitStats

> GitHub Repository Analytics CLI

**GitStats** é uma ferramenta de linha de comando (CLI), desenvolvida em Go, para análise, consulta e visualização de informações de repositórios hospedados no GitHub.

A aplicação permite que desenvolvedores obtenham informações sobre seus projetos diretamente pelo terminal, incluindo estatísticas de commits, contribuidores, linguagens de programação, atividade e comparação entre repositórios.

---

# 1. Objetivo do projeto

Desenvolver uma CLI capaz de consumir a GitHub REST API, processar os dados retornados e apresentar informações relevantes sobre repositórios de maneira organizada, intuitiva e acessível.

O projeto tem como principais objetivos:

- Facilitar a análise de repositórios GitHub pelo terminal.
- Apresentar métricas sobre desenvolvimento e colaboração.
- Identificar padrões de atividade e contribuição.
- Permitir a comparação entre diferentes repositórios.
- Gerar relatórios em diferentes formatos.
- Oferecer uma ferramenta que possa ser utilizada tanto por desenvolvedores individuais quanto por equipes.

# 2. Tecnologias

| Tecnologia | Finalidade |
|---|---|
| Go (Golang) | Linguagem principal |
| Cobra | Gerenciamento de comandos e subcomandos da CLI |
| encoding/json | Serialização e desserialização JSON |
| Lip Gloss | Estilização e formatação do terminal |
| Viper | Gerenciamento de configurações |
| SQLite | Persistência local e histórico de consultas |
| Testify | Testes automatizados |
| GoReleaser | Compilação e distribuição de binários |
| GitHub Actions | Automação de testes e CI/CD |

**Fonte de dados:** GitHub REST API.

Documentação oficial: https://docs.github.com/en/rest

# 3. Funcionalidades

## 3.1. Repository Overview

Apresenta uma visão geral de um repositório GitHub.

**Comando:**

```bash
gitstats repo <owner>/<repository>
```

**Exemplo:**

```bash
gitstats repo TechAbraao/jobs
```

**Informações apresentadas:**

- Nome e descrição do repositório.
- Proprietário.
- Visibilidade (público ou privado).
- Data de criação.
- Última atualização.
- Branch padrão.
- Linguagem principal.
- Quantidade de stars.
- Quantidade de forks.
- Quantidade de watchers.
- Quantidade de issues abertas.
- URL do repositório.

**Opções:**

```bash
gitstats repo TechAbraao/jobs --json
```

Permite exportar as informações em JSON.

## 3.2. Commit Analysis

Analisa o histórico de commits de um repositório.

**Comando:**

```bash
gitstats commits <owner>/<repository>
```

**Exemplos:**

```bash
gitstats commits TechAbraao/jobs

gitstats commits TechAbraao/jobs --limit 20

gitstats commits TechAbraao/jobs --author TechAbraao

gitstats commits TechAbraao/jobs --since 2026-01-01

gitstats commits TechAbraao/jobs --until 2026-12-31
```

**Informações apresentadas:**

- Quantidade de commits consultados.
- Autor de cada commit.
- Hash (SHA) do commit.
- Mensagem do commit.
- Data e horário.
- Autores mais ativos.
- Distribuição de commits por período.
- Primeiro e último commit do intervalo consultado.

**Análises adicionais:**

- Média de commits por dia.
- Média de commits por semana.
- Distribuição de commits por mês.
- Dias com maior atividade.

**Observação:** as métricas de commits devem considerar o período e os filtros utilizados. A quantidade de resultados retornados não deve ser confundida com o total histórico do repositório.

## 3.3. Contributors Analysis

Identifica e apresenta os principais contribuidores de um repositório.

**Comando:**

```bash
gitstats contributors <owner>/<repository>
```

**Exemplos:**

```bash
gitstats contributors TechAbraao/jobs

gitstats contributors TechAbraao/jobs --limit 10
```

**Informações apresentadas:**

- Nome de usuário.
- Avatar.
- Perfil do GitHub.
- Quantidade de contribuições registradas.
- Posição no ranking de contribuidores.

**Análises adicionais:**

- Ranking dos contribuidores mais ativos.
- Percentual de contribuições por usuário.
- Quantidade total de contribuidores retornados pela API.

A análise deve considerar as limitações do endpoint de contribuidores e a possibilidade de resultados anonimizados.

## 3.4. Languages Analysis

Identifica as linguagens de programação utilizadas no repositório.

**Comando:**

```bash
gitstats languages <owner>/<repository>
```

**Exemplo:**

```bash
gitstats languages TechAbraao/jobs
```

**Informações apresentadas:**

- Linguagens identificadas.
- Quantidade de bytes por linguagem.
- Percentual de utilização de cada linguagem.
- Linguagem predominante.

**Exemplo de saída:**

```text
Languages

Python       65.4%  █████████████
TypeScript   22.1%  ████
HTML          8.5%  ██
CSS           4.0%  █
```

O percentual será calculado com base na quantidade de bytes atribuída a cada linguagem pela GitHub API, não na quantidade de arquivos.

## 3.5. Activity Analysis

Apresenta informações sobre a atividade de desenvolvimento do repositório ao longo do tempo.

**Comando:**

```bash
gitstats activity <owner>/<repository>
```

**Exemplos:**

```bash
gitstats activity TechAbraao/jobs

gitstats activity TechAbraao/jobs --period 12m

gitstats activity TechAbraao/jobs --period 30d

gitstats activity TechAbraao/jobs --period 90d
```

**Informações apresentadas:**

- Quantidade de commits por semana.
- Quantidade de commits por mês.
- Evolução da atividade.
- Períodos com maior atividade.
- Períodos sem atividade registrada.

**Exemplo de visualização:**

```text
Monthly Activity

Jan  ████████        18
Feb  ███████████     25
Mar  █████            12
Apr  █████████        21
May  █████████████    30
Jun  ████              9

Total commits: 115
```

A visualização será apresentada diretamente no terminal.

## 3.6. Repository Comparison

Permite comparar dois ou mais repositórios GitHub.

**Comando:**

```bash
gitstats compare <owner>/<repository-1> <owner>/<repository-2>
```

**Exemplo:**

```bash
gitstats compare TechAbraao/jobs TechAbraao/techabraao.com
```

**Critérios de comparação:**

- Stars.
- Forks.
- Watchers.
- Issues abertas.
- Quantidade de contribuidores.
- Quantidade de commits no período selecionado.
- Linguagens utilizadas.
- Data da última atualização.

A comparação deverá apresentar os dados lado a lado, facilitando a identificação das diferenças entre os projetos.

**Opções:**

```bash
gitstats compare TechAbraao/jobs TechAbraao/techabraao.com --json
```

# 4. Autenticação

A aplicação deverá permitir autenticação com o GitHub por meio de Personal Access Token (PAT).

A autenticação será opcional para consultas de repositórios públicos, mas recomendada para aumentar os limites de requisições e permitir acesso a recursos privados autorizados.

**Comandos:**

```bash
gitstats auth login
gitstats auth status
gitstats auth logout
```

**Funcionalidades:**

- Configuração do token.
- Validação das credenciais.
- Consulta do usuário autenticado.
- Verificação do status de autenticação.
- Remoção das credenciais armazenadas.

O token não deverá ser exibido no terminal nem armazenado em texto puro sem proteção.

# 5. Configurações

A aplicação deverá permitir personalização por meio de arquivos de configuração e flags.

**Localização padrão:**

```text
~/.config/gitstats/config.yaml
```

**Exemplo de configuração:**

```yaml
github:
  base_url: "https://api.github.com"
  timeout: 30

output:
  format: "table"
  color: true

cache:
  enabled: true
  ttl: 300
```

As configurações poderão ser sobrescritas por argumentos de linha de comando.

# 6. Formatos de saída

O GitStats deverá oferecer diferentes formatos de apresentação.

## 6.1. Table

Formato padrão, com informações organizadas em tabelas e elementos visuais no terminal.

```bash
gitstats repo TechAbraao/jobs
```

## 6.2. JSON

Permite integração com outras ferramentas e scripts.

```bash
gitstats repo TechAbraao/jobs --json
```

## 6.3. CSV

Permite exportação de dados para análise em planilhas.

```bash
gitstats contributors TechAbraao/jobs --format csv
```

## 6.4. Output to file

Permite salvar os resultados em arquivos.

```bash
gitstats contributors TechAbraao/jobs --format json --output contributors.json

gitstats activity TechAbraao/jobs --format csv --output activity.csv
```

# 7. Histórico e cache

O GitStats deverá utilizar SQLite para armazenar informações localmente.

**Funcionalidades:**

- Histórico de consultas.
- Registro de repositórios analisados.
- Armazenamento temporário de resultados.
- Redução de requisições desnecessárias à GitHub API.
- Consulta de análises anteriores.

**Comandos planejados:**

```bash
gitstats history

gitstats history --limit 10

gitstats history clear

gitstats cache clear

gitstats cache status
```

O cache deverá possuir tempo de expiração configurável e permitir atualização manual dos dados.

# 8. Integração com GitHub REST API

A aplicação utilizará os seguintes endpoints:

| Funcionalidade | Método e endpoint |
|---|---|
| Informações do repositório | `GET /repos/{owner}/{repo}` |
| Commits | `GET /repos/{owner}/{repo}/commits` |
| Contribuidores | `GET /repos/{owner}/{repo}/contributors` |
| Linguagens | `GET /repos/{owner}/{repo}/languages` |
| Atividade de commits | `GET /repos/{owner}/{repo}/stats/commit_activity` |
| Estatísticas semanais | `GET /repos/{owner}/{repo}/stats/participation` |
| Usuário autenticado | `GET /user` |

**Requisitos da integração:**

- Tratamento de respostas HTTP.
- Desserialização de JSON.
- Paginação de resultados.
- Controle de timeout.
- Tratamento de rate limits.
- Autenticação por token.
- Tratamento de respostas incompletas.
- Cache de resultados.

**Atenção às estatísticas:** alguns endpoints de estatísticas do GitHub podem retornar HTTP 202 enquanto os dados estão sendo calculados. A aplicação deverá tratar esse comportamento, evitando interpretar uma resposta ainda indisponível como ausência de atividade.

# 9. Estrutura de diretórios

A organização do projeto será a seguinte:

```text
gitstats/
│
├── assets/
│   ├── logo.png
│   └── screenshots/
│
├── gitstats/
│   │
│   ├── cmd/
│   │   ├── root.go
│   │   ├── repo.go
│   │   ├── commits.go
│   │   ├── contributors.go
│   │   ├── languages.go
│   │   ├── activity.go
│   │   ├── compare.go
│   │   ├── auth.go
│   │   ├── history.go
│   │   └── cache.go
│   │
│   └── internal/
│       │
│       ├── github/
│       │   ├── client.go
│       │   ├── repositories.go
│       │   ├── commits.go
│       │   ├── contributors.go
│       │   ├── languages.go
│       │   └── activity.go
│       │
│       ├── analyzer/
│       │   ├── commits.go
│       │   ├── contributors.go
│       │   ├── languages.go
│       │   └── activity.go
│       │
│       ├── config/
│       │   ├── config.go
│       │   └── auth.go
│       │
│       ├── output/
│       │   ├── table.go
│       │   ├── terminal.go
│       │   ├── json.go
│       │   └── csv.go
│       │
│       ├── cache/
│       │   └── cache.go
│       │
│       └── storage/
│           ├── database.go
│           └── history.go
│
├── scripts/
│   └── build.sh
│
├── .gitignore
├── LICENSE
├── Makefile
├── README.md
├── go.mod
├── go.sum
└── main.go
```

## Responsabilidades

**`assets/`**

Armazena recursos estáticos do projeto, como logotipo e imagens para documentação.

**`gitstats/cmd/`**

Contém a implementação dos comandos e subcomandos da CLI. Cada arquivo será responsável por definir argumentos, flags, validações e execução de uma funcionalidade.

**`gitstats/internal/github/`**

Responsável pela comunicação com a GitHub REST API, incluindo requisições, autenticação, paginação e interpretação das respostas.

**`gitstats/internal/analyzer/`**

Contém a lógica responsável por processar os dados recebidos e calcular estatísticas.

**`gitstats/internal/config/`**

Gerencia as configurações da aplicação e as credenciais de autenticação.

**`gitstats/internal/output/`**

Responsável pela apresentação e exportação dos resultados em diferentes formatos.

**`gitstats/internal/cache/`**

Gerencia o armazenamento temporário de respostas da API.

**`gitstats/internal/storage/`**

Gerencia o banco de dados SQLite e o histórico de consultas.

**`scripts/`**

Armazena scripts auxiliares para compilação e automação.

**`main.go`**

Ponto de entrada da aplicação, responsável por inicializar a CLI.

# 10. Tratamento de erros

A aplicação deverá possuir tratamento consistente de erros.

Os principais cenários são:

| Cenário | Comportamento esperado |
|---|---|
| Repositório inexistente | Informar que o repositório não foi encontrado |
| Nome inválido | Informar o formato esperado `owner/repository` |
| Sem conexão | Informar falha de comunicação |
| Rate limit excedido | Informar limite de requisições e horário de recuperação, quando disponível |
| Token inválido | Solicitar nova autenticação |
| Sem permissão | Informar que o usuário não possui acesso |
| Timeout | Informar que a requisição excedeu o tempo limite |
| Resposta inesperada | Apresentar erro sem encerrar abruptamente a aplicação |
| Estatísticas indisponíveis | Informar que o GitHub ainda está processando os dados |

A aplicação deverá apresentar mensagens claras e evitar expor detalhes sensíveis.

# 11. Concorrência e performance

Como Go possui suporte nativo à concorrência, o GitStats poderá utilizar goroutines e channels para executar operações independentes simultaneamente.

Exemplos de utilização:

- Consultar informações de múltiplos repositórios.
- Processar diferentes conjuntos de dados.
- Executar análises independentes em paralelo.
- Comparar repositórios simultaneamente.

A concorrência deverá respeitar os limites da GitHub API, utilizando controle de requisições e evitando sobrecarga desnecessária.

# 12. Testes

O projeto deverá possuir testes automatizados para garantir a confiabilidade das funcionalidades.

**Testes unitários:**

- Validação dos argumentos.
- Cálculo de percentuais de linguagens.
- Cálculo de estatísticas de commits.
- Processamento de contribuidores.
- Formatação de resultados.

**Testes de integração:**

- Comunicação com a GitHub API.
- Desserialização de respostas.
- Paginação.
- Autenticação.
- Integração com SQLite.

**Testes adicionais:**

- Repositórios inexistentes.
- Respostas HTTP inválidas.
- Rate limits.
- Timeout.
- Dados vazios.
- Cache.

As chamadas HTTP deverão ser testáveis sem depender constantemente da API real, utilizando mocks ou servidores HTTP de teste.

# 13. Distribuição

O GitStats deverá ser distribuído como um executável independente para diferentes sistemas operacionais.

**Sistemas suportados:**

- Linux.
- Windows.
- macOS.

**Arquiteturas previstas:**

- AMD64 (x86_64).
- ARM64.

**Exemplos de utilização após instalação:**

```bash
gitstats --version

gitstats --help

gitstats repo TechAbraao/jobs
```

A distribuição será automatizada com GoReleaser e GitHub Actions.

# 14. Roadmap de desenvolvimento

## Fase 1 — Estrutura inicial

- [ ] Inicializar Go Modules.
- [ ] Configurar Cobra.
- [ ] Criar estrutura de diretórios.
- [ ] Implementar comando raiz.
- [ ] Implementar comando `repo`.
- [ ] Configurar tratamento básico de erros.

## Fase 2 — Integração com GitHub

- [ ] Implementar cliente HTTP.
- [ ] Consumir endpoint de repositórios.
- [ ] Criar estruturas de dados (structs).
- [ ] Implementar desserialização JSON.
- [ ] Implementar paginação.
- [ ] Tratar rate limits e timeouts.

## Fase 3 — Análises

- [ ] Implementar `commits`.
- [ ] Implementar `contributors`.
- [ ] Implementar `languages`.
- [ ] Implementar `activity`.
- [ ] Implementar `compare`.

## Fase 4 — Interface e exportação

- [ ] Implementar tabelas no terminal.
- [ ] Adicionar estilização com Lip Gloss.
- [ ] Implementar barras de progresso.
- [ ] Implementar saída JSON.
- [ ] Implementar saída CSV.
- [ ] Implementar exportação para arquivos.

## Fase 5 — Recursos adicionais

- [ ] Implementar autenticação.
- [ ] Implementar configurações com Viper.
- [ ] Implementar cache.
- [ ] Implementar histórico com SQLite.
- [ ] Adicionar processamento concorrente.

## Fase 6 — Qualidade e distribuição

- [ ] Criar testes unitários.
- [ ] Criar testes de integração.
- [ ] Configurar GitHub Actions.
- [ ] Criar Makefile.
- [ ] Configurar GoReleaser.
- [ ] Gerar binários multiplataforma.
- [ ] Finalizar documentação.

# 15. Padrões de utilização

A CLI deverá seguir algumas convenções:

**Sintaxe padrão:**

```bash
gitstats <command> <arguments> [flags]
```

**Ajuda:**

```bash
gitstats --help

gitstats repo --help

gitstats commits --help
```

**Versão:**

```bash
gitstats --version
```

**Formato de saída:**

```bash
gitstats repo TechAbraao/jobs --format json
```

**Atualização de dados:**

```bash
gitstats repo TechAbraao/jobs --refresh
```

**Regras gerais:**

- Os comandos deverão possuir nomes intuitivos.
- Os argumentos deverão ser validados.
- As flags deverão possuir descrições.
- A saída padrão deverá ser amigável para leitura humana.
- Os formatos de exportação deverão ser consistentes.
- Erros deverão possuir mensagens claras.
- Os comandos deverão ser documentados individualmente.

# 16. Escopo do MVP

A primeira versão funcional do GitStats não precisará implementar todas as funcionalidades previstas.

O MVP será composto por:

1. Estrutura inicial da aplicação.
2. CLI utilizando Cobra.
3. Integração com GitHub REST API utilizando `net/http`.
4. Consulta de informações de repositórios.
5. Análise de linguagens.
6. Análise de commits.
7. Análise de contribuidores.
8. Apresentação dos resultados no terminal.
9. Tratamento de erros.
10. Testes automatizados básicos.

Os recursos de autenticação, cache, SQLite, comparação avançada e distribuição serão implementados nas próximas versões.

---

# 17. Resultado esperado

Ao final do desenvolvimento, o GitStats deverá ser uma ferramenta de linha de comando funcional, capaz de consultar, analisar, comparar e apresentar informações de repositórios GitHub.

A aplicação deverá oferecer uma experiência simples de utilização, com comandos intuitivos, resultados organizados, possibilidade de exportação e arquitetura modular.

**GitStats — Understand your repositories.**