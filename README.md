# godeclutter
Distribui URLs de maneira extremamente rápida e flexível, para melhorar a entrada para automações de hackers da Web, como rastreadores e varreduras de vulnerabilidade.

O godeclutter é uma ferramenta muito simples que pegará uma lista de URLs, limpará essas URLs e produzirá as URLs exclusivas que estão presentes. Isso reduzirá o número de solicitações que você terá que fazer ao seu site de destino e também filtrará URLs que provavelmente são desinteressantes.

Características
O godeclutter executará as seguintes etapas na seção de host da URL:

Limpe http:// URLs que apontam para a porta SSL padrão (443) e vice-versa, já que são principalmente páginas de erro CDN.
Notação de porta limpa de URLs apontando para as portas de protocolo padrão, uma vez que essas portas já estão implícitas no esquema de protocolo. (como :443 e :80)
Limpe http:// URLs se um https:// para o mesmo host e porta estiver presente, já que 99,9% desses casos serão apenas um redirecionamento para https://.
Remova extensões de mídia desinteressantes, como png, jpg, css, etc. (Este manterá .js arquivos, uma vez que esses às vezes são interessantes)
Remova urls com palavras desinteressantes, como bootstrap, jquery, node_modules, etc.
Classificar parâmetros de consulta
Minúsculas todos os esquemas e nomes de host, uma vez que a caixa maiúscula é irrelevante para esses.
Substitua todos os escapes de codificação de URI minúsculos para maiúsculas, para manter um padrão.
Decodifice escapes desnecessários para caracteres que não são especiais no contexto da URL (ou seja, http://example.com/%41 ).
Remover cadeias de caracteres de consulta vazias (ou seja, http://example.com/? )
Remova as barras à direita, isso é bastante agressivo, mas filtra muitas duplicatas desinteressantes na maioria dos casos. (ou seja, http://host/path/ -> http://host/path)
Normalize os segmentos de pontos, também bastante agressivos, mas úteis ao trabalhar com fontes sujas. (http://host/path/./a/b/../c -> http://host/path/a/c)
Instalar
go install github.com/c3l3si4n/godeclutter@HEAD
Uso Básico
Você pode enviar URLs enviando-os para stdin.

user@arch ~/D/g/godeclutter (main)> cat test_urls.txt 
https://1.1.1.1:443/
http://1.1.1.1/
http://1.1.1.1:80/
https://1.1.1.1:443/
https://1.1.1.1:80/
https://1.1.1.1:8443/
https://1.1.1.1:8443/
https://1.1.1.1:443/
https://1.1.1.1:8443/
https://1.1.1.1:8443/
https://1.1.1.1:443/?
http://1.1.1.1:443/?
https://1.1.1.1:80/?a
https://1.1.1.1:443/?
https://1.1.1.1:443/?
https://1.1.1.1:443/?1=1
https://1.1.1.1:443/?a=a&b=1
https://1.1.1.1:443/?a=a&b=1
https://1.1.1.1:443/a.js?a=a&b=1
https://1.1.1.1:443/fiqef.html?a=a&b=1
https://1.1.1.1:443/fmef.jpg?b=1&a=a
https://1.1.1.1:443/?b=1&a=a
https://1.1.1.1:443/a.js?
https://1.1.1.1:443/a.jpg
https://1.1.1.1:8443/
https://1.1.1.1:443/
https://1.1.1.1:443/node_modules/
https://1.1.1.1:443/path/scripts/jquery.js
user@arch ~/D/g/godeclutter (main)> cat test_urls.txt | ./godeclutter -bw -be -c -p
https://1.1.1.1/
https://1.1.1.1:8443/
https://1.1.1.1/?1=1
https://1.1.1.1/?a=a&b=1
https://1.1.1.1/a.js?a=a&b=1
https://1.1.1.1/fiqef.html?a=a&b=1
https://1.1.1.1/a.js
Argumentos
$> ./godeclutter -h
Usage of ./godeclutter:
  -be
    	Blacklist Extensions - clean some uninteresting extensions. (default true)
  -bec string
    	Blacklist Extensions - Specify additional extensions separated by commas to be cleared along the default ones.
  -bw
    	Blacklist Words - clean some uninteresting words. (default true)
  -bwc string
    	Blacklist Words - Specify additional words separated by commas  to be cleared along the default ones.
  -bwl string
    	Blacklist Words - Defines the level of word blocking. Values can be: minimal,aggressive (default "minimal")
  -c	Clean URLs - Aggressively clean/normalize URLs before outputting them. (default true)
  -p	Prefer HTTPS - If there's a https url present, don't print the http for it. (since it will probably just redirect to https)
