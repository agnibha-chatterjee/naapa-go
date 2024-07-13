package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type ApiMessage struct {
	Role      string          `json:"role"`
	Content   string          `json:"content,omitempty"`
	ImageName string          `json:"imageName,omitempty"`
	Image     map[string]byte `json:"image,omitempty"`
}

func collectMapValues(m map[string]byte) []byte {
	values := make([]byte, len(m))

	for _, value := range m {
		values = append(values, value)
	}

	return values
}

func saveImages(msgs []ApiMessage) {
	for _, msg := range msgs {
		if msg.Role == "user" {
			msg.saveImage()
		}
	}
}

func (m *ApiMessage) saveImage() error {
	imageData := collectMapValues(m.Image)

	fmt.Printf("%+v", imageData)

	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create(m.ImageName + ".jpeg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 1

	err = jpeg.Encode(out, img, &opts)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}

func (m *ApiMessage) UnmarshalJSON(b []byte) error {

	var msg map[string]any

	if err := json.Unmarshal(b, &msg); err != nil {
		return err
	}

	if role, ok := msg["role"].(string); ok {
		m.Role = role
	}

	if content, ok := msg["content"].(string); ok {
		m.Content = content
	}

	if imageName, ok := msg["imageName"].(string); ok {
		m.ImageName = imageName
	}

	if image, ok := msg["image"].(map[string]interface{}); ok {
		m.Image = make(map[string]byte)
		for k, v := range image {
			if floatVal, ok := v.(float64); ok {
				m.Image[k] = byte(floatVal)
			}
		}
	}

	return nil
}

func HandleExecuteTask(w http.ResponseWriter, r *http.Request) {

	var messages []ApiMessage

	err := json.NewDecoder(r.Body).Decode(&messages)

	defer r.Body.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	saveImages(messages)

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

	sysMsgFromReq := messages[0]

	sysMsg := &genai.Content{
		Role: sysMsgFromReq.Role,
		Parts: []genai.Part{
			genai.Text(sysMsgFromReq.Content),
		},
	}
	model.SystemInstruction = sysMsg

	resp, err := model.GenerateContent(ctx, genai.Text("Oh it's such a hot day jeez"))
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

	return res
}
