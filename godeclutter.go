pacote principal

importação (
	"bufio"
	"bandeira"
	"fmt"
	"rede"
	"net/url"
	"os"
	"cordas"

	"github.com/PuerkitoBio/purell"
)

var strFlag = sinalizador. String("long-string", "", "Descrição")
var preferHttpsFlag = sinalizador. Bool("p", false, "Prefira HTTPS - Se houver uma url https presente, não imprima o http para ela. (já que provavelmente apenas redirecionará para https)")
var normalizeURLFlag = sinalizador. Bool("c", true, "URLs limpos - Limpe/normalize agressivamente URLs antes de emiti-los".)
var blacklistExtensionsFlag = sinalizador. Bool("ser", verdade, "Blacklist Extensions - limpe algumas extensões desinteressantes.")
var customBlacklistExtensionsFlag = sinalizador. String("bec", "", "Blacklist Extensions - Especifique extensões adicionais separadas por vírgulas a serem limpas ao longo das extensões padrão.")

var blacklistWordsFlag = sinalizador. Bool("bw", verdade, "Blacklist Words - limpe algumas palavras desinteressantes".)
var blacklistedPresetFlag = sinalizador. String("bwl", "minimal", "Blacklist Words - Define o nível de bloqueio de palavras. Os valores podem ser: mínimos, agressivos")
var customBlacklistWordsFlag = sinalizador. String("bwc", "", "Blacklist Words - Especifique palavras adicionais separadas por vírgulas a serem limpas ao longo das palavras padrão.")

var blacklistedExtensions = []string{"css", "scss", "png", "jpg", "jpeg", "img", "svg", "ico", "webp", "webm", "tif", "ttf", "tiff", "otf", "woff", "woff2", "gif", "pdf", "bmp", "eot" , "mp3", "mp4", "m4a", "m4p", "avi", "flv", "swf", "eot" }

func iterInput(c chan string) {
	scanner := bufio. NewScanner(os. Stdin)
	para scanner. Varredura() {
		c <- scanner. Texto()
	}

	fechar(c)
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	lista := []string{}
	para _, item := intervalo strSlice {
		if _, value := allKeys[item]; ! valor {
			allKeys[item] = verdadeiro
			list = append(lista, item)
		}
	}
	lista de retornos
}

func stringInSlice(uma cadeia de caracteres, lista []string) (int, bool) {
	para i, b := lista de intervalos {
		se b == a {
			return i, true
		}
	}
	return 0, false
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	retorno s[:len(s)-1 ]
}

func normalizeURL (cadeia de caracteres url) string {
	normalizado := purel. MustNormalizeURLString(url, purell. FlagLowercaseScheme| purejar. SinalizarLowercaseHost| purejar. FlagUppercaseEscapes| purejar. FlagDecodeUnnecessaryEscapes| purejar. FlagEncodeNecessaryEscapes| purejar. FlagRemoveDefaultPort| purejar. FlagRemoveEmptyQuerySeparator| purejar. FlagRemoveDotSegments| purejar. SinalizarRemoverDuplicateSlashs| purejar. FlagSortQuery)
	retorno normalizado

}

func principal() {

	bandeira. Analisar()
	var blacklistedWords []string

	if *blacklistedPresetFlag == "agressivo" {
		blacklistedWords = []string{"/node_modules/", "wp-includes", "jquery", "/bootstrap", "/webpack-runtime", "accessChatPublic"}
	} else if *blacklistedPresetFlag == "mínimo" {
		blacklistedWords = []string{"/node_modules/", "/wp-includes/", "/jquery", "/webpack-runtime", "accessChatPublic"}
	}

	if *customBlacklistWordsFlag != "" {
		additionalWords := cadeias de caracteres. Dividir(*customBlacklistWordsFlag, ",")
		para _, palavra := intervalo adicionalPalavras {
			blacklistedWords = append(blacklistedPalavras, palavra)
		}
	}

	if *customBlacklistExtensionsFlag != "" {
		additionalExtensions := cadeias de caracteres. Dividir(*customBlacklistExtensionsFlag, ",")
		para _, extensão := intervalo additionalExtensions {
			blacklistedExtensions = append(blacklistedExtensions, extensão)
		}
	}

	var c  chan string = make(chan string )
	ir iterInput(c)

	var processedUrls []string
	processedUrlMap := make(map[string]string)

	para a linha := intervalo c {

		u, err := url. ParseRequestURI(linha)
		se errar != nulo {
			continuar
		}

		Remover falsos positivos http-on-https
		se u. Esquema == "http" && u. Port() == "443" {
			continuar
		}
		se u. Esquema == "https" && u. Port() == "80" {
			continuar
		}

		Escape redundant port syntax (para portas padrão)
		se strings. Contém(u. Anfitrião, ":") {
			host, port, err := net. SplitHostPort(u. Anfitrião)
			se errar != nulo {
				continuar
			}

			se porta == "443" && u. Esquema == "https" {
				u. Host = host
			} else if port == "80" && u. Esquema == "http" {
				u. Host = host
			} else se a porta == "" {
				fmt. Imprimir(host)
				u. Host = host
			} mais {
				u. Host = host  + ":" + porta
			}

		}

		Prefira https
		if *preferHttpsFlag {
			se u. Esquema == "https" {
				scheme, hostname_found := processedUrlMap[u. Anfitrião]
				se hostname_found {
					se esquema == "http" {
						check_u, _ := url. Analisar(u. Corda())
						check_u. Esquema = "http"
						found_index, encontrado := stringInSlice(check_u. String(), processedUrls)
						se encontrado {
							processedUrls = remove(processedUrls, found_index)
						}
						processedUrlMap[u. Host] = u. Esquema
					}
				} mais {
					processedUrlMap[u.Host] = u.Scheme
				}

			} else if u.Scheme == "http" {
				_, hostname_found := processedUrlMap[u.Host]
				if hostname_found {
					continue
				}
			}
		}

		if *blacklistExtensionsFlag {
			foundBlacklistedExtension := false
			for _, ext := range blacklistedExtensions {
				if strings.HasSuffix(u.Path, ext) {
					foundBlacklistedExtension = true
					break
				}
			}
			if foundBlacklistedExtension {
				continuar
			}
		}

		se *blacklistWordsFlag {
			foundBlacklistedWord := falso
			para _, palavra := intervalo na lista negraPalavras {
				se strings. Contém(u. Caminho, palavra) {
					foundBlacklistedWord = verdadeiro
					quebrar
				}
			}
			se encontradoBlacklistedWord {
				continuar
			}
		}

		if *normalizeURLFlag {
			u_str := normalizeURL(u. Corda())
			processedUrls = append(processedUrls, u_str)
		} mais {
			processedUrls = append(processedUrls, u. Corda())

		}

	}

	filteredProcessedUrls := removeDuplicateStr(processedUrls)

	para _, url := intervalo filtradoProcessedUrls {
		fmt. Println(url)
	}

}
