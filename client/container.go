package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

func GetContainers() []map[string]interface{} {
	// 도커 소켓 경로
	dockerSocketPath := "/var/run/docker.sock"

	// Unix 도메인 소켓에 연결
	conn, err := net.Dial("unix", dockerSocketPath)
	if err != nil {
		fmt.Println("Error connecting to Docker socket:", err)
	}
	defer conn.Close()

	// HTTP 요청 생성
	request, err := http.NewRequest("GET", "/containers/json", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	request.Host = "localhost" // 필요한 경우 호스트 설정

	// 요청 전송
	if err := request.Write(conn); err != nil {
		fmt.Println("Error sending request:", err)
	}

	// 응답 읽기
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
	defer response.Body.Close()

	// 응답 출력
	fmt.Println("Response Status:", response.Status)
	fmt.Println("Response Headers:", response.Header)
	fmt.Println("Response body:", response.Body)


	// 응답 본문 파싱 및 출력
	var containers []map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&containers); err != nil {
		fmt.Println("Error decoding response:", err)
	}

	return containers
	/*
	// 컨테이너 목록 출력
	for _, container := range containers {
		fmt.Println("Container ID:", container["Id"])
		fmt.Println("Container Name:", container["Names"])
		fmt.Println("Container Image:", container["Image"])
		fmt.Println("---------")
	}
	*/
}