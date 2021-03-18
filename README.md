# httpserv

Eu me encontro precisando de um servidor de arquivos HTTP simples com frequência e `python -m SimpleHTTPServer` simplesmente não é mais adequado para mim, já que sempre roda em` 0.0.0.0`.

A ferramenta `httpserv` leva as coisas apenas uma ideia adiante, permitindo que os aspectos mais básicos (host, porta, diretório) sejam configurados.

Use o seguinte comando para baixar o servidor (requer um ambiente Go funcional):

```sh
GO111MODULE=on go get https://github.com/isaccanedo/httpserv.git
```
