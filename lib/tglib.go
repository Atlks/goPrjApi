package lib

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)

type Button struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data,omitempty"`
	URL          string `json:"url,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]Button `json:"inline_keyboard"`
}

func main333() {
	htmlStr := `
	<table>
		<tr>
			<td>
				<button data-callback_data="daigou">
					代购须知
				</button>
				<button data-callback_data="daishou">
					代收须知
				</button>
			</td>
		</tr>
		<tr>
			<td>
				<button data-url="https://t.me/LianXin_ShangWu">
					联系代购/付客服
				</button>
			</td>
		</tr>
	</table>`

	jsonResult, err := ConvertHTMLToJSON(htmlStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonResult))
}

func ConvertHTMLToJSON(htmlStr string) ([]byte, error) {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}

	var buttons [][]Button
	var currentRow []Button

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "button" {
			var button Button
			for _, attr := range n.Attr {
				if attr.Key == "data-callback_data" {
					button.CallbackData = attr.Val
				} else if attr.Key == "data-url" {
					button.URL = attr.Val
				}
			}
			if n.FirstChild != nil {
				button.Text = strings.TrimSpace(n.FirstChild.Data)
			}
			currentRow = append(currentRow, button)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

		if n.Type == html.ElementNode && n.Data == "tr" && len(currentRow) > 0 {
			buttons = append(buttons, currentRow)
			currentRow = nil
		}
	}
	f(doc)

	markup := InlineKeyboardMarkup{InlineKeyboard: buttons}
	return json.MarshalIndent(markup, "", "  ")
}
