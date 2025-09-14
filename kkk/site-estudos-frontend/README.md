# Site Estudos Frontend

Este é um projeto de frontend desenvolvido em React, destinado a fornecer uma plataforma de estudos com recursos de pesquisa e recomendação de cursos.

## Estrutura do Projeto

O projeto possui a seguinte estrutura de diretórios:

```
site-estudos-frontend
├── public
│   └── index.html          # Ponto de entrada do aplicativo
├── src
│   ├── components          # Componentes reutilizáveis
│   │   ├── SearchBar.tsx   # Componente de barra de pesquisa
│   │   ├── RecommendedCourses.tsx # Componente para cursos recomendados
│   │   └── CourseCard.tsx  # Componente para exibir informações de um curso
│   ├── pages               # Páginas do aplicativo
│   │   └── Home.tsx        # Página principal
│   ├── App.tsx             # Componente principal do aplicativo
│   ├── index.tsx           # Ponto de entrada do React
│   └── styles              # Estilos do aplicativo
│       └── global.css      # Estilos globais
├── package.json            # Configuração do npm
├── tsconfig.json           # Configuração do TypeScript
└── README.md               # Documentação do projeto
```

## Instalação

Para instalar as dependências do projeto, execute o seguinte comando na raiz do diretório do projeto:

```
npm install
```

## Uso

Para iniciar o servidor de desenvolvimento, utilize o comando:

```
npm start
```

O aplicativo estará disponível em `http://localhost:3000`.

## Contribuição

Sinta-se à vontade para contribuir com melhorias ou correções. Para isso, faça um fork do repositório e envie um pull request com suas alterações.

## Licença

Este projeto está licenciado sob a MIT License. Veja o arquivo LICENSE para mais detalhes.