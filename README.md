# realchat
---
## 🚀 Prévia

Imagine um chat simples, direto e funcional, estilo "sala pública", onde todo mundo que acessa pode trocar mensagens em tempo real,  é só fazer login e conversar. Ele tem uma interface amigável, mensagens instantâneas e tudo isso rodando com a performance do Go por trás.

---
## 🔐 Usuários para login

O sistema já vem com alguns usuários cadastrados para testes:

- **vinicius** – senha: `171`  
- **gabi** – senha: `6969`  
- **percy** – senha: `123`  

🔧 Se quiser, você pode adicionar novos usuários diretamente no backend.

---
## 💬 Múltiplas Salas de Chat

Além do chat principal, o sistema também tem suporte a **salas separadas**!

- Ao entrar, você está na **sala pública**.
- Para trocar de sala, digite o nome da sala no **prompt que aparece**.
- Exemplo: digite `private` para ir para a **sala privada**.

Você pode criar e entrar em qualquer sala só escrevendo o nome. Se a sala não existir, ela será criada automaticamente quando você entrar nela.

---
## 📸 Demonstração

[🎥 Clique para ver o vídeo](./docs/video.mp4)

---
## 🧰 Tecnologias Usadas

- **Go (Golang)** — servidor backend com WebSockets
- **JavaScript** — lógica do frontend
- **HTML + CSS** — estrutura e estilo da interface
---
## 🛠️ Como rodar o projeto

1. Clone o repositório:

```bash
git clone https://github.com/ViniciusDadalte/realchat.git
cd realchat
```

2. Execute o arquivo certgen.bash para poder ter os arquivos essenciais para iniciar o servidor:
```bash
bash certgen.bash
```

3. Execute o projeto:
```bash
go run *.go
```

4. Acesse no navegador:
```bash
http://localhost:8080
```

5. Faça login com um dos usuários e:

	- Envie mensagens e veja tudo rolando em tempo real!
	
	- Caso queira mudar de sala digite no prompt `private` para entrar na sala privada, ou o nome da sala que quer entrar ou criar.
---
## ✅ Funcionalidades

- 🔐 Login com usuários pré-cadastrados
    
- 🔄 Envio e recebimento de mensagens em tempo real
    
- 👥 Múltiplas salas de chat
    
- 📩 Broadcast por sala (mensagens separadas por grupo)
    
- 💻 Interface leve
    
- 🚀 Backend rápido e eficiente com Go

---
## 👨‍💻 Autor

Isso foi feito com 💻 e ☕ por mim, e é isso, espero que tenham gostado :)
