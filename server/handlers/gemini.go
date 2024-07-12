package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func HandleExecuteTask(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-pro")
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
	}
	resp, err := model.GenerateContent(ctx, genai.Text("I want to be bad. Please help."))
	if err != nil {
		log.Fatal(err)
	}

	res := printResponse(resp)

	w.Write([]byte(res))
}

func printResponse(resp *genai.GenerateContentResponse) string {

	res := ""

	for _, c := range resp.Candidates {
		if c.Content != nil {
			for _, part := range c.Content.Parts {
				res += fmt.Sprint(part)
			}
		}
	}
	fmt.Println("---")

	return res
}
