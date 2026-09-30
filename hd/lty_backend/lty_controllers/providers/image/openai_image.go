package image

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"lty_backend/lty_models"
)

func ProcessOpenAIImageTask(
	upstream *lty_models.ModelUpstream,
	prompt, sizeForAPI, resolutionForAPI, referenceImage string,
	useURL bool,
	updateLog func(string),
	index, totalImages int,
	downloadImageToLocal func(string, *http.Request) (string, error),
	getBaseURL func(*http.Request) string,
) ([]string, error) {
	reqURL := upstream.ConnectionURL

	// OpenAI 标准协议处理: 补全 URL
	parsedURL, parseErr := url.Parse(reqURL)
	if parseErr == nil && (parsedURL.Path == "" || parsedURL.Path == "/" || parsedURL.Path == "/v1" || parsedURL.Path == "/v1/") {
		suffix := "images/generations"
		isAgnes := strings.Contains(strings.ToLower(upstream.LogicalModel), "agnes")
		if !useURL && referenceImage != "" && !isAgnes {
			suffix = "images/edits"
		}
		if strings.HasSuffix(reqURL, "/v1") || strings.HasSuffix(reqURL, "/v1/") {
			if strings.HasSuffix(reqURL, "/") {
				reqURL = reqURL + suffix
			} else {
				reqURL = reqURL + "/" + suffix
			}
		} else {
			if strings.HasSuffix(reqURL, "/") {
				reqURL = reqURL + "v1/" + suffix
			} else {
				reqURL = reqURL + "/v1/" + suffix
			}
		}
	}

	var currentGenerated []string
	reqFormat := "b64_json" // 默认优先请求 b64_json，避免下载 URL 缓慢导致重试

	for attempt := 0; attempt <= 1; attempt++ {
		var req *http.Request
		var reqErr error
		var reqBodyStr string

		// 判断是否使用 Multipart Form-Data (edits 接口)
		isAgnes := strings.Contains(strings.ToLower(upstream.LogicalModel), "agnes")
	
	if !useURL && referenceImage != "" && strings.Contains(reqURL, "/v1/images/edits") && !isAgnes {
		var b bytes.Buffer
		w := multipart.NewWriter(&b)

		w.WriteField("model", upstream.LogicalModel)
		w.WriteField("prompt", prompt)
		w.WriteField("n", "1")

		if sizeForAPI != "" && sizeForAPI != "auto" {
			w.WriteField("size", sizeForAPI)
		}
		
		w.WriteField("response_format", reqFormat) // 明确要求上游优先返回 URL 直链或 b64

		isAgnes := strings.Contains(strings.ToLower(upstream.LogicalModel), "agnes")
		if resolutionForAPI != "" && resolutionForAPI != "auto" && !isAgnes {
			w.WriteField("resolution", resolutionForAPI)
		}

		var refImages []string
		if strings.HasPrefix(referenceImage, "[") && strings.HasSuffix(referenceImage, "]") {
			json.Unmarshal([]byte(referenceImage), &refImages)
		} else if referenceImage != "" {
			refImages = append(refImages, referenceImage)
		}

		for i, refImg := range refImages {
			if idx := strings.Index(refImg, "/uploads/"); idx != -1 {
				localPath := refImg[idx+1:]
				fileBytes, err := os.ReadFile(localPath)
				if err == nil {
					isZhimeng := strings.Contains(reqURL, "zmoapi.cn") || strings.Contains(reqURL, "zhimeng")
					if isZhimeng {
						// 知梦 API 专属逻辑：严格只发送 image[]，避免与 image 字段冲突导致其只读取单图
						if fwArr, errArr := w.CreateFormFile("image[]", filepath.Base(localPath)); errArr == nil {
							fwArr.Write(fileBytes)
						}
					} else {
						// 1. 兼容常规的 image, image2, image3 等格式
						fieldName := "image"
						if i > 0 {
							fieldName = fmt.Sprintf("image%d", i+1)
						}
						if fw, err := w.CreateFormFile(fieldName, filepath.Base(localPath)); err == nil {
							fw.Write(fileBytes)
						}
						
						// 2. 专门针对类似 zhimengapi 的 image[] 数组格式
						if fwArr, errArr := w.CreateFormFile("image[]", filepath.Base(localPath)); errArr == nil {
							fwArr.Write(fileBytes)
						}
						
						// 3. 兼容 images[] 格式
						if fwArr2, errArr2 := w.CreateFormFile("images[]", filepath.Base(localPath)); errArr2 == nil {
							fwArr2.Write(fileBytes)
						}
						
						// 4. 兼容 images 格式 (同名字段传多次)
						if fwArr3, errArr3 := w.CreateFormFile("images", filepath.Base(localPath)); errArr3 == nil {
							fwArr3.Write(fileBytes)
						}
						
						// 5. 对于第一张图，提供 file 字段以防极个别 API 需要
						if i == 0 {
							if fwFile, errFile := w.CreateFormFile("file", filepath.Base(localPath)); errFile == nil {
								fwFile.Write(fileBytes)
							}
						}
					}
				} else {
					updateLog("Failed to read local image for multipart conversion: " + err.Error())
				}
			}
		}
		w.Close()

		reqBodyStr = fmt.Sprintf("Multipart Form-Data: model=%s, prompt=%s, n=1, response_format=%s, size=%s, resolution=%s", upstream.LogicalModel, prompt, reqFormat, sizeForAPI, resolutionForAPI)
		req, reqErr = http.NewRequest("POST", reqURL, &b)
		if reqErr == nil {
			req.Header.Set("Content-Type", w.FormDataContentType())
		}
	} else {
		// 常规 JSON 格式 (generations 接口)
		requestBody := map[string]interface{}{
			"model":           upstream.LogicalModel,
			"prompt":          prompt,
			"n":               1,
			"response_format": reqFormat, // 动态使用 url 或 b64_json
		}

		if referenceImage != "" {
			var refImages []string
			if strings.HasPrefix(referenceImage, "[") && strings.HasSuffix(referenceImage, "]") {
				json.Unmarshal([]byte(referenceImage), &refImages)
			} else {
				refImages = append(refImages, referenceImage)
			}

			if len(refImages) > 0 {
				var promptURLs []string
				var base64Data []string
				var firstImage string

				for i, refImg := range refImages {
					var currentImg string
					if idx := strings.Index(refImg, "/uploads/"); idx != -1 {
						if !useURL {
							// 非 URL 模式：将本地图片路径转换为 Base64
							localPath := refImg[idx+1:]
							fileBytes, err := os.ReadFile(localPath)
							if err == nil {
								ext := strings.ToLower(filepath.Ext(localPath))
								mimeType := "image/png"
								if ext == ".jpg" || ext == ".jpeg" {
									mimeType = "image/jpeg"
								} else if ext == ".webp" {
									mimeType = "image/webp"
								}
								base64Str := base64.StdEncoding.EncodeToString(fileBytes)
								currentImg = fmt.Sprintf("data:%s;base64,%s", mimeType, base64Str)
								base64Data = append(base64Data, currentImg)
							} else {
								updateLog("Failed to read local image for base64 conversion: " + err.Error())
							}
						} else {
							// URL 模式：强制替换为当前最新的公网域名
							baseURL := getBaseURL(nil)
							currentImg = baseURL + refImg[idx:]
							// 知梦 API 强制要求 HTTPS 链接，如果当前环境为 HTTP，且上游为 zhimeng，则尝试替换为 HTTPS
							isZhimeng := strings.Contains(reqURL, "zmoapi.cn") || strings.Contains(reqURL, "zhimeng")
							if isZhimeng && strings.HasPrefix(currentImg, "http://") {
							    // 只有在使用默认端口或其他明显是公网域名时才替换
							    // 对于内网 IP 如 127.0.0.1 或者局域网 IP，如果强转 HTTPS 会导致图片无法访问
							    // 但因为上游强制要求 HTTPS，所以我们只能无条件替换
							    currentImg = strings.Replace(currentImg, "http://", "https://", 1)
							}
							promptURLs = append(promptURLs, currentImg)
						}
					} else {
						currentImg = refImg
						if strings.HasPrefix(refImg, "http") {
						    isZhimeng := strings.Contains(reqURL, "zmoapi.cn") || strings.Contains(reqURL, "zhimeng")
							if isZhimeng && strings.HasPrefix(currentImg, "http://") {
							    currentImg = strings.Replace(currentImg, "http://", "https://", 1)
							}
							promptURLs = append(promptURLs, currentImg)
						} else if strings.HasPrefix(refImg, "data:") {
							base64Data = append(base64Data, refImg)
						}
					}
					
					if i == 0 && currentImg != "" {
						firstImage = currentImg
					}
					refImages[i] = currentImg // 更新回原数组以便后续传递
				}

				// 只有 URL 模式的图片才能拼接到 prompt 里，Base64 绝对不能拼接进 prompt 否则会导致 Request Entity Too Large
				if len(promptURLs) > 0 {
					promptPrefix := strings.Join(promptURLs, " ")
					requestBody["prompt"] = promptPrefix + " " + prompt
				} else {
					requestBody["prompt"] = prompt
				}
				
				if len(refImages) == 1 {
					requestBody["image_url"] = firstImage
					requestBody["image"] = firstImage
				} else {
					requestBody["image_url"] = refImages
					requestBody["image"] = refImages
				}
				
				// Midjourney Proxy 专用多图参数 (必须是 Base64 格式数组)
				if len(base64Data) > 0 {
					requestBody["base64Array"] = base64Data
				}
				
				if len(refImages) > 1 {
					requestBody["images_flat"] = refImages
					requestBody["image_urls"] = refImages
				}
				
				// 兼容 zhimengapi/OpenAI 的 Images 数组对象格式 [{"image_url": "..."}]
				var zhimengImages []map[string]string
				for _, img := range refImages {
				    zhimengImages = append(zhimengImages, map[string]string{"image_url": img})
				}
				requestBody["images"] = zhimengImages
			}
		}

		if sizeForAPI != "" && sizeForAPI != "auto" {
			requestBody["size"] = sizeForAPI
		}
		
		isAgnes := strings.Contains(strings.ToLower(upstream.LogicalModel), "agnes")
		if resolutionForAPI != "" && resolutionForAPI != "auto" && !isAgnes {
			requestBody["resolution"] = resolutionForAPI
		}

		reqBytes, _ := json.Marshal(requestBody)
		reqBodyStr = string(reqBytes)
		req, reqErr = http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBytes))
		if reqErr == nil {
			req.Header.Set("Content-Type", "application/json")
		}
	}

		if reqErr != nil {
			updateLog("Failed to create request to upstream API: " + reqErr.Error())
			return nil, reqErr
		}

		if upstream.APIKey != "" {
			req.Header.Set("Authorization", "Bearer "+upstream.APIKey)
		}

		logReqBody := reqBodyStr
		if len(logReqBody) > 2000 {
			logReqBody = logReqBody[:2000] + "... [truncated]"
		}
		updateLog(fmt.Sprintf("Sending request to upstream API (Task %d/%d): %s\nBody: %s", index+1, totalImages, reqURL, logReqBody))

		timeout := time.Duration(upstream.TimeoutSeconds) * time.Second
		if timeout <= 0 {
			timeout = 300 * time.Second
		}
		client := &http.Client{Timeout: timeout}

		resp, err := client.Do(req)
		if err != nil {
			updateLog(fmt.Sprintf("Upstream API request failed (Task %d): %v", index+1, err))
			return nil, err
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		logBody := string(body)
		if len(logBody) > 2000 {
			logBody = logBody[:2000] + "... [truncated]"
		}

		// 将上游返回的真实完整数据记录到日志中，供前端任务详情查看
		updateLog(fmt.Sprintf("OpenAI upstream response (Task %d): %s", index+1, logBody))

		if resp.StatusCode != http.StatusOK {
			errMsg := fmt.Sprintf("Upstream API returned error status (Task %d): %d\nBody: %s", index+1, resp.StatusCode, logBody)
			
			// 如果是因为请求了 b64_json 导致上游报错，则自动切换为 url 格式重试一次
			if attempt == 0 && reqFormat == "b64_json" {
				updateLog(fmt.Sprintf("Upstream may not support b64_json format, automatically switching to 'url' format and retrying... (Task %d)", index+1))
				reqFormat = "url"
				continue
			}

			updateLog(errMsg)
			return nil, fmt.Errorf("upstream returned %d", resp.StatusCode)
		}

		// 解析 OpenAI 响应
		var apiRes struct {
			Data []struct {
				URL     string `json:"url"`
				B64Json string `json:"b64_json"`
			} `json:"data"`
		}

		if err := json.Unmarshal(body, &apiRes); err != nil {
			errMsg := fmt.Sprintf("Failed to parse upstream response (Task %d): %v", index+1, err)
			updateLog(errMsg)
			return nil, err
		}

		if len(apiRes.Data) == 0 {
			updateLog(fmt.Sprintf("Upstream returned empty data (Task %d)", index+1))
			return nil, fmt.Errorf("empty response")
		}

		for _, item := range apiRes.Data {
			var urlStr string
			// 优先使用请求的格式
			if reqFormat == "url" && item.URL != "" {
				urlStr = item.URL
			} else if reqFormat == "b64_json" && item.B64Json != "" {
				urlStr = "data:image/png;base64," + item.B64Json
			} else if item.URL != "" {
				urlStr = item.URL
			} else if item.B64Json != "" {
				urlStr = "data:image/png;base64," + item.B64Json
			}

			if urlStr != "" {
				localUrl, err := downloadImageToLocal(urlStr, nil)
				if err == nil {
					currentGenerated = append(currentGenerated, localUrl)
				} else {
					updateLog(fmt.Sprintf("Failed to download/save image to local server (Task %d, possibly due to 15s timeout): %v", index+1, err))
					currentGenerated = append(currentGenerated, urlStr)
				}
			}
		}

		if len(currentGenerated) == 0 {
			return nil, fmt.Errorf("no valid images found")
		}

		return currentGenerated, nil
	}

	return nil, fmt.Errorf("task failed after retries")
}
